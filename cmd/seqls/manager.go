package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"text/tabwriter"

	fileseq "github.com/justinfx/gofileseq/v3"
	"github.com/justinfx/gofileseq/v3/cmd/seqls/internal/fastwalk"
)

// number of goroutines to spawn for processing directories
// and building FileSequence results
var numWorkers = 50

// workManager contains the input and output channels
// for all work, as well as the options for processing
type workManager struct {
	inDirs   chan string
	inSeqs   chan *fileseq.FileSequence
	outSeqs  chan fileseq.FileSequences
	fileOpts []fileseq.FileOption
	hasRun   bool
}

// Create a new workManager, with input and output channels,
// with a given list of options
func NewWorkManager() *workManager {
	var fileopts []fileseq.FileOption
	if Options.AllFiles {
		fileopts = append(fileopts, fileseq.HiddenFiles)
	}
	if !Options.SeqsOnly {
		fileopts = append(fileopts, fileseq.SingleFiles)
	}
	if Options.PadHash1 {
		fileopts = append(fileopts, fileseq.FileOptPadStyleHash1)
	}
	if Options.StrictPad {
		fileopts = append(fileopts, fileseq.StrictPadding)
	}
	s := &workManager{
		inDirs:   make(chan string),
		inSeqs:   make(chan *fileseq.FileSequence),
		outSeqs:  make(chan fileseq.FileSequences),
		fileOpts: fileopts,
	}
	return s
}

// Process method starts the workManager, by processing
// the rootPaths, and resolving file/sequence listings
// from them. Output is printed according to application
// Options.
// It is an error to call this method more than once.
func (w *workManager) Process(rootPaths []string) error {
	if w.hasRun {
		return errors.New("Process() cannot be called more than once")
	}

	w.hasRun = true

	var errCount uint64

	// Start the workers to find sequences
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			// Processes work from the input chans
			// and places them into the output chans
			numErrs := w.processSources()
			atomic.AddUint64(&errCount, numErrs)
			wg.Done()
		}()
	}

	// Load the root paths into the source channel
	go func() {
		var numErrs uint64
		if Options.Recurse {
			// Perform a recursive walk on all paths
			numErrs = w.loadRecursive(rootPaths)
		} else {
			// Load each root path directly into chan
			numErrs = w.load(rootPaths)
		}
		atomic.AddUint64(&errCount, numErrs)
		w.closeInputs()
	}()

	// Will close the output channel when no more source
	// paths are being added
	go func() {
		wg.Wait()
		w.closeOutput()
	}()

	// Pull out all processed sequences and print
	w.processResults()

	if errCount > 0 {
		return fmt.Errorf("Error: finished with %d path error(s)", errCount)
	}
	return nil
}

// processSources pulls from input channels, and processes
// them into the output channel until there is no more work
func (w *workManager) processSources() uint64 {
	var (
		ok       bool
		path     string
		errCount uint64
		seq      *fileseq.FileSequence
	)

	fileopts := w.fileOpts

	inDirs := w.inDirs
	inSeqs := w.inSeqs
	outSeqs := w.outSeqs

	isDone := func() bool {
		return (inDirs == nil && inSeqs == nil)
	}

	for !isDone() {
		select {

		// Directory paths will be scanned for contents
		case path, ok = <-inDirs:
			if !ok {
				inDirs = nil
				continue
			}
			seqs, err := fileseq.FindSequencesOnDisk(path, fileopts...)
			if err != nil {
				errCount++
				fmt.Fprintf(errOut, "%s %q: %s\n", ErrorPath, path, err)
				continue
			}
			outSeqs <- seqs

		// Sequence paths will be scanned for a direct match
		// against the sequence pattern
		case seq, ok = <-inSeqs:
			if !ok {
				inSeqs = nil
				continue
			}

			path, err := seq.Format("{{dir}}{{base}}{{pad}}{{ext}}")
			if err != nil {
				fmt.Fprintf(errOut, "%s %q: Not a valid path\n", ErrorPattern, path)
				continue
			}

			seq, err := fileseq.FindSequenceOnDisk(path, fileopts...)
			if err != nil {
				if !os.IsNotExist(err) {
					fmt.Fprintf(errOut, "%s %q: %s\n", ErrorPattern, path, err)
				}
				continue
			}

			if seq != nil {
				outSeqs <- fileseq.FileSequences{seq}
			}
		}
	}
	return errCount
}

// Pull from the output channel until there are no more results,
// and print the results accordingly, dictated by the main options
type WriterFlusher interface {
	io.Writer
	Flush() error
}

func (w *workManager) processResults() {
	var writ WriterFlusher
	var printer func(io.Writer, *fileseq.FileSequence)

	if Options.LongList {
		writ = tabwriter.NewWriter(os.Stdout, 5, 0, 2, ' ', 0)
		printer = printLongListing

	} else {
		writ = bufio.NewWriter(os.Stdout)
		printer = printShortListing
	}

	// Start processing the results that are feeding in the channel
	for seqs := range w.outSeqs {
		for _, s := range seqs {
			printer(writ, s)
		}
	}

	writ.Flush()
}

// Returns true if the input channels are nil
func (w *workManager) isInputDone() bool {
	if w.inDirs != nil {
		return false
	}
	if w.inSeqs != nil {
		return false
	}
	return true
}

// CloseInputs closes the input channels indicating
// that no more paths will be loaded.
func (w *workManager) closeInputs() {
	if w.inDirs != nil {
		close(w.inDirs)

	}
	if w.inSeqs != nil {
		close(w.inSeqs)
	}
}

// Close the output channel indicating that no more
// output is to be expected
func (w *workManager) closeOutput() {
	if w.outSeqs != nil {
		close(w.outSeqs)
	}
}

// loadStandard takes paths and loads them into the
// dir input channel for processing
func (w *workManager) load(paths []string) uint64 {
	dirs, seqs, errCount := preparePaths(paths)

	for _, s := range seqs {
		w.inSeqs <- s
	}

	for _, r := range dirs {
		w.inDirs <- r
	}

	return errCount
}

// Parallel walk the root paths and populate the path
// channel for the worker routines to consume.
func (w *workManager) loadRecursive(paths []string) uint64 {
	cache := map[string]struct{}{}
	var mu sync.RWMutex

	walkFn := func(path string, typ os.FileMode) error {
		var isDir bool
		var ret error

		if typ.IsDir() {
			isDir = true

		} else if typ&os.ModeSymlink != 0 {
			if info, err := os.Stat(path); err == nil && info.IsDir() {
				tgt, err := filepath.EvalSymlinks(path)
				if err != nil {
					return err
				}

				isDir = true
				ret = fastwalk.TraverseLink

				// prevent symlink cycles.
				// don't traverse the same symlink path if a subdirectory
				// has a reference back to a link we have already seen

				// check with read lock
				mu.RLock()
				_, exists := cache[tgt]
				mu.RUnlock()
				if exists {
					ret = fastwalk.SkipFiles
				} else {
					// attempt to cache with write lock
					mu.Lock()
					if _, exists = cache[tgt]; exists {
						ret = fastwalk.SkipFiles
					} else {
						cache[tgt] = struct{}{}
						cache[path] = struct{}{}
					}
					mu.Unlock()
				}

				// We will still allow the immediate file listing
				// of this symlink no matter what. But we may be
				// instructing the walk not to recurse any further.
			}
		}

		if !isDir {
			return ret
		}

		if !Options.AllFiles {
			// Skip traversing into hidden dirs
			name := filepath.Base(path)
			if len(name) > 1 && strings.HasPrefix(name, ".") {
				return filepath.SkipDir
			}
		}

		// Add the path to the input channel for sequence scanning
		w.inDirs <- path
		return ret
	}

	dirs, seqs, errCount := preparePaths(paths)

	for _, s := range seqs {
		w.inSeqs <- s
	}

	for _, r := range dirs {
		if err := fastwalk.Walk(r, walkFn); err != nil {
			if err != fastwalk.SkipFiles && err != fastwalk.TraverseLink {
				atomic.AddUint64(&errCount, 1)
				fmt.Fprintf(errOut, "%s %q: %s\n", ErrorPath, r, err)
			}
		}
	}

	return errCount
}

// Take a list of paths and reduce them to cleaned
// and unique paths. Return two slices, separated by
// directory paths, and sequence patterns
func preparePaths(paths []string) ([]string, fileseq.FileSequences, uint64) {
	var (
		fi       os.FileInfo
		errCount uint64
		err      error
	)

	dirs := make([]string, 0)
	seqs := make(fileseq.FileSequences, 0)
	previous := make(map[string]struct{})

	for _, p := range paths {
		p := strings.TrimSpace(filepath.Clean(p))
		if p == "" {
			continue
		}

		if _, seen := previous[p]; seen {
			continue
		}
		previous[p] = struct{}{}

		if fi, err = os.Stat(p); err != nil {

			// If the path doesn't exist, test it for
			// a valid fileseq pattern
			if seq, err := fileseq.NewFileSequence(p); err == nil {
				seqs = append(seqs, seq)
				continue
			}

			errCount++
			fmt.Fprintf(errOut, "%s %q: %s\n", ErrorPath, p, err)
			continue
		}

		if !fi.IsDir() {
			continue
		}

		dirs = append(dirs, p)
	}

	return dirs, seqs, errCount
}

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
	"runtime"
	"sync/atomic"
	"text/tabwriter"

	fileseq "github.com/justinfx/gofileseq/v3"
)

// number of goroutines to spawn for processing directories
// and building FileSequence results
var numWorkers = 50

// workManager contains the input and output channels
// for all work, as well as the options for processing
type workManager struct {
	inFiles  chan []string
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
		inFiles:  make(chan []string, numWorkers),
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
		files    []string
		errCount uint64
		seq      *fileseq.FileSequence
	)

	fileopts := w.fileOpts

	inFiles := w.inFiles
	inSeqs := w.inSeqs
	outSeqs := w.outSeqs

	isDone := func() bool {
		return (inFiles == nil && inSeqs == nil)
	}

	for !isDone() {
		select {

		// Pre-read file lists are collapsed into sequences without a second dir read.
		case files, ok = <-inFiles:
			if !ok {
				inFiles = nil
				continue
			}
			seqs, err := fileseq.FindSequencesInList(files, fileopts...)
			if err != nil {
				errCount++
				fmt.Fprintf(errOut, "%s: %s\n", ErrorPath, err)
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

// CloseInputs closes the input channels indicating
// that no more paths will be loaded.
func (w *workManager) closeInputs() {
	if w.inFiles != nil {
		close(w.inFiles)
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

// load reads each root directory once and sends its file list for sequence detection.
func (w *workManager) load(paths []string) uint64 {
	dirs, seqs, errCount := preparePaths(paths)

	for _, s := range seqs {
		w.inSeqs <- s
	}

	for _, dir := range dirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			atomic.AddUint64(&errCount, 1)
			fmt.Fprintf(errOut, "%s %q: %s\n", ErrorPath, dir, err)
			continue
		}
		var filePaths []string
		for _, e := range entries {
			if e.IsDir() {
				continue
			}
			typ := e.Type()
			full := filepath.Join(dir, e.Name())
			if typ&os.ModeSymlink != 0 {
				info, err := os.Stat(full)
				if err != nil || info.IsDir() {
					continue
				}
			}
			filePaths = append(filePaths, full)
		}
		if len(filePaths) > 0 {
			w.inFiles <- filePaths
		}
	}

	return errCount
}

// loadRecursive walks root paths concurrently using os.ReadDir, reading each
// directory exactly once. File paths are collected and sent directly to workers
// via inFiles; no second directory read is needed.
func (w *workManager) loadRecursive(paths []string) uint64 {
	var errCount uint64

	sem := make(chan struct{}, runtime.NumCPU()*4)

	// symlink cycle detection
	var mu sync.Mutex
	seen := map[string]struct{}{}

	var wg sync.WaitGroup

	var walkDir func(dir string)
	walkDir = func(dir string) {
		sem <- struct{}{}
		defer func() { <-sem }()
		defer wg.Done()

		entries, err := os.ReadDir(dir)
		if err != nil {
			atomic.AddUint64(&errCount, 1)
			fmt.Fprintf(errOut, "%s %q: %s\n", ErrorPath, dir, err)
			return
		}

		var filePaths []string
		for _, e := range entries {
			name := e.Name()
			full := filepath.Join(dir, name)
			typ := e.Type()

			if typ.IsDir() {
				if !Options.AllFiles && len(name) > 1 && strings.HasPrefix(name, ".") {
					continue
				}
				wg.Add(1)
				go walkDir(full)
				continue
			}

			if typ&os.ModeSymlink != 0 {
				info, err := os.Stat(full)
				if err != nil {
					continue
				}
				if info.IsDir() {
					if !Options.AllFiles && len(name) > 1 && strings.HasPrefix(name, ".") {
						continue
					}
					tgt, err := filepath.EvalSymlinks(full)
					if err != nil {
						continue
					}
					mu.Lock()
					_, exists := seen[tgt]
					if !exists {
						seen[tgt] = struct{}{}
						seen[full] = struct{}{}
					}
					mu.Unlock()
					if !exists {
						wg.Add(1)
						go walkDir(full)
					}
					continue
				}
				// symlink to a regular file — include it
			}

			filePaths = append(filePaths, full)
		}

		if len(filePaths) > 0 {
			w.inFiles <- filePaths
		}
	}

	dirs, seqs, errCount0 := preparePaths(paths)
	atomic.AddUint64(&errCount, errCount0)

	for _, s := range seqs {
		w.inSeqs <- s
	}

	for _, dir := range dirs {
		wg.Add(1)
		go walkDir(dir)
	}

	wg.Wait()
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

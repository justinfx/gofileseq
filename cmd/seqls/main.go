/*
seqls - list directory contents, rolled up into file sequences
*/
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"text/tabwriter"
	"time"

	"github.com/MichaelTJones/walk"
	"github.com/jessevdk/go-flags"
	"github.com/justinfx/gofileseq"
)

const Version = "0.9.0"

// number of goroutines to spawn for processing directories
// and building FileSequence results
var numWorkers = 50

var opts struct {
	Recurse   bool `short:"r" long:"recurse" description:"recursively scan all sub-directories"`
	AllFiles  bool `short:"a" long:"all" description:"list all files, including those without frame patterns"`
	LongList  bool `short:"l" long:"long" description:"long listing; include extra stat information"`
	AbsPath   bool `short:"f" long:"full" description:"show absolute paths"`
	HumanSize bool `short:"H" long:"human" description:"when using long listing, show human-readable file size units"`
}

const DateFmt = `Jan _2 15:04`

func init() {
	// If $GOMAXPROCS isn't set, use the full capacity of the machine.
	// For small machines, use at least 4 threads.
	if os.Getenv("GOMAXPROCS") == "" {
		n := runtime.NumCPU()
		if n < 4 {
			n = 4
		}
		runtime.GOMAXPROCS(n)
	}
}

var usage = fmt.Sprintf(`[options] [path [...]]

Version: %s 

List directory contents, rolled up into file sequences.

One or more paths may be provided. If no path is provided, the current
working directory will be scanned.

Only files are shown in listing results.
`, Version)

func main() {
	parser := flags.NewParser(&opts, flags.HelpFlag|flags.PrintErrors)
	parser.Name = "seqls"
	parser.Usage = usage
	roots, err := parser.Parse()

	if err != nil {
		os.Exit(1)
	}

	if len(roots) == 0 {
		// Use the current directory if specific dirs were not passed
		roots = []string{"."}
	}

	var wg sync.WaitGroup

	pathChan := make(chan string)
	seqChan := make(chan fileseq.FileSequences)

	// fmt.Printf("launching %d workers\n", numWorkers)

	var listerFn func(string) (fileseq.FileSequences, error)
	if opts.AllFiles {
		listerFn = fileseq.ListFiles
	} else {
		listerFn = fileseq.FindSequencesOnDisk
	}

	// Start the workers to find sequences
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			for path := range pathChan {
				seqs, err := listerFn(path)
				if err != nil {
					// Bail out of the app for any path error
					fmt.Fprintf(os.Stderr, "Error finding sequence in dir %q: %s\n", path, err)
					os.Exit(255)
				}
				seqChan <- seqs
			}
			wg.Done()
		}()
	}

	// Load the root paths into the source channel
	unique := cleanDirs(roots)
	go func() {
		if opts.Recurse {
			loadRecursive(unique, pathChan)
		} else {
			for r := range unique {
				pathChan <- r
			}
		}
		close(pathChan)
	}()

	// Will close the output channel when no more source
	// paths are being added
	go func() {
		wg.Wait()
		close(seqChan)
	}()

	// Pull out all processed sequences and print
	var w io.Writer
	if opts.LongList {
		w = tabwriter.NewWriter(os.Stdout, 5, 0, 2, ' ', 0)
	} else {
		w = os.Stdout
	}

	var printer func(io.Writer, *fileseq.FileSequence)
	if opts.LongList {
		printer = printLongListing
	} else {
		printer = printShortListing
	}

	// Start processing the results that are feeding in the channel
	for seqs := range seqChan {
		for _, s := range seqs {
			printer(w, s)
		}
	}

	if opts.LongList {
		// Flush the tabwriter buffer
		w.(*tabwriter.Writer).Flush()
	}

	os.Exit(0)
}

// Parallel walk the root paths and populate the path
// channel for the worker routines to consume.
func loadRecursive(roots uniquePaths, out chan string) {
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			out <- path

		} else if fi, err := os.Stat(path); err == nil && fi.IsDir() {
			out <- path
		}
		return nil
	}
	for r := range roots {
		r := r
		if err := walk.Walk(r, walkFn); err != nil {
			fmt.Fprintf(os.Stderr, "Error: Failed to read dir %q: %s\n", r, err)
		}
	}
}

type uniquePaths map[string]struct{}

// Take a list of paths and reduce them to cleaned
// and unique paths.
func cleanDirs(paths []string) uniquePaths {
	u := make(uniquePaths)
	var (
		fi  os.FileInfo
		err error
	)
	for _, p := range paths {
		p := strings.TrimSpace(filepath.Clean(p))
		if p == "" {
			continue
		}
		if fi, err = os.Stat(p); err != nil {
			fmt.Fprintf(os.Stderr, "Error: Failed to read path %q: %s\n", p, err)
			continue
		}
		if !fi.IsDir() {
			continue
		}
		u[p] = struct{}{}
	}
	return u
}

func printShortListing(w io.Writer, fs *fileseq.FileSequence) {
	str := fs.String()
	if opts.AbsPath {
		abs, err := filepath.Abs(str)
		if err == nil {
			str = abs
		}
	}
	fmt.Fprintln(w, str)
}

func printLongListing(w io.Writer, fs *fileseq.FileSequence) {
	var (
		err      error
		byteSize int64
		stat     os.FileInfo
		str      string
	)

	count := fs.Len()
	if count == 1 {
		// Only a single file
		str = fs.Index(0)
		stat, err = os.Stat(str)
		if err == nil {
			byteSize = stat.Size()
		}

	} else {
		str = fs.String()
		// For a sequence of files, we want to get the most
		// recent modtime, as well as the total size of the
		// combined files.
		mod := time.Unix(0, 0)
		var thisStat os.FileInfo
		var thisStr string
		for i := 0; i < count; i++ {
			thisStr = fs.Index(i)
			thisStat, err = os.Stat(thisStr)
			if err != nil {
				continue
			}
			byteSize += thisStat.Size()
			if thisStat.ModTime().After(mod) {
				stat = thisStat
				mod = thisStat.ModTime()
			}
		}
	}

	if opts.AbsPath {
		abs, err := filepath.Abs(str)
		if err == nil {
			str = abs
		}
	}

	if err != nil {
		fmt.Println(str)
		fmt.Fprintf(os.Stderr, "Error stating file: %s\n", err.Error())
		return
	}

	usr, gid := uidGidFromFileInfo(stat)

	var size interface{}
	if opts.HumanSize {
		size = ByteSize(byteSize)
	} else {
		size = byteSize
	}

	fmt.Fprintf(w, "%s\t%s\t%s\t%v\t%s\t%s\n",
		stat.Mode(),
		usr,
		gid,
		size,
		stat.ModTime().Format(DateFmt),
		str)

}

type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
)

func (b ByteSize) String() string {
	switch {
	case b >= TB:
		return fmt.Sprintf("%.1fT", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.1fG", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.1fM", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.1fK", b/KB)
	case b == 0:
		return strconv.Itoa(0)
	}
	return fmt.Sprintf("%.0f", b)
}

/*
seqls - list directory contents, rolled up into file sequences
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/MichaelTJones/walk"
	"github.com/justinfx/gofileseq"
)

var (
	numWorkers = 50
	recurse    = flag.Bool("r", false, `recursively scan all sub-directories`)
)

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

const usage = `seqls: list directory contents, rolled up into file sequences

Usage:  seqls [options] [path [...]]

One or more paths may be provided. If no path is provided, the current
working directory will be scanned.
`

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usage)
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\n")
		os.Exit(1)
	}
	flag.Parse()

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	var wg sync.WaitGroup

	pathChan := make(chan string)
	seqChan := make(chan fileseq.FileSequences)

	// fmt.Printf("launching %d workers\n", numWorkers)

	// Start the workers to find sequences
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			for path := range pathChan {
				seqs, err := fileseq.FindSequencesOnDisk(path)
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
		if *recurse {
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
	for seqs := range seqChan {
		for _, s := range seqs {
			fmt.Println(s)
		}
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

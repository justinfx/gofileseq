/*
seqls - list directory contents, rolled up into file sequences
*/
package main // import "github.com/justinfx/gofileseq/v2/cmd/seqls"

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/jessevdk/go-flags"
	fileseq "github.com/justinfx/gofileseq/v2"
)

var Options struct {
	Recurse   bool `short:"r" long:"recurse" description:"Recursively scan all sub-directories"`
	AllFiles  bool `short:"a" long:"all" description:"List all files, including hidden files/dirs and single (non-sequence) files"`
	SeqsOnly  bool `short:"s" long:"seqs" description:"Only return sequences, and not single files that didn't have a frame number"`
	LongList  bool `short:"l" long:"long" description:"Long listing; include extra stat information"`
	AbsPath   bool `short:"f" long:"full" description:"Show absolute paths"`
	HumanSize bool `short:"H" long:"human" description:"When using long listing, show human-readable file size units"`
	StrictPad bool `short:"S" long:"strict"`
	PadHash1  bool `long:"hash1" description:"Use a pad width of 1 for the # char (like Nuke) instead of 4 (Katana)"`
	Quiet     bool `short:"q" long:"quiet" description:"Don't print errors encountered when reading the file system"`
}

var Version = fileseq.Version

const (
	DateFmt      = `Jan _2 15:04`
	ErrorPath    = `Error: Failed reading path`
	ErrorPattern = `Error: Failed matching pattern`
)

// The default error output goes to stderr
var errOut io.Writer = os.Stderr

var usage = fmt.Sprintf(`[options] [path [...]]

Version: %s 

List directory contents, rolled up into file sequences.

One or more paths may be provided. If no path is provided, the current
working directory will be scanned.

Only files are shown in listing results. By default, hidden files are 
skipped, and hidden directories are ignored. Flags may be passed to control 
showing hidden files, and files in hidden directories. 
`, Version)

func main() {
	// If $GOMAXPROCS isn't set, use the full capacity of the machine.
	// For small machines, use at least 4 threads.
	if os.Getenv("GOMAXPROCS") == "" {
		n := runtime.NumCPU()
		if n < 4 {
			n = 4
		}
		runtime.GOMAXPROCS(n)
	}

	parser := flags.NewParser(&Options, flags.HelpFlag|flags.PrintErrors)
	parser.Name = "seqls"
	parser.Usage = usage
	if opt := parser.FindOptionByLongName("strict"); opt != nil {
		opt.Description = `Use exact length of search string padding chars` +
			` instead of treating as a general wildcard for any frame number`
	}
	roots, err := parser.Parse()

	if err != nil {
		os.Exit(1)
	}

	if Options.Quiet {
		errOut = ioutil.Discard
	}

	if len(roots) == 0 {
		// Use the current directory if specific dirs were not passed
		roots = []string{"."}
	}

	manager := NewWorkManager()
	manager.Process(roots)
}

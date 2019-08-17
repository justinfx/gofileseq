package main // import "github.com/justinfx/gofileseq/v2/cmd/seqinfo"

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"

	"github.com/jessevdk/go-flags"
	"github.com/justinfx/gofileseq/v2"
)

var Version = fileseq.Version

// OptionsType collects flag options
type OptionsType struct {
	Dirname  string `short:"d" long:"dirname" description:"Set a new dir name"`
	Basename string `short:"b" long:"basename" description:"Set a new base name"`
	Range    string `short:"r" long:"range" description:"Set a new frame range (i.e. '1-100')"`
	Padding  string `short:"p" long:"padding" description:"Set new padding characters (i.e. '##')"`
	Ext      string `short:"e" long:"ext" description:"Set a new ext"`
	Format   string `long:"format" description:"Reformat using a template"`

	Index *int `short:"i" long:"index" description:"Return the frame at the given index of the sequence"`
	Frame *int `short:"f" long:"frame" description:"Return a path with the given frame replaced into the pattern"`

	Inverted bool `long:"inverted" description:"Invert the frame range"`

	Json  bool `short:"j" long:"json" description:"Output info in json format"`
	Hash1 bool `long:"hash1" description:"Use a pad width of 1 for the # char (like Nuke) instead of 4 (Katana)"`
}

var Options OptionsType

var usage = fmt.Sprintf(`[options] [[file seq pattern [...]]

Version: %s

Parse and list information about given file sequence patterns.
Patterns can be supplied either as arguments, or via stdin (newline separated).

Examples:
	seqinfo file.1-100#.ext
	echo "file.1-100#.ext\nfile2.1-100x2#.ext" | seqinfo -j

Format (--format) Options:
    Utilizes Go text/template format syntax. Available functions include:

    dir      - the directory name.
    base     - the basename of the sequence (leading up to the frame range).
    ext      - the file extension of the sequence including leading period.
    startf   - the start frame.
    endf     - the end frame.
    len      - the length of the frame range.
    pad      - the detected padding characters (i.e. # , @@@ , ...).
    frange   - the frame range.
    inverted - the inverted frame range. (returns empty string if none)
    zfill    - the int width of the frame padding

    Example:
        {{dir}}{{base}}{{frange}}{{pad}}{{ext}}
`, Version)

func main() {
	parser := flags.NewParser(&Options, flags.HelpFlag|flags.PrintErrors)
	parser.Name = "seqinfo"
	parser.Usage = usage
	patterns, err := parser.Parse()

	if err != nil {
		os.Exit(1)
	}

	// See if we are being given paths from stdin
	stat, err := os.Stdin.Stat()
	if err == nil && stat.Mode()&os.ModeCharDevice == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				continue
			}
			patterns = append(patterns, line)
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error: Failed to read from stdin:", err)
			os.Exit(1)
		}
	}

	n := len(patterns)
	if n == 0 {
		parser.WriteHelp(os.Stderr)
		fmt.Fprintln(os.Stderr, "\nError: At least 1 file sequence pattern must be given")
		os.Exit(1)
	}

	// Decide how to output the results
	var printer ResultPrinter = printPlainResults
	if Options.Json {
		printer = printJsonResults
	}

	out := make(chan *Result, n)
	results := make(Results, n)

	// Process all strings, concurrently
	for _, pat := range patterns {
		go func(pat string) {
			out <- parse(pat, &Options)
		}(pat)
	}

	// Collect
	for i := 0; i < n; i++ {
		res := <-out
		results[res.origString] = res
	}

	// Report
	if err = printer(results); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}

// Result is the data to be gathered from a FileSequence
type Result struct {
	origString string
	Error      string `json:"error"`
	String     string `json:"string"`
	Dirname    string `json:"dir"`
	Basename   string `json:"base"`
	Range      string `json:"range"`
	Padding    string `json:"pad"`
	Ext        string `json:"ext"`
	Start      int    `json:"start"`
	End        int    `json:"end"`
	Len        int    `json:"length"`
	ZFill      int    `json:"zfill"`
	HasRange   bool   `json:"hasRange"`
}

// Results is a collection of all parsed sequences
type Results map[string]*Result

// ResultPrinter is a function that can print results
type ResultPrinter func(Results) error

// printPlainResults prints plain-text output for results
func printPlainResults(results Results) error {
	for _, res := range results {
		// Explicitely start with the string and error output
		fmt.Printf("Source = %s\n", res.origString)
		fmt.Printf("    String = %s\n", res.String)

		if res.Error != "" {
			fmt.Printf("    Error = %s\n", res.Error)
			continue
		}

		// Dynamically loop over the rest of the fields
		typ := reflect.TypeOf(*res)
		val := reflect.ValueOf(*res)

		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			if field.Name == "Error" || field.Name == "String" {
				continue
			}
			if field.PkgPath != "" {
				// ignore unexported fields
				continue
			}
			fmt.Printf("    %s = %v\n", field.Name, val.Field(i).Interface())
		}

		fmt.Print("\n")
	}
	return nil
}

// printJsonResults prints json-formatted output for results
func printJsonResults(results Results) error {
	data, err := json.MarshalIndent(results, "", "    ")
	if err != nil {
		return fmt.Errorf("Failed to convert results to JSON: %s", err.Error())
	}

	if _, err = io.Copy(os.Stdout, bytes.NewReader(data)); err != nil {
		return fmt.Errorf("Failed to write json output: %s", err.Error())
	}

	fmt.Print("\n")
	return nil
}

// parse takes a file sequence pattern and returns the parsed
// data as a Result
func parse(pattern string, opts *OptionsType) *Result {
	res := &Result{String: pattern, origString: pattern}

	padStyle := fileseq.PadStyleHash4
	if opts.Hash1 {
		padStyle = fileseq.PadStyleHash1
	}

	fs, err := fileseq.NewFileSequencePad(pattern, padStyle)
	if err != nil {
		res.Error = err.Error()
		return res
	}

	// Handle any of the optional modifications to the seq
	if opts.Format != "" {
		refmt, err := fs.Format(opts.Format)
		if err != nil {
			res.Error = err.Error()
			return res
		}
		fs, err = fileseq.NewFileSequencePad(refmt, padStyle)
		if err != nil {
			res.Error = err.Error()
			return res
		}
	}

	if opts.Dirname != "" {
		fs.SetDirname(opts.Dirname)
	}

	if opts.Basename != "" {
		fs.SetBasename(opts.Basename)
	}

	if opts.Ext != "" {
		fs.SetExt(opts.Ext)
	}

	if opts.Padding != "" {
		fs.SetPadding(opts.Padding)
	}

	if opts.Range != "" {
		if err := fs.SetFrameRange(opts.Range); err != nil {
			res.Error = err.Error()
			return res
		}
	}

	if opts.Inverted {
		frange := fs.InvertedFrameRange()
		if frange != "" {
			_ = fs.SetFrameRange(frange)
		} else {
			fs.SetFrameSet(nil)
		}
	}

	if opts.Index != nil {
		frame := fs.Index(*opts.Index)
		if frame == "" {
			res.Error = fmt.Sprintf("Index %d out of range %q", *opts.Index, fs.FrameRange())
			return res
		}
		fs, _ = fileseq.NewFileSequencePad(frame, padStyle)
		_ = fs.SetFrameRange(strconv.Itoa(fs.Start()))
	}

	if opts.Frame != nil {
		frame, _ := fs.Frame(*opts.Frame)
		fs, _ = fileseq.NewFileSequencePad(frame, padStyle)
		_ = fs.SetFrameRange(strconv.Itoa(fs.Start()))
	}

	// Final setting of results
	res.String = fs.String()
	res.Dirname = fs.Dirname()
	res.Basename = fs.Basename()
	res.Ext = fs.Ext()
	res.Start = fs.Start()
	res.End = fs.End()
	res.Len = fs.Len()
	res.Padding = fs.Padding()
	res.ZFill = fs.ZFill()
	res.Range = fs.FrameRange()
	res.HasRange = (fs.FrameSet() != nil)

	return res
}

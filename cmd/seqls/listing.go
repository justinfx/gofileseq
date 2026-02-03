package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/justinfx/gofileseq/v3"
)

func printShortListing(w io.Writer, fs *fileseq.FileSequence) {
	str := fs.String()
	if Options.AbsPath {
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

	if Options.AbsPath {
		abs, err := filepath.Abs(str)
		if err == nil {
			str = abs
		}
	}

	if err != nil {
		fmt.Println(str)
		fmt.Fprintf(errOut, "%s: %s\n", ErrorPath, err.Error())
		return
	}

	var size interface{}
	if Options.HumanSize {
		size = ByteSize(byteSize)
	} else {
		size = byteSize
	}

	printDetails(w, str, stat, size)
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

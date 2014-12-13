// +build windows

package main

import (
	"fmt"
	"io"
	"os"
)

func printDetails(w io.Writer, path string, stat os.FileInfo, totalSize interface{}) {
	fmt.Fprintf(w, "%s\t%v\t%s\t%s\n",
		stat.Mode(),
		totalSize,
		stat.ModTime().Format(DateFmt),
		path)
}

// +build !windows

package main

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

func printDetails(w io.Writer, path string, stat os.FileInfo, totalSize interface{}) {
	stat_t := stat.Sys().(*syscall.Stat_t)
	usr := strconv.Itoa(int(stat_t.Uid))
	gid := "-"

	uObj, err := user.LookupId(usr)
	if err == nil {
		usr = uObj.Username
		gid = uObj.Gid
	}

	fmt.Fprintf(w, "%s\t%s\t%s\t%v\t%s\t%s\n",
		stat.Mode(),
		usr,
		gid,
		totalSize,
		stat.ModTime().Format(DateFmt),
		path)
}

// +build !windows

package main

import (
	"os"
	"os/user"
	"strconv"
	"syscall"
)

func uidGidFromFileInfo(stat os.FileInfo) (string, string) {
	stat_t := stat.Sys().(*syscall.Stat_t)
	usr := strconv.Itoa(int(stat_t.Uid))
	gid := "-"

	uObj, err := user.LookupId(usr)
	if err == nil {
		usr = uObj.Username
		gid = uObj.Gid
	}

	return usr, gid
}

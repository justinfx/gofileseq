// +build windows

package main

import "os"

func uidGidFromFileInfo(stat os.FileInfo) (string, string) {
	return "", ""
}

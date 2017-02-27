package scp

import (
	"os"
	"syscall"
	"time"
)

func newFileInfoFromOS(fi os.FileInfo, replaceName string) *FileInfo {
	var name string
	if replaceName == "" {
		name = fi.Name()
	} else {
		name = replaceName
	}

	modTime := fi.ModTime()

	var accessTime time.Time
	sysStat, ok := fi.Sys().(*syscall.Win32FileAttributeData)
	if ok {
		accessTime = time.Unix(0, sysStat.LastAccessTime.Nanoseconds())
	}

	return NewFileInfo(name, fi.Size(), fi.Mode(), modTime, accessTime)
}

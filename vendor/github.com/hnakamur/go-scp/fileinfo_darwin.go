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
	sysStat, ok := fi.Sys().(*syscall.Stat_t)
	if ok {
		sec, nsec := sysStat.Atimespec.Unix()
		accessTime = time.Unix(sec, nsec)
	}

	return NewFileInfo(name, fi.Size(), fi.Mode(), modTime, accessTime)
}

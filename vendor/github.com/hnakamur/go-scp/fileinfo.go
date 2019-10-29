package scp

import (
	"os"
	"path/filepath"
	"time"
)

// FileInfo represents a file or a directory information.
type FileInfo struct {
	name       string
	size       int64
	mode       os.FileMode
	modTime    time.Time
	accessTime time.Time
}

// NewFileInfo creates a file information. The filepath.Base(name) is
// used as the name and the mode is masked with (os.ModePerm | os.ModeDir).
func NewFileInfo(name string, size int64, mode os.FileMode, modTime, accessTime time.Time) *FileInfo {
	return &FileInfo{
		name:       filepath.Base(name),
		size:       size,
		mode:       mode & (os.ModePerm | os.ModeDir),
		modTime:    modTime,
		accessTime: accessTime,
	}
}

// Name returns base name of the file.
func (i *FileInfo) Name() string { return i.name }

// Size length in bytes for regular files; system-dependent for others.
func (i *FileInfo) Size() int64 { return i.size }

// Mode returns file mode bits.
func (i *FileInfo) Mode() os.FileMode { return i.mode }

// ModTime returns modification time.
func (i *FileInfo) ModTime() time.Time { return i.modTime }

// IsDir is abbreviation for Mode().IsDir().
func (i *FileInfo) IsDir() bool { return i.Mode().IsDir() }

// Sys returns underlying data source (can return nil).
func (i *FileInfo) Sys() interface{} { return i }

// AccessTime returns access time.
func (i *FileInfo) AccessTime() time.Time { return i.accessTime }

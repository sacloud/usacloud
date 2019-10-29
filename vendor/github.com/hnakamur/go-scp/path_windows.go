// +build windows

package scp

import (
	"path/filepath"
	"strings"
)

func realPath(path string) string {
	return strings.Replace(path, string(filepath.Separator), "/", -1)
}

package scp

import "strings"

func escapeShellArg(arg string) string {
	return "'" + strings.Replace(arg, "'", `'\''`, -1) + "'"
}

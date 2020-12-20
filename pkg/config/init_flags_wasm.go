// +build wasm

package config

import "github.com/spf13/pflag"

// InitConfig 指定のFlagSetにConfigへ値を設定するためのフラグを登録する
func InitConfig(flags *pflag.FlagSet) {
	initCredentialConfig(flags)
	initOutputConfig(flags)
	initDebugConfig(flags)
	// misc flags
	flags.IntP("process-timeout-sec", "", 0, "number of seconds before the command execution is timed out")
	flags.BoolP("version", "v", false, "show version info")
}

func initCredentialConfig(fs *pflag.FlagSet) {
	fs.StringP("token", "", "", "the API token used when calling SAKURA Cloud API")
	fs.StringP("secret", "", "", "the API secret used when calling SAKURA Cloud API")
}

func initOutputConfig(fs *pflag.FlagSet) {
	fs.BoolP("no-color", "", false, "disable ANSI color output")
}

func initDebugConfig(fs *pflag.FlagSet) {
	fs.BoolP("fake", "", false, "enable fake API driver")
}

// Copyright 2017-2021 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build !wasm
// +build !wasm

package config

import "github.com/spf13/pflag"

// InitConfig 指定のFlagSetにConfigへ値を設定するためのフラグを登録する
func InitConfig(flags *pflag.FlagSet) {
	initCredentialConfig(flags)
	initOutputConfig(flags)
	initDebugConfig(flags)
	initBehaviorConfig(flags)
	// misc flags
	flags.BoolP("version", "v", false, "show version info")
}

func initCredentialConfig(fs *pflag.FlagSet) {
	fs.StringP("profile", "", "", "the name of saved credentials")
	fs.StringP("token", "", "", "the API token used when calling SAKURA Cloud API")
	fs.StringP("secret", "", "", "the API secret used when calling SAKURA Cloud API")
	fs.StringSliceP("zones", "", []string{}, "permitted zone names")
}

func initOutputConfig(fs *pflag.FlagSet) {
	fs.BoolP("no-color", "", false, "disable ANSI color output")
}

func initDebugConfig(fs *pflag.FlagSet) {
	fs.BoolP("trace", "", false, "enable trace logs for API calling")
	fs.BoolP("fake", "", false, "enable fake API driver")
	fs.StringP("fake-store", "", "", "path to file store used by the fake API driver")
}

func initBehaviorConfig(fs *pflag.FlagSet) {
	fs.IntP("process-timeout-sec", "", 0, "number of seconds before the command execution is timed out")
	fs.StringP("argument-match-mode", "", "", "how to compare the argument and resource name when identifying the resource to be manipulated  options: [partial/exact]")
}

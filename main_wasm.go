// Copyright 2017-2025 The sacloud/usacloud Authors
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

//go:build js && wasm

package main

import (
	"bytes"
	"os"
	"syscall/js"

	"github.com/sacloud/usacloud/pkg"
)

func main() {
	js.Global().Set("usacloudSetEnv", js.FuncOf(setEnv))
	js.Global().Set("usacloudRun", js.FuncOf(runCommand))
	<-make(chan struct{})
}

// setEnv はJavaScript側から環境変数を設定できるようにする。
// 使用例: usacloudSetEnv("SAKURACLOUD_ACCESS_TOKEN", "your-token")
func setEnv(this js.Value, args []js.Value) interface{} {
	if len(args) < 2 {
		return nil
	}
	os.Setenv(args[0].String(), args[1].String())
	return nil
}

// runCommand はJavaScript側からusacloudコマンドを実行する。
// goroutineで非同期実行し、Promiseを返す。
//
// 使用例: await usacloudRun(["server", "list"], onStdout, onStderr)
//
// args[0]: コマンド引数の配列 (JS Array of strings)
// args[1]: stdout コールバック function(string) (optional)
// args[2]: stderr コールバック function(string) (optional)
//
// 返り値: Promise<number> 終了コード
func runCommand(this js.Value, args []js.Value) interface{} {
	var cmdArgs []string
	if len(args) > 0 && !args[0].IsNull() && !args[0].IsUndefined() {
		length := args[0].Length()
		for i := 0; i < length; i++ {
			cmdArgs = append(cmdArgs, args[0].Index(i).String())
		}
	}

	var onStdout, onStderr js.Value
	if len(args) > 1 {
		onStdout = args[1]
	}
	if len(args) > 2 {
		onStderr = args[2]
	}

	// Promiseを返し、goroutineで非同期実行する。
	// js.FuncOf のコールバックは JSイベントループをブロックするため、
	// I/Oを伴うコマンド実行はgoroutineに委譲する必要がある。
	promiseCtor := js.Global().Get("Promise")
	return promiseCtor.New(js.FuncOf(func(this js.Value, resolveArgs []js.Value) interface{} {
		resolve := resolveArgs[0]
		go func() {
			var outBuf, errBuf bytes.Buffer
			exitCode := pkg.RunWasm(cmdArgs, &outBuf, &errBuf)

			if !onStdout.IsNull() && !onStdout.IsUndefined() && outBuf.Len() > 0 {
				onStdout.Invoke(outBuf.String())
			}
			if !onStderr.IsNull() && !onStderr.IsUndefined() && errBuf.Len() > 0 {
				onStderr.Invoke(errBuf.String())
			}
			resolve.Invoke(exitCode)
		}()
		return nil
	}))
}

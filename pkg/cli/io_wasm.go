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

//go:build js

package cli

import (
	"io"
	"os"
)

// wasmOut/wasmErr は RunWasm から SetWasmIO() で設定される。
// コマンドごとに書き換えられるため、並列実行は想定しない。
var wasmOut io.Writer = os.Stdout
var wasmErr io.Writer = os.Stderr

// SetWasmIO はWASMコマンド実行前に出力先を設定する。
func SetWasmIO(out, err io.Writer) {
	wasmOut = out
	wasmErr = err
}

func newIO() IO {
	return &cliIO{
		in:       os.Stdin,
		out:      wasmOut,
		progress: io.Discard, // WASMではプログレス表示なし
		err:      wasmErr,
	}
}

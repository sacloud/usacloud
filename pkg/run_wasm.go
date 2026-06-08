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

package pkg

import (
	"fmt"
	"io"
	"log"
	"os"
	"slices"

	iaas "github.com/sacloud/iaas-api-go"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/commands/root"
	"github.com/sacloud/usacloud/pkg/config"
)

// RunWasm はWASM環境でusacloudコマンドを実行する。
// os.Exitは呼ばず、終了コードを返す。
func RunWasm(args []string, outW, errW io.Writer) (exitCode int) {
	os.Args = append([]string{"usacloud"}, args...)

	// ブラウザWASMにはホームディレクトリがないため、プロファイルディレクトリを仮設定する。
	// saclient-go の lookupProfileDir が HOME or SAKURACLOUD_PROFILE_DIR を要求するため。
	if os.Getenv("HOME") == "" && os.Getenv("SAKURACLOUD_PROFILE_DIR") == "" {
		os.Setenv("SAKURACLOUD_PROFILE_DIR", "/nonexistent-wasm-profile")
	}

	// SAKURACLOUD_API_ROOT 環境変数でAPIのベースURLを上書きできる。
	// ローカル開発時にプロキシサーバー経由にするために使う。
	// 例: usacloudSetEnv("SAKURACLOUD_API_ROOT", "http://localhost:8080/cloud/zone")
	if v := os.Getenv("SAKURACLOUD_API_ROOT"); v != "" {
		iaas.SakuraCloudAPIRoot = v
	}

	// WASM環境ではinit()でSetEnvironを呼べないため、コマンド実行前に呼ぶ。
	// usacloudSetEnv() で設定された環境変数がここで反映される。
	if err := config.TheClient.SetEnviron(slices.Clone(os.Environ())); err != nil {
		log.Printf("Failed to load environment variables: %s", err)
	}

	// cli.IO の出力先を設定する。
	// newIO() が colorable.NewColorableStdout() ではなくこちらを使うようになる。
	cli.SetWasmIO(outW, errW)

	root.Command.SetOut(outW)
	root.Command.SetErr(errW)

	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(errW, "panic: %v\n", r)
			exitCode = 1
		}
		root.Command.SetOut(nil)
		root.Command.SetErr(nil)
	}()

	initCommands()
	if err := root.Command.Execute(); err != nil {
		return 1
	}
	return 0
}

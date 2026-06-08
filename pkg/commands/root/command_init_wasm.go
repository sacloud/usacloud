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

package root

import (
	"flag"

	"github.com/sacloud/usacloud/pkg/config"
)

func init() {
	Command.Flags().SortFlags = false
	Command.PersistentFlags().SortFlags = false

	// WASM環境ではバージョン確認を無効化する。
	// alertNewVersionReleased はファイルシステムアクセスを行うが、
	// js.FuncOf の同期コールバック実行中はJSイベントループが止まるためデッドロックする。
	Command.PersistentPreRun = nil

	// SetEnviron はホームディレクトリを参照するため、WASM環境では RunWasm 内で呼ぶ。
	// ここではフラグの初期化だけ行う。
	config.InitConfig(Command.PersistentFlags())
	Command.PersistentFlags().AddGoFlagSet(config.TheClient.FlagSet(flag.ContinueOnError))
}

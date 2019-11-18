// Copyright 2016-2019 The Libsacloud Authors
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

// Package builder is the High Level APIs for creating resources on SakuraCloud.
//
// さくらのクラウドでのリソース作成用の高レベルAPIです。
// サーバー/ディスク作成時の手順を単純化します。
//
// Building resources
//
// リソースの作成は以下のように行います。
//	import (
//		"github.com/sacloud/libsacloud/api"
//		"github.com/sacloud/libsacloud/builder"
//		"github.com/sacloud/libsacloud/sacloud/ostype"
//	)
//
//	func main() {
//
//		// APIクライアントの作成
//		client := api.NewClient("PUT-YOUR-TOKEN", "PUT-YOUR-SECRET", "tk1a")
//
//		// パブリックアーカイブ(CentOS)から作成するビルダー、共有セグメントに接続、以外はデフォルト値で作成
//      b := builder.serverPublicArchiveUnix(builder.NewAPIClient(client), ostype.CentOS, "ServerName", "Password")
//      b.AddPublicNWConnectedNIC()
//		res , err := b.WithAddPublicNWConnectedNIC().Build()
//
//		if err != nil {
//			panic(err)
//		}
//		fmt.Printf("%v" , res.Server)
//	}
// 1) 作成したいサーバーのディスク/ソースアーカイブの種類ごとにビルダーを作成します。
//
// 2) 必要に応じてNICやディスク、サーバースペックなどをビルダーのメソッドで設定します。
//
// 3) Buildメソッドを呼び出すことでサーバーが作成されます。
//
//
// Server builder types
//
// ビルダーはディスク構成やソースアーカイブ/ディスクにより以下のような種類に分かれています。
//
// それぞれに対応するビルダー作成用関数を用意しています。
//
// - Linux(Unix)系パブリックアーカイブ
//	// ビルダー
//	type PublicArchiveUnixServerBuilder interface { ... }
//
//	// ビルダー作成用関数
//	func serverPublicArchiveUnix(client APIClient, os ostype.ArchiveOSTypes, name string, password string) PublicArchiveUnixServerBuilder
//
// - Windows系パブリックアーカイブ
//	// ビルダー
//	type PublicArchiveWindowsServerBuilder interface { ... }
//
//	// ビルダー作成用関数
//	func serverPublicArchiveWindows(client APIClient, name string, archiveID int64) PublicArchiveWindowsServerBuilder
//
// - 汎用
//	// ビルダー
//	type CommonServerBuilder interface { ... }
//
//	// ビルダー作成用関数(アーカイブから作成)
//	func serverFromArchive(client APIClient, name string, sourceArchiveID int64) CommonServerBuilder
//
//	// ビルダー作成用関数(ディスクから作成)
//	func serverFromDisk(client APIClient, name string, sourceDiskID int64) CommonServerBuilder
//
// - ディスクレス
//	// ビルダー
//	type DisklessServerBuilder interface { ... }
//
//	// ビルダー作成用関数
//	func ServerDiskless(client APIClient, name string) DisklessServerBuilder
//
// - 空のディスク
//	// ビルダー
//	type BlankDiskServerBuilder interface { ... }
//
//	// ビルダー作成用関数
//	func ServerBlankDisk(client APIClient, name string) BlankDiskServerBuilder
//
//
//
// Event handling
//
// ビルダーでは各タイミングごとにイベントハンドラ(コールバック)が利用可能です。
//	func main() {
//		// APIクライアントの作成
//		client := api.NewClient("PUT-YOUR-TOKEN", "PUT-YOUR-SECRET", "tk1a")
//
//		// ディスクレスビルダー、イベントハンドラ(ServerBuildOnComplete)を登録
//		b := builder.ServerDiskless(builder.NewAPIClient(client), "example")
//		b.SetEventHandler(builder.ServerBuildOnComplete, callbackFunc).
//		b.Build()
//	}
//
//	func callbackFunc(value *builder.ServerBuildValue, result *builder.ServerBuildResult) {
//		// Do someting here
//	}
//
package builder

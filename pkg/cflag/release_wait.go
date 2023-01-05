// Copyright 2017-2023 The sacloud/usacloud Authors
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

package cflag

type WaitForReleaseParameter struct {
	WaitForRelease        bool `cli:",category=wait,order=10"` // trueの場合、他リソースから参照されている間は削除を待ち合わせし続ける
	WaitForReleaseTimeout int  `cli:",category=wait,order=20"` // WaitForReleaseがtrueの場合の待ち時間タイムアウト(デフォルト:1時間)
	WaitForReleaseTick    int  `cli:",category=wait,order=30"` // WaitForReleaseがtrueの場合の待ち処理のポーリング間隔(デフォルト:5秒)
}

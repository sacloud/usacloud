// Copyright 2017-2022 The Usacloud Authors
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

package vdef

// FlagOptionsMap CLIで指定するフラグでの静的な候補値一覧(cliタグで指定する)
//
// Note: コード生成で利用されるため実行時に動的に変化する項目には利用できない
var FlagOptionsMap = map[string][]string{
	"os_type_simple": {
		"centos",
		"centos8stream",
		"centos8",
		"ubuntu",
		"ubuntu2004",
		"debian",
		"debian11",
		"rancheros",
		"k3os",
		"...",
	},
}

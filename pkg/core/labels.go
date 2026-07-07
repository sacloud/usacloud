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

package core

// Labels usacloud で扱うリソースを識別するためのラベル情報。
//
// シェル補完や引数を ID or Name or Tags にマッチさせるために利用される。
type Labels struct {
	Id   string
	Name string
	Tags []string
}

// LabelsExtractors サービスから返された値から Labels を抽出する関数のリスト。
//
// 各プラットフォームパッケージ（例: pkg/commands/iaas, pkg/commands/webaccel）は
// init() でプラットフォーム単位の汎用 extractor を登録する。
var LabelsExtractors []func(v interface{}) *Labels

func extractLabels(v interface{}) *Labels {
	for _, extractor := range LabelsExtractors {
		if l := extractor(v); l != nil {
			return l
		}
	}
	return nil
}

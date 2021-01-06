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

package validate

import (
	"strings"

	"github.com/sacloud/usacloud/pkg/util"
)

type Target struct {
	FlagName string
	Value    interface{}
}

func ConflictWith(targets ...*Target) error {
	hasValueCount := 0
	for _, v := range targets {
		if !util.IsEmpty(v.Value) {
			hasValueCount++
		}
	}
	if hasValueCount > 1 { // 値が設定されているのが2つ以上ある場合はエラー
		return NewFlagError(buildFlagName(targets...), "only one of them can be specified")
	}

	return nil
}

func buildFlagName(targets ...*Target) string {
	var names []string
	for _, v := range targets {
		name := v.FlagName
		if !strings.HasPrefix(name, "--") {
			name = "--" + name
		}
		names = append(names, name)
	}
	return strings.Join(names, " & ")
}

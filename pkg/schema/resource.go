// Copyright 2017-2020 The Usacloud Authors
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

package schema

import (
	"sort"
	"strings"
)

type Resource struct {
	Aliases             []string
	Usage               string
	Commands            map[string]*Command
	DefaultCommand      string // 空の場合は`resource -h`
	AltResource         string // 空の場合はResourceのキーをCamelizeしてsacloud.XXXを対象とする。
	ListResultFieldName string
	CommandCategories   []Category
	ResourceCategory    Category
	SkipApplyConfigFile bool
	ExperimentWarning   string
	IsGlobal            bool // グローバルリソースか(API呼び出し時にゾーン指定が必要か)
}

func (r *Resource) CommandCategory(key string) *Category {
	if key == "" {
		return DefaultCommandCategory
	}

	if len(r.CommandCategories) == 0 {
		return &Category{
			Key:         key,
			DisplayName: strings.Title(key),
			Order:       1,
		}
	}

	for _, cat := range r.CommandCategories {
		if cat.Key == key {
			return &cat
		}
	}

	return nil
}

func (r *Resource) SortedCommands() SortableCommands {
	params := SortableCommands{}
	for k, v := range r.Commands {
		params = append(params, SortableCommand{
			CommandKey: k,
			Command:    v,
			Category:   r.CommandCategory(v.Category),
		})
	}

	sort.Sort(params)
	return params
}

type SortableCommand struct {
	Category   *Category
	Command    *Command
	CommandKey string
}

type SortableCommands []SortableCommand

func (s SortableCommands) Len() int {
	return len(s)
}

func (s SortableCommands) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortableCommands) Less(i, j int) bool {
	if s[i].Category.Order == s[j].Category.Order {
		if s[i].Command.Order == s[j].Command.Order {
			return s[i].CommandKey < s[j].CommandKey
		}
		return s[i].Command.Order < s[j].Command.Order
	}
	return s[i].Category.Order < s[j].Category.Order
}

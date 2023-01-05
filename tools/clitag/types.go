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

package clitag

import (
	"fmt"
	"reflect"
	"strings"
)

// Field reflect.StructField + Tag
type StructField struct {
	reflect.StructField
	Tag
}

// Tag structにつけられたftagの値
//
// 例: タグが`example,aliases=foo bar,short=e,desc=foobar`の場合
// - FlagName: example
// - Aliases: foo, bar
// - Shorthand: e
// - Description: foobar
// - Squash: false
// - Ignore: false
// - Category: (empty)
// - Order: 0
type Tag struct {
	FlagName       string
	FieldName      string
	Aliases        []string
	Shorthand      string
	Description    string
	Squash         bool
	Ignore         bool
	Category       string
	Order          int
	Options        []string // 設定可能な値のリスト
	DisplayOptions []string // 表示用
}

// LongDescription DescriptionにAliasesとOptionsを連結して返す
func (t Tag) LongDescription() string {
	var tokens []string

	if t.Description != "" {
		tokens = append(tokens, t.Description)
	}

	options := t.OptionsString()
	if options != "" {
		tokens = append(tokens, options)
	}

	aliases := t.AliasesString()
	if aliases != "" {
		tokens = append(tokens, aliases)
	}

	return strings.Join(tokens, " ")
}

func (t Tag) AliasesString() string {
	if len(t.Aliases) == 0 {
		return ""
	}

	var aliases []string
	for _, a := range t.Aliases {
		aliases = append(aliases, "--"+a)
	}

	return fmt.Sprintf("(aliases: %s)", strings.Join(aliases, ", "))
}

func (t Tag) OptionsString() string {
	if len(t.Options) == 0 {
		return ""
	}

	options := t.Options
	if len(t.DisplayOptions) > 0 {
		options = t.DisplayOptions
	}
	return fmt.Sprintf("options: [%s]", strings.Join(options, "/"))
}

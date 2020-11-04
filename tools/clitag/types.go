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

package clitag

import "reflect"

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
	FlagName    string
	Aliases     []string
	Shorthand   string
	Description string
	Squash      bool
	Ignore      bool
	Category    string
	Order       int
}

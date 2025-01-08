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

package clitag

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	aliasesKey        = "aliases"
	shortHandKey      = "short"
	descKey           = "desc"
	squashKey         = "squash"
	categoryKey       = "category"
	orderKey          = "order"
	optionsKey        = "options"
	displayOptionsKey = "display_options"
)

func (p *Parser) parseTag(t string) (Tag, error) {
	tag := Tag{}
	t = strings.TrimSpace(t)
	if t == "" {
		return tag, nil
	}

	tokens := strings.Split(t, `,`) // 1つ以上の要素を含むスライスを返す
	name := strings.TrimSpace(tokens[0])
	if name == "-" {
		tag.Ignore = true
		return tag, nil
	}

	if name != "" {
		tag.FlagName = name
	}
	if len(tokens) > 1 {
		for _, token := range tokens[1:] {
			token = strings.TrimSpace(token)
			kv := strings.Split(token, `=`)

			key := strings.TrimSpace(kv[0])
			val := ""
			if key != squashKey {
				if len(kv) != 2 {
					return tag, fmt.Errorf("got invalid tag value: %q", token)
				}
				val = strings.TrimSpace(kv[1])
			}

			switch key {
			case squashKey:
				// NOTE: squashに値があっても無視する(ex. squash=foobar)
				tag.Squash = true
			case aliasesKey:
				names := strings.Split(val, ` `)
				for _, n := range names {
					if n != "" {
						tag.Aliases = append(tag.Aliases, n)
					}
				}
			case shortHandKey:
				if len(val) != 1 {
					return tag, fmt.Errorf("got invalid tag value: key 'short' must have only 1 character: %q", token)
				}
				tag.Shorthand = val
			case descKey:
				tag.Description = val
			case categoryKey:
				tag.Category = val
			case orderKey:
				order, err := strconv.ParseInt(val, 10, 64)
				if err != nil {
					return tag, fmt.Errorf("got invalid tag value: key 'order' must have valid number: %q", token)
				}
				tag.Order = int(order)
			case optionsKey:
				options := strings.Split(val, " ")
				for _, o := range options {
					if registered, ok := p.Config.OptionsMap[o]; ok {
						tag.Options = append(tag.Options, registered...)
						continue
					}
					// 登録済みでなければキーをそのまま登録
					tag.Options = append(tag.Options, o)
				}
			case displayOptionsKey:
				options := strings.Split(val, " ")
				for _, o := range options {
					if registered, ok := p.Config.OptionsMap[o]; ok {
						tag.DisplayOptions = append(tag.DisplayOptions, registered...)
						continue
					}
					// 登録済みでなければキーをそのまま登録
					tag.DisplayOptions = append(tag.DisplayOptions, o)
				}
			default:
				return tag, fmt.Errorf("got invalid tag key: %q", token)
			}
		}
	}
	return tag, nil
}

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

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/sacloud/usacloud/pkg/naming"
)

// DefaultTagName ftagのデフォルト名
const DefaultTagName = "cli"

// ParserConfig パーサ設定
type ParserConfig struct {
	TagName    string
	OptionsMap map[string][]string // 指定可能な値(オプション)参照用のマップ
}

// Parser ftagのパーサー
type Parser struct {
	Config *ParserConfig
}

// Parse デフォルトのParser(タグ名:cli)でftagをパースする
func Parse(v interface{}) ([]StructField, error) {
	parser := &Parser{Config: &ParserConfig{TagName: DefaultTagName}}
	return parser.Parse(v)
}

// Parse ftagをパースする
func (p *Parser) Parse(v interface{}) ([]StructField, error) {
	if v == nil {
		return nil, errors.New("value required")
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Ptr:
		return p.Parse(rv.Elem().Interface()) // dereference pointer
	case reflect.Struct:
		return p.parseFields("", reflect.TypeOf(v))
	default:
		return nil, fmt.Errorf("unsupported value: %#v", v)
	}
}

func (p *Parser) parseFields(prefix string, tp reflect.Type) ([]StructField, error) {
	var fields []StructField
	for i := 0; i < tp.NumField(); i++ {
		f := tp.Field(i)
		if f.PkgPath == "" { // exported?
			parsed, err := p.parseField(prefix, f)
			if err != nil {
				return nil, err
			}
			fields = append(fields, parsed...)
		}
	}
	return fields, nil
}

func (p *Parser) parseField(prefix string, f reflect.StructField) ([]StructField, error) {
	tag, err := p.parseTag(f.Tag.Get(p.Config.TagName))
	if err != nil {
		return nil, err
	}
	if tag.Ignore {
		return nil, err
	}

	// handle tag values
	if tag.FlagName == "" {
		tag.FlagName = naming.ToKebabCase(f.Name)
	}
	if !tag.Squash {
		if prefix != "" && tag.FlagName != "" {
			tag.FlagName = fmt.Sprintf("%s-%s", prefix, tag.FlagName)
		}
		prefix = tag.FlagName
	}

	kind := f.Type.Kind()
	switch kind {
	case reflect.Ptr:
		if f.Type.Elem().Kind() == reflect.Struct {
			return p.parseFields(prefix, f.Type.Elem())
		}
	case reflect.Struct:
		return p.parseFields(prefix, f.Type)
	}

	return []StructField{{StructField: f, Tag: tag}}, nil
}

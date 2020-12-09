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

var DefaultParser = &Parser{Config: &ParserConfig{TagName: DefaultTagName}}

// Parse デフォルトのParser(タグ名:cli)でftagをパースする
func Parse(v interface{}) ([]StructField, error) {
	return DefaultParser.Parse(v)
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
		return p.ParseFields("", "", reflect.TypeOf(v))
	default:
		return nil, fmt.Errorf("unsupported value: %#v", v)
	}
}

func (p *Parser) ParseFields(flagPrefix, fieldPrefix string, tp reflect.Type) ([]StructField, error) {
	var fields []StructField
	for i := 0; i < tp.NumField(); i++ {
		f := tp.Field(i)
		if f.PkgPath == "" { // exported?
			parsed, err := p.ParseField(flagPrefix, fieldPrefix, f)
			if err != nil {
				return nil, err
			}
			fields = append(fields, parsed...)
		}
	}
	return fields, nil
}

func (p *Parser) ParseField(flagPrefix, fieldPrefix string, f reflect.StructField) ([]StructField, error) {
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
	if tag.FieldName == "" {
		tag.FieldName = f.Name
	}

	if !tag.Squash {
		if flagPrefix != "" && tag.FlagName != "" {
			tag.FlagName = fmt.Sprintf("%s-%s", flagPrefix, tag.FlagName)
		}
		flagPrefix = tag.FlagName
	}

	if !f.Anonymous {
		if fieldPrefix != "" && tag.FieldName != "" {
			tag.FieldName = fmt.Sprintf("%s.%s", fieldPrefix, tag.FieldName)
		}
		fieldPrefix = tag.FieldName
	}

	kind := f.Type.Kind()
	switch kind {
	case reflect.Ptr:
		if f.Type.Elem().Kind() == reflect.Struct {
			return p.ParseFields(flagPrefix, fieldPrefix, f.Type.Elem())
		}
	case reflect.Struct:
		return p.ParseFields(flagPrefix, fieldPrefix, f.Type)
	}

	return []StructField{{StructField: f, Tag: tag}}, nil
}

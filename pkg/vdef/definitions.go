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

package vdef

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/sacloud/libsacloud/v2/pkg/mapconv"
	"github.com/sacloud/libsacloud/v2/sacloud/ostype"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type definition struct {
	key   interface{}
	value interface{}
}

// definitions usacloudで使う名称(key)/値(value)のペア
var definitions = map[string][]*definition{
	"disk_plan": {
		{key: "ssd", value: types.DiskPlans.SSD},
		{key: "hdd", value: types.DiskPlans.HDD},
	},
	"disk_connection": {
		{key: types.DiskConnections.VirtIO.String(), value: types.DiskConnections.VirtIO.String()},
		{key: types.DiskConnections.IDE.String(), value: types.DiskConnections.IDE.String()},
	},
	"scope": {
		{key: types.Scopes.User.String(), value: types.Scopes.User},
		{key: types.Scopes.Shared.String(), value: types.Scopes.Shared},
	},
	"os_type": ostypeDefinition(),
}

func ostypeDefinition() []*definition {
	var def []*definition
	for _, name := range ostype.OSTypeShortNames {
		def = append(def, &definition{key: name, value: ostype.StrToOSType(name)})
	}
	return def
}

func init() {
	registerFunctions()
}

func registerFunctions() {
	// definitionsから各種定義を登録(Note: 同名のものがあった場合は上書き)
	registerConverterFilters()
	registerTemplateFuncMap()
	registerValidators()
	registerCLITagOptions()
}

func registerConverterFilters() {
	for name, defs := range definitions {
		ConverterFilters[name+"_to_value"] = convertFuncToValue(name, defs)
		ConverterFilters[name+"_to_key"] = convertFuncToKey(name, defs)
	}
}

func registerTemplateFuncMap() {
	for name, defs := range definitions {
		TemplateFuncMap[name+"_to_value"] = templateFuncToValue(defs)
		TemplateFuncMap[name+"_to_key"] = templateFuncToKey(defs)
	}
}

func registerValidators() {
	// definitionsの各値からキーを取り出し、"oneof=keyのスペース区切り"というルールを登録する
	for name, defs := range definitions {
		var allows []string
		for _, def := range defs {
			switch s := def.key.(type) {
			case string:
				allows = append(allows, s)
			case fmt.Stringer:
				allows = append(allows, s.String())
			}
		}
		validatorAliases[name] = fmt.Sprintf("oneof=%s", joinWithSpace(allows))
	}
}

func registerCLITagOptions() {
	// definitionsの各値からキーを取り出し、FlagOptionsMapに登録する
	for name, defs := range definitions {
		var allows []string
		for _, def := range defs {
			switch s := def.key.(type) {
			case string:
				allows = append(allows, s)
			case fmt.Stringer:
				allows = append(allows, s.String())
			}
		}
		FlagOptionsMap[name] = allows
	}
}

func convertFuncToValue(defName string, defs []*definition) mapconv.FilterFunc {
	return func(v interface{}) (interface{}, error) {
		var result interface{}
		for _, def := range defs {
			if reflect.DeepEqual(v, def.key) {
				result = def.value
				break
			}
		}
		if result == nil {
			return nil, fmt.Errorf("key %v not found in %s", v, defName)
		}
		return result, nil
	}
}

func convertFuncToKey(defName string, defs []*definition) mapconv.FilterFunc {
	return func(v interface{}) (interface{}, error) {
		var result interface{}
		for _, def := range defs {
			if reflect.DeepEqual(v, def.value) {
				result = def.key
				break
			}
		}
		if result == nil {
			return nil, fmt.Errorf("value %v not found in %s", v, defName)
		}
		return result, nil
	}
}

func templateFuncToValue(defs []*definition) func(interface{}) interface{} {
	return func(raw interface{}) interface{} {
		in := raw
		if v, ok := raw.(json.Number); ok {
			in = v.String()
		}
		var result interface{}
		for _, def := range defs {
			switch ky := def.key.(type) {
			case fmt.Stringer:
				if reflect.DeepEqual(in, ky.String()) {
					result = def.value
					break
				}
			default:
				if reflect.DeepEqual(in, def.key) {
					result = def.value
					break
				}
			}
		}
		return result
	}
}

func templateFuncToKey(defs []*definition) func(interface{}) interface{} {
	return func(raw interface{}) interface{} {
		in := raw
		if v, ok := raw.(json.Number); ok {
			in = v.String()
		}
		var result interface{}
		for _, def := range defs {
			switch val := def.value.(type) {
			case fmt.Stringer:
				if reflect.DeepEqual(in, val.String()) {
					result = def.key
					break
				}
			default:
				if reflect.DeepEqual(in, def.value) {
					result = def.key
					break
				}
			}
		}
		return result
	}
}

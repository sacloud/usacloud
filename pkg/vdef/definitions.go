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
	"backup_start_minute": {
		{key: 0, value: 0},
		{key: 15, value: 15},
		{key: 30, value: 30},
		{key: 45, value: 45},
	},
	"cdrom_sizes": {
		{key: 5, value: 5},
		{key: 10, value: 10},
	},
	"container_registry_access_levels": {
		{key: types.ContainerRegistryAccessLevels.ReadWrite.String(), value: types.ContainerRegistryAccessLevels.ReadWrite},
		{key: types.ContainerRegistryAccessLevels.ReadOnly.String(), value: types.ContainerRegistryAccessLevels.ReadOnly},
		{key: types.ContainerRegistryAccessLevels.None.String(), value: types.ContainerRegistryAccessLevels.None},
	},
	"database_plan": {
		{key: "10g", value: types.DatabasePlans.DB10GB},
		{key: "30g", value: types.DatabasePlans.DB30GB},
		{key: "90g", value: types.DatabasePlans.DB90GB},
		{key: "240g", value: types.DatabasePlans.DB240GB},
		{key: "500g", value: types.DatabasePlans.DB500GB},
		{key: "1t", value: types.DatabasePlans.DB1TB},
	},
	"database_type": {
		{key: "postgresql", value: "postgresql"},
		{key: "mariadb", value: "mariadb"},
	},
	"disk_plan": {
		{key: "ssd", value: types.DiskPlans.SSD},
		{key: "hdd", value: types.DiskPlans.HDD},
	},
	"disk_connection": {
		{key: types.DiskConnections.VirtIO.String(), value: types.DiskConnections.VirtIO.String()},
		{key: types.DiskConnections.IDE.String(), value: types.DiskConnections.IDE.String()},
	},
	"gslb_protocol": {
		{key: types.GSLBHealthCheckProtocols.HTTP.String(), value: types.GSLBHealthCheckProtocols.HTTP},
		{key: types.GSLBHealthCheckProtocols.HTTPS.String(), value: types.GSLBHealthCheckProtocols.HTTPS},
		{key: types.GSLBHealthCheckProtocols.Ping.String(), value: types.GSLBHealthCheckProtocols.Ping},
		{key: types.GSLBHealthCheckProtocols.TCP.String(), value: types.GSLBHealthCheckProtocols.TCP},
	},
	"internet_network_mask_len": {
		{key: 28, value: 28},
		{key: 27, value: 27},
		{key: 26, value: 26},
		{key: 25, value: 25},
		{key: 24, value: 24},
	},
	"internet_bandwidth": {
		{key: 100, value: 100},
		{key: 250, value: 250},
		{key: 500, value: 500},
		{key: 1000, value: 1000},
		{key: 1500, value: 1500},
		{key: 2000, value: 2000},
		{key: 2500, value: 2500},
		{key: 3000, value: 3000},
		{key: 5000, value: 5000},
	},
	"loadbalancer_plan": {
		{key: "standard", value: types.LoadBalancerPlans.Standard},
		{key: "highspec", value: types.LoadBalancerPlans.HighSpec},
	},
	"scope": {
		{key: types.Scopes.User.String(), value: types.Scopes.User},
		{key: types.Scopes.Shared.String(), value: types.Scopes.Shared},
	},
	"os_type": ostypeDefinition(),
	"nfs_plan": {
		{key: "ssd", value: types.NFSPlans.SSD},
		{key: "hdd", value: types.NFSPlans.HDD},
	},
	"note_class": {
		{key: "shell", value: "shell"},
		{key: "yaml_cloud_config", value: "sheyaml_cloud_configll"},
	},
	"weekdays": {
		{key: "all", value: "all"},
		{key: types.BackupSpanWeekdays.Sunday.String(), value: types.BackupSpanWeekdays.Sunday},
		{key: types.BackupSpanWeekdays.Monday.String(), value: types.BackupSpanWeekdays.Monday},
		{key: types.BackupSpanWeekdays.Tuesday.String(), value: types.BackupSpanWeekdays.Tuesday},
		{key: types.BackupSpanWeekdays.Wednesday.String(), value: types.BackupSpanWeekdays.Wednesday},
		{key: types.BackupSpanWeekdays.Thursday.String(), value: types.BackupSpanWeekdays.Thursday},
		{key: types.BackupSpanWeekdays.Friday.String(), value: types.BackupSpanWeekdays.Friday},
		{key: types.BackupSpanWeekdays.Saturday.String(), value: types.BackupSpanWeekdays.Saturday},
	},
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
		if _, ok := ConverterFilters[name+"_to_value"]; !ok {
			ConverterFilters[name+"_to_value"] = convertFuncToValue(name, defs)
		}
		if _, ok := ConverterFilters[name+"_to_key"]; !ok {
			ConverterFilters[name+"_to_key"] = convertFuncToKey(name, defs)
		}
	}
}

func registerTemplateFuncMap() {
	for name, defs := range definitions {
		if _, ok := TemplateFuncMap[name+"_to_value"]; !ok {
			TemplateFuncMap[name+"_to_value"] = templateFuncToValue(defs)
		}
		if _, ok := TemplateFuncMap[name+"_to_key"]; !ok {
			TemplateFuncMap[name+"_to_key"] = templateFuncToKey(defs)
		}
	}
}

func registerValidators() {
	// definitionsの各値からキーを取り出し、"oneof=keyのスペース区切り"というルールを登録する
	for name, defs := range definitions {
		if _, ok := validatorAliases[name]; ok {
			continue
		}
		var allows []string
		for _, def := range defs {
			switch s := def.key.(type) {
			case string:
				allows = append(allows, s)
			case fmt.Stringer:
				allows = append(allows, s.String())
			default:
				allows = append(allows, fmt.Sprintf("%v", s))
			}
		}
		validatorAliases[name] = fmt.Sprintf("oneof=%s", joinWithSpace(allows))
	}
}

func registerCLITagOptions() {
	// definitionsの各値からキーを取り出し、FlagOptionsMapに登録する
	for name, defs := range definitions {
		if _, ok := FlagOptionsMap[name]; ok {
			continue
		}
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
	var fn func(interface{}) (interface{}, error)
	fn = func(v interface{}) (interface{}, error) {
		// スライスの場合は再帰処理
		vt := reflect.ValueOf(v)
		if vt.Kind() == reflect.Slice || vt.Kind() == reflect.Array {
			var results []interface{}
			for i := 0; i < vt.Len(); i++ {
				res, err := fn(vt.Index(i).Interface())
				if err != nil {
					return nil, err
				}
				results = append(results, res)
			}
			return results, nil
		}

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
	return fn
}

func convertFuncToKey(defName string, defs []*definition) mapconv.FilterFunc {
	var fn func(interface{}) (interface{}, error)
	fn = func(v interface{}) (interface{}, error) {
		// スライスの場合は再帰処理
		vt := reflect.ValueOf(v)
		if vt.Kind() == reflect.Slice || vt.Kind() == reflect.Array {
			var results []interface{}
			for i := 0; i < vt.Len(); i++ {
				res, err := fn(vt.Index(i).Interface())
				if err != nil {
					return nil, err
				}
				results = append(results, res)
			}
			return results, nil
		}

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
	return fn
}

func templateFuncToValue(defs []*definition) func(interface{}) interface{} {
	return func(raw interface{}) interface{} {
		in := raw
		if v, ok := raw.(fmt.Stringer); ok {
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
		if v, ok := raw.(fmt.Stringer); ok {
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

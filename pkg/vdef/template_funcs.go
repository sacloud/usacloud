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
	"strings"
	"text/template"

	"github.com/sacloud/libsacloud/v2/sacloud"

	"github.com/sacloud/libsacloud/v2/sacloud/types"

	"github.com/sacloud/libsacloud/v2/pkg/size"
	"github.com/sacloud/usacloud/pkg/util"
)

var TemplateFuncMap = template.FuncMap{
	// definitionsから実行時に構築される、独自に追加してもOK
	"gib_to_mib": func(value interface{}) int {
		num := 0
		switch v := value.(type) {
		case int:
			num = v
		case int64:
			num = int(v)
		case json.Number:
			n, _ := v.Int64() // この段階でのエラーは握り潰す
			num = int(n)
		}
		return size.GiBToMiB(num)
	},
	"mib_to_gib": func(value interface{}) int {
		num := 0
		switch v := value.(type) {
		case int:
			num = v
		case int64:
			num = int(v)
		case json.Number:
			n, _ := v.Int64() // この段階でのエラーは握り潰す
			num = int(n)
		}
		return size.MiBToGiB(num)
	},
	"ignore_empty": func(value interface{}) interface{} {
		if util.IsEmpty(value) {
			return ""
		}
		return value
	},
	"ellipsis": ellipsis,
	"to_single_line": func(value interface{}) interface{} {
		v := fmt.Sprintf("%v", value)
		return strings.ReplaceAll(v, "\n", "\\n")
	},
	"weekdays": func(value interface{}) interface{} {
		if value == nil {
			return nil
		}
		weekdays, ok := value.([]types.EBackupSpanWeekday)
		if !ok {
			return nil
		}
		if len(weekdays) == 7 {
			return "all"
		}

		var results []string
		for _, d := range weekdays {
			results = append(results, d.String())
		}
		return results
	},
	"switch_type": func(value interface{}) string {
		if value == nil {
			return "unknown"
		}
		v, ok := value.(*sacloud.Switch)
		if !ok {
			return "unknown"
		}
		if len(v.Subnets) > 0 && v.Subnets[0].Internet != nil {
			return "switch+router"
		}
		return "switch"
	},
}

func ellipsis(length int, value interface{}) interface{} {
	runes := []rune(fmt.Sprintf("%v", value))
	if len(runes) > length {
		return string(runes[0:length]) + "..."
	}
	return string(runes)
}

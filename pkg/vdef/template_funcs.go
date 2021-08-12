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

package vdef

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/sacloud/libsacloud/v2/pkg/size"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
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
	"first_line": func(value interface{}) interface{} {
		if value == nil {
			return nil
		}
		if v, ok := value.([]string); ok {
			if len(v) > 0 {
				return v[0]
			}
		}
		return ""
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
	"user_friendly_price": userFriendlyPriceString,
	"join": func(sep string, values []string) string {
		return strings.Join(values, sep)
	},
	"first_non_empty": func(values ...interface{}) interface{} {
		for _, v := range values {
			if !util.IsEmpty(v) {
				return v
			}
		}
		return nil
	},
	"unix_time_to_date": func(value int64) *time.Time {
		if value <= 0 {
			return nil
		}
		v := time.Unix(value/1000, 0)
		return &v
	},
	"file": func(path string) string {
		if path == "" {
			return ""
		}

		poc, err := homedir.Expand(path)
		if err != nil {
			return ""
		}

		data, err := os.ReadFile(poc)
		if err != nil {
			return ""
		}
		return string(data)
	},
	"trim_space": func(s string) string {
		return strings.TrimSpace(s)
	},
}

func ellipsis(length int, value interface{}) interface{} {
	runes := []rune(fmt.Sprintf("%v", value))
	if len(runes) > length {
		return string(runes[0:length]) + "..."
	}
	return string(runes)
}

func userFriendlyPriceString(value interface{}) string {
	if value == nil {
		return ""
	}
	v, ok := value.(*sacloud.Price)
	if !ok {
		return ""
	}

	var results []string

	if v.Base > 0 {
		results = append(results, fmt.Sprintf("Base:%d", v.Base))
	}
	if v.Daily > 0 {
		results = append(results, fmt.Sprintf("Daily:%d", v.Daily))
	}
	if v.Hourly > 0 {
		results = append(results, fmt.Sprintf("Hourly:%d", v.Hourly))
	}
	if v.Monthly > 0 {
		results = append(results, fmt.Sprintf("Monthly:%d", v.Monthly))
	}
	if v.PerUse > 0 {
		results = append(results, fmt.Sprintf("PerUse:%d", v.PerUse))
	}
	if v.Basic > 0 {
		results = append(results, fmt.Sprintf("Basic:%d", v.Basic))
	}
	if v.Traffic > 0 {
		results = append(results, fmt.Sprintf("Traffic:%d", v.Traffic))
	}
	if v.DocomoTraffic > 0 {
		results = append(results, fmt.Sprintf("DocomoTraffic:%d", v.DocomoTraffic))
	}
	if v.KddiTraffic > 0 {
		results = append(results, fmt.Sprintf("KddiTraffic:%d", v.KddiTraffic))
	}
	if v.SbTraffic > 0 {
		results = append(results, fmt.Sprintf("SbTraffic:%d", v.SbTraffic))
	}
	if v.SimSheet > 0 {
		results = append(results, fmt.Sprintf("SimSheet:%d", v.SimSheet))
	}
	if v.Zone != "" {
		results = append(results, fmt.Sprintf("Zone:%s", v.Zone))
	}

	return strings.Join(results, " / ")
}

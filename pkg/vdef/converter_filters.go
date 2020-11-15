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
	"time"

	"github.com/sacloud/libsacloud/v2/pkg/mapconv"
	"github.com/sacloud/libsacloud/v2/sacloud/ostype"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// ConverterFilters mapconvでの変換時に利用されるフィルターの定義
var ConverterFilters = map[string]mapconv.FilterFunc{
	"disk_plan_to_id": diskPlanToID,
	"os_type":         strToOSType,
	"rfc3339":         strToTime,
}

func diskPlanToID(v interface{}) (interface{}, error) {
	s, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("invalid disk plan: %v", v)
	}
	for diskPlanName, id := range types.DiskPlanIDMap {
		if diskPlanName == s {
			return id, nil
		}
	}
	return nil, fmt.Errorf("disk plan %s not found", s)
}

func strToOSType(v interface{}) (interface{}, error) {
	s, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("invalid os type: %v", v)
	}
	if s == "" {
		return ostype.Custom, nil
	}

	ot := ostype.StrToOSType(s)
	if ot == ostype.Custom {
		return nil, fmt.Errorf("os type %s not found", s)
	}
	return ot, nil
}

func strToTime(v interface{}) (interface{}, error) {
	s, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("invalid time format: %v", v)
	}
	if s == "" {
		return time.Time{}, nil
	}

	allowDatetimeFormatList := []string{
		time.RFC3339,
	}
	for _, format := range allowDatetimeFormatList {
		d, err := time.Parse(format, s)
		if err == nil {
			// success
			return d, nil
		}
	}
	return nil, fmt.Errorf("invalid time format: %v", v)
}

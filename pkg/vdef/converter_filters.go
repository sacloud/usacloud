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
)

// ConverterFilters mapconvでの変換時に利用されるフィルターの定義、definitionsに登録したものは実行時に動的に追加される
var ConverterFilters = map[string]mapconv.FilterFunc{
	"rfc3339": strToTime,
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

// Copyright 2017-2023 The sacloud/usacloud Authors
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
	"encoding/base64"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/sacloud/iaas-api-go/mapconv"
	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/packages-go/pointer"
	"github.com/sacloud/packages-go/size"
	"github.com/sacloud/usacloud/pkg/util"
)

// ConverterFilters mapconvでの変換時に利用されるフィルターの定義、definitionsに登録したものは実行時に動的に追加される
var ConverterFilters = map[string]mapconv.FilterFunc{
	"rfc3339":         strToTime,
	"path_to_reader":  pathToReader,
	"path_to_writer":  pathToWriter,
	"path_or_content": pathOrContent,
	"weekdays":        weekdaysFilter,
	"dereference":     dereferenceFilter,
	"base64encode":    base64Encode,
	"gib_to_mib": func(v interface{}) (interface{}, error) {
		if v == nil {
			return nil, nil
		}
		sv, ok := v.(int)
		if !ok {
			return nil, fmt.Errorf("invalid value: %v", v)
		}
		return size.GiBToMiB(sv), nil
	},
}

func strToTime(v interface{}) (interface{}, error) {
	if v == nil {
		return nil, nil
	}
	s, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("invalid time value: %v", v)
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

// pathToReader ファイルパスから*os.File(O_RDONLY)を返す
//
// Note: ファイルはここではクローズされないため、このフィルタを適用する先のリクエストでCloseを適切に呼ぶようにする
// libsacloud serviceの場合はservice内でcloseが呼ばれる
func pathToReader(v interface{}) (interface{}, error) {
	if v == nil {
		return nil, nil
	}
	s, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("invalid filepath value: %v", v)
	}
	if s == "" {
		return nil, nil
	}

	file, err := os.Open(s)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// pathToWriter ファイルパスから*os.File(O_RDWR|O_CREATE|O_TRUNC、パーミッション:0666)を返す
//
// Note: os.Create(path)を使用するため、バリデーションで上書き確認を行うこと
// Note: ファイルはここではクローズされないため、このフィルタを適用する先のリクエストでCloseを適切に呼ぶようにする
// libsacloud serviceの場合はservice内でcloseが呼ばれる
func pathToWriter(v interface{}) (interface{}, error) {
	if v == nil {
		return nil, nil
	}
	s, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("invalid filepath value: %v", v)
	}
	if s == "" {
		return nil, nil
	}

	file, err := os.Create(s)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// pathOrContent 値がファイルだった場合はファイルの内容を、そうでない場合は値をそのまま返す
func pathOrContent(v interface{}) (interface{}, error) {
	if v == nil {
		return nil, nil
	}
	s, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("invalid filepath value: %v", v)
	}
	if s == "" {
		return nil, nil
	}

	return util.StringFromPathOrContent(s)
}

func weekdaysFilter(v interface{}) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	var days *[]string
	switch v := v.(type) {
	case []string:
		days = pointer.NewStringSlice(v)
	case *[]string:
		days = v
	default:
		return nil, fmt.Errorf("invalid weekdays value: %v", v)
	}

	var results []types.EDayOfTheWeek
	for _, d := range *days {
		// 途中に"all"が見つかった場合は全曜日とする
		if d == "all" {
			return []types.EDayOfTheWeek{
				types.DaysOfTheWeek.Sunday,
				types.DaysOfTheWeek.Monday,
				types.DaysOfTheWeek.Tuesday,
				types.DaysOfTheWeek.Wednesday,
				types.DaysOfTheWeek.Thursday,
				types.DaysOfTheWeek.Friday,
				types.DaysOfTheWeek.Saturday,
			}, nil
		}
		results = append(results, types.EDayOfTheWeek(d))
	}
	return results, nil
}

func dereferenceFilter(v interface{}) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	vt := reflect.ValueOf(v)
	if vt.Kind() != reflect.Ptr {
		return v, nil
	}
	return vt.Elem().Interface(), nil
}

func base64Encode(v interface{}) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	var data []byte
	switch value := v.(type) {
	case string:
		data = []byte(value)
	case []byte:
		data = value
	default:
		return nil, fmt.Errorf("invalid based64 target: %v", v)
	}

	return base64.StdEncoding.EncodeToString(data), nil
}

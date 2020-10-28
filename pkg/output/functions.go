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

package output

import (
	"reflect"

	"github.com/sacloud/usacloud/pkg/util"
)

func toSlice(v interface{}) interface{} {
	if util.IsEmpty(v) {
		return nil
	}
	objValue := reflect.ValueOf(v)

	switch objValue.Kind() {
	case reflect.Array, reflect.Slice:
		return v
	default:
		return []interface{}{v}
	}
}

func sliceLen(v interface{}) int {
	if util.IsEmpty(v) {
		return 0
	}
	objValue := reflect.ValueOf(v)

	switch objValue.Kind() {
	case reflect.Array, reflect.Slice:
		return objValue.Len()
	default:
		return 1
	}
}

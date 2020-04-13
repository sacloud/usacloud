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

package validation

import "fmt"

func SliceLenBetween(fieldName string, object interface{}, min int, max int) []error {

	if object == nil {
		object = []int64{}
	}

	isSlice := func(object interface{}) bool {
		_, ok1 := object.([]int64)
		_, ok2 := object.([]string)

		return ok1 || ok2
	}

	if isSlice(object) {
		sliceLen := 0
		if s, ok := object.([]int64); ok {
			sliceLen = len(s)
		} else {
			s := object.([]string)
			sliceLen = len(s)
		}

		if max <= 0 {
			if sliceLen < min {
				return []error{fmt.Errorf("%q: slice length must be %d or more", fieldName, min)}
			}
		} else {
			if !(min <= sliceLen && sliceLen <= max) {
				return []error{fmt.Errorf("%q: slice length must be beetween %d and %d", fieldName, min, max)}
			}

		}
	}

	return []error{}
}

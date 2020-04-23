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

import (
	"fmt"

	"github.com/sacloud/usacloud/pkg/util"
)

func IntBetween(fieldName string, object interface{}, min int, max int) []error {
	// if target is empty, return OK(Use required attr if necessary)
	if util.IsEmpty(object) {
		return []error{}
	}

	v, ok := object.(int)
	if !ok {
		return []error{fmt.Errorf("%q: must be int", fieldName)}
	}

	if !(min <= v && v <= max) {
		return []error{fmt.Errorf("%q: must be between %d and %d", fieldName, min, max)}
	}

	return []error{}
}

func Int64Between(fieldName string, object interface{}, min int64, max int64) []error {
	// if target is empty, return OK(Use required attr if necessary)
	if util.IsEmpty(object) {
		return []error{}
	}

	v, ok := object.(int64)
	if !ok {
		return []error{fmt.Errorf("%q: must be int64", fieldName)}
	}

	if !(min <= v && v <= max) {
		return []error{fmt.Errorf("%q: must be between %d and %d", fieldName, min, max)}
	}

	return []error{}
}

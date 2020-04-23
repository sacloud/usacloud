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

	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/util"
)

func ValidID(fieldName string, object interface{}) []error {
	var res []error

	// if target is empty, return OK(Use required attr if necessary)
	if util.IsEmpty(object) {
		return res
	}

	var id types.ID
	switch v := object.(type) {
	case int64:
		id = types.ID(v)
	case string:
		id = types.StringID(v)
	default:
		res = append(res, fmt.Errorf("%q: Resource ID must be valid format", fieldName))
		return res
	}

	if id.IsEmpty() {
		res = append(res, fmt.Errorf("%q: Resource ID must be valid format", fieldName))
	}

	return res
}

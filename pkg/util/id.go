// Copyright 2017-2025 The sacloud/usacloud Authors
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

package util

import (
	"github.com/sacloud/iaas-api-go/types"
)

func UniqIDs(elements []types.ID) []types.ID {
	encountered := map[types.ID]bool{}
	result := []types.ID{}
	for v := range elements {
		if !encountered[elements[v]] {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}

// StringIDs iaas.IDスライスを文字列のスライスに変換する
func StringIDs(ids []types.ID) []string {
	var strIDs []string

	for _, v := range ids {
		if v != 0 {
			strIDs = append(strIDs, v.String())
		}
	}

	return strIDs
}

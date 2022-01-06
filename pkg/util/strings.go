// Copyright 2017-2022 The Usacloud Authors
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

func FirstNonEmptyString(values ...string) string {
	if len(values) == 0 {
		return ""
	}
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return values[len(values)-1]
}

func UniqStrings(elements []string) []string {
	encountered := map[string]bool{}
	var result []string
	for v := range elements {
		if !encountered[elements[v]] {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}

func RemoveStringsFromSlice(elements []string, remove []string) []string {
	var results []string
	for _, e := range elements {
		exists := false
		for _, r := range remove {
			if e == r {
				exists = true
				break
			}
		}
		if !exists {
			results = append(results, e)
		}
	}
	return results
}

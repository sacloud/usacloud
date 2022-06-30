// Copyright 2017-2022 The sacloud/usacloud Authors
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

import "testing"

func TestConverterFilters_fromDefinitions(t *testing.T) {
	// definitionsから動的にフィルタ登録されているか?
	expectFilters := []string{
		"disk_plan_to_value", "disk_plan_to_key",
	}

	for _, name := range expectFilters {
		if _, ok := ConverterFilters[name]; !ok {
			t.Fatalf("ConverterFilters[%s] not exists", name)
		}
	}
}

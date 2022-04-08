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

package core

import "sort"

type Resources []*Resource

func (r Resources) CategorizedResources() []*CategorizedResources {
	categories := ResourceCategories
	sort.Slice(categories, func(i, j int) bool {
		return categories[i].Order < categories[j].Order
	})

	var results []*CategorizedResources
	for _, c := range categories {
		result := &CategorizedResources{
			Category:  c,
			Resources: []*Resource{},
		}
		for _, resource := range r {
			if c.Equals(&resource.Category) {
				result.Resources = append(result.Resources, resource)
			}
		}

		if len(result.Resources) > 0 {
			results = append(results, result)
		}
	}
	return results
}

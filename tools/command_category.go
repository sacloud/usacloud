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

package tools

import (
	"sort"

	"github.com/sacloud/usacloud/pkg/category"
)

type CategorizedParameterFields struct {
	*category.Category
	Fields []Field
}

func (c *Command) CategorizedParameterFields() []*CategorizedParameterFields {
	if c.ParameterInitializer == nil {
		return nil
	}

	m := map[string]*CategorizedParameterFields{}
	for _, f := range c.Fields() {
		cp, ok := m[f.Category]
		if !ok {
			cp = &CategorizedParameterFields{
				Category: c.Command.ParameterCategoryBy(f.Category),
			}
		}
		cp.Fields = append(cp.Fields, f)
		m[f.Category] = cp
	}
	var categorizedFields []*CategorizedParameterFields
	for _, cat := range m {
		sort.Slice(cat.Fields, func(i, j int) bool {
			if cat.Fields[i].Order == cat.Fields[j].Order {
				return cat.Fields[i].FlagName < cat.Fields[j].FlagName
			}
			return cat.Fields[i].Order < cat.Fields[j].Order
		})
		categorizedFields = append(categorizedFields, cat)
	}
	sort.Slice(categorizedFields, func(i, j int) bool {
		if categorizedFields[i].Order == categorizedFields[j].Order {
			return categorizedFields[i].Key < categorizedFields[j].Key
		}
		return categorizedFields[i].Order < categorizedFields[j].Order
	})

	return categorizedFields
}

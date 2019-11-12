// Copyright 2017-2019 The Usacloud Authors
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

package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func ZoneResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find"},
			Params:             zoneListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: zoneListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"read": {
			Type:          schema.CommandManipulateIDOnly,
			Params:        zoneReadParam(),
			IncludeFields: zoneDetailIncludes(),
			ExcludeFields: zoneDetailExcludes(),
			Category:      "basics",
			Order:         20,
		},
	}

	return &schema.Resource{
		Commands:         commands,
		DefaultCommand:   "list",
		ResourceCategory: CategoryInformation,
	}
}

func zoneListParam() map[string]*schema.Schema {
	return CommonListParam
}

func zoneListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{Name: "Description"},
		{
			Name:    "Region",
			Sources: []string{"Region.Name", "Region.ID"},
			Format:  "%s(%s)",
		},
	}
}

func zoneDetailIncludes() []string {
	return []string{}
}

func zoneDetailExcludes() []string {
	return []string{}
}

func zoneReadParam() map[string]*schema.Schema {
	id := getParamResourceShortID("resource ID", 5)
	id.Hidden = true
	return map[string]*schema.Schema{
		"id": id,
	}
}

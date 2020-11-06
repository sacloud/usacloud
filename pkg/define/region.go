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

package define

import (
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/schema"
)

func RegionResource() *schema.Resource {
	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find"},
			Params:             regionListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: regionListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"read": {
			Type:          schema.CommandManipulateIDOnly,
			Params:        regionReadParam(),
			IncludeFields: regionDetailIncludes(),
			ExcludeFields: regionDetailExcludes(),
			Category:      "basics",
			Order:         20,
		},
	}

	return &schema.Resource{
		Commands:         commands,
		DefaultCommand:   "list",
		ResourceCategory: CategoryInformation,
		IsGlobal:         true,
	}
}

func regionListParam() map[string]*schema.Parameter {
	return CommonListParam
}

func regionListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{Name: "Description"},
		{
			Name:    "NameServers",
			Sources: []string{"NameServers.0", "NameServers.1"},
			Format:  "%s,%s",
		},
	}
}

func regionDetailIncludes() []string {
	return []string{}
}

func regionDetailExcludes() []string {
	return []string{}
}

func regionReadParam() map[string]*schema.Parameter {
	id := getParamResourceShortID("resource ID", 3)
	id.Hidden = true
	return map[string]*schema.Parameter{
		"id": id,
	}
}

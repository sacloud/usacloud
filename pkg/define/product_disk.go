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

func ProductDiskResource() *schema.Resource {
	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find"},
			Params:             productDiskListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: productDiskListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"read": {
			Type:          schema.CommandManipulateIDOnly,
			Params:        productDiskReadParam(),
			IncludeFields: productDiskDetailIncludes(),
			ExcludeFields: productDiskDetailExcludes(),
			Category:      "basics",
			Order:         20,
		},
	}

	return &schema.Resource{
		Commands:            commands,
		DefaultCommand:      "list",
		Aliases:             []string{"disk-plan"},
		ResourceCategory:    CategoryInformation,
		ListResultFieldName: "DiskPlans",
		AltResource:         "DiskPlan",
	}
}

func productDiskListParam() map[string]*schema.Parameter {
	return CommonListParam
}

func productDiskListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
	}
}

func productDiskDetailIncludes() []string {
	return []string{}
}

func productDiskDetailExcludes() []string {
	return []string{}
}

func productDiskReadParam() map[string]*schema.Parameter {
	id := getParamResourceShortID("resource ID", 1)
	id.Hidden = true
	return map[string]*schema.Parameter{
		"id": id,
	}
}

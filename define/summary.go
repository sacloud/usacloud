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
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func SummaryResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"show": {
			Type:               schema.CommandList,
			Params:             showSummaryParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: summaryShowColumns(),
			UseCustomCommand:   true,
		},
	}

	return &schema.Resource{
		Commands:         commands,
		DefaultCommand:   "show",
		ResourceCategory: CategorySummary,
		Usage:            "Show summary of resource usage",
	}
}

func showSummaryParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"paid-resources-only": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"paid"},
			Description: "Show paid-resource only",
		},
	}
}

func summaryShowColumns() []output.ColumnDef {
	defs := []output.ColumnDef{
		{Name: "Name"},
	}

	for _, zone := range AllowZones {
		defs = append(defs, output.ColumnDef{Name: zone})
	}

	defs = append(defs, output.ColumnDef{Name: "Total"})

	return defs
}

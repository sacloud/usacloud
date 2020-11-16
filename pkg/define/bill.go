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
	"fmt"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/schema"
)

func BillResource() *schema.Resource {
	commands := map[string]*schema.Command{
		"list": {
			Type:    schema.CommandCustom,
			Aliases: []string{"ls", "find"},
			Params:  billListParam(),
			// TableType:          output.TableSimple,
			TableColumnDefines: billListColumns(),
			UseCustomCommand:   true,
			NeedlessConfirm:    true,
		},
		"csv": {
			Type:             schema.CommandCustom,
			Params:           billReadParam(),
			UseCustomCommand: true,
			NoOutput:         true, // doing manual output to GlobalOption.Out
			NeedlessConfirm:  true,
		},
	}

	return &schema.Resource{
		Commands:         commands,
		DefaultCommand:   "list",
		ResourceCategory: CategoryBilling,
	}
}

func billListParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"year": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set year",
			ValidateFunc: validateIntRange(2000, 9999),
			Category:     "filter",
			Order:        10,
		},
		"month": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set month",
			ValidateFunc: validateIntRange(1, 12),
			Category:     "filter",
			Order:        20,
		},
	}
}

func billListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "BillID"},
		{
			Name: "Date",
			FormatFunc: func(values map[string]string) string {
				if strDate, ok := values["Date"]; ok {
					t, err := time.Parse(types.DatetimeLayout, strDate)
					if err != nil {
						return ""
					}
					return fmt.Sprintf("%d/%02d", t.Year(), t.Month())
				}
				return ""
			},
		},
		{Name: "Paid"},
		{Name: "Amount"},
	}
}

func billReadParam() map[string]*schema.Parameter {
	id := getParamResourceShortID("bill ID", 8)
	id.Required = false
	return map[string]*schema.Parameter{
		"bill-id": id,
		"no-header": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set output header flag",
			Category:    "output",
			Order:       10,
		},
		"bill-output": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"file"},
			Description: "set bill-detail output path",
			Category:    "output",
			Order:       20,
		},
	}
}

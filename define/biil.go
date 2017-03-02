package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func BillResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandCustom,
			Aliases:            []string{"l", "ls", "find"},
			Params:             billListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: billListColumns(),
			UseCustomCommand:   true,
		},
		"csv": {
			Type:                   schema.CommandRead,
			Params:                 billReadParam(),
			UseCustomCommand:       true,
			UseCustomArgCompletion: true,
		},
	}

	return &schema.Resource{
		Commands:         commands,
		ResourceCategory: CategoryBilling,
	}
}

func billListParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"year": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set year",
			ValidateFunc: validateIntRange(2000, 9999),
		},
		"month": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set month",
			ValidateFunc: validateIntRange(1, 12),
		},
	}
}

func billListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "BillID"},
		{Name: "Date"},
		{Name: "PayLimit"},
	}
}

func billReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": getParamResourceShortID("bill ID", 8),
		"no-header": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set output header flag",
		},
		"bill-output": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"file"},
			Description: "set bill-detail output path",
		},
	}
}

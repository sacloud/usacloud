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
		ResourceCategory: CategoryInformation,
		Usage:            "Show summary of resource usage",
	}
}

func showSummaryParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"paid-resources-only": {
			Type:         schema.TypeBool,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"paid"},
			Description:  "Show paid-resource only",
			DefaultValue: false,
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

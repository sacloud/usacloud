package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func SummaryResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"show": {
			Type:               schema.CommandList,
			Params:             summaryShowParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: summaryShowColumns(),
			UseCustomCommand:   true,
		},
	}

	return &schema.Resource{
		Commands:            commands,
		DefaultCommand:      "show",
		ListResultFieldName: "ServiceClasses",
		ResourceCategory:    CategoryInformation,
		Usage:               "Show summary of resource usage",
	}
}

func summaryShowParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func summaryShowColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ResourceName"},
		{Name: "tk1a"},
		{Name: "is1a"},
		{Name: "is1b"},
		{Name: "Total"},
	}
}

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
		},
		"read": {
			Type:          schema.CommandManipulateIDOnly,
			Params:        zoneReadParam(),
			IncludeFields: zoneDetailIncludes(),
			ExcludeFields: zoneDetailExcludes(),
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
	return map[string]*schema.Schema{
		"id": getParamResourceShortID("resource ID", 5),
	}
}

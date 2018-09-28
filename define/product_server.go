package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func ProductServerResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find"},
			Params:             productServerListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: productServerListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"read": {
			Type:          schema.CommandManipulateIDOnly,
			Params:        productServerReadParam(),
			IncludeFields: productServerDetailIncludes(),
			ExcludeFields: productServerDetailExcludes(),
			Category:      "basics",
			Order:         10,
		},
	}

	return &schema.Resource{
		Commands:            commands,
		DefaultCommand:      "list",
		Aliases:             []string{"server-plan"},
		ResourceCategory:    CategoryInformation,
		ListResultFieldName: "ServerPlans",
	}
}

func productServerListParam() map[string]*schema.Schema {
	return CommonListParam
}

func productServerListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{Name: "CPU"},
		{
			Name:    "Memory",
			Sources: []string{"MemoryMB"},
			Format:  "%sMB",
		},
		{Name: "Generation"},
	}
}

func productServerDetailIncludes() []string {
	return []string{}
}

func productServerDetailExcludes() []string {
	return []string{}
}

func productServerReadParam() map[string]*schema.Schema {
	id := getParamResourceShortID("resource ID", 9)
	id.Hidden = true
	return map[string]*schema.Schema{
		"id": id,
	}
}

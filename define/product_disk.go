package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func ProductDiskResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"l", "ls", "find"},
			Params:             productDiskListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: productDiskListColumns(),
		},
		"read": {
			Type:          schema.CommandManipulateIDOnly,
			Aliases:       []string{"r"},
			Params:        productDiskReadParam(),
			IncludeFields: productDiskDetailIncludes(),
			ExcludeFields: productDiskDetailExcludes(),
		},
	}

	return &schema.Resource{
		Commands:            commands,
		DefaultCommand:      "list",
		Aliases:             []string{"disk-plan"},
		ResourceCategory:    CategoryInformation,
		ListResultFieldName: "DiskPlans",
	}
}

func productDiskListParam() map[string]*schema.Schema {
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

func productDiskReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": getParamResourceShortID("resource ID", 1),
	}
}

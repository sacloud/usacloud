package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func ProductInternetResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find"},
			Params:             productInternetListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: productInternetListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"read": {
			Type:          schema.CommandManipulateIDOnly,
			Params:        productInternetReadParam(),
			IncludeFields: productInternetDetailIncludes(),
			ExcludeFields: productInternetDetailExcludes(),
			Category:      "basics",
			Order:         20,
		},
	}

	return &schema.Resource{
		Commands:            commands,
		DefaultCommand:      "list",
		Aliases:             []string{"internet-plan"},
		ResourceCategory:    CategoryInformation,
		ListResultFieldName: "InternetPlans",
	}
}

func productInternetListParam() map[string]*schema.Schema {
	return CommonListParam
}

func productInternetListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{
			Name:    "BandWidth",
			Sources: []string{"BandWidthMbps"},
			Format:  "%sMbps",
		},
	}
}

func productInternetDetailIncludes() []string {
	return []string{}
}

func productInternetDetailExcludes() []string {
	return []string{}
}

func productInternetReadParam() map[string]*schema.Schema {
	id := getParamResourceShortID("resource ID", 4)
	id.Hidden = true
	return map[string]*schema.Schema{
		"id": id,
	}
}

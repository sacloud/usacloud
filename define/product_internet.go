package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func ProductInternetResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"l", "ls", "find"},
			Params:             productInternetListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: productInternetListColumns(),
		},
		"read": {
			Type:          schema.CommandRead,
			Aliases:       []string{"r"},
			Params:        productInternetReadParam(),
			IncludeFields: productInternetDetailIncludes(),
			ExcludeFields: productInternetDetailExcludes(),
		},
	}

	return &schema.Resource{
		Commands:            commands,
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
	return map[string]*schema.Schema{
		"id": getParamResourceShortID("resource ID", 4),
	}
}

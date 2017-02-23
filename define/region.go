package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func RegionResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:                schema.CommandList,
			ListResultFieldName: "Regions",
			Aliases:             []string{"l", "ls", "find"},
			Params:              regionListParam(),
			TableType:           output.TableSimple,
			TableColumnDefines:  regionListColumns(),
		},
		"read": {
			Type:          schema.CommandRead,
			Aliases:       []string{"r"},
			Params:        regionReadParam(),
			IncludeFields: regionDetailIncludes(),
			ExcludeFields: regionDetailExcludes(),
		},
	}

	return &schema.Resource{
		Commands: commands,
	}
}

func regionListParam() map[string]*schema.Schema {
	return CommonListParam
}

func regionListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{Name: "Description"},
		{
			Name:    "NameServers",
			Sources: []string{"NameServers.0", "NameServers.1"},
			Format:  "%s,%s",
		},
	}
}

func regionDetailIncludes() []string {
	return []string{}
}

func regionDetailExcludes() []string {
	return []string{}
}

func regionReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": getParamResourceShortID("resource ID", 3),
	}
}

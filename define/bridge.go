package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func BridgeResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:                schema.CommandList,
			ListResultFieldName: "Bridges",
			Aliases:             []string{"l", "ls", "find"},
			Params:              bridgeListParam(),
			TableType:           output.TableSimple,
			TableColumnDefines:  bridgeListColumns(),
		},
		"create": {
			Type:          schema.CommandCreate,
			Aliases:       []string{"c"},
			Params:        bridgeCreateParam(),
			IncludeFields: bridgeDetailIncludes(),
			ExcludeFields: bridgeDetailExcludes(),
		},
		"read": {
			Type:          schema.CommandRead,
			Aliases:       []string{"r"},
			Params:        bridgeReadParam(),
			IncludeFields: bridgeDetailIncludes(),
			ExcludeFields: bridgeDetailExcludes(),
		},
		"update": {
			Type:          schema.CommandUpdate,
			Aliases:       []string{"u"},
			Params:        bridgeUpdateParam(),
			IncludeFields: bridgeDetailIncludes(),
			ExcludeFields: bridgeDetailExcludes(),
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"d", "rm"},
			Params:        bridgeDeleteParam(),
			IncludeFields: bridgeDetailIncludes(),
			ExcludeFields: bridgeDetailExcludes(),
		},
	}

	return &schema.Resource{
		Commands: commands,
	}
}

func bridgeListParam() map[string]*schema.Schema {
	return CommonListParam
}

func bridgeListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{
			Name:    "Switch(this zone only)",
			Sources: []string{"SwitchInZone.ID", "SwitchInZone.Name"},
			Format:  "%s:%s",
		},
	}
}

func bridgeDetailIncludes() []string {
	return []string{}
}

func bridgeDetailExcludes() []string {
	return []string{
		"Region.Description",
		"Region.NameServers",
	}
}

func bridgeCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramName,
		"description": paramDescription,
	}
}

func bridgeReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func bridgeUpdateParam() map[string]*schema.Schema {
	updateParam := map[string]*schema.Schema{
		"id": paramID,
	}
	return mergeParameterMap(bridgeCreateParam(), updateParam)
}

func bridgeDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

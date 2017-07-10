package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func BridgeResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find"},
			Params:             bridgeListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: bridgeListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"create": {
			Type:          schema.CommandCreate,
			Params:        bridgeCreateParam(),
			IncludeFields: bridgeDetailIncludes(),
			ExcludeFields: bridgeDetailExcludes(),
			Category:      "basics",
			Order:         20,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        bridgeReadParam(),
			IncludeFields: bridgeDetailIncludes(),
			ExcludeFields: bridgeDetailExcludes(),
			NoSelector:    true,
			Category:      "basics",
			Order:         30,
		},
		"update": {
			Type:          schema.CommandUpdate,
			Params:        bridgeUpdateParam(),
			IncludeFields: bridgeDetailIncludes(),
			ExcludeFields: bridgeDetailExcludes(),
			NoSelector:    true,
			Category:      "basics",
			Order:         40,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"rm"},
			Params:        bridgeDeleteParam(),
			IncludeFields: bridgeDetailIncludes(),
			ExcludeFields: bridgeDetailExcludes(),
			NoSelector:    true,
			Category:      "basics",
			Order:         50,
		},
	}

	return &schema.Resource{
		Commands:         commands,
		ResourceCategory: CategoryNetworking,
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
			Format:  "%s(%s)",
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
		"name":        paramRequiredName,
		"description": paramDescription,
	}
}

func bridgeReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func bridgeUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramName,
		"description": paramDescription,
	}
}

func bridgeDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

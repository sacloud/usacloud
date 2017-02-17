package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func SwitchResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:                schema.CommandList,
			ListResultFieldName: "Switches",
			Aliases:             []string{"l", "ls", "find"},
			Params:              switchListParam(),
			TableType:           output.TableSimple,
			TableColumnDefines:  switchListColumns(),
		},
		"create": {
			Type:          schema.CommandCreate,
			Aliases:       []string{"c"},
			Params:        switchCreateParam(),
			IncludeFields: switchDetailIncludes(),
			ExcludeFields: switchDetailExcludes(),
		},
		"read": {
			Type:          schema.CommandRead,
			Aliases:       []string{"r"},
			Params:        switchReadParam(),
			IncludeFields: switchDetailIncludes(),
			ExcludeFields: switchDetailExcludes(),
		},
		"update": {
			Type:          schema.CommandUpdate,
			Aliases:       []string{"u"},
			Params:        switchUpdateParam(),
			IncludeFields: switchDetailIncludes(),
			ExcludeFields: switchDetailExcludes(),
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"d", "rm"},
			Params:        switchDeleteParam(),
			IncludeFields: switchDetailIncludes(),
			ExcludeFields: switchDetailExcludes(),
		},
		"bridge-connect": {
			Type:             schema.CommandManipulate,
			UseCustomCommand: true,
			Params:           switchConnectBridgeParam(),
			IncludeFields:    switchDetailIncludes(),
			ExcludeFields:    switchDetailExcludes(),
		},
		"bridge-disconnect": {
			Type:             schema.CommandManipulate,
			UseCustomCommand: true,
			Params:           switchDisconnectBridgeParam(),
			IncludeFields:    switchDetailIncludes(),
			ExcludeFields:    switchDetailExcludes(),
		},
	}

	return &schema.Resource{
		Commands: commands,
	}
}

func switchListParam() map[string]*schema.Schema {
	return CommonListParam
}

func switchListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{
			Name:    "Server",
			Sources: []string{"ServerCount"},
		},
		{
			Name:    "Appliance",
			Sources: []string{"ApplianceCount"},
		},
		{
			Name:    "Bridge",
			Sources: []string{"Bridge.Name"},
		},
		{
			Name:    "Gateway",
			Sources: []string{"Subnets.0.DefaultRoute"},
		},
		{
			Name: "Network",
			Sources: []string{
				"Subnets.0.NetworkAddress",
				"Subnets.0.NetworkMaskLen",
			},
			Format: "%s/%s",
		},
		{
			Name:    "BandWidth",
			Sources: []string{"Subnets.0.Internet.BandWidthMbps"},
			Format:  "%sMbps",
		},
	}
}

func switchDetailIncludes() []string {
	return []string{}
}

func switchDetailExcludes() []string {
	return []string{
		"Bridge.SwitchInZone",
		"Bridge.Info.Switches",
	}
}

func switchCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramRequiredName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon":        getParamSubResourceID("Icon"),
	}
}

func switchReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func switchUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id":          paramID,
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon":        getParamSubResourceID("Icon"),
	}
}

func switchDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func switchConnectBridgeParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id":        paramID,
		"bridge-id": getParamResourceID("bridge ID"),
	}
}

func switchDisconnectBridgeParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

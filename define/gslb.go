package define

import (
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func GSLBResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:                schema.CommandList,
			ListResultFieldName: "CommonServiceGSLBItems",
			Aliases:             []string{"l", "ls", "find"},
			Params:              gslbListParam(),
			TableType:           output.TableSimple,
			TableColumnDefines:  gslbListColumns(),
		},
		"create": {
			Type:             schema.CommandCreate,
			Aliases:          []string{"c"},
			Params:           gslbCreateParam(),
			IncludeFields:    gslbDetailIncludes(),
			ExcludeFields:    gslbDetailExcludes(),
			UseCustomCommand: true,
		},
		"read": {
			Type:          schema.CommandRead,
			Aliases:       []string{"r"},
			Params:        gslbReadParam(),
			IncludeFields: gslbDetailIncludes(),
			ExcludeFields: gslbDetailExcludes(),
		},
		"update": {
			Type:             schema.CommandUpdate,
			Aliases:          []string{"u"},
			Params:           gslbUpdateParam(),
			IncludeFields:    gslbDetailIncludes(),
			ExcludeFields:    gslbDetailExcludes(),
			UseCustomCommand: true,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"d", "rm"},
			Params:        gslbDeleteParam(),
			IncludeFields: gslbDetailIncludes(),
			ExcludeFields: gslbDetailExcludes(),
		},
		"server-list": {
			Type:               schema.CommandManipulate,
			Params:             gslbServerListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: gslbServerListColumns(),
			UseCustomCommand:   true,
		},
		"server-add": {
			Type:               schema.CommandManipulate,
			Params:             gslbServerAddParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: gslbServerListColumns(),
			UseCustomCommand:   true,
		},
		"server-update": {
			Type:               schema.CommandManipulate,
			Params:             gslbServerUpdateParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: gslbServerListColumns(),
			UseCustomCommand:   true,
		},
		"server-delete": {
			Type:               schema.CommandManipulate,
			Params:             gslbServerDeleteParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: gslbServerListColumns(),
			UseCustomCommand:   true,
		},
	}

	return &schema.Resource{
		Commands:         commands,
		ResourceCategory: CategoryCommonServiceItem,
	}
}

func gslbListParam() map[string]*schema.Schema {
	return CommonListParam
}

func gslbListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{
			Name:    "Health-Check",
			Sources: []string{"Settings.GSLB.HealthCheck.Protocol"},
		},
		{
			Name:    "FQDN",
			Sources: []string{"Status.FQDN"},
		},
		{
			Name:    "Sorry-Server",
			Sources: []string{"Settings.GSLB.SorryServer"},
		},
	}
}

func gslbServerListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "__ORDER__"}, // magic column name(generated on demand)
		{Name: "IPAddress"},
		{Name: "Weight"},
		{Name: "Enabled"},
	}
}

func gslbDetailIncludes() []string {
	return []string{}
}

func gslbDetailExcludes() []string {
	return []string{}
}

func gslbCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramRequiredName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     getParamSubResourceID("Icon"),
		"protocol": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set healthcheck protocol[http/https/ping/tcp]",
			DefaultValue: "ping",
			Required:     true,
			ValidateFunc: validateInStrValues(sacloud.AllowGSLBHealthCheckProtocol()...),
		},
		"host-header": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set host header of http/https healthcheck request",
		},
		"path": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set path of http/https healthcheck request",
			DefaultValue: "/",
		},
		"response-code": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set response-code of http/https healthcheck request",
			DefaultValue: 200,
		},
		"port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set port of tcp healthcheck",
			ValidateFunc: validateIntRange(1, 65535),
		},
		"delay-loop": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set delay-loop of healthcheck",
			ValidateFunc: validateIntRange(10, 60),
			Required:     true,
			DefaultValue: 10,
		},
		"weighted": {
			Type:         schema.TypeBool,
			HandlerType:  schema.HandlerNoop,
			Description:  "enable weighted",
			DefaultValue: true,
		},
		"sorry-server": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set sorry-server hostname/ipaddress",
		},
	}
}

func gslbReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func gslbUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id":          paramID,
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     getParamSubResourceID("Icon"),
		"protocol": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set healthcheck protocol[http/https/ping/tcp]",
			ValidateFunc: validateInStrValues(sacloud.AllowGSLBHealthCheckProtocol()...),
		},
		"host-header": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set host header of http/https healthcheck request",
		},
		"path": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set path of http/https healthcheck request",
		},
		"response-code": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "set response-code of http/https healthcheck request",
		},
		"port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set port of tcp healthcheck",
			ValidateFunc: validateIntRange(1, 65535),
		},
		"delay-loop": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set delay-loop of healthcheck",
			ValidateFunc: validateIntRange(10, 60),
		},
		"weighted": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "enable weighted",
		},
		"sorry-server": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set sorry-server hostname/ipaddress",
		},
	}
}

func gslbDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func gslbServerListParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func gslbServerAddParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"ipaddress": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target ipaddress",
			ValidateFunc: validateIPv4Address(),
		},
		"enabled": {
			Type:         schema.TypeBool,
			HandlerType:  schema.HandlerNoop,
			Description:  "set enabled",
			DefaultValue: true,
		},
		"weight": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set weight",
			ValidateFunc: validateIntRange(1, 10000),
		},
	}
}
func gslbServerUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target server",
			Required:    true,
		},
		"ipaddress": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target ipaddress",
			ValidateFunc: validateIPv4Address(),
		},
		"enabled": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set enabled",
		},
		"weight": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set weight",
			ValidateFunc: validateIntRange(1, 10000),
		},
	}
}
func gslbServerDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target server",
			Required:    true,
		},
	}
}

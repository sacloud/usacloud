package define

import (
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func GSLBResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find", "selector"},
			Params:             gslbListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: gslbListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"create": {
			Type:             schema.CommandCreate,
			Params:           gslbCreateParam(),
			IncludeFields:    gslbDetailIncludes(),
			ExcludeFields:    gslbDetailExcludes(),
			UseCustomCommand: true,
			Category:         "basics",
			Order:            20,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        gslbReadParam(),
			IncludeFields: gslbDetailIncludes(),
			ExcludeFields: gslbDetailExcludes(),
			Category:      "basics",
			Order:         30,
		},
		"update": {
			Type:             schema.CommandUpdate,
			Params:           gslbUpdateParam(),
			IncludeFields:    gslbDetailIncludes(),
			ExcludeFields:    gslbDetailExcludes(),
			UseCustomCommand: true,
			Category:         "basics",
			Order:            40,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"rm"},
			Params:        gslbDeleteParam(),
			IncludeFields: gslbDetailIncludes(),
			ExcludeFields: gslbDetailExcludes(),
			Category:      "basics",
			Order:         50,
		},
		"server-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             gslbServerListParam(),
			Aliases:            []string{"server-list"},
			TableType:          output.TableSimple,
			TableColumnDefines: gslbServerListColumns(),
			UseCustomCommand:   true,
			NeedlessConfirm:    true,
			Category:           "servers",
			Order:              10,
		},
		"server-add": {
			Type:               schema.CommandManipulateSingle,
			Params:             gslbServerAddParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: gslbServerListColumns(),
			UseCustomCommand:   true,
			Category:           "servers",
			Order:              20,
		},
		"server-update": {
			Type:               schema.CommandManipulateSingle,
			Params:             gslbServerUpdateParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: gslbServerListColumns(),
			UseCustomCommand:   true,
			Category:           "servers",
			Order:              30,
		},
		"server-delete": {
			Type:               schema.CommandManipulateSingle,
			Params:             gslbServerDeleteParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: gslbServerListColumns(),
			UseCustomCommand:   true,
			ConfirmMessage:     "delete server",
			Category:           "servers",
			Order:              40,
		},
	}

	return &schema.Resource{
		Commands:            commands,
		ResourceCategory:    CategoryCommonServiceItem,
		ListResultFieldName: "CommonServiceGSLBItems",
	}
}

func gslbListParam() map[string]*schema.Schema {
	return mergeParameterMap(CommonListParam, paramTagsCond)
}

func gslbListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{
			Name:    "HealthCheck",
			Sources: []string{"Settings.GSLB.HealthCheck.Protocol"},
		},
		{
			Name:    "FQDN",
			Sources: []string{"Status.FQDN"},
		},
		{
			Name:    "SorryServer",
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
		"protocol": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set healthcheck protocol[http/https/ping/tcp]",
			DefaultValue: "ping",
			Required:     true,
			ValidateFunc: validateInStrValues(sacloud.AllowGSLBHealthCheckProtocol()...),
			CompleteFunc: completeInStrValues(sacloud.AllowGSLBHealthCheckProtocol()...),
			Category:     "GSLB",
			Order:        10,
		},
		"host-header": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set host header of http/https healthcheck request",
			Category:    "GSLB",
			Order:       20,
		},
		"path": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set path of http/https healthcheck request",
			DefaultValue: "/",
			Category:     "GSLB",
			Order:        30,
		},
		"response-code": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set response-code of http/https healthcheck request",
			DefaultValue: 200,
			Category:     "GSLB",
			Order:        40,
		},
		"port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set port of tcp healthcheck",
			ValidateFunc: validateIntRange(1, 65535),
			Category:     "GSLB",
			Order:        50,
		},
		"delay-loop": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set delay-loop of healthcheck",
			ValidateFunc: validateIntRange(10, 60),
			Required:     true,
			DefaultValue: 10,
			Category:     "GSLB",
			Order:        60,
		},
		"weighted": {
			Type:         schema.TypeBool,
			HandlerType:  schema.HandlerNoop,
			Description:  "enable weighted",
			DefaultValue: true,
			Category:     "GSLB",
			Order:        70,
		},
		"sorry-server": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set sorry-server hostname/ipaddress",
			Category:    "GSLB",
			Order:       80,
		},
		"name":        paramRequiredName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func gslbReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func gslbUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"protocol": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set healthcheck protocol[http/https/ping/tcp]",
			ValidateFunc: validateInStrValues(sacloud.AllowGSLBHealthCheckProtocol()...),
			CompleteFunc: completeInStrValues(sacloud.AllowGSLBHealthCheckProtocol()...),
			Category:     "GSLB",
			Order:        10,
		},
		"host-header": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set host header of http/https healthcheck request",
			Category:    "GSLB",
			Order:       20,
		},
		"path": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set path of http/https healthcheck request",
			Category:    "GSLB",
			Order:       30,
		},
		"response-code": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "set response-code of http/https healthcheck request",
			Category:    "GSLB",
			Order:       40,
		},
		"port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set port of tcp healthcheck",
			ValidateFunc: validateIntRange(1, 65535),
			Category:     "GSLB",
			Order:        50,
		},
		"delay-loop": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set delay-loop of healthcheck",
			ValidateFunc: validateIntRange(10, 60),
			Category:     "GSLB",
			Order:        60,
		},
		"weighted": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "enable weighted",
			Category:    "GSLB",
			Order:       70,
		},
		"sorry-server": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set sorry-server hostname/ipaddress",
			Category:    "GSLB",
			Order:       80,
		},
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func gslbDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func gslbServerListParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func gslbServerAddParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ipaddress": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target ipaddress",
			ValidateFunc: validateIPv4Address(),
			Category:     "server",
			Order:        10,
		},
		"enabled": {
			Type:         schema.TypeBool,
			HandlerType:  schema.HandlerNoop,
			Description:  "set enabled",
			DefaultValue: true,
			Category:     "server",
			Order:        20,
		},
		"weight": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set weight",
			ValidateFunc: validateIntRange(1, 10000),
			Category:     "server",
			Order:        30,
		},
	}
}
func gslbServerUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target server",
			Required:    true,
			Category:    "server",
			Order:       1,
		},
		"ipaddress": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target ipaddress",
			ValidateFunc: validateIPv4Address(),
			Category:     "server",
			Order:        10,
		},
		"enabled": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set enabled",
			Category:    "server",
			Order:       20,
		},
		"weight": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set weight",
			ValidateFunc: validateIntRange(1, 10000),
			Category:     "server",
			Order:        30,
		},
	}
}
func gslbServerDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target server",
			Required:    true,
			Category:    "server",
			Order:       1,
		},
	}
}

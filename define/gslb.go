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
			Aliases:            []string{"ls", "find"},
			Params:             gslbListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: gslbListColumns(),
		},
		"create": {
			Type:             schema.CommandCreate,
			Params:           gslbCreateParam(),
			IncludeFields:    gslbDetailIncludes(),
			ExcludeFields:    gslbDetailExcludes(),
			UseCustomCommand: true,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        gslbReadParam(),
			IncludeFields: gslbDetailIncludes(),
			ExcludeFields: gslbDetailExcludes(),
		},
		"update": {
			Type:             schema.CommandUpdate,
			Params:           gslbUpdateParam(),
			IncludeFields:    gslbDetailIncludes(),
			ExcludeFields:    gslbDetailExcludes(),
			UseCustomCommand: true,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"rm"},
			Params:        gslbDeleteParam(),
			IncludeFields: gslbDetailIncludes(),
			ExcludeFields: gslbDetailExcludes(),
		},
		"server-list": {
			Type:               schema.CommandManipulateSingle,
			Params:             gslbServerListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: gslbServerListColumns(),
			UseCustomCommand:   true,
			NeedlessConfirm:    true,
		},
		"server-add": {
			Type:               schema.CommandManipulateSingle,
			Params:             gslbServerAddParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: gslbServerListColumns(),
			UseCustomCommand:   true,
		},
		"server-update": {
			Type:               schema.CommandManipulateSingle,
			Params:             gslbServerUpdateParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: gslbServerListColumns(),
			UseCustomCommand:   true,
		},
		"server-delete": {
			Type:               schema.CommandManipulateSingle,
			Params:             gslbServerDeleteParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: gslbServerListColumns(),
			UseCustomCommand:   true,
			ConfirmMessage:     "delete server",
		},
	}

	return &schema.Resource{
		Commands:            commands,
		ResourceCategory:    CategoryCommonServiceItem,
		ListResultFieldName: "CommonServiceGSLBItems",
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
		"name":        paramRequiredName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
		"protocol": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set healthcheck protocol[http/https/ping/tcp]",
			DefaultValue: "ping",
			Required:     true,
			ValidateFunc: validateInStrValues(sacloud.AllowGSLBHealthCheckProtocol()...),
			CompleteFunc: completeInStrValues(sacloud.AllowGSLBHealthCheckProtocol()...),
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
	return map[string]*schema.Schema{}
}

func gslbUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
		"protocol": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set healthcheck protocol[http/https/ping/tcp]",
			ValidateFunc: validateInStrValues(sacloud.AllowGSLBHealthCheckProtocol()...),
			CompleteFunc: completeInStrValues(sacloud.AllowGSLBHealthCheckProtocol()...),
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
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target server",
			Required:    true,
		},
	}
}

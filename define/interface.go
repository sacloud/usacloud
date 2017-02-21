package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func InterfaceResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:                schema.CommandList,
			ListResultFieldName: "Interfaces",
			Aliases:             []string{"l", "ls", "find"},
			Params:              interfaceListParam(),
			TableType:           output.TableSimple,
			TableColumnDefines:  interfaceListColumns(),
			UseCustomCommand:    true, // to ignore appliance(system servers)
		},
		"create": {
			Type:          schema.CommandCreate,
			Aliases:       []string{"c"},
			Params:        interfaceCreateParam(),
			IncludeFields: interfaceDetailIncludes(),
			ExcludeFields: interfaceDetailExcludes(),
		},
		"read": {
			Type:          schema.CommandRead,
			Aliases:       []string{"r"},
			Params:        interfaceReadParam(),
			IncludeFields: interfaceDetailIncludes(),
			ExcludeFields: interfaceDetailExcludes(),
		},
		"update": {
			Type:          schema.CommandUpdate,
			Aliases:       []string{"u"},
			Params:        interfaceUpdateParam(),
			IncludeFields: interfaceDetailIncludes(),
			ExcludeFields: interfaceDetailExcludes(),
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"d", "rm"},
			Params:        interfaceDeleteParam(),
			IncludeFields: interfaceDetailIncludes(),
			ExcludeFields: interfaceDetailExcludes(),
		},
		"packet-filter-connect": {
			Type:             schema.CommandManipulate,
			Params:           interfacePacketFilterConnectParam(),
			UseCustomCommand: true,
		},
		"packet-filter-disconnect": {
			Type:             schema.CommandManipulate,
			Params:           interfacePacketFilterDisconnectParam(),
			UseCustomCommand: true,
		},
	}

	return &schema.Resource{
		Commands: commands,
	}
}

func interfaceListParam() map[string]*schema.Schema {
	return CommonListParam
}

func interfaceListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{
			Name:    "User-IPAddress",
			Sources: []string{"UserIPAddress"},
			Format:  "%s",
		},
		{Name: "MACAddress"},
		{
			Name:    "Gateway",
			Sources: []string{"Switch.Subnets.0.DefaultRoute"},
		},
		{
			Name: "Network",
			Sources: []string{
				"Switch.Subnets.0.NetworkAddress",
				"Switch.Subnets.0.NetworkMaskLen",
			},
			Format: "%s/%s",
		},
		{
			Name:    "BandWidth",
			Sources: []string{"Switch.Subnets.0.Internet.BandWidthMbps"},
			Format:  "%sMbps",
		},
		{
			Name:    "Switch-ID",
			Sources: []string{"Switch.ID"},
			Format:  "%s",
		},
		{
			Name:    "Server-ID",
			Sources: []string{"Server.ID"},
			Format:  "%s",
		},
	}
}

func interfaceDetailIncludes() []string {
	return []string{}
}

func interfaceDetailExcludes() []string {
	return []string{}
}

func interfaceCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"server-id": {
			Type:            schema.TypeInt64,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetServerID",
			Description:     "set server ID",
			Required:        true,
			ValidateFunc:    validateSakuraID(),
		},
	}
}

func interfaceReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func interfaceUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"user-ipaddress": {
			Type:            schema.TypeString,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetUserIPAddress",
			Description:     "set user-ipaddress",
			ValidateFunc:    validateIPv4Address(),
		},
	}
}

func interfaceDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func interfacePacketFilterConnectParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"packet-filter-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set packet filter ID",
			Required:     true,
			ValidateFunc: validateSakuraID(),
		},
	}
}

func interfacePacketFilterDisconnectParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"packet-filter-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set packet filter ID",
			Required:     true,
			ValidateFunc: validateSakuraID(),
		},
	}
}

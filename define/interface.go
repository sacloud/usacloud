package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func InterfaceResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find"},
			Params:             interfaceListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: interfaceListColumns(),
			UseCustomCommand:   true, // to ignore appliance(system servers)
			Category:           "basics",
			Order:              10,
		},
		"create": {
			Type:          schema.CommandCreate,
			Params:        interfaceCreateParam(),
			IncludeFields: interfaceDetailIncludes(),
			ExcludeFields: interfaceDetailExcludes(),
			Category:      "basics",
			Order:         20,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        interfaceReadParam(),
			IncludeFields: interfaceDetailIncludes(),
			ExcludeFields: interfaceDetailExcludes(),
			Category:      "basics",
			Order:         30,
		},
		"update": {
			Type:          schema.CommandUpdate,
			Params:        interfaceUpdateParam(),
			IncludeFields: interfaceDetailIncludes(),
			ExcludeFields: interfaceDetailExcludes(),
			Category:      "basics",
			Order:         40,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"rm"},
			Params:        interfaceDeleteParam(),
			IncludeFields: interfaceDetailIncludes(),
			ExcludeFields: interfaceDetailExcludes(),
			Category:      "basics",
			Order:         50,
		},
		"packet-filter-connect": {
			Type:             schema.CommandManipulateMulti,
			Params:           interfacePacketFilterConnectParam(),
			UseCustomCommand: true,
			NoOutput:         true,
			Category:         "packer-filter",
			Order:            10,
		},
		"packet-filter-disconnect": {
			Type:             schema.CommandManipulateMulti,
			Params:           interfacePacketFilterDisconnectParam(),
			UseCustomCommand: true,
			NoOutput:         true,
			Category:         "packer-filter",
			Order:            20,
		},
	}

	return &schema.Resource{
		Commands:         commands,
		ResourceCategory: CategoryNetworking,
	}
}

func interfaceListParam() map[string]*schema.Schema {
	return CommonListParam
}

func interfaceListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{
			Name:    "UserIPAddress",
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
			Name:    "SwitchID",
			Sources: []string{"Switch.ID"},
			Format:  "%s",
		},
		{
			Name:    "ServerID",
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
			CompleteFunc:    completeServerID(),
			Category:        "interface",
			Order:           10,
		},
	}
}

func interfaceReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func interfaceUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"user-ipaddress": {
			Type:            schema.TypeString,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetUserIPAddress",
			Description:     "set user-ipaddress",
			ValidateFunc:    validateIPv4Address(),
			Category:        "interface",
			Order:           10,
		},
	}
}

func interfaceDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func interfacePacketFilterConnectParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"packet-filter-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set packet filter ID",
			Required:     true,
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completePacketFilterID(),
			Category:     "packet-filter",
			Order:        10,
		},
	}
}

func interfacePacketFilterDisconnectParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"packet-filter-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set packet filter ID",
			Required:     true,
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completePacketFilterID(),
			Category:     "packet-filter",
			Order:        10,
		},
	}
}

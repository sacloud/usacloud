package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func IPv4Resource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find"},
			Params:             ipv4ListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: ipv4ListColumns(),
			UseCustomCommand:   true,
			ArgsUsage:          "IPAddress",
			Category:           "basics",
			Order:              10,
		},
		"ptr-add": {
			Type:             schema.CommandCustom,
			Params:           ipv4PTRCreateParam(),
			IncludeFields:    ipv4DetailIncludes(),
			ExcludeFields:    ipv4DetailExcludes(),
			UseCustomCommand: true,
			ArgsUsage:        "IPAddress",
			Category:         "basics",
			Order:            20,
		},
		"ptr-read": {
			Type:             schema.CommandCustom,
			Params:           ipv4PTRReadParam(),
			IncludeFields:    ipv4DetailIncludes(),
			ExcludeFields:    ipv4DetailExcludes(),
			NeedlessConfirm:  true,
			UseCustomCommand: true,
			ArgsUsage:        "IPAddress",
			Category:         "basics",
			Order:            30,
		},
		"ptr-update": {
			Type:             schema.CommandCustom,
			Params:           ipv4PTRUpdateParam(),
			IncludeFields:    ipv4DetailIncludes(),
			ExcludeFields:    ipv4DetailExcludes(),
			UseCustomCommand: true,
			ArgsUsage:        "IPAddress",
			Category:         "basics",
			Order:            40,
		},
		"ptr-delete": {
			Type:             schema.CommandCustom,
			Params:           ipv4PTRDeleteParam(),
			IncludeFields:    ipv4DetailIncludes(),
			ExcludeFields:    ipv4DetailExcludes(),
			UseCustomCommand: true,
			ArgsUsage:        "IPAddress",
			Category:         "basics",
			Order:            50,
		},
	}

	return &schema.Resource{
		Commands:         commands,
		ResourceCategory: CategoryNetworking,
	}
}

func ipv4ListParam() map[string]*schema.Schema {
	return CommonListParam
}

func ipv4ListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "IPAddress"},
		{Name: "HostName"},
		{
			Name:    "Interface-ID",
			Sources: []string{"Interface.ID"},
		},
		{
			Name:    "MACAddress",
			Sources: []string{"Interface.MACAddress"},
		},
		{
			Name:    "Server-ID",
			Sources: []string{"Interface.Server.ID"},
		},
		{
			Name:    "Server-Name",
			Sources: []string{"Interface.Server.Name"},
		},
		{
			Name:    "Server-HostName",
			Sources: []string{"Interface.Server.HostName"},
		},
	}
}

func ipv4DetailIncludes() []string {
	return []string{}
}

func ipv4DetailExcludes() []string {
	return []string{}
}

func ipv4PTRCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"hostname": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set server hostname",
			Required:    true,
			Category:    "ipv4",
			Order:       10,
		},
	}
}

func ipv4PTRReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func ipv4PTRUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"hostname": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set server hostname",
			Required:    true,
			Category:    "ipv4",
			Order:       10,
		},
	}
}

func ipv4PTRDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

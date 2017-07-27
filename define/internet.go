package define

import (
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func InternetResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find", "selector"},
			Params:             internetListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: internetListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"create": {
			Type:          schema.CommandCreate,
			Params:        internetCreateParam(),
			IncludeFields: internetDetailIncludes(),
			ExcludeFields: internetDetailExcludes(),
			Category:      "basics",
			Order:         20,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        internetReadParam(),
			IncludeFields: internetDetailIncludes(),
			ExcludeFields: internetDetailExcludes(),
			Category:      "basics",
			Order:         30,
		},
		"update": {
			Type:          schema.CommandUpdate,
			Params:        internetUpdateParam(),
			IncludeFields: internetDetailIncludes(),
			ExcludeFields: internetDetailExcludes(),
			Category:      "basics",
			Order:         40,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"rm"},
			Params:        internetDeleteParam(),
			IncludeFields: internetDetailIncludes(),
			ExcludeFields: internetDetailExcludes(),
			Category:      "basics",
			Order:         50,
		},
		"update-bandwidth": {
			Type:             schema.CommandManipulateMulti,
			Params:           internetUpdateBandWidthParam(),
			IncludeFields:    internetDetailIncludes(),
			ExcludeFields:    internetDetailExcludes(),
			UseCustomCommand: true,
			Category:         "spec",
			Order:            10,
		},
		"monitor": {
			Type:               schema.CommandRead,
			Params:             internetMonitorParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: internetMonitorColumns(),
			UseCustomCommand:   true,
			Category:           "monitor",
			Order:              10,
		},
	}

	return &schema.Resource{
		Commands:            commands,
		ResourceCategory:    CategoryNetworking,
		ListResultFieldName: "Internet",
	}
}

var internetCommandCategories = []schema.Category{
	{
		Key:         "basic",
		DisplayName: "Basics",
		Order:       10,
	},
	{
		Key:         "spec",
		DisplayName: "Router spec",
		Order:       20,
	},
	{
		Key:         "monitor",
		DisplayName: "Monitoring",
		Order:       30,
	},
}

func internetListParam() map[string]*schema.Schema {
	return mergeParameterMap(CommonListParam, paramTagsCond)
}

func internetListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
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
			Sources: []string{"BandWidthMbps"},
			Format:  "%sMbps",
		},
		{
			Name: "IPv6Prefix",
			Sources: []string{
				"Switch.IPv6Nets.0.IPv6Prefix",
				"Switch.IPv6Nets.0.IPv6PrefixLen",
			},
			Format: "%s/%s",
		},
	}
}

func internetDetailIncludes() []string {
	return []string{}
}

func internetDetailExcludes() []string {
	return []string{
		"Switch.UserSubnet",
	}
}

func internetCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"nw-masklen": {
			Type:            schema.TypeInt,
			HandlerType:     schema.HandlerPathThrough,
			Aliases:         []string{"network-masklen"},
			Description:     "set Global-IPAddress prefix",
			DestinationProp: "SetNetworkMaskLen",
			Required:        true,
			DefaultValue:    28,
			ValidateFunc:    validateInIntValues(sacloud.AllowInternetNetworkMaskLen()...),
			Category:        "router",
			Order:           10,
		},
		"band-width": {
			Type:            schema.TypeInt,
			HandlerType:     schema.HandlerPathThrough,
			Description:     "set band-width(Mbpm)",
			DestinationProp: "SetBandWidthMbps",
			Required:        true,
			DefaultValue:    100,
			ValidateFunc:    validateInIntValues(sacloud.AllowInternetBandWidth()...),
			CompleteFunc:    completeInIntValues(sacloud.AllowInternetBandWidth()...),
			Category:        "router",
			Order:           20,
		},
		"name":        paramRequiredName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func internetReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func internetUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"band-width": {
			Type:            schema.TypeInt,
			HandlerType:     schema.HandlerPathThrough,
			Description:     "set band-width(Mbpm)",
			DestinationProp: "SetBandWidthMbps",
			ValidateFunc:    validateInIntValues(sacloud.AllowInternetBandWidth()...),
			CompleteFunc:    completeInIntValues(sacloud.AllowInternetBandWidth()...),
			Category:        "router",
			Order:           20,
		},
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func internetDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func internetUpdateBandWidthParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"band-width": {
			Type:            schema.TypeInt,
			HandlerType:     schema.HandlerPathThrough,
			Description:     "set band-width(Mbpm)",
			DestinationProp: "SetBandWidthMbps",
			Required:        true,
			DefaultValue:    100,
			ValidateFunc:    validateInIntValues(sacloud.AllowInternetBandWidth()...),
			CompleteFunc:    completeInIntValues(sacloud.AllowInternetBandWidth()...),
			Category:        "router",
			Order:           20,
		},
	}
}

func internetMonitorParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"start": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set start-time",
			ValidateFunc: validateDateTimeString(),
			Category:     "monitor",
			Order:        10,
		},
		"end": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set end-time",
			ValidateFunc: validateDateTimeString(),
			Category:     "monitor",
			Order:        20,
		},
		"key-format": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set monitoring value key-format",
			DefaultValue: "sakuracloud.internet.{{.ID}}.nic",
			Required:     true,
			Category:     "monitor",
			Order:        30,
		},
	}
}

func internetMonitorColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Key"},
		{Name: "TimeStamp"},
		{Name: "UnixTime"},
		{Name: "In"},
		{Name: "Out"},
	}
}

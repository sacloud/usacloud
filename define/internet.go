package define

import (
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func InternetResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:                schema.CommandList,
			ListResultFieldName: "Internet",
			Aliases:             []string{"l", "ls", "find"},
			Params:              internetListParam(),
			TableType:           output.TableSimple,
			TableColumnDefines:  internetListColumns(),
		},
		"create": {
			Type:          schema.CommandCreate,
			Aliases:       []string{"c"},
			Params:        internetCreateParam(),
			IncludeFields: internetDetailIncludes(),
			ExcludeFields: internetDetailExcludes(),
		},
		"read": {
			Type:          schema.CommandRead,
			Aliases:       []string{"r"},
			Params:        internetReadParam(),
			IncludeFields: internetDetailIncludes(),
			ExcludeFields: internetDetailExcludes(),
		},
		"update": {
			Type:          schema.CommandUpdate,
			Aliases:       []string{"u"},
			Params:        internetUpdateParam(),
			IncludeFields: internetDetailIncludes(),
			ExcludeFields: internetDetailExcludes(),
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"d", "rm"},
			Params:        internetDeleteParam(),
			IncludeFields: internetDetailIncludes(),
			ExcludeFields: internetDetailExcludes(),
		},
		"update-bandwidth": {
			Type:             schema.CommandManipulate,
			Params:           internetUpdateBandWidthParam(),
			IncludeFields:    internetDetailIncludes(),
			ExcludeFields:    internetDetailExcludes(),
			UseCustomCommand: true,
		},
	}

	return &schema.Resource{
		Commands: commands,
	}
}

func internetListParam() map[string]*schema.Schema {
	return CommonListParam
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
		"name":        paramRequiredName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon":        getParamSubResourceID("Icon"),
		"nw-masklen": {
			Type:            schema.TypeInt,
			HandlerType:     schema.HandlerPathThrough,
			Description:     "set Global-IPAddress prefix",
			DestinationProp: "SetNetworkMaskLen",
			Required:        true,
			DefaultValue:    28,
			ValidateFunc:    validateInIntValues(sacloud.AllowInternetNetworkMaskLen()...),
		},
	}
}

func internetReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func internetUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id":          paramID,
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon":        getParamSubResourceID("Icon"),
		"band-width": {
			Type:            schema.TypeInt,
			HandlerType:     schema.HandlerPathThrough,
			Description:     "set band-width(Mbpm)",
			DestinationProp: "SetBandWidthMbps",
			Required:        true,
			DefaultValue:    100,
			ValidateFunc:    validateInIntValues(sacloud.AllowInternetBandWidth()...),
		},
	}
}

func internetDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func internetUpdateBandWidthParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"band-width": {
			Type:            schema.TypeInt,
			HandlerType:     schema.HandlerPathThrough,
			Description:     "set band-width(Mbpm)",
			DestinationProp: "SetBandWidthMbps",
			Required:        true,
			DefaultValue:    100,
			ValidateFunc:    validateInIntValues(sacloud.AllowInternetBandWidth()...),
		},
	}
}

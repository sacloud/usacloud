// Copyright 2017-2020 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package define

import (
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/schema"
)

func InternetResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find", "select"},
			Params:             internetListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: internetListColumns(),
			Category:           "basic",
			Order:              10,
		},
		"create": {
			Type:          schema.CommandCreate,
			Params:        internetCreateParam(),
			IncludeFields: internetDetailIncludes(),
			ExcludeFields: internetDetailExcludes(),
			Category:      "basic",
			Order:         20,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        internetReadParam(),
			IncludeFields: internetDetailIncludes(),
			ExcludeFields: internetDetailExcludes(),
			Category:      "basic",
			Order:         30,
		},
		"update": {
			Type:          schema.CommandUpdate,
			Params:        internetUpdateParam(),
			IncludeFields: internetDetailIncludes(),
			ExcludeFields: internetDetailExcludes(),
			Category:      "basic",
			Order:         40,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"rm"},
			Params:        internetDeleteParam(),
			IncludeFields: internetDetailIncludes(),
			ExcludeFields: internetDetailExcludes(),
			Category:      "basic",
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
		"subnet-info": {
			Type:               schema.CommandManipulateMulti,
			Params:             internetSubnetInfoParam(),
			IncludeFields:      internetDetailIncludes(),
			ExcludeFields:      internetDetailExcludes(),
			TableType:          output.TableSimple,
			TableColumnDefines: internetSubnetInfoColumns(),
			UseCustomCommand:   true,
			Category:           "subnet",
			NeedlessConfirm:    true,
			Order:              10,
		},
		"subnet-add": {
			Type:               schema.CommandManipulateMulti,
			Params:             internetSubnetAddParam(),
			IncludeFields:      internetDetailIncludes(),
			ExcludeFields:      internetDetailExcludes(),
			TableType:          output.TableSimple,
			TableColumnDefines: internetSubnetInfoColumns(),
			UseCustomCommand:   true,
			Category:           "subnet",
			Order:              20,
		},
		"subnet-update": {
			Type:               schema.CommandManipulateMulti,
			Params:             internetSubnetUpdateParam(),
			IncludeFields:      internetDetailIncludes(),
			ExcludeFields:      internetDetailExcludes(),
			TableType:          output.TableSimple,
			TableColumnDefines: internetSubnetInfoColumns(),
			UseCustomCommand:   true,
			Category:           "subnet",
			Order:              20,
		},
		"subnet-delete": {
			Type:             schema.CommandManipulateMulti,
			Params:           internetSubnetDeleteParam(),
			IncludeFields:    internetDetailIncludes(),
			ExcludeFields:    internetDetailExcludes(),
			UseCustomCommand: true,
			NoOutput:         true,
			Category:         "subnet",
			Order:            20,
		},
		"ipv6-info": {
			Type:               schema.CommandManipulateMulti,
			Params:             internetIPv6EnableParam(),
			IncludeFields:      internetDetailIncludes(),
			ExcludeFields:      internetDetailExcludes(),
			TableType:          output.TableSimple,
			TableColumnDefines: internetIPv6NetInfoColumns(),
			UseCustomCommand:   true,
			Category:           "ipv6",
			NeedlessConfirm:    true,
			Order:              10,
		},
		"ipv6-enable": {
			Type:               schema.CommandManipulateMulti,
			Params:             internetIPv6EnableParam(),
			IncludeFields:      internetDetailIncludes(),
			ExcludeFields:      internetDetailExcludes(),
			TableType:          output.TableSimple,
			TableColumnDefines: internetIPv6NetInfoColumns(),
			UseCustomCommand:   true,
			Category:           "ipv6",
			Order:              20,
		},
		"ipv6-disable": {
			Type:             schema.CommandManipulateMulti,
			Params:           internetIPv6EnableParam(),
			IncludeFields:    internetDetailIncludes(),
			ExcludeFields:    internetDetailExcludes(),
			UseCustomCommand: true,
			Category:         "ipv6",
			NoOutput:         true,
			Order:            30,
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
		CommandCategories:   internetCommandCategories,
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
		Key:         "subnet",
		DisplayName: "Router Subnet Management",
		Order:       30,
	},
	{
		Key:         "ipv6",
		DisplayName: "Router IPv6 Network Management",
		Order:       40,
	},
	{
		Key:         "monitor",
		DisplayName: "Monitoring",
		Order:       90,
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

func internetSubnetInfoColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{
			Name:    "NetworkAddress",
			Sources: []string{"NetworkAddress", "NetworkMaskLen"},
			Format:  "%s/%s",
		},
		{Name: "DefaultRoute"},
		{Name: "NextHop"},
		{
			Name:    "IPAddress-Range",
			Sources: []string{"IPAddressRangeStart", "IPAddressRangeEnd"},
			Format:  "%s - %s",
		},
	}
}

func internetIPv6NetInfoColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{
			Name:    "Prefix",
			Sources: []string{"IPv6Prefix"},
		},
		{
			Name:    "PrefixLen",
			Sources: []string{"IPv6PrefixLen"},
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
			ValidateFunc:    validateInIntValues(types.InternetNetworkMaskLengths...),
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
			ValidateFunc:    validateInIntValues(types.InternetBandWidths...),
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
			ValidateFunc:    validateInIntValues(types.InternetBandWidths...),
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
			ValidateFunc:    validateInIntValues(types.InternetBandWidths...),
			Category:        "router",
			Order:           20,
		},
	}
}

func internetSubnetInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func internetSubnetAddParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"nw-masklen": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"network-masklen"},
			Description:  "set Global-IPAddress(subnet) prefix",
			Required:     true,
			DefaultValue: 28,
			ValidateFunc: validateInIntValues(types.InternetNetworkMaskLengths...),
			Category:     "router",
			Order:        10,
		},
		"next-hop": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set NextHop IPAddress",
			ValidateFunc: validateIPv4Address(),
			Required:     true,
			Category:     "router",
			Order:        10,
		},
	}
}

func internetSubnetUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"subnet-id": {
			Type:         schema.TypeId,
			HandlerType:  schema.HandlerNoop,
			Description:  "set Target Subnet ID",
			ValidateFunc: validateSakuraShortID(12),
			Category:     "router",
			Order:        10,
		},
		"next-hop": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set NextHop IPAddress",
			ValidateFunc: validateIPv4Address(),
			Required:     true,
			Category:     "router",
			Order:        10,
		},
	}
}

func internetSubnetDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"subnet-id": {
			Type:         schema.TypeId,
			HandlerType:  schema.HandlerNoop,
			Description:  "set Target Subnet ID",
			ValidateFunc: validateSakuraShortID(12),
			Category:     "router",
			Order:        10,
		},
	}
}

func internetIPv6EnableParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
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

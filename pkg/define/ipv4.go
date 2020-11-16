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
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/schema"
)

func IPv4Resource() *schema.Resource {
	commands := map[string]*schema.Command{
		"list": {
			Type:    schema.CommandList,
			Aliases: []string{"ls", "find"},
			Params:  ipv4ListParam(),
			// TableType:          output.TableSimple,
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

func ipv4ListParam() map[string]*schema.Parameter {
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

func ipv4PTRCreateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
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

func ipv4PTRReadParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func ipv4PTRUpdateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
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

func ipv4PTRDeleteParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

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
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func IPv6Resource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find"},
			Params:             ipv6ListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: ipv6ListColumns(),
			UseCustomCommand:   true,
			ArgsUsage:          "IPAddress",
			Category:           "basics",
			Order:              10,
		},
		"ptr-add": {
			Type:             schema.CommandCustom,
			Params:           ipv6PTRCreateParam(),
			IncludeFields:    ipv6DetailIncludes(),
			ExcludeFields:    ipv6DetailExcludes(),
			UseCustomCommand: true,
			ArgsUsage:        "IPAddress",
			Category:         "basics",
			Order:            20,
		},
		"ptr-read": {
			Type:             schema.CommandCustom,
			Params:           ipv6PTRReadParam(),
			IncludeFields:    ipv6DetailIncludes(),
			ExcludeFields:    ipv6DetailExcludes(),
			NeedlessConfirm:  true,
			UseCustomCommand: true,
			ArgsUsage:        "IPAddress",
			Category:         "basics",
			Order:            30,
		},
		"ptr-update": {
			Type:             schema.CommandCustom,
			Params:           ipv6PTRUpdateParam(),
			IncludeFields:    ipv6DetailIncludes(),
			ExcludeFields:    ipv6DetailExcludes(),
			UseCustomCommand: true,
			ArgsUsage:        "IPAddress",
			Category:         "basics",
			Order:            40,
		},
		"ptr-delete": {
			Type:             schema.CommandCustom,
			Params:           ipv6PTRDeleteParam(),
			IncludeFields:    ipv6DetailIncludes(),
			ExcludeFields:    ipv6DetailExcludes(),
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

func ipv6ListParam() map[string]*schema.Schema {
	return mergeParameterMap(CommonListParam, paramIPv6NetIDCond, paramInternetIDCond)
}

var paramIPv6NetIDCond = map[string]*schema.Schema{
	"ipv6net-id": {
		Type:         schema.TypeId,
		HandlerType:  schema.HandlerFilterFunc,
		FilterFunc:   filterByIPv6NetID,
		Description:  "set filter by ipv6net-id",
		Category:     "filter",
		ValidateFunc: validateSakuraShortID(12),
		Order:        6,
	},
}

var paramInternetIDCond = map[string]*schema.Schema{
	"internet-id": {
		Type:         schema.TypeId,
		HandlerType:  schema.HandlerFilterFunc,
		FilterFunc:   filterByInternetID,
		Description:  "set filter by internet-id",
		Category:     "filter",
		ValidateFunc: validateSakuraID(),
		Order:        7,
	},
}

func filterByIPv6NetID(_ []interface{}, item interface{}, param interface{}) bool {

	type idHandler interface {
		GetIPv6NetID() int64
	}

	idHolder, ok := item.(idHandler)
	if !ok {
		return false
	}

	id := param.(int64)
	if id == 0 {
		return true
	}

	return idHolder.GetIPv6NetID() == id
}

func filterByInternetID(_ []interface{}, item interface{}, param interface{}) bool {

	type idHandler interface {
		GetInternetID() int64
	}

	idHolder, ok := item.(idHandler)
	if !ok {
		return false
	}

	id := param.(int64)
	if id == 0 {
		return true
	}

	return idHolder.GetInternetID() == id
}

func ipv6ListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{
			Name:    "IPAddress",
			Sources: []string{"IPv6Addr"},
		},
		{Name: "HostName"},
		{
			Name:    "IPv6Net-ID",
			Sources: []string{"IPv6Net.ID"},
		},
		{
			Name:    "IPv6Net-Prefix",
			Sources: []string{"IPv6Net.IPv6Prefix", "IPv6Net.IPv6PrefixLen"},
			Format:  "%s/%s",
		},
		{
			Name:    "Internet-ID",
			Sources: []string{"IPv6Net.Switch.Internet.ID"},
		},
		{
			Name:    "Internet-Name",
			Sources: []string{"IPv6Net.Switch.Internet.Name"},
		},
	}
}

func ipv6DetailIncludes() []string {
	return []string{}
}

func ipv6DetailExcludes() []string {
	return []string{}
}

func ipv6PTRCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"hostname": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set server hostname",
			Required:    true,
			Category:    "ipv6",
			Order:       10,
		},
	}
}

func ipv6PTRReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func ipv6PTRUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"hostname": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set server hostname",
			Required:    true,
			Category:    "ipv6",
			Order:       10,
		},
	}
}

func ipv6PTRDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

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

func GSLBResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find", "select"},
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
		ListResultFieldName: "GSLBs",
		IsGlobal:            true,
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
			ValidateFunc: validateInStrValues(types.GSLBHealthCheckProtocolStrings...),
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
			ValidateFunc: validateInStrValues(types.GSLBHealthCheckProtocolStrings...),
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
		"disabled": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set disabled",
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
		"disabled": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set disabled",
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

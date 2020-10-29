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

func DNSResource() *schema.Resource {
	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find", "select"},
			Params:             dnsListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: dnsListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"create": {
			Type:          schema.CommandCreate,
			Params:        dnsCreateParam(),
			IncludeFields: dnsDetailIncludes(),
			ExcludeFields: dnsDetailExcludes(),
			Category:      "basics",
			Order:         20,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        dnsReadParam(),
			IncludeFields: dnsDetailIncludes(),
			ExcludeFields: dnsDetailExcludes(),
			Category:      "basics",
			Order:         30,
		},
		"update": {
			Type:             schema.CommandUpdate,
			Params:           dnsUpdateParam(),
			IncludeFields:    dnsDetailIncludes(),
			ExcludeFields:    dnsDetailExcludes(),
			UseCustomCommand: true,
			Category:         "basics",
			Order:            40,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"rm"},
			Params:        dnsDeleteParam(),
			IncludeFields: dnsDetailIncludes(),
			ExcludeFields: dnsDetailExcludes(),
			Category:      "basics",
			Order:         50,
		},
		"record-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             dnsRecordListParam(),
			Aliases:            []string{"record-list"},
			TableType:          output.TableSimple,
			TableColumnDefines: dnsRecordListColumns(),
			UseCustomCommand:   true,
			NeedlessConfirm:    true,
			Category:           "records",
			Order:              10,
		},
		"record-bulk-update": {
			Type:             schema.CommandManipulateSingle,
			Params:           dnsRecordBulkUpdateParam(),
			UseCustomCommand: true,
			Category:         "records",
			Order:            15,
		},
		"record-add": {
			Type:               schema.CommandManipulateSingle,
			Params:             dnsRecordAddParam(),
			ParamCategories:    dnsCommandParamCategories,
			TableType:          output.TableSimple,
			TableColumnDefines: dnsRecordListColumns(),
			UseCustomCommand:   true,
			Category:           "records",
			Order:              20,
		},
		"record-update": {
			Type:               schema.CommandManipulateSingle,
			Params:             dnsRecordUpdateParam(),
			ParamCategories:    dnsCommandParamCategories,
			TableType:          output.TableSimple,
			TableColumnDefines: dnsRecordListColumns(),
			UseCustomCommand:   true,
			Category:           "records",
			Order:              30,
		},
		"record-delete": {
			Type:               schema.CommandManipulateSingle,
			Params:             dnsRecordDeleteParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: dnsRecordListColumns(),
			UseCustomCommand:   true,
			ConfirmMessage:     "delete record",
			Category:           "records",
			Order:              40,
		},
	}

	return &schema.Resource{
		Commands:            commands,
		ResourceCategory:    CategoryCommonServiceItem,
		ListResultFieldName: "DNS",
		IsGlobal:            true,
	}
}

var dnsCommandParamCategories = []schema.Category{
	{
		Key:         "record",
		DisplayName: "Common record options",
		Order:       10,
	},
	{
		Key:         "MX",
		DisplayName: "MX record options",
		Order:       20,
	},
	{
		Key:         "SRV",
		DisplayName: "SRV record options",
		Order:       30,
	},
}

func dnsListParam() map[string]*schema.Schema {
	return mergeParameterMap(CommonListParam, paramTagsCond)
}

func dnsListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{
			Name:    "NameServers",
			Sources: []string{"Status.NS.0", "Status.NS.1"},
			Format:  "%s / %s",
		},
	}
}

func dnsRecordListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Index"},
		{Name: "Type"},
		{Name: "Name"},
		{Name: "RData"},
	}
}

func dnsDetailIncludes() []string {
	return []string{}
}

func dnsDetailExcludes() []string {
	return []string{}
}

func dnsCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerPathThrough,
			Description:  "set DNS zone name",
			Required:     true,
			ValidateFunc: validateStrLen(2, 63),
			Category:     "common",
			Order:        500,
		},
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func dnsReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func dnsUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func dnsDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func dnsRecordListParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set name",
			ValidateFunc: validateStrLen(1, 63),
			Category:     "record",
			Order:        10,
		},
		"type": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set record type[A/AAAA/ALIAS/NS/CNAME/MX/TXT/SRV/CAA/PTR]",
			ValidateFunc: validateInStrValues(allowDNSTypes...),
			Category:     "record",
			Order:        20,
		},
	}
}

var allowDNSTypes = []string{
	"a", "aaaa", "alias", "ns", "cname", "mx", "txt", "srv", "caa", "ptr",
	"A", "AAAA", "ALIAS", "NS", "CNAME", "MX", "TXT", "SRV", "CAA", "PTR",
}

func dnsRecordBulkUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"file": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set name",
			Required:     true,
			ValidateFunc: validateFileExists(),
			Category:     "record",
			Order:        10,
		},
		"mode": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set name",
			Required:     true,
			DefaultValue: "upsert-only",
			ValidateFunc: validateInStrValues("upsert-only", "sync"),
			Category:     "record",
			Order:        20,
		},
	}
}

func dnsRecordAddParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set name",
			Required:     true,
			ValidateFunc: validateStrLen(1, 63),
			Category:     "record",
			Order:        10,
		},
		"type": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set record type[A/AAAA/ALIAS/NS/CNAME/MX/TXT/SRV/CAA/PTR]",
			Required:     true,
			ValidateFunc: validateInStrValues(allowDNSTypes...),
			Category:     "record",
			Order:        20,
		},
		"value": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set record data",
			Category:    "record",
			Order:       30,
		},
		"ttl": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set ttl",
			DefaultValue: 3600,
			ValidateFunc: validateIntRange(10, 3600000),
			Category:     "record",
			Order:        40,
		},
		"mx-priority": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set MX priority",
			DefaultValue: 10,
			ValidateFunc: validateIntRange(1, 65535),
			Category:     "MX",
			Order:        10,
		},
		"srv-priority": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set SRV priority",
			DefaultValue: 0,
			ValidateFunc: validateIntRange(0, 65535),
			Category:     "SRV",
			Order:        10,
		},
		"srv-weight": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set SRV priority",
			DefaultValue: 0,
			ValidateFunc: validateIntRange(0, 65535),
			Category:     "SRV",
			Order:        20,
		},
		"srv-port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set SRV priority",
			DefaultValue: 0,
			ValidateFunc: validateIntRange(1, 65535),
			Category:     "SRV",
			Order:        30,
		},
		"srv-target": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set SRV priority",
			ValidateFunc: validateStrLen(1, 254),
			Category:     "SRV",
			Order:        40,
		},
	}
}
func dnsRecordUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target record",
			Required:    true,
			Category:    "record",
			Order:       1,
		},
		"name": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set name",
			ValidateFunc: validateStrLen(1, 63),
			Category:     "record",
			Order:        10,
		},
		"type": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set record type[A/AAAA/ALIAS/NS/CNAME/MX/TXT/SRV/CAA/PTR]",
			ValidateFunc: validateInStrValues(allowDNSTypes...),
			Category:     "record",
			Order:        20,
		},
		"value": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set record data",
			Category:    "record",
			Order:       30,
		},
		"ttl": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set ttl",
			ValidateFunc: validateIntRange(10, 3600000),
			Category:     "record",
			Order:        40,
		},
		"mx-priority": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set MX priority",
			ValidateFunc: validateIntRange(1, 65535),
			Category:     "MX",
			Order:        10,
		},
		"srv-priority": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set SRV priority",
			ValidateFunc: validateIntRange(0, 65535),
			Category:     "SRV",
			Order:        10,
		},
		"srv-weight": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set SRV priority",
			ValidateFunc: validateIntRange(0, 65535),
			Category:     "SRV",
			Order:        20,
		},
		"srv-port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set SRV priority",
			ValidateFunc: validateIntRange(1, 65535),
			Category:     "SRV",
			Order:        30,
		},
		"srv-target": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set SRV priority",
			ValidateFunc: validateStrLen(1, 254),
			Category:     "SRV",
			Order:        40,
		},
	}
}
func dnsRecordDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target record",
			Required:    true,
			Category:    "record",
			Order:       10,
		},
	}
}

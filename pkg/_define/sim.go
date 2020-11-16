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

package _define

import (
	"math"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/schema"
)

func SIMResource() *schema.Resource {
	commands := map[string]*schema.Command{
		"list": {
			Type:    schema.CommandList,
			Aliases: []string{"ls", "find", "select"},
			Params:  simListParam(),
			// TableType:          output.TableSimple,
			TableColumnDefines: simListColumns(),
			UseCustomCommand:   true,
			Category:           "basic",
			Order:              10,
		},
		"create": {
			Type:             schema.CommandCreate,
			Params:           simCreateParam(),
			IncludeFields:    simDetailIncludes(),
			ExcludeFields:    simDetailExcludes(),
			Category:         "basic",
			Order:            20,
			UseCustomCommand: true,
		},
		"read": {
			Type:             schema.CommandRead,
			Params:           simReadParam(),
			IncludeFields:    simDetailIncludes(),
			ExcludeFields:    simDetailExcludes(),
			Category:         "basic",
			Order:            30,
			UseCustomCommand: true,
		},
		"update": {
			Type:          schema.CommandUpdate,
			Params:        simUpdateParam(),
			IncludeFields: simDetailIncludes(),
			ExcludeFields: simDetailExcludes(),
			Category:      "basic",
			Order:         40,
		},
		"delete": {
			Type:             schema.CommandDelete,
			Aliases:          []string{"rm"},
			Params:           simDeleteParam(),
			IncludeFields:    simDetailIncludes(),
			ExcludeFields:    simDetailExcludes(),
			Category:         "basic",
			Order:            50,
			UseCustomCommand: true,
			NoOutput:         true,
		},
		"carrier-info": {
			Type:    schema.CommandManipulateSingle,
			Params:  simCarrierInfoParam(),
			Aliases: []string{"carrier-list"},
			// TableType:          output.TableSimple,
			TableColumnDefines: simCarrierInfoColumns(),
			UseCustomCommand:   true,
			NeedlessConfirm:    true,
			Category:           "carrier",
			Order:              10,
		},
		"carrier-update": {
			Type:             schema.CommandManipulateMulti,
			Params:           simCarrierUpdateParam(),
			UseCustomCommand: true,
			Category:         "carrier",
			Order:            15,
			NoOutput:         true,
		},
		"activate": {
			Type:             schema.CommandManipulateMulti,
			Params:           simActivateParam(),
			UseCustomCommand: true,
			Category:         "activate",
			Order:            10,
			NoOutput:         true,
		},
		"deactivate": {
			Type:             schema.CommandManipulateMulti,
			Params:           simDeactivateParam(),
			UseCustomCommand: true,
			Category:         "activate",
			Order:            20,
			NoOutput:         true,
		},
		"ip-add": {
			Type:             schema.CommandManipulateSingle,
			Params:           simIPAddParam(),
			UseCustomCommand: true,
			Category:         "ip",
			Order:            10,
			NoOutput:         true,
		},
		"ip-delete": {
			Type:             schema.CommandManipulateSingle,
			Params:           simIPDeleteParam(),
			Aliases:          []string{"ip-del"},
			UseCustomCommand: true,
			Category:         "ip",
			Order:            20,
			NoOutput:         true,
		},
		"imei-lock": {
			Type:             schema.CommandManipulateMulti,
			Params:           simIMEILockParams(),
			UseCustomCommand: true,
			Category:         "imei",
			Order:            10,
			NoOutput:         true,
		},
		"imei-unlock": {
			Type:             schema.CommandManipulateMulti,
			Params:           simIMEIUnlockParams(),
			UseCustomCommand: true,
			Category:         "imei",
			Order:            20,
			NoOutput:         true,
		},
		"logs": {
			Type:   schema.CommandRead,
			Params: simLogsParam(),
			// TableType:          output.TableSimple,
			TableColumnDefines: simLogsColumns(),
			Order:              10,
			UseCustomCommand:   true,
			NeedlessConfirm:    true,
			Category:           "monitor",
		},
		"monitor": {
			Type:   schema.CommandRead,
			Params: simMonitorParam(),
			// TableType:          output.TableSimple,
			TableColumnDefines: simMonitorColumns(),
			Order:              20,
			UseCustomCommand:   true,
			Category:           "monitor",
		},
	}

	return &schema.Resource{
		Commands:            commands,
		ResourceCategory:    CategoryCommonServiceItem,
		CommandCategories:   SIMCommandCategories,
		ListResultFieldName: "SIMs",
		IsGlobal:            true,
	}
}

var SIMCommandCategories = []schema.Category{
	{
		Key:         "basic",
		DisplayName: "Basics",
		Order:       10,
	},
	{
		Key:         "carrier",
		DisplayName: "Carrier",
		Order:       20,
	},
	{
		Key:         "activate",
		DisplayName: "Activate/Deactivate",
		Order:       25,
	},
	{
		Key:         "imei",
		DisplayName: "IMEI lock/unlock",
		Order:       30,
	},
	{
		Key:         "ip",
		DisplayName: "IPAddress Management",
		Order:       30,
	},
	{
		Key:         "monitor",
		DisplayName: "Monitoring",
		Order:       40,
	},
	{
		Key:         "other",
		DisplayName: "Other",
		Order:       1000,
	},
}

func simListParam() map[string]*schema.Parameter {
	return mergeParameterMap(CommonListParam, paramTagsCond)
}

func simListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{
			Name:    "ICCID",
			Sources: []string{"Status.ICCID"},
		},
		{
			Name:    "IP",
			Sources: []string{"Status.sim.ip"},
		},
		{
			Name:    "Status",
			Sources: []string{"Status.sim.session_status"},
		},
		{
			Name:       "IMEI Lock",
			FormatFunc: formatBoolFunc("Status.sim.imei_lock"),
		},
		{
			Name:       "Registered",
			FormatFunc: formatBoolFunc("Status.sim.registered"),
		},
		{
			Name:       "Activated",
			FormatFunc: formatBoolFunc("Status.sim.activated"),
		},
	}
}

func simCarrierInfoColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "name"},
		{Name: "country_code"},
		{Name: "allow"},
	}
}

func simLogsColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "__ORDER__"}, // magic column name(generated on demand)
		{Name: "Date"},
		{Name: "SessionStatus"},
		{Name: "ResourceID"},
		{Name: "IMEI"},
		{Name: "IMSI"},
	}
}

func simMonitorColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Key"},
		{Name: "TimeStamp"},
		{Name: "UnixTime"},
		{Name: "UplinkBPS"},
		{Name: "DownlinkBPS"},
	}
}

func simDetailIncludes() []string {
	return []string{}
}

func simDetailExcludes() []string {
	return []string{}
}

func simCreateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"iccid": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Required:    true,
			Category:    "sim",
			Order:       10,
		},
		"passcode": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Required:    true,
			Category:    "sim",
			Order:       20,
		},
		"disabled": {
			Type:         schema.TypeBool,
			HandlerType:  schema.HandlerNoop,
			DefaultValue: false,
			Category:     "sim",
			Order:        30,
		},
		"imei": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Category:    "sim",
			Order:       40,
		},
		"carrier": {
			Type:         schema.TypeStringList,
			HandlerType:  schema.HandlerNoop,
			Category:     "sim",
			Required:     true,
			Order:        50,
			ValidateFunc: validateStringSlice(validateInStrValues(types.SIMOperatorShortNames()...)),
			MinItems:     1,
			MaxItems:     3,
		},
		"name": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set resource display name",
			Required:     true,
			ValidateFunc: validateStrLen(1, 64),
			Category:     "common",
			Order:        500,
		},
		"description": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set resource description",
			Aliases:      []string{"desc"},
			ValidateFunc: validateStrLen(0, 254),
			Category:     "common",
			Order:        510,
		},
		"tags": {
			Type:         schema.TypeStringList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set resource tags",
			ValidateFunc: validateStringSlice(validateStrLen(1, 32)),
			Category:     "common",
			Order:        520,
		},
		"icon-id": {
			Type:         schema.TypeId,
			HandlerType:  schema.HandlerNoop,
			Description:  "set Icon ID",
			ValidateFunc: validateSakuraID(),
			Category:     "common",
			Order:        530,
		},
	}
}

func simReadParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func simUpdateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func simDeleteParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"force": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"f"},
			Description: "forced-delete flag if SIM is still activating",
			Category:    "operation",
			Order:       10,
		},
	}
}

func simCarrierInfoParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func simCarrierUpdateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"carrier": {
			Type:         schema.TypeStringList,
			HandlerType:  schema.HandlerNoop,
			ValidateFunc: validateStringSlice(validateInStrValues(types.SIMOperatorShortNames()...)),
			MinItems:     1,
			MaxItems:     3,
			Required:     true,
			Order:        10,
		},
	}
}

func simActivateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func simDeactivateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func simIPAddParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"ip": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			ValidateFunc: validateIPv4Address(),
			Required:     true,
			Order:        10,
		},
	}
}

func simIPDeleteParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func simIMEILockParams() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"imei": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Required:    true,
			Order:       10,
		},
	}
}

func simIMEIUnlockParams() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func simLogsParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"follow": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "follow log output",
			Aliases:     []string{"f"},
			Category:    "monitor",
			Order:       10,
		},
		"refresh-interval": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			ValidateFunc: validateIntRange(1, math.MaxInt32),
			DefaultValue: int64(3),
			Description:  "log refresh interval second",
			Category:     "monitor",
			Order:        20,
		},
	}
}

func simMonitorParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
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
			DefaultValue: "sakuracloud.sim.{{.ID}}",
			Required:     true,
			Category:     "monitor",
			Order:        30,
		},
	}
}

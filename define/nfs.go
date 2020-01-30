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
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func NFSResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find", "selector"},
			Params:             nfsListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: nfsListColumns(),
			UseCustomCommand:   true,
			Category:           "basic",
			Order:              10,
		},
		"create": {
			Type:             schema.CommandCreate,
			Params:           nfsCreateParam(),
			ParamCategories:  nfsParamsCategories,
			IncludeFields:    nfsDetailIncludes(),
			ExcludeFields:    nfsDetailExcludes(),
			Category:         "basic",
			Order:            20,
			UseCustomCommand: true,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        nfsReadParam(),
			IncludeFields: nfsDetailIncludes(),
			ExcludeFields: nfsDetailExcludes(),
			Category:      "basic",
			Order:         30,
		},
		"update": {
			Type:          schema.CommandUpdate,
			Params:        nfsUpdateParam(),
			IncludeFields: nfsDetailIncludes(),
			ExcludeFields: nfsDetailExcludes(),
			Category:      "basic",
			Order:         40,
		},
		"delete": {
			Type:             schema.CommandDelete,
			Aliases:          []string{"rm"},
			Params:           nfsDeleteParam(),
			IncludeFields:    nfsDetailIncludes(),
			ExcludeFields:    nfsDetailExcludes(),
			Category:         "basic",
			Order:            50,
			UseCustomCommand: true,
		},
		"boot": {
			Type:             schema.CommandManipulateMulti,
			Aliases:          []string{"power-on"},
			Params:           nfsPowerOnParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            10,
			NoOutput:         true,
		},
		"shutdown": {
			Type:             schema.CommandManipulateMulti,
			Aliases:          []string{"power-off"},
			Params:           nfsPowerOffParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            20,
			NoOutput:         true,
		},
		"shutdown-force": {
			Type:             schema.CommandManipulateMulti,
			Aliases:          []string{"stop"},
			Params:           nfsPowerOffParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            25,
			NoOutput:         true,
		},
		"reset": {
			Type:             schema.CommandManipulateMulti,
			Params:           nfsResetParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            30,
			NoOutput:         true,
		},
		"wait-for-boot": {
			Type:             schema.CommandManipulateMulti,
			Params:           nfsWaitForParams(),
			Usage:            "Wait until boot is completed",
			UseCustomCommand: true,
			Category:         "power",
			Order:            40,
			NoOutput:         true,
			NeedlessConfirm:  true,
		},
		"wait-for-down": {
			Type:             schema.CommandManipulateMulti,
			Params:           nfsWaitForParams(),
			Usage:            "Wait until shutdown is completed",
			UseCustomCommand: true,
			Category:         "power",
			Order:            50,
			NoOutput:         true,
			NeedlessConfirm:  true,
		},
		"monitor-nic": {
			Type:               schema.CommandRead,
			Params:             nfsMonitorParam("nic"),
			Usage:              "Collect NIC(s) monitor values",
			TableType:          output.TableSimple,
			TableColumnDefines: nfsMonitorNICColumns(),
			UseCustomCommand:   true,
			Category:           "monitor",
			Order:              10,
		},
		"monitor-free-disk-size": {
			Type:               schema.CommandRead,
			Params:             nfsMonitorParam("free-disk-size"),
			Usage:              "Collect system-disk monitor values(IO)",
			TableType:          output.TableSimple,
			TableColumnDefines: nfsMonitorSizeColumns(),
			UseCustomCommand:   true,
			Category:           "monitor",
			Order:              20,
		},
	}

	return &schema.Resource{
		Commands:            commands,
		ListResultFieldName: "NFS",
		ResourceCategory:    CategoryAppliance,
		CommandCategories:   NFSCommandCategories,
	}
}

var NFSCommandCategories = []schema.Category{
	{
		Key:         "basic",
		DisplayName: "Basics",
		Order:       10,
	},
	{
		Key:         "power",
		DisplayName: "Power Management",
		Order:       20,
	},
	{
		Key:         "vip",
		DisplayName: "VirtualIPAddress Management",
		Order:       30,
	},
	{
		Key:         "monitor",
		DisplayName: "Monitoring",
		Order:       50,
	},
	{
		Key:         "other",
		DisplayName: "Other",
		Order:       1000,
	},
}

var nfsParamsCategories = []schema.Category{
	{
		Key:         "nfs",
		DisplayName: "NFS options",
		Order:       10,
	},
	{
		Key:         "network",
		DisplayName: "Network options",
		Order:       20,
	},
	{
		Key:         "common",
		DisplayName: "Common options",
		Order:       1000,
	},
}

func nfsListParam() map[string]*schema.Schema {
	return mergeParameterMap(CommonListParam, paramTagsCond)
}

func nfsListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{
			Name:    "Power",
			Sources: []string{"Instance.Status"},
		},
		{
			Name:    "Plan",
			Sources: []string{"PlanName"},
		},
		{
			Name:    "Size",
			Sources: []string{"Size"},
			Format:  "%sGB",
		},
		{
			Name: "IPAddress",
			FormatFunc: func(values map[string]string) string {
				if ip, ok := values["Remark.Servers.0.IPAddress"]; ok {
					format := "%s/%s"
					return fmt.Sprintf(format,
						ip,
						values["Remark.Network.NetworkMaskLen"],
					)
				}

				return ""
			},
		},
		{
			Name:    "DefaultRoute",
			Sources: []string{"Remark.Network.DefaultRoute"},
		},
	}
}

func nfsDetailIncludes() []string {
	return []string{}
}

func nfsDetailExcludes() []string {
	return []string{}
}

func nfsCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"switch-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Required:     true,
			Description:  "set connect switch ID",
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeSwitchID(),
			Category:     "nfs",
			Order:        10,
		},
		"plan": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Required:     true,
			DefaultValue: "hdd",
			Description:  "set plan[ssd/hdd]",
			ValidateFunc: validateInStrValues("ssd", "hdd"),
			CompleteFunc: completeInStrValues("ssd", "hdd"),
			Category:     "nfs",
			Order:        40,
		},
		"size": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Required:     true,
			DefaultValue: 100,
			Description:  "set plan[100/500/1024/2048/4096/8192/12288]",
			ValidateFunc: validateInIntValues(sacloud.AllowNFSNormalPlanSizes()...),
			CompleteFunc: completeInIntValues(sacloud.AllowNFSNormalPlanSizes()...),
			Category:     "nfs",
			Order:        45,
		},
		"ipaddress": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"ip"},
			Description:  "set ipaddress(#)",
			ValidateFunc: validateIPv4Address(),
			Required:     true,
			Category:     "network",
			Order:        10,
		},
		"nw-mask-len": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set network mask length",
			Required:     true,
			ValidateFunc: validateIntRange(8, 29),
			Category:     "network",
			Order:        30,
		},
		"default-route": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set default route",
			ValidateFunc: validateIPv4Address(),
			Category:     "network",
			Order:        40,
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
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set Icon ID",
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeIconID(),
			Category:     "common",
			Order:        530,
		},
	}
}

func nfsReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func nfsUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func nfsDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"force": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"f"},
			Description: "forced-shutdown flag if server is running",
			Category:    "nfs",
			Order:       10,
		},
	}
}

func nfsPowerOnParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func nfsPowerOffParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func nfsResetParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func nfsWaitForParams() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func nfsMonitorParam(key string) map[string]*schema.Schema {
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
			DefaultValue: fmt.Sprintf("sakuracloud.disk.{{.ID}}.%s", key),
			Required:     true,
			Category:     "monitor",
			Order:        30,
		},
	}
}

func nfsMonitorNICColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Key"},
		{Name: "TimeStamp"},
		{Name: "UnixTime"},
		{Name: "Receive"},
		{Name: "Send"},
	}
}

func nfsMonitorSizeColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Key"},
		{Name: "TimeStamp"},
		{Name: "UnixTime"},
		{Name: "Free"},
	}
}

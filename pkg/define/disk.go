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
	"strings"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/schema"
)

func DiskResource() *schema.Resource {
	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find", "select"},
			Params:             diskListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: diskListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"create": {
			Type:             schema.CommandCreate,
			Params:           diskCreateParam(),
			IncludeFields:    diskDetailIncludes(),
			ExcludeFields:    diskDetailExcludes(),
			UseCustomCommand: true,
			Category:         "basics",
			Order:            20,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        diskReadParam(),
			IncludeFields: diskDetailIncludes(),
			ExcludeFields: diskDetailExcludes(),
			Category:      "basics",
			Order:         30,
		},
		"update": {
			Type:             schema.CommandUpdate,
			Params:           diskUpdateParam(),
			IncludeFields:    diskDetailIncludes(),
			ExcludeFields:    diskDetailExcludes(),
			UseCustomCommand: true,
			Category:         "basics",
			Order:            40,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"rm"},
			Params:        diskDeleteParam(),
			IncludeFields: diskDetailIncludes(),
			ExcludeFields: diskDetailExcludes(),
			Category:      "basics",
			Order:         50,
		},
		"edit": {
			Type:             schema.CommandManipulateMulti,
			Aliases:          []string{"config"},
			Params:           diskConfigParam(),
			IncludeFields:    diskDetailIncludes(),
			ExcludeFields:    diskDetailExcludes(),
			UseCustomCommand: true,
			Category:         "edit",
			Order:            10,
		},
		"resize-partition": {
			Type:             schema.CommandManipulateMulti,
			Params:           diskResizePartitionParam(),
			IncludeFields:    diskDetailIncludes(),
			ExcludeFields:    diskDetailExcludes(),
			UseCustomCommand: true,
			Category:         "edit",
			Order:            20,
		},
		"reinstall-from-archive": {
			Type:             schema.CommandManipulateMulti,
			Params:           diskReinstallFromArchiveParam(),
			IncludeFields:    diskDetailIncludes(),
			ExcludeFields:    diskDetailExcludes(),
			UseCustomCommand: true,
			ConfirmMessage:   "re-install from archive",
			NoOutput:         true,
			Category:         "re-install",
			Order:            20,
		},
		"reinstall-from-disk": {
			Type:             schema.CommandManipulateMulti,
			Params:           diskReinstallFromDiskParam(),
			IncludeFields:    diskDetailIncludes(),
			ExcludeFields:    diskDetailExcludes(),
			UseCustomCommand: true,
			ConfirmMessage:   "re-install from disk",
			NoOutput:         true,
			Category:         "re-install",
			Order:            30,
		},
		"reinstall-to-blank": {
			Type:             schema.CommandManipulateMulti,
			Params:           diskReinstallToBlankParam(),
			IncludeFields:    diskDetailIncludes(),
			ExcludeFields:    diskDetailExcludes(),
			UseCustomCommand: true,
			ConfirmMessage:   "re-install to blank",
			NoOutput:         true,
			Category:         "re-install",
			Order:            40,
		},
		"server-connect": {
			Type:             schema.CommandManipulateMulti,
			Params:           diskServerConnectParam(),
			IncludeFields:    diskDetailIncludes(),
			ExcludeFields:    diskDetailExcludes(),
			UseCustomCommand: true,
			NoOutput:         true,
			Category:         "server",
			Order:            10,
		},
		"server-disconnect": {
			Type:             schema.CommandManipulateMulti,
			Params:           diskServerDisconnectParam(),
			IncludeFields:    diskDetailIncludes(),
			ExcludeFields:    diskDetailExcludes(),
			UseCustomCommand: true,
			NoOutput:         true,
			Category:         "server",
			Order:            20,
		},
		"monitor": {
			Type:               schema.CommandRead,
			Params:             diskMonitorParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: diskMonitorColumns(),
			UseCustomCommand:   true,
			Category:           "monitor",
			Order:              10,
		},
		"wait-for-copy": {
			Type:             schema.CommandManipulateMulti,
			Aliases:          []string{"wait"},
			Params:           diskWaitForCopyParam(),
			IncludeFields:    diskDetailIncludes(),
			ExcludeFields:    diskDetailExcludes(),
			UseCustomCommand: true,
			NoOutput:         true,
			NeedlessConfirm:  true,
			Category:         "other",
			Order:            10,
		},
	}

	return &schema.Resource{
		Commands:          commands,
		ResourceCategory:  CategoryStorage,
		CommandCategories: DiskCommandCategories,
	}
}

var DiskCommandCategories = []schema.Category{
	{
		Key:         "basics",
		DisplayName: "Basics",
		Order:       10,
	},
	{
		Key:         "edit",
		DisplayName: "Disk Edit",
		Order:       20,
	},
	{
		Key:         "re-install",
		DisplayName: "Re-Install",
		Order:       25,
	},
	{
		Key:         "server",
		DisplayName: "Server Connection Management",
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

func diskListParam() map[string]*schema.Schema {
	return mergeParameterMap(
		CommonListParam,
		paramScopeCond,
		paramTagsCond,
		paramSourceArchiveIDCond,
		paramSourceDiskCond,
		map[string]*schema.Schema{
			"storage": {
				Type:        schema.TypeString,
				HandlerType: schema.HandlerFilterFunc,
				FilterFunc: func(_ []interface{}, item interface{}, param interface{}) bool {
					if param == nil {
						return true
					}
					storageName := param.(string)
					if storageName == "" {
						return true
					}
					if disk, ok := item.(*sacloud.Disk); ok {
						return strings.Contains(disk.Storage.Name, storageName)
					}
					return true
				},
				Description: "set filter by storage-name",
				Category:    "filter",
				Order:       10,
			},
		},
	)
}

func diskListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{
			Name:    "Server",
			Sources: []string{"Server.ID", "Server.Name"},
			Format:  "%s(%s)",
		},
		{
			Name:    "Plan",
			Sources: []string{"Plan.ID"},
			ValueMapping: []map[string]string{
				{
					"4": "ssd",
					"2": "hdd",
				},
			},
		},
		{
			Name:    "Size",
			Sources: []string{"SizeMB"},
			Format:  "%sMB",
		},
		{Name: "Connection"},
		{
			Name:    "Storage",
			Sources: []string{"Storage.Name"},
		},
	}
}

func diskDetailIncludes() []string {
	return []string{}
}

func diskDetailExcludes() []string {
	return []string{
		"SourceArchive.Storage.",
		"Storage.",
	}
}

func diskCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramRequiredName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
		"plan": {
			Type:            schema.TypeString,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetDiskPlan",
			Required:        true,
			DefaultValue:    "ssd",
			Description:     "set disk plan('hdd' or 'ssd')",
			ValidateFunc:    validateInStrValues(types.DiskPlanStrings...),
			Category:        "disk",
			Order:           10,
		},
		"connection": {
			Type:            schema.TypeString,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetDiskConnectionByStr",
			Required:        true,
			DefaultValue:    "virtio",
			Description:     "set disk connection('virtio' or 'ide')",
			ValidateFunc:    validateInStrValues(types.DiskConnectionStrings...),
			Category:        "disk",
			Order:           20,
		},
		"source-archive-id": {
			Type:            schema.TypeId,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetSourceArchive",
			Description:     "set source disk ID",
			ValidateFunc:    validateSakuraID(),
			ConflictsWith:   []string{"source-disk-id"},
			Category:        "disk",
			Order:           30,
		},
		"source-disk-id": {
			Type:            schema.TypeId,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetSourceDisk",
			Description:     "set source disk ID",
			ValidateFunc:    validateSakuraID(),
			ConflictsWith:   []string{"source-archive-id"},
			Category:        "disk",
			Order:           40,
		},
		"size": {
			Type:            schema.TypeInt,
			HandlerType:     schema.HandlerPathThrough,
			Description:     "set disk size(GB)",
			DestinationProp: "SetSizeGB",
			DefaultValue:    20,
			Required:        true,
			Category:        "disk",
			Order:           50,
		},
		"distant-from": {
			Type:         schema.TypeIdList,
			HandlerType:  schema.HandlerPathThrough,
			Description:  "set distant from disk IDs",
			ValidateFunc: validateIntSlice(validateSakuraID()),
			Category:     "disk",
			Order:        60,
		},
	}
}

func diskReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func diskUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
		"connection": {
			Type:            schema.TypeString,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetDiskConnectionByStr",
			Description:     "set disk connection('virtio' or 'ide')",
			Category:        "disk",
			Order:           20,
		},
	}
}

func diskDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func diskConfigParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"hostname": {
			Type:            schema.TypeString,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetHostName",
			Description:     "set hostname",
			Category:        "edit",
			Order:           10,
		},
		"password": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerPathThrough,
			Description: "set password",
			Category:    "edit",
			Order:       20,
		},
		"ssh-key-ids": {
			Type:            schema.TypeIdList,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetSSHKeys",
			Description:     "set ssh-key ID(s)",
			ValidateFunc:    validateIntSlice(validateSakuraID()),
			Category:        "edit",
			Order:           30,
		},
		"disable-password-auth": {
			Type:            schema.TypeBool,
			Aliases:         []string{"disable-pw-auth"},
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetDisablePWAuth",
			Description:     "disable password auth on SSH",
			Category:        "edit",
			Order:           35,
		},
		"ipaddress": {
			Type:            schema.TypeString,
			Aliases:         []string{"ip"},
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetUserIPAddress",
			Description:     "set ipaddress",
			Category:        "edit",
			Order:           40,
		},
		"default-route": {
			Type:            schema.TypeString,
			Aliases:         []string{"gateway"},
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetDefaultRoute",
			Description:     "set default gateway",
			Category:        "edit",
			Order:           41,
		},
		"nw-masklen": {
			Type:            schema.TypeInt,
			HandlerType:     schema.HandlerPathThrough,
			Aliases:         []string{"network-masklen"},
			Description:     "set ipaddress  prefix",
			DestinationProp: "SetNetworkMaskLen",
			DefaultValue:    24,
			ValidateFunc:    validateIntRange(8, 29),
			Category:        "edit",
			Order:           42,
		},
		"startup-script-ids": {
			Type:            schema.TypeIdList,
			Aliases:         []string{"note-ids"},
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetNotes",
			Description:     "set startup-script ID(s)",
			ValidateFunc:    validateIntSlice(validateSakuraID()),
			Category:        "edit",
			Order:           50,
		},
	}
}

func diskResizePartitionParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func diskWaitForCopyParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func diskReinstallFromArchiveParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"source-archive-id": {
			Type:         schema.TypeId,
			HandlerType:  schema.HandlerNoop,
			Description:  "set source archive ID",
			Required:     true,
			ValidateFunc: validateSakuraID(),
			Category:     "install",
			Order:        10,
		},
		"distant-from": {
			Type:         schema.TypeIdList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set distant from disk IDs",
			ValidateFunc: validateIntSlice(validateSakuraID()),
			Category:     "install",
			Order:        20,
		},
	}
}

func diskReinstallFromDiskParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"source-disk-id": {
			Type:         schema.TypeId,
			HandlerType:  schema.HandlerNoop,
			Description:  "set source disk ID",
			Required:     true,
			ValidateFunc: validateSakuraID(),
			Category:     "install",
			Order:        10,
		},
		"distant-from": {
			Type:         schema.TypeIdList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set distant from disk IDs",
			ValidateFunc: validateIntSlice(validateSakuraID()),
			Category:     "install",
			Order:        20,
		},
	}
}

func diskReinstallToBlankParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"distant-from": {
			Type:         schema.TypeIdList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set distant from disk IDs",
			ValidateFunc: validateIntSlice(validateSakuraID()),
			Category:     "install",
			Order:        10,
		},
	}
}

func diskServerConnectParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"server-id": {
			Type:         schema.TypeId,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target server ID",
			Required:     true,
			ValidateFunc: validateSakuraID(),
			Category:     "connect",
			Order:        10,
		},
	}
}

func diskServerDisconnectParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func diskMonitorParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"start": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set start-time",
			ValidateFunc: validateDateTimeString(),
		},
		"end": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set end-time",
			ValidateFunc: validateDateTimeString(),
		},
		// TODO キーフォーマットはv2でサポートすべきか検討
		"key-format": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set monitoring value key-format",
			DefaultValue: "sakuracloud.disk.{{.ID}}.disk",
			Required:     true,
		},
	}
}

func diskMonitorColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Key"},
		{Name: "TimeStamp"},
		{Name: "UnixTime"},
		{Name: "Read"},
		{Name: "Write"},
	}
}

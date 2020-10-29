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
	"math"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/schema"
)

func DatabaseResource() *schema.Resource {
	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find", "select"},
			Params:             databaseListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: databaseListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"create": {
			Type:             schema.CommandCreate,
			Params:           databaseCreateParam(),
			ParamCategories:  databaseParamsCategories,
			IncludeFields:    databaseDetailIncludes(),
			ExcludeFields:    databaseDetailExcludes(),
			Category:         "basics",
			Order:            20,
			UseCustomCommand: true,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        databaseReadParam(),
			IncludeFields: databaseDetailIncludes(),
			ExcludeFields: databaseDetailExcludes(),
			Category:      "basics",
			Order:         30,
		},
		"update": {
			Type:             schema.CommandUpdate,
			Params:           databaseUpdateParam(),
			IncludeFields:    databaseDetailIncludes(),
			ExcludeFields:    databaseDetailExcludes(),
			Category:         "basics",
			Order:            40,
			UseCustomCommand: true,
		},
		"delete": {
			Type:             schema.CommandDelete,
			Aliases:          []string{"rm"},
			Params:           databaseDeleteParam(),
			IncludeFields:    databaseDetailIncludes(),
			ExcludeFields:    databaseDetailExcludes(),
			Category:         "basics",
			Order:            50,
			UseCustomCommand: true,
		},
		"boot": {
			Type:             schema.CommandManipulateMulti,
			Aliases:          []string{"power-on"},
			Params:           databasePowerOnParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            10,
			NoOutput:         true,
		},
		"shutdown": {
			Type:             schema.CommandManipulateMulti,
			Aliases:          []string{"power-off"},
			Params:           databasePowerOffParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            20,
			NoOutput:         true,
		},
		"shutdown-force": {
			Type:             schema.CommandManipulateMulti,
			Aliases:          []string{"stop"},
			Params:           databasePowerOffParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            25,
			NoOutput:         true,
		},
		"reset": {
			Type:             schema.CommandManipulateMulti,
			Params:           databaseResetParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            30,
			NoOutput:         true,
		},
		"wait-for-boot": {
			Type:             schema.CommandManipulateMulti,
			Params:           databaseWaitForParam(),
			Usage:            "Wait until boot is completed",
			UseCustomCommand: true,
			Category:         "power",
			Order:            40,
			NoOutput:         true,
			NeedlessConfirm:  true,
		},
		"wait-for-down": {
			Type:             schema.CommandManipulateMulti,
			Params:           databaseWaitForParam(),
			Usage:            "Wait until shutdown is completed",
			UseCustomCommand: true,
			Category:         "power",
			Order:            50,
			NoOutput:         true,
			NeedlessConfirm:  true,
		},
		"backup-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             databaseBackupListParam(),
			Aliases:            []string{"backups", "backup-list"},
			Usage:              "Show information of backup",
			TableType:          output.TableSimple,
			TableColumnDefines: databaseBackupListColumns(),
			UseCustomCommand:   true,
			Category:           "backup",
			Order:              10,
			NeedlessConfirm:    true,
		},
		"backup-create": {
			Type:               schema.CommandManipulateSingle,
			Params:             databaseBackupCreateParam(),
			Usage:              "Make new database backup",
			TableType:          output.TableSimple,
			TableColumnDefines: databaseBackupListColumns(),
			UseCustomCommand:   true,
			NoSelector:         true,
			Category:           "backup",
			Order:              20,
		},
		"backup-restore": {
			Type:               schema.CommandManipulateSingle,
			Params:             databaseBackupManipulateParam(),
			Usage:              "Restore database from backup",
			TableType:          output.TableSimple,
			TableColumnDefines: databaseBackupListColumns(),
			UseCustomCommand:   true,
			NoSelector:         true,
			Category:           "backup",
			Order:              30,
		},
		"backup-lock": {
			Type:               schema.CommandManipulateSingle,
			Params:             databaseBackupLockParam(),
			Usage:              "Lock backup",
			TableType:          output.TableSimple,
			TableColumnDefines: databaseBackupListColumns(),
			UseCustomCommand:   true,
			NoSelector:         true,
			Category:           "backup",
			Order:              40,
		},
		"backup-unlock": {
			Type:               schema.CommandManipulateSingle,
			Params:             databaseBackupLockParam(),
			Usage:              "Unlock backup",
			TableType:          output.TableSimple,
			TableColumnDefines: databaseBackupListColumns(),
			UseCustomCommand:   true,
			NoSelector:         true,
			Category:           "backup",
			Order:              50,
		},
		"backup-remove": {
			Type:               schema.CommandManipulateSingle,
			Params:             databaseBackupManipulateParam(),
			Usage:              "Remove backup",
			TableType:          output.TableSimple,
			TableColumnDefines: databaseBackupListColumns(),
			UseCustomCommand:   true,
			NoSelector:         true,
			Category:           "backup",
			Order:              60,
		},
		"clone": {
			Type:             schema.CommandManipulateSingle,
			Params:           databaseCloneParam(),
			Usage:            "Create clone instance",
			TableType:        output.TableDetail,
			UseCustomCommand: true,
			NoSelector:       true,
			Category:         "clone",
			Order:            10,
		},
		"replica-create": {
			Type:             schema.CommandManipulateSingle,
			Params:           databaseReplicaCreateParam(),
			Usage:            "Create replication slave instance",
			TableType:        output.TableDetail,
			UseCustomCommand: true,
			NoSelector:       true,
			Category:         "replication",
			Order:            10,
		},
		"monitor-cpu": {
			Type:               schema.CommandRead,
			Params:             databaseMonitorParam("cpu"),
			Usage:              "Collect CPU monitor values",
			TableType:          output.TableSimple,
			TableColumnDefines: databaseMonitorCPUColumns(),
			UseCustomCommand:   true,
			Category:           "monitor",
			Order:              10,
		},
		"monitor-memory": {
			Type:               schema.CommandRead,
			Params:             databaseMonitorParam("memory"),
			Usage:              "Collect memory monitor values",
			TableType:          output.TableSimple,
			TableColumnDefines: databaseMonitorSizeColumns(),
			UseCustomCommand:   true,
			Category:           "monitor",
			Order:              20,
		},
		"monitor-nic": {
			Type:               schema.CommandRead,
			Params:             databaseMonitorParam("nic"),
			Usage:              "Collect NIC(s) monitor values",
			TableType:          output.TableSimple,
			TableColumnDefines: databaseMonitorNICColumns(),
			UseCustomCommand:   true,
			Category:           "monitor",
			Order:              30,
		},
		"monitor-system-disk": {
			Type:               schema.CommandRead,
			Params:             databaseMonitorParam("disk1"),
			Usage:              "Collect system-disk monitor values(IO)",
			TableType:          output.TableSimple,
			TableColumnDefines: databaseMonitorDiskColumns(),
			UseCustomCommand:   true,
			Category:           "monitor",
			Order:              40,
		},
		"monitor-backup-disk": {
			Type:               schema.CommandRead,
			Params:             databaseMonitorParam("disk2"),
			Usage:              "Collect backup-disk monitor values(IO)",
			TableType:          output.TableSimple,
			TableColumnDefines: databaseMonitorDiskColumns(),
			UseCustomCommand:   true,
			Category:           "monitor",
			Order:              50,
		},

		"monitor-system-disk-size": {
			Type:               schema.CommandRead,
			Params:             databaseMonitorParam("disk1"),
			Usage:              "Collect system-disk monitor values(usage)",
			TableType:          output.TableSimple,
			TableColumnDefines: databaseMonitorSizeColumns(),
			UseCustomCommand:   true,
			Category:           "monitor",
			Order:              60,
		},
		"monitor-backup-disk-size": {
			Type:               schema.CommandRead,
			Params:             databaseMonitorParam("disk2"),
			Usage:              "Collect backup-disk monitor values(usage)",
			TableType:          output.TableSimple,
			TableColumnDefines: databaseMonitorSizeColumns(),
			UseCustomCommand:   true,
			Category:           "monitor",
			Order:              70,
		},
		"logs": {
			Type:             schema.CommandRead,
			Params:           databaseLogParam(),
			UseCustomCommand: true,
			Order:            100,
			Category:         "monitor",
			NoOutput:         true,
		},
	}

	return &schema.Resource{
		Commands:          commands,
		ResourceCategory:  CategoryAppliance,
		CommandCategories: DatabaseCommandCategories,
	}
}

var DatabaseCommandCategories = []schema.Category{
	{
		Key:         "basics",
		DisplayName: "Basics",
		Order:       10,
	},
	{
		Key:         "power",
		DisplayName: "Power Management",
		Order:       20,
	},
	{
		Key:         "backup",
		DisplayName: "Backup Management",
		Order:       30,
	},
	{
		Key:         "clone",
		DisplayName: "Clone Instance Management",
		Order:       40,
	},
	{
		Key:         "replication",
		DisplayName: "Replica Instance Management",
		Order:       45,
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

var databaseParamsCategories = []schema.Category{
	{
		Key:         "database",
		DisplayName: "Database options",
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
		Order:       100,
	},
}

func databaseListParam() map[string]*schema.Schema {
	return mergeParameterMap(CommonListParam, paramTagsCond)
}

func databaseListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{
			Name:    "Power",
			Sources: []string{"Instance.Status"},
		},
		{
			Name:    "Database",
			Sources: []string{"Remark.DBConf.Common.DatabaseTitle"},
		},
		{
			Name:    "Plan",
			Sources: []string{"Plan.ID"},
			ValueMapping: []map[string]string{
				{
					"10":  "10GB",
					"30":  "30GB",
					"90":  "90GB",
					"240": "240GB",
				},
			},
		},
		{
			Name:    "DefaultUser",
			Sources: []string{"Settings.DBConf.Common.DefaultUser"},
		},
		{
			Name:    "Port",
			Sources: []string{"Settings.DBConf.Common.ServicePort"},
		},
		{
			Name:    "IPAddress",
			Sources: []string{"Remark.Servers.0.IPAddress", "Remark.Network.NetworkMaskLen"},
			Format:  "%s/%s",
		},
		{
			Name:    "DefaultRoute",
			Sources: []string{"Remark.Network.DefaultRoute"},
		},
	}
}

func databaseBackupListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "__ORDER__"}, // magic column name(generated on demand)
		{Name: "createdat"},
		{Name: "recoveredat"},
		{
			Name:    "Size(MB)",
			Sources: []string{"SizeMB"},
		},
		{
			Name:    "Locked",
			Sources: []string{"availability"},
			ValueMapping: []map[string]string{
				{
					"available":    "true",
					"discontinued": "false",
				},
			},
		},
	}
}

func databaseDetailIncludes() []string {
	return []string{}
}

func databaseDetailExcludes() []string {
	return []string{}
}

func databaseCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"switch-id": {
			Type:         schema.TypeId,
			HandlerType:  schema.HandlerNoop,
			Description:  "set connect switch ID",
			ValidateFunc: validateSakuraID(),
			Required:     true,
			Category:     "database",
			Order:        10,
		},
		"plan": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Required:     true,
			DefaultValue: 10,
			Description:  "set plan[10/30/90/240/500/1000]",
			ValidateFunc: validateInIDs(types.DatabasePlanIDs...),
			Category:     "database",
			Order:        20,
		},
		"database": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"db"},
			Description:  "set database type[postgresql/mariadb]",
			Required:     true,
			ValidateFunc: validateInStrValues("postgresql", "mariadb"),
			Category:     "database",
			Order:        30,
		},
		"username": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set database default user name",
			Required:     true,
			ValidateFunc: validateStrLen(4, 20),
			Category:     "database",
			Order:        40,
		},
		"password": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set database default user password",
			Required:     true,
			ValidateFunc: validateStrLen(8, 30),
			Category:     "database",
			Order:        50,
		},
		"replica-user-password": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set database replica user password",
			ValidateFunc: validateStrLen(8, 30),
			Category:     "database",
			Order:        55,
		},
		"source-networks": {
			Type:         schema.TypeStringList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set network of allow connection",
			ValidateFunc: validateStringSlice(validateIPv4AddressWithPrefixOption()),
			Category:     "database",
			Order:        60,
		},
		"enable-web-ui": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "enable web-ui",
			Category:    "database",
			Order:       70,
		},
		"enable-backup": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "enable backup",
			Category:    "database",
			Order:       80,
		},
		"backup-weekdays": {
			Type:        schema.TypeStringList,
			HandlerType: schema.HandlerNoop,
			Description: "set backup target weekdays[all or mon/tue/wed/thu/fri/sat/sun]",
			ValidateFunc: validateStringSlice(
				validateInStrValues(append(types.BackupWeekdayStrings, "all")...),
			),
			DefaultValue: []string{"all"},
			Category:     "database",
			Order:        85,
		},
		"backup-time": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set backup start time",
			ValidateFunc: validateBackupTime(),
			Category:     "database",
			Order:        88,
		},
		"port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set database port",
			DefaultText:  "PostgreSQL:5432, MariaDB:3306",
			ValidateFunc: validateIntRange(1024, 65535),
			Category:     "network",
			Order:        10,
		},
		"ipaddress1": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"ip1", "ipaddress", "ip"},
			Description:  "set ipaddress(#1)",
			ValidateFunc: validateIPv4Address(),
			Required:     true,
			Category:     "network",
			Order:        20,
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
			Required:     true,
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
			Type:         schema.TypeId,
			HandlerType:  schema.HandlerNoop,
			Description:  "set Icon ID",
			ValidateFunc: validateSakuraID(),
			Category:     "common",
			Order:        530,
		},
	}
}

func databaseReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func databaseUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"password": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set database default user password",
			ValidateFunc: validateStrLen(8, 30),
			Category:     "database",
			Order:        50,
		},
		"replica-user-password": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set database replica user password",
			ValidateFunc: validateStrLen(8, 30),
			Category:     "database",
			Order:        55,
		},
		"enable-replication": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "enable replication",
			Category:    "database",
			Order:       58,
		},
		"port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set database port",
			DefaultText:  "PostgreSQL:5432, MariaDB:3306",
			ValidateFunc: validateIntRange(1024, 65535),
			Category:     "database",
			Order:        60,
		},
		"source-networks": {
			Type:         schema.TypeStringList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set network of allow connection",
			ValidateFunc: validateStringSlice(validateIPv4AddressWithPrefixOption()),
			Category:     "database",
			Order:        100,
		},
		"enable-web-ui": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "enable web-ui",
			Category:    "database",
			Order:       110,
		},
		"enable-backup": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "enable backup",
			Category:    "database",
			Order:       120,
		},
		"backup-weekdays": {
			Type:        schema.TypeStringList,
			HandlerType: schema.HandlerNoop,
			Description: "set backup target weekdays[all or mon/tue/wed/thu/fri/sat/sun]",
			ValidateFunc: validateStringSlice(
				validateInStrValues(append(types.BackupWeekdayStrings, "all")...),
			),
			DefaultValue: []string{"all"},
			Category:     "database",
			Order:        125,
		},
		"backup-time": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set backup start time",
			ValidateFunc: validateBackupTime(),
			Category:     "database",
			Order:        128,
		},
		"name": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set resource display name",
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

func databaseDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"force": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"f"},
			Description: "forced-shutdown flag if database is running",
		},
	}
}

func databasePowerOnParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func databasePowerOffParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func databaseResetParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func databaseWaitForParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func databaseBackupListParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func databaseBackupCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func databaseBackupLockParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "index of target backup",
			Required:     true,
			ValidateFunc: validateIntRange(1, 8),
			Category:     "backup",
			Order:        1,
		},
	}
}

func databaseBackupManipulateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "index of target backup",
			Required:     true,
			ValidateFunc: validateIntRange(1, 16),
			Category:     "backup",
			Order:        1,
		},
	}
}

func databaseCloneParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"switch-id": {
			Type:         schema.TypeId,
			HandlerType:  schema.HandlerNoop,
			Description:  "set connect switch ID",
			ValidateFunc: validateSakuraID(),
			Category:     "database",
			Order:        10,
		},
		"plan": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Required:     true,
			DefaultValue: 10,
			Description:  "set plan[10/30/90/240/500/1000]",
			ValidateFunc: validateInIDs(types.DatabasePlanIDs...),
			Category:     "database",
			Order:        20,
		},
		"replica-user-password": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set database replica user password",
			ValidateFunc: validateStrLen(8, 30),
			Category:     "database",
			Order:        55,
		},
		"source-networks": {
			Type:         schema.TypeStringList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set network of allow connection",
			ValidateFunc: validateStringSlice(validateIPv4AddressWithPrefixOption()),
			Category:     "database",
			Order:        60,
		},
		"enable-web-ui": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "enable web-ui",
			Category:    "database",
			Order:       70,
		},
		"enable-backup": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "enable backup",
			Category:    "database",
			Order:       80,
		},
		"backup-weekdays": {
			Type:        schema.TypeStringList,
			HandlerType: schema.HandlerNoop,
			Description: "set backup target weekdays[all or mon/tue/wed/thu/fri/sat/sun]",
			ValidateFunc: validateStringSlice(
				validateInStrValues(append(types.BackupWeekdayStrings, "all")...),
			),
			DefaultValue: []string{"all"},
			Category:     "database",
			Order:        85,
		},
		"backup-time": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set backup start time",
			ValidateFunc: validateBackupTime(),
			Category:     "database",
			Order:        88,
		},
		"port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set database port",
			DefaultText:  "PostgreSQL:5432, MariaDB:3306",
			ValidateFunc: validateIntRange(1024, 65535),
			Category:     "network",
			Order:        10,
		},
		"ipaddress1": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"ip1", "ipaddress", "ip"},
			Description:  "set ipaddress(#1)",
			ValidateFunc: validateIPv4Address(),
			Required:     true,
			Category:     "network",
			Order:        20,
		},
		"nw-mask-len": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set network mask length",
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
			Type:         schema.TypeId,
			HandlerType:  schema.HandlerNoop,
			Description:  "set Icon ID",
			ValidateFunc: validateSakuraID(),
			Category:     "common",
			Order:        530,
		},
	}
}

func databaseReplicaCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"switch-id": {
			Type:         schema.TypeId,
			HandlerType:  schema.HandlerNoop,
			Description:  "set connect switch ID",
			ValidateFunc: validateSakuraID(),
			Category:     "database",
			Order:        10,
		},
		"ipaddress1": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"ip1", "ipaddress", "ip"},
			Description:  "set ipaddress(#1)",
			ValidateFunc: validateIPv4Address(),
			Required:     true,
			Category:     "network",
			Order:        20,
		},
		"nw-mask-len": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set network mask length",
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
			Type:         schema.TypeId,
			HandlerType:  schema.HandlerNoop,
			Description:  "set Icon ID",
			ValidateFunc: validateSakuraID(),
			Category:     "common",
			Order:        530,
		},
	}
}

func databaseMonitorParam(key string) map[string]*schema.Schema {
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
			DefaultValue: fmt.Sprintf("sakuracloud.database.{{.ID}}.%s", key),
			Required:     true,
			Category:     "monitor",
			Order:        30,
		},
	}
}

func databaseMonitorCPUColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Key"},
		{Name: "TimeStamp"},
		{Name: "UnixTime"},
		{Name: "CPUTime"},
	}
}

func databaseMonitorNICColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Key"},
		{Name: "TimeStamp"},
		{Name: "UnixTime"},
		{Name: "Receive"},
		{Name: "Send"},
	}
}

func databaseMonitorDiskColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Key"},
		{Name: "TimeStamp"},
		{Name: "UnixTime"},
		{Name: "Read"},
		{Name: "Write"},
	}
}

func databaseMonitorSizeColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Key"},
		{Name: "TimeStamp"},
		{Name: "UnixTime"},
		{Name: "Used"},
		{Name: "Total"},
	}
}

func databaseLogParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"log-name": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"name"},
			Description:  "set target logfile name",
			DefaultValue: "all",
			Category:     "monitor",
			Order:        10,
		},
		"follow": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "follow log output",
			Aliases:     []string{"f"},
			Category:    "monitor",
			Order:       20,
		},
		"refresh-interval": {
			Type:         schema.TypeId,
			HandlerType:  schema.HandlerNoop,
			ValidateFunc: validateIntRange(1, math.MaxInt32),
			DefaultValue: int64(3),
			Description:  "log refresh interval second",
			Category:     "monitor",
			Order:        30,
		},
		"list-log-names": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "show log-name list",
			Category:    "monitor",
			Order:       40,
		},
	}
}

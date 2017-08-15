package define

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
	"math"
)

func DatabaseResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find", "selector"},
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
			Usage:              "Collect Disk(s) monitor values",
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
			Usage:              "Collect Disk(s) monitor values",
			TableType:          output.TableSimple,
			TableColumnDefines: databaseMonitorDiskColumns(),
			UseCustomCommand:   true,
			Category:           "monitor",
			Order:              40,
		},
		"monitor-backup-disk": {
			Type:               schema.CommandRead,
			Params:             databaseMonitorParam("disk2"),
			Usage:              "Collect Disk(s) monitor values",
			TableType:          output.TableSimple,
			TableColumnDefines: databaseMonitorDiskColumns(),
			UseCustomCommand:   true,
			Category:           "monitor",
			Order:              50,
		},

		"monitor-system-disk-size": {
			Type:               schema.CommandRead,
			Params:             databaseMonitorParam("disk1"),
			Usage:              "Collect Disk(s) monitor values",
			TableType:          output.TableSimple,
			TableColumnDefines: databaseMonitorSizeColumns(),
			UseCustomCommand:   true,
			Category:           "monitor",
			Order:              60,
		},
		"monitor-backup-disk-size": {
			Type:               schema.CommandRead,
			Params:             databaseMonitorParam("disk2"),
			Usage:              "Collect Disk(s) monitor values",
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
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set connect switch ID",
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeSwitchID(),
			Required:     true,
			Category:     "database",
			Order:        10,
		},
		"plan": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Required:     true,
			DefaultValue: 10,
			Description:  "set plan[10/30/90/240]",
			ValidateFunc: validateInIntValues(sacloud.AllowDatabasePlans()...),
			CompleteFunc: completeInIntValues(sacloud.AllowDatabasePlans()...),
			Category:     "database",
			Order:        20,
		},
		"database": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"db"},
			Description:  "set database type",
			Required:     true,
			ValidateFunc: validateInStrValues("postgresql", "mariadb"),
			CompleteFunc: completeInStrValues("postgresql", "mariadb"),
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
		"backup-time": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set backup start time",
			ValidateFunc: validateBackupTime(),
			CompleteFunc: completeBackupTime(),
			Category:     "database",
			Order:        80,
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
		"backup-time": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set backup start time",
			ValidateFunc: validateBackupTime(),
			CompleteFunc: completeBackupTime(),
			Category:     "database",
			Order:        120,
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

var databaseLogNameCompletions = []string{
	"all",
	"systemctl",
	"mariadb.log",
}

func databaseLogParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"log-name": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"name"},
			Description:  "set target logfile name",
			DefaultValue: "all",
			CompleteFunc: completeInStrValues(databaseLogNameCompletions...),
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
			Type:         schema.TypeInt64,
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

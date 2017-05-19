package define

import (
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func DatabaseResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find"},
			Params:             databaseListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: databaseListColumns(),
			Category:           "basic",
			Order:              10,
		},
		"create": {
			Type:             schema.CommandCreate,
			Params:           databaseCreateParam(),
			IncludeFields:    databaseDetailIncludes(),
			ExcludeFields:    databaseDetailExcludes(),
			Category:         "basic",
			Order:            20,
			UseCustomCommand: true,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        databaseReadParam(),
			IncludeFields: databaseDetailIncludes(),
			ExcludeFields: databaseDetailExcludes(),
			Category:      "basic",
			Order:         30,
		},
		"update": {
			Type:             schema.CommandUpdate,
			Params:           databaseUpdateParam(),
			IncludeFields:    databaseDetailIncludes(),
			ExcludeFields:    databaseDetailExcludes(),
			Category:         "basic",
			Order:            40,
			UseCustomCommand: true,
		},
		"delete": {
			Type:             schema.CommandDelete,
			Aliases:          []string{"rm"},
			Params:           databaseDeleteParam(),
			IncludeFields:    databaseDetailIncludes(),
			ExcludeFields:    databaseDetailExcludes(),
			Category:         "basic",
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
			Params:           databaseWaitForParams(),
			Usage:            "Wait until boot is completed",
			UseCustomCommand: true,
			Category:         "power",
			Order:            40,
			NoOutput:         true,
			NeedlessConfirm:  true,
		},
		"wait-for-down": {
			Type:             schema.CommandManipulateMulti,
			Params:           databaseWaitForParams(),
			Usage:            "Wait until shutdown is completed",
			UseCustomCommand: true,
			Category:         "power",
			Order:            50,
			NoOutput:         true,
			NeedlessConfirm:  true,
		},
	}

	return &schema.Resource{
		Commands:          commands,
		ResourceCategory:  CategoryNetworking,
		CommandCategories: DatabaseCommandCategories,
	}
}

var DatabaseCommandCategories = []schema.Category{
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
		Key:         "monitor",
		DisplayName: "Monitoring",
		Order:       30,
	},
	{
		Key:         "other",
		DisplayName: "Other",
		Order:       1000,
	},
}

func databaseListParam() map[string]*schema.Schema {
	return CommonListParam
}

func databaseListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
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
		},
		"plan": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Required:     true,
			DefaultValue: 10,
			Description:  "set plan[10/30/90/240]",
			ValidateFunc: validateInIntValues(sacloud.AllowDatabasePlans()...),
			CompleteFunc: completeInIntValues(sacloud.AllowDatabasePlans()...),
		},
		"database": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"db"},
			Description:  "set database type",
			Required:     true,
			ValidateFunc: validateInStrValues("postgresql", "mariadb"),
			CompleteFunc: completeInStrValues("postgresql", "mariadb"),
		},
		"username": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set database default user name",
			Required:     true,
			ValidateFunc: validateStrLen(4, 20),
		},
		"password": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set database default user password",
			Required:     true,
			ValidateFunc: validateStrLen(8, 30),
		},
		"port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set database port",
			ValidateFunc: validateIntRange(1024, 65535),
		},
		"ipaddress1": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"ip1", "ipaddress", "ip"},
			Description:  "set ipaddress(#1)",
			ValidateFunc: validateIPv4Address(),
			Required:     true,
		},
		"nw_mask_len": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set network mask length",
			Required:     true,
			ValidateFunc: validateIntRange(8, 29),
		},
		"default_route": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set default route",
			ValidateFunc: validateIPv4Address(),
			Required:     true,
		},
		"source-networks": {
			Type:         schema.TypeStringList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set network of allow connection",
			ValidateFunc: validateStringSlice(validateStrLen(1, 32)),
		},
		"enable-web-ui": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "enable web-ui",
		},
		"backup-time": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set backup start time",
			ValidateFunc: validateBackupTime(),
			CompleteFunc: completeBackupTime(),
		},
		"name": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set resource display name",
			Required:     true,
			ValidateFunc: validateStrLen(1, 64),
		},
		"description": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set resource description",
			Aliases:      []string{"desc"},
			ValidateFunc: validateStrLen(0, 254),
		},
		"tags": {
			Type:         schema.TypeStringList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set resource tags",
			ValidateFunc: validateStringSlice(validateStrLen(1, 32)),
		},
		"icon-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set Icon ID",
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeIconID(),
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
		},
		"port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set database port",
			ValidateFunc: validateIntRange(1024, 65535),
		},
		"source-networks": {
			Type:         schema.TypeStringList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set network of allow connection",
			ValidateFunc: validateStringSlice(validateStrLen(1, 32)),
		},
		"enable-web-ui": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "enable web-ui",
		},
		"backup-time": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set backup start time",
			ValidateFunc: validateBackupTime(),
			CompleteFunc: completeBackupTime(),
		},
		"name": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set resource display name",
			ValidateFunc: validateStrLen(1, 64),
		},
		"description": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set resource description",
			Aliases:      []string{"desc"},
			ValidateFunc: validateStrLen(0, 254),
		},
		"tags": {
			Type:         schema.TypeStringList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set resource tags",
			ValidateFunc: validateStringSlice(validateStrLen(1, 32)),
		},
		"icon-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set Icon ID",
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeIconID(),
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

func databaseWaitForParams() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func databaseMonitorParam() map[string]*schema.Schema {
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
		"key-format": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set monitoring value key-format",
			DefaultValue: "sakuracloud.{{.ID}}.internet",
			Required:     true,
		},
	}
}

func databaseMonitorColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Key"},
		{Name: "TimeStamp"},
		{Name: "UnixTime"},
		{Name: "Receive"},
		{Name: "Send"},
	}
}

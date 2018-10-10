package define

import (
	"math"

	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func SIMResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find", "selector"},
			Params:             simListParam(),
			TableType:          output.TableSimple,
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
			Type:          schema.CommandDelete,
			Aliases:       []string{"rm"},
			Params:        simDeleteParam(),
			IncludeFields: simDetailIncludes(),
			ExcludeFields: simDetailExcludes(),
			Category:      "basic",
			Order:         50,
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
			Type:             schema.CommandRead,
			Params:           simLogsParam(),
			Order:            10,
			UseCustomCommand: true,
			NeedlessConfirm:  true,
			NoOutput:         true,
			Category:         "monitor",
		},
		"monitor": {
			Type:               schema.CommandRead,
			Params:             simMonitorParam(),
			TableType:          output.TableSimple,
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
		ListResultFieldName: "CommonServiceSIMItems",
	}
}

var SIMCommandCategories = []schema.Category{
	{
		Key:         "basic",
		DisplayName: "Basics",
		Order:       10,
	},
	{
		Key:         "activate",
		DisplayName: "Activate/Deactivate",
		Order:       20,
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

func simListParam() map[string]*schema.Schema {
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

func simCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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

func simReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func simInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func simUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func simDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func simActivateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func simDeactivateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func simIPAddParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ip": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			ValidateFunc: validateIPv4Address(),
			Required:     true,
			Order:        10,
		},
	}
}

func simIPDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func simIMEILockParams() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"imei": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Required:    true,
			Order:       10,
		},
	}
}

func simIMEIUnlockParams() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func simLogsParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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

func simMonitorParam() map[string]*schema.Schema {
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
			DefaultValue: "sakuracloud.sim.{{.ID}}",
			Required:     true,
			Category:     "monitor",
			Order:        30,
		},
	}
}

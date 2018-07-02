package define

import (
	"fmt"
	"math"

	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func MobileGatewayResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find", "selector"},
			Params:             mobileGatewayListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: mobileGatewayListColumns(),
			Category:           "basic",
			Order:              10,
		},
		"create": {
			Type:             schema.CommandCreate,
			Params:           mobileGatewayCreateParam(),
			IncludeFields:    mobileGatewayDetailIncludes(),
			ExcludeFields:    mobileGatewayDetailExcludes(),
			UseCustomCommand: true,
			Category:         "basic",
			Order:            20,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        mobileGatewayReadParam(),
			IncludeFields: mobileGatewayDetailIncludes(),
			ExcludeFields: mobileGatewayDetailExcludes(),
			Category:      "basic",
			Order:         30,
		},
		"update": {
			Type:             schema.CommandUpdate,
			Params:           mobileGatewayUpdateParam(),
			IncludeFields:    mobileGatewayDetailIncludes(),
			ExcludeFields:    mobileGatewayDetailExcludes(),
			UseCustomCommand: true,
			Category:         "basic",
			Order:            40,
		},
		"delete": {
			Type:             schema.CommandDelete,
			Aliases:          []string{"rm"},
			Params:           mobileGatewayDeleteParam(),
			IncludeFields:    mobileGatewayDetailIncludes(),
			ExcludeFields:    mobileGatewayDetailExcludes(),
			Category:         "basic",
			Order:            50,
			UseCustomCommand: true,
		},
		"boot": {
			Type:             schema.CommandManipulateMulti,
			Aliases:          []string{"power-on"},
			Params:           mobileGatewayPowerOnParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            10,
			NoOutput:         true,
		},
		"shutdown": {
			Type:             schema.CommandManipulateMulti,
			Aliases:          []string{"power-off"},
			Params:           mobileGatewayPowerOffParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            20,
			NoOutput:         true,
		},
		"shutdown-force": {
			Type:             schema.CommandManipulateMulti,
			Aliases:          []string{"stop"},
			Params:           mobileGatewayPowerOffParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            25,
			NoOutput:         true,
		},
		"reset": {
			Type:             schema.CommandManipulateMulti,
			Params:           mobileGatewayResetParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            30,
			NoOutput:         true,
		},
		"wait-for-boot": {
			Type:             schema.CommandManipulateMulti,
			Params:           mobileGatewayWaitForParams(),
			Usage:            "Wait until boot is completed",
			UseCustomCommand: true,
			Category:         "power",
			Order:            40,
			NoOutput:         true,
			NeedlessConfirm:  true,
		},
		"wait-for-down": {
			Type:             schema.CommandManipulateMulti,
			Params:           mobileGatewayWaitForParams(),
			Usage:            "Wait until shutdown is completed",
			UseCustomCommand: true,
			Category:         "power",
			Order:            50,
			NoOutput:         true,
			NeedlessConfirm:  true,
		},
		"interface-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             mobileGatewayInterfaceInfoParam(),
			Aliases:            []string{"interface-list"},
			Usage:              "Show information of NIC(s) connected to mobile-gateway",
			TableType:          output.TableSimple,
			TableColumnDefines: mobileGatewayInterfaceListColumns(),
			UseCustomCommand:   true,
			Category:           "nic",
			Order:              10,
			NeedlessConfirm:    true,
		},
		"interface-connect": {
			Type:             schema.CommandManipulateSingle,
			Params:           mobileGatewayInterfaceConnectParam(),
			Usage:            "Connected to switch",
			UseCustomCommand: true,
			Category:         "nic",
			Order:            20,
			NoOutput:         true,
		},
		"interface-update": {
			Type:             schema.CommandManipulateSingle,
			Params:           mobileGatewayInterfaceUpdateParam(),
			Usage:            "Update interface",
			UseCustomCommand: true,
			Category:         "nic",
			Order:            30,
			NoOutput:         true,
		},
		"interface-disconnect": {
			Type:             schema.CommandManipulateSingle,
			Params:           mobileGatewayInterfaceDisconnectParam(),
			Usage:            "Disconnected to switch",
			UseCustomCommand: true,
			Category:         "nic",
			Order:            40,
			NoOutput:         true,
		},
		"static-route-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             mobileGatewayStaticRouteInfoParam(),
			Aliases:            []string{"static-route-list"},
			Usage:              "Show information of static-routes",
			TableType:          output.TableSimple,
			TableColumnDefines: mobileGatewayStaticRouteListColumns(),
			UseCustomCommand:   true,
			Category:           "static-route",
			Order:              10,
			NeedlessConfirm:    true,
		},
		"static-route-add": {
			Type:             schema.CommandManipulateSingle,
			Params:           mobileGatewayStaticRouteAddParam(),
			Usage:            "Add static-route",
			UseCustomCommand: true,
			Category:         "static-route",
			Order:            20,
			NoOutput:         true,
		},
		"static-route-update": {
			Type:             schema.CommandManipulateSingle,
			Params:           mobileGatewayStaticRouteUpdateParam(),
			Usage:            "Update static-route",
			UseCustomCommand: true,
			Category:         "static-route",
			Order:            30,
			NoOutput:         true,
		},
		"static-route-delete": {
			Type:             schema.CommandManipulateSingle,
			Params:           mobileGatewayStaticRouteDeleteParam(),
			Usage:            "Delete static-route",
			UseCustomCommand: true,
			Category:         "static-route",
			Order:            40,
			NoOutput:         true,
		},
		"sim-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             mobileGatewaySIMInfoParam(),
			Aliases:            []string{"interface-list"},
			Usage:              "Show information of NIC(s) connected to mobile-gateway",
			TableType:          output.TableSimple,
			TableColumnDefines: mobileGatewaySIMInfoColumns(),
			UseCustomCommand:   true,
			Category:           "sim",
			Order:              10,
			NeedlessConfirm:    true,
		},
		"sim-add": {
			Type:             schema.CommandManipulateSingle,
			Params:           mobileGatewaySIMAddParam(),
			Usage:            "Connected to switch",
			UseCustomCommand: true,
			Category:         "sim",
			Order:            20,
			NoOutput:         true,
		},
		"sim-update": {
			Type:             schema.CommandManipulateSingle,
			Params:           mobileGatewaySIMUpdateParam(),
			Usage:            "Connected to switch",
			UseCustomCommand: true,
			Category:         "sim",
			Order:            30,
			NoOutput:         true,
		},
		"sim-delete": {
			Type:             schema.CommandManipulateSingle,
			Params:           mobileGatewaySIMDeleteParam(),
			Usage:            "Disconnected to switch",
			UseCustomCommand: true,
			Category:         "sim",
			Order:            40,
			NoOutput:         true,
		},
		"dns-update": {
			Type:             schema.CommandManipulateSingle,
			Params:           mobileGatewayDNSUpdateParam(),
			Usage:            "Update interface",
			UseCustomCommand: true,
			Category:         "dns",
			Order:            10,
			NoOutput:         true,
		},
		"logs": {
			Type:             schema.CommandRead,
			Params:           mobileGatewayLogsParam(),
			UseCustomCommand: true,
			NoOutput:         true,
			Category:         "monitor",
		},
	}

	return &schema.Resource{
		Commands:          commands,
		Aliases:           []string{"mgw"},
		ResourceCategory:  CategoryAppliance,
		CommandCategories: MobileGatewayCommandCategories,
	}
}

var MobileGatewayCommandCategories = []schema.Category{
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
		Key:         "nic",
		DisplayName: "Interface Management",
		Order:       30,
	},
	{
		Key:         "static-route",
		DisplayName: "StaticRoute Management",
		Order:       40,
	},
	{
		Key:         "sim",
		DisplayName: "SIM Management",
		Order:       50,
	},
	{
		Key:         "dns",
		DisplayName: "DNS Management",
		Order:       60,
	},
	{
		Key:         "monitor",
		DisplayName: "Monitoring",
		Order:       70,
	},
	{
		Key:         "other",
		DisplayName: "Other",
		Order:       1000,
	},
}

func mobileGatewayListParam() map[string]*schema.Schema {
	return mergeParameterMap(CommonListParam, paramTagsCond)
}

func mobileGatewayListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{
			Name:    "Power",
			Sources: []string{"Instance.Status"},
		},
		{
			Name: "IPAddress(Public)",
			FormatFunc: func(values map[string]string) string {
				if ip, ok := values["Interfaces.0.IPAddress"]; ok {
					format := "%s/%s"
					return fmt.Sprintf(format,
						ip,
						values["Interfaces.0.Switch.Subnet.NetworkMaskLen"],
					)
				}
				return ""
			},
		},
		{
			Name: "IPAddress(Private)",
			FormatFunc: func(values map[string]string) string {
				if ip, ok := values["Settings.MobileGateway.Interfaces.1.IPAddress.0"]; ok {
					format := "%s/%s"
					return fmt.Sprintf(format,
						ip,
						values["Settings.MobileGateway.Interfaces.1.NetworkMaskLen"],
					)
				}
				return ""
			},
		},
	}
}

func mobileGatewayInterfaceListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Type"},
		{Name: "Switch"},
		{Name: "IPAddress"},
		{Name: "NetworkMaskLen"},
	}
}

func mobileGatewayStaticRouteListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "__ORDER__"}, // magic column name(generated on demand)
		{Name: "Prefix"},
		{Name: "NextHop"},
	}
}

func mobileGatewaySIMInfoColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{
			Name:    "ID",
			Sources: []string{"resource_id"},
		},
		{
			Name:    "ICCID",
			Sources: []string{"iccid"},
		},
		{
			Name:    "IP",
			Sources: []string{"ip"},
		},
		{
			Name:    "Status",
			Sources: []string{"session_status"},
		},
		{
			Name:       "IMEI Lock",
			FormatFunc: formatBoolFunc("imei_lock"),
		},
		{
			Name:       "Registered",
			FormatFunc: formatBoolFunc("registered"),
		},
		{
			Name:       "Activated",
			FormatFunc: formatBoolFunc("activated"),
		},
	}
}

func mobileGatewayLogsColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "__ORDER__"}, // magic column name(generated on demand)
		{Name: "Date"},
		{Name: "SessionStatus"},
		{Name: "ResourceID"},
		{Name: "IMEI"},
		{Name: "IMSI"},
	}
}

func mobileGatewayDetailIncludes() []string {
	return []string{}
}

func mobileGatewayDetailExcludes() []string {
	return []string{}
}

func mobileGatewayCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
		"internet-connection": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "connect to internet",
			Category:    "network",
			Order:       10,
		},
	}
}

func mobileGatewayReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func mobileGatewayUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
		"internet-connection": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "connect to internet",
			Category:    "network",
			Order:       10,
		},
	}
}

func mobileGatewayDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"force": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"f"},
			Description: "forced-shutdown flag if mobile-gateway is running",
			Category:    "mobile-gateway",
			Order:       10,
		},
	}
}

func mobileGatewayPowerOnParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func mobileGatewayPowerOffParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func mobileGatewayResetParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func mobileGatewayWaitForParams() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func mobileGatewayInterfaceInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func mobileGatewayInterfaceConnectParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"switch-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set connect switch ID",
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeSwitchID(),
			Required:     true,
			Category:     "interface",
			Order:        20,
		},
		"ipaddress": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"ip"},
			Description:  "set ipaddress",
			ValidateFunc: validateIPv4Address(),
			Required:     true,
			Category:     "network",
			Order:        10,
		},
		"nw-masklen": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"network-masklen"},
			Description:  "set ipaddress prefix",
			DefaultValue: 24,
			ValidateFunc: validateIntRange(8, 29),
			Category:     "network",
			Order:        40,
		},
	}
}

func mobileGatewayInterfaceUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ipaddress": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"ip"},
			Description:  "set ipaddress",
			ValidateFunc: validateIPv4Address(),
			Category:     "network",
			Order:        10,
		},
		"nw-masklen": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"network-masklen"},
			Description:  "set ipaddress prefix",
			DefaultValue: 24,
			ValidateFunc: validateIntRange(8, 29),
			Category:     "network",
			Order:        50,
		},
	}
}

func mobileGatewayInterfaceDisconnectParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func mobileGatewayDNSUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"dns1": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set DNS server address",
			Required:     true,
			ValidateFunc: validateIPv4Address(),
			Category:     "dns",
			Order:        10,
		},
		"dns2": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set DNS server address",
			Required:     true,
			ValidateFunc: validateIPv4Address(),
			Category:     "dns",
			Order:        15,
		},
	}
}

func mobileGatewayLogsParam() map[string]*schema.Schema {
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

func mobileGatewayStaticRouteInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func mobileGatewayStaticRouteAddParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"prefix": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set prefix",
			Required:     true,
			ValidateFunc: validateIPv4AddressWithPrefix(),
			Category:     "Static-Route",
			Order:        10,
		},
		"next-hop": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set next-hop",
			Required:     true,
			ValidateFunc: validateIPv4Address(),
			Category:     "Static-Route",
			Order:        20,
		},
	}
}

func mobileGatewayStaticRouteUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target static-route",
			Required:    true,
			Category:    "Static-Route",
			Order:       1,
		},
		"prefix": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set prefix",
			ValidateFunc: validateIPv4AddressWithPrefix(),
			Category:     "Static-Route",
			Order:        10,
		},
		"next-hop": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set next-hop",
			ValidateFunc: validateIPv4Address(),
			Category:     "Static-Route",
			Order:        20,
		},
	}
}

func mobileGatewayStaticRouteDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target static-route",
			Required:    true,
			Category:    "Static-Route",
			Order:       1,
		},
	}
}

func mobileGatewaySIMInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func mobileGatewaySIMAddParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"sim-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeSIMID(),
			Required:     true,
			Category:     "interface",
			Order:        10,
		},
		"ipaddress": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"ip"},
			Description:  "set ipaddress",
			ValidateFunc: validateIPv4Address(),
			Required:     true,
			Category:     "network",
			Order:        10,
		},
	}
}

func mobileGatewaySIMUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"sim-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeSIMID(),
			Required:     true,
			Category:     "interface",
			Order:        10,
		},
		"ipaddress": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"ip"},
			Description:  "set ipaddress",
			ValidateFunc: validateIPv4Address(),
			Category:     "network",
			Order:        10,
		},
	}
}

func mobileGatewaySIMDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"sim-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeSIMID(),
			Required:     true,
			Category:     "interface",
			Order:        10,
		},
	}
}

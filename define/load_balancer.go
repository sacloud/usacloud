package define

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func LoadBalancerResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find", "selector"},
			Params:             loadBalancerListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: loadBalancerListColumns(),
			Category:           "basic",
			Order:              10,
		},
		"create": {
			Type:             schema.CommandCreate,
			Params:           loadBalancerCreateParam(),
			ParamCategories:  loadBalancerParamsCategories,
			IncludeFields:    loadBalancerDetailIncludes(),
			ExcludeFields:    loadBalancerDetailExcludes(),
			Category:         "basic",
			Order:            20,
			UseCustomCommand: true,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        loadBalancerReadParam(),
			IncludeFields: loadBalancerDetailIncludes(),
			ExcludeFields: loadBalancerDetailExcludes(),
			Category:      "basic",
			Order:         30,
		},
		"update": {
			Type:          schema.CommandUpdate,
			Params:        loadBalancerUpdateParam(),
			IncludeFields: loadBalancerDetailIncludes(),
			ExcludeFields: loadBalancerDetailExcludes(),
			Category:      "basic",
			Order:         40,
		},
		"delete": {
			Type:             schema.CommandDelete,
			Aliases:          []string{"rm"},
			Params:           loadBalancerDeleteParam(),
			IncludeFields:    loadBalancerDetailIncludes(),
			ExcludeFields:    loadBalancerDetailExcludes(),
			Category:         "basic",
			Order:            50,
			UseCustomCommand: true,
		},
		"boot": {
			Type:             schema.CommandManipulateMulti,
			Aliases:          []string{"power-on"},
			Params:           loadBalancerPowerOnParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            10,
			NoOutput:         true,
		},
		"shutdown": {
			Type:             schema.CommandManipulateMulti,
			Aliases:          []string{"power-off"},
			Params:           loadBalancerPowerOffParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            20,
			NoOutput:         true,
		},
		"shutdown-force": {
			Type:             schema.CommandManipulateMulti,
			Aliases:          []string{"stop"},
			Params:           loadBalancerPowerOffParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            25,
			NoOutput:         true,
		},
		"reset": {
			Type:             schema.CommandManipulateMulti,
			Params:           loadBalancerResetParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            30,
			NoOutput:         true,
		},
		"wait-for-boot": {
			Type:             schema.CommandManipulateMulti,
			Params:           loadBalancerWaitForParams(),
			Usage:            "Wait until boot is completed",
			UseCustomCommand: true,
			Category:         "power",
			Order:            40,
			NoOutput:         true,
			NeedlessConfirm:  true,
		},
		"wait-for-down": {
			Type:             schema.CommandManipulateMulti,
			Params:           loadBalancerWaitForParams(),
			Usage:            "Wait until shutdown is completed",
			UseCustomCommand: true,
			Category:         "power",
			Order:            50,
			NoOutput:         true,
			NeedlessConfirm:  true,
		},
		"vip-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             loadBalancerVIPInfoParam(),
			Usage:              "Show information of VIP(s)",
			TableType:          output.TableSimple,
			TableColumnDefines: loadBalancerVIPListColumns(),
			Category:           "vip",
			Order:              10,
			UseCustomCommand:   true,
			NeedlessConfirm:    true,
		},
		"vip-add": {
			Type:               schema.CommandManipulateSingle,
			Params:             loadBalancerVIPAddParam(),
			Usage:              "Add VIP to LoadBalancer",
			TableType:          output.TableSimple,
			TableColumnDefines: loadBalancerVIPListColumns(),
			Category:           "vip",
			Order:              20,
			UseCustomCommand:   true,
			NoOutput:           true,
		},
		"vip-update": {
			Type:               schema.CommandManipulateSingle,
			Params:             loadBalancerVIPUpdateParam(),
			Usage:              "Update VIP",
			TableType:          output.TableSimple,
			TableColumnDefines: loadBalancerVIPListColumns(),
			Category:           "vip",
			Order:              30,
			UseCustomCommand:   true,
			NoOutput:           true,
		},
		"vip-delete": {
			Type:               schema.CommandManipulateSingle,
			Params:             loadBalancerVIPDeleteParam(),
			Usage:              "Delete VIP from LoadBalancer",
			TableType:          output.TableSimple,
			TableColumnDefines: loadBalancerVIPListColumns(),
			Category:           "vip",
			Order:              40,
			UseCustomCommand:   true,
			NoOutput:           true,
		},
		"server-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             loadBalancerServerInfoParam(),
			Usage:              "Show servers under VIP(s)",
			TableType:          output.TableSimple,
			TableColumnDefines: loadBalancerServerListColumns(),
			Category:           "servers",
			Order:              10,
			UseCustomCommand:   true,
			NeedlessConfirm:    true,
		},
		"server-add": {
			Type:               schema.CommandManipulateSingle,
			Params:             loadBalancerServerAddParam(),
			Usage:              "Add server under VIP(s)",
			TableType:          output.TableSimple,
			TableColumnDefines: loadBalancerServerListColumns(),
			Category:           "servers",
			Order:              20,
			UseCustomCommand:   true,
			NoOutput:           true,
		},
		"server-update": {
			Type:               schema.CommandManipulateSingle,
			Params:             loadBalancerServerUpdateParam(),
			Usage:              "Update server under VIP(s)",
			TableType:          output.TableSimple,
			TableColumnDefines: loadBalancerServerListColumns(),
			Category:           "servers",
			Order:              30,
			UseCustomCommand:   true,
			NoOutput:           true,
		},
		"server-delete": {
			Type:               schema.CommandManipulateSingle,
			Params:             loadBalancerServerDeleteParam(),
			Usage:              "Delete server under VIP(s)",
			TableType:          output.TableSimple,
			TableColumnDefines: loadBalancerServerListColumns(),
			Category:           "servers",
			Order:              40,
			UseCustomCommand:   true,
			NoOutput:           true,
		},
		"monitor": {
			Type:               schema.CommandRead,
			Params:             loadBalancerMonitorParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: loadBalancerMonitorColumns(),
			UseCustomCommand:   true,
			Category:           "monitor",
		},
	}

	return &schema.Resource{
		Commands:          commands,
		ResourceCategory:  CategoryAppliance,
		CommandCategories: LoadBalancerCommandCategories,
	}
}

var LoadBalancerCommandCategories = []schema.Category{
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
		Key:         "servers",
		DisplayName: "Servers under VIP Management",
		Order:       40,
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

var loadBalancerParamsCategories = []schema.Category{
	{
		Key:         "load-balancer",
		DisplayName: "LoadBalancer options",
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

func loadBalancerListParam() map[string]*schema.Schema {
	return mergeParameterMap(CommonListParam, paramTagsCond)
}

func loadBalancerListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{
			Name:    "Power",
			Sources: []string{"Instance.Status"},
		},
		{
			Name:    "VRID",
			Sources: []string{"Remark.VRRP.VRID"},
		},
		{
			Name:    "Plan",
			Sources: []string{"Plan.ID"},
			ValueMapping: []map[string]string{
				{
					"1": "standard",
					"2": "highspec",
				},
			},
		},
		{
			Name: "HA",
			FormatFunc: func(values map[string]string) string {
				if ip, ok := values["Remark.Servers.1.IPAddress"]; ok {
					if ip != "" {
						return "true"
					}
				}

				return "false"
			},
		},
		{
			Name: "IPAddress(#1)",
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
			Name: "IPAddress(#2)",
			FormatFunc: func(values map[string]string) string {
				if ip, ok := values["Remark.Servers.1.IPAddress"]; ok {
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

func loadBalancerVIPListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "__ORDER__"}, // magic column name(generated on demand)
		{
			Name:    "VIP",
			Sources: []string{"VirtualIPAddress"},
		},
		{Name: "Port"},
		{Name: "DelayLoop"},
		{Name: "SorryServer"},
	}
}

func loadBalancerServerListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "__ORDER__"}, // magic column name(generated on demand)
		{Name: "IPAddress"},
		{Name: "Port"},
		{Name: "Enabled"},
		{
			Name:    "Protocol",
			Sources: []string{"HealthCheck.Protocol"},
		},
		{
			Name:    "Path",
			Sources: []string{"HealthCheck.Path"},
		},
		{
			Name:    "ResponseCode",
			Sources: []string{"HealthCheck.Status"},
		},
	}
}

func loadBalancerDetailIncludes() []string {
	return []string{}
}

func loadBalancerDetailExcludes() []string {
	return []string{}
}

func loadBalancerCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"switch-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set connect switch ID",
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeSwitchID(),
			Category:     "load-balancer",
			Order:        10,
		},
		"vrid": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"VRID"},
			Description:  "set VRID",
			DefaultValue: 1,
			Required:     true,
			Category:     "load-balancer",
			Order:        20,
		},
		"high-availability": {
			Type:         schema.TypeBool,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"ha"},
			Description:  "use HA(High-Availability) mode",
			DefaultValue: false,
			Category:     "load-balancer",
			Order:        30,
		},
		"plan": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Required:     true,
			DefaultValue: "standard",
			Description:  "set plan[standard/highspec]",
			ValidateFunc: validateInStrValues("standard", "highspec"),
			CompleteFunc: completeInStrValues("standard", "highspec"),
			Category:     "load-balancer",
			Order:        40,
		},
		"ipaddress1": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"ip1"},
			Description:  "set ipaddress(#1)",
			ValidateFunc: validateIPv4Address(),
			Required:     true,
			Category:     "network",
			Order:        10,
		},
		"ipaddress2": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"ip2"},
			Description:  "set ipaddress(#2)",
			ValidateFunc: validateIPv4Address(),
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

func loadBalancerReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func loadBalancerUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func loadBalancerDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"force": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"f"},
			Description: "forced-shutdown flag if load-balancer is running",
			Category:    "load-balancer",
			Order:       10,
		},
	}
}

func loadBalancerPowerOnParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func loadBalancerPowerOffParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func loadBalancerResetParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func loadBalancerWaitForParams() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func loadBalancerVIPInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func loadBalancerVIPAddParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"vip": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set VirtualIPAddress",
			ValidateFunc: validateIPv4Address(),
			Required:     true,
			Category:     "vip",
			Order:        10,
		},
		"port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set port",
			ValidateFunc: validateIntRange(1, 65535),
			Required:     true,
			Category:     "vip",
			Order:        20,
		},
		"delay-loop": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set delay-loop",
			ValidateFunc: validateIntRange(10, 2147483647),
			DefaultValue: 10,
			Category:     "vip",
			Order:        30,
		},
		"sorry-server": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set IPAddress of sorry-server",
			ValidateFunc: validateIPv4Address(),
			Category:     "vip",
			Order:        40,
		},
	}
}

func loadBalancerVIPUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target VIP",
			Required:    true,
			Category:    "vip",
			Order:       1,
		},
		"vip": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set VirtualIPAddress",
			ValidateFunc: validateIPv4Address(),
			Category:     "vip",
			Order:        10,
		},
		"port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set port",
			ValidateFunc: validateIntRange(1, 65535),
			Category:     "vip",
			Order:        20,
		},
		"delay-loop": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set delay-loop",
			ValidateFunc: validateIntRange(10, 2147483647),
			DefaultValue: 10,
			Category:     "vip",
			Order:        30,
		},
		"sorry-server": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set IPAddress of sorry-server",
			ValidateFunc: validateIPv4Address(),
			Category:     "vip",
			Order:        40,
		},
	}
}

func loadBalancerVIPDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target VIP",
			Required:    true,
			Category:    "vip",
			Order:       1,
		},
	}
}

func loadBalancerServerInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"vip-index": {
			Type:          schema.TypeInt,
			HandlerType:   schema.HandlerNoop,
			Description:   "index of target VIP",
			ConflictsWith: []string{"vip", "port"},
			Category:      "server",
			Order:         1,
		},
		"vip": {
			Type:          schema.TypeString,
			HandlerType:   schema.HandlerNoop,
			Description:   "set VirtualIPAddress",
			ValidateFunc:  validateIPv4Address(),
			ConflictsWith: []string{"vip-index"},
			Category:      "server",
			Order:         2,
		},
		"port": {
			Type:          schema.TypeInt,
			HandlerType:   schema.HandlerNoop,
			Description:   "set port",
			ValidateFunc:  validateIntRange(1, 65535),
			ConflictsWith: []string{"vip-index"},
			Category:      "server",
			Order:         3,
		},
	}
}

func loadBalancerServerAddParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"vip-index": {
			Type:          schema.TypeInt,
			HandlerType:   schema.HandlerNoop,
			Description:   "index of target VIP",
			ConflictsWith: []string{"vip", "port"},
			Category:      "server",
			Order:         1,
		},
		"vip": {
			Type:          schema.TypeString,
			HandlerType:   schema.HandlerNoop,
			Description:   "set VirtualIPAddress",
			ValidateFunc:  validateIPv4Address(),
			ConflictsWith: []string{"vip-index"},
			Category:      "server",
			Order:         2,
		},
		"port": {
			Type:          schema.TypeInt,
			HandlerType:   schema.HandlerNoop,
			Description:   "set port",
			ValidateFunc:  validateIntRange(1, 65535),
			ConflictsWith: []string{"vip-index"},
			Category:      "server",
			Order:         3,
		},
		"ipaddress": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"ip"},
			Description:  "set real server IPAddress",
			ValidateFunc: validateIPv4Address(),
			Required:     true,
			Category:     "server",
			Order:        10,
		},
		"protocol": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set health check protocol[http/https/ping/tcp]",
			ValidateFunc: validateInStrValues(sacloud.AllowLoadBalancerHealthCheckProtocol()...),
			CompleteFunc: completeInStrValues(sacloud.AllowLoadBalancerHealthCheckProtocol()...),
			Required:     true,
			DefaultValue: "ping",
			Category:     "server",
			Order:        20,
		},
		"path": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set path of http/https health check request",
			Category:    "server",
			Order:       30,
		},
		"response-code": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "set expect response-code of http/https health check request",
			Category:    "server",
			Order:       40,
		},
		"disabled": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set disable",
			Category:    "server",
			Order:       50,
		},
	}
}

func loadBalancerServerUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"vip-index": {
			Type:          schema.TypeInt,
			HandlerType:   schema.HandlerNoop,
			Description:   "index of target VIP",
			ConflictsWith: []string{"vip", "port"},
			Category:      "server",
			Order:         1,
		},
		"vip": {
			Type:          schema.TypeString,
			HandlerType:   schema.HandlerNoop,
			Description:   "set VirtualIPAddress",
			ValidateFunc:  validateIPv4Address(),
			ConflictsWith: []string{"vip-index"},
			Category:      "server",
			Order:         2,
		},
		"port": {
			Type:          schema.TypeInt,
			HandlerType:   schema.HandlerNoop,
			Description:   "set port",
			ValidateFunc:  validateIntRange(1, 65535),
			ConflictsWith: []string{"vip-index"},
			Category:      "server",
			Order:         3,
		},
		"ipaddress": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"ip"},
			Description:  "set real server IPAddress",
			ValidateFunc: validateIPv4Address(),
			Required:     true,
			Category:     "server",
			Order:        10,
		},
		"protocol": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set health check protocol[http/https/ping/tcp]",
			ValidateFunc: validateInStrValues(sacloud.AllowLoadBalancerHealthCheckProtocol()...),
			CompleteFunc: completeInStrValues(sacloud.AllowLoadBalancerHealthCheckProtocol()...),
			Category:     "server",
			Order:        20,
		},
		"path": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set path of http/https health check request",
			Category:    "server",
			Order:       30,
		},
		"response-code": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "set expect response-code of http/https health check request",
			Category:    "server",
			Order:       40,
		},
		"disabled": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set enable/disable",
			Category:    "server",
			Order:       50,
		},
	}
}

func loadBalancerServerDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"vip-index": {
			Type:          schema.TypeInt,
			HandlerType:   schema.HandlerNoop,
			Description:   "index of target VIP",
			ConflictsWith: []string{"vip", "port"},
			Category:      "server",
			Order:         1,
		},
		"vip": {
			Type:          schema.TypeString,
			HandlerType:   schema.HandlerNoop,
			Description:   "set VirtualIPAddress",
			ValidateFunc:  validateIPv4Address(),
			ConflictsWith: []string{"vip-index"},
			Category:      "server",
			Order:         2,
		},
		"port": {
			Type:          schema.TypeInt,
			HandlerType:   schema.HandlerNoop,
			Description:   "set port",
			ValidateFunc:  validateIntRange(1, 65535),
			ConflictsWith: []string{"vip-index"},
			Category:      "server",
			Order:         3,
		},
		"ipaddress": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"ip"},
			Description:  "set real server IPAddress",
			ValidateFunc: validateIPv4Address(),
			Required:     true,
			Category:     "server",
			Order:        10,
		},
	}
}

func loadBalancerMonitorParam() map[string]*schema.Schema {
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
			DefaultValue: "sakuracloud.loadbalancer.{{.ID}}.nic",
			Required:     true,
			Category:     "monitor",
			Order:        30,
		},
	}
}

func loadBalancerMonitorColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Key"},
		{Name: "TimeStamp"},
		{Name: "UnixTime"},
		{Name: "Receive"},
		{Name: "Send"},
	}
}

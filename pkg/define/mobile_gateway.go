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

	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/schema"
)

func MobileGatewayResource() *schema.Resource {
	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find", "select"},
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
		"traffic-control-info": {
			Type:               schema.CommandRead,
			Params:             mobileGatewayTrafficControlInfoParam(),
			Usage:              "Show information of traffic-control",
			TableType:          output.TableSimple,
			TableColumnDefines: mobileGatewayTrafficControlInfoColumns(),
			UseCustomCommand:   true,
			Category:           "traffic-control",
			Order:              10,
		},
		"traffic-control-enable": {
			Type:             schema.CommandUpdate,
			Params:           mobileGatewayTrafficControlCreateParam(),
			Usage:            "Enable traffic-control",
			NoOutput:         true,
			UseCustomCommand: true,
			Category:         "traffic-control",
			Order:            20,
		},
		"traffic-control-update": {
			Type:             schema.CommandUpdate,
			Params:           mobileGatewayTrafficControlUpdateParam(),
			Usage:            "Update traffic-control config",
			NoOutput:         true,
			UseCustomCommand: true,
			Category:         "traffic-control",
			Order:            30,
		},
		"traffic-control-disable": {
			Type:             schema.CommandUpdate,
			Aliases:          []string{"traffic-control-delete"},
			Params:           mobileGatewayTrafficControlDeleteParam(),
			Usage:            "Disable traffic-control config",
			NoOutput:         true,
			UseCustomCommand: true,
			Category:         "traffic-control",
			Order:            40,
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
		"sim-route-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             mobileGatewaySIMRouteInfoParam(),
			Aliases:            []string{"sim-route-list"},
			Usage:              "Show information of sim-routes",
			TableType:          output.TableSimple,
			TableColumnDefines: mobileGatewaySIMRouteListColumns(),
			UseCustomCommand:   true,
			Category:           "sim-route",
			Order:              10,
			NeedlessConfirm:    true,
		},
		"sim-route-add": {
			Type:             schema.CommandManipulateSingle,
			Params:           mobileGatewaySIMRouteAddParam(),
			Usage:            "Add sim-route",
			UseCustomCommand: true,
			Category:         "sim-route",
			Order:            20,
			NoOutput:         true,
		},
		"sim-route-update": {
			Type:             schema.CommandManipulateSingle,
			Params:           mobileGatewaySIMRouteUpdateParam(),
			Usage:            "Update sim-route",
			UseCustomCommand: true,
			Category:         "sim-route",
			Order:            30,
			NoOutput:         true,
		},
		"sim-route-delete": {
			Type:             schema.CommandManipulateSingle,
			Params:           mobileGatewaySIMRouteDeleteParam(),
			Usage:            "Delete sim-route",
			UseCustomCommand: true,
			Category:         "sim-route",
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
		Key:         "traffic-control",
		DisplayName: "Traffic Control",
		Order:       40,
	},
	{
		Key:         "static-route",
		DisplayName: "StaticRoute Management",
		Order:       50,
	},
	{
		Key:         "sim",
		DisplayName: "SIM Management",
		Order:       60,
	},
	{
		Key:         "sim-route",
		DisplayName: "SIM Route Management",
		Order:       65,
	},
	{
		Key:         "dns",
		DisplayName: "DNS Management",
		Order:       70,
	},
	{
		Key:         "monitor",
		DisplayName: "Monitoring",
		Order:       80,
	},
	{
		Key:         "other",
		DisplayName: "Other",
		Order:       1000,
	},
}

func mobileGatewayListParam() map[string]*schema.Parameter {
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

func mobileGatewayTrafficControlInfoColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{
			Name:    "Quota(MB)",
			Sources: []string{"config.traffic_quota_in_mb"},
		},
		{
			Name:    "Limit(Kbps)",
			Sources: []string{"config.bandwidth_limit_in_kbps"},
		},
		{
			Name:    "Email",
			Sources: []string{"config.email_config.enabled"},
		},
		{
			Name:    "Slack",
			Sources: []string{"config.slack_config.slack_url"},
		},
		{
			Name:    "TrafficShaping",
			Sources: []string{"config.auto_traffic_shaping"},
		},
		{
			Name:    "UplinkBPS",
			Sources: []string{"status.uplink_bytes"},
		},
		{
			Name:    "DownlinkBPS",
			Sources: []string{"status.downlink_bytes"},
		},
	}
}

func mobileGatewayStaticRouteListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "__ORDER__"}, // magic column name(generated on demand)
		{Name: "Prefix"},
		{Name: "NextHop"},
	}
}

func mobileGatewaySIMRouteListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "__ORDER__"}, // magic column name(generated on demand)
		{
			Name:    "Prefix",
			Sources: []string{"prefix"},
		},
		{
			Name:    "SIM ID",
			Sources: []string{"resource_id"},
		},
		{
			Name:    "ICCID",
			Sources: []string{"iccid"},
		},
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

func mobileGatewayDetailIncludes() []string {
	return []string{}
}

func mobileGatewayDetailExcludes() []string {
	return []string{}
}

func mobileGatewayCreateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
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
		"internet-connection": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "connect to internet",
			Category:    "network",
			Order:       10,
		},
	}
}

func mobileGatewayReadParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func mobileGatewayUpdateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
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

func mobileGatewayDeleteParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
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

func mobileGatewayPowerOnParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func mobileGatewayPowerOffParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func mobileGatewayResetParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func mobileGatewayWaitForParams() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func mobileGatewayInterfaceInfoParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func mobileGatewayInterfaceConnectParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"switch-id": {
			Type:         schema.TypeId,
			HandlerType:  schema.HandlerNoop,
			Description:  "set connect switch ID",
			ValidateFunc: validateSakuraID(),
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

func mobileGatewayInterfaceUpdateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
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

func mobileGatewayInterfaceDisconnectParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func mobileGatewayTrafficControlInfoParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func mobileGatewayTrafficControlCreateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"quota": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Required:     true,
			DefaultValue: 512,
			ValidateFunc: validateIntRange(1, math.MaxInt32),
			Category:     "traffic-control",
			Order:        10,
		},
		"band-width-limit": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			ValidateFunc: validateIntRange(1, math.MaxInt32),
			Category:     "traffic-control",
			Order:        20,
		},
		"enable-email": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Category:    "traffic-control",
			Order:       30,
		},
		"slack-webhook-url": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Category:     "traffic-control",
			ValidateFunc: validateSlackWebhookURL(),
			Order:        40,
		},
		"auto-traffic-shaping": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Category:    "traffic-control",
			Order:       50,
		},
	}
}

func mobileGatewayTrafficControlUpdateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"quota": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			ValidateFunc: validateIntRange(1, math.MaxInt32),
			Category:     "traffic-control",
			Order:        10,
		},
		"band-width-limit": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			ValidateFunc: validateIntRange(1, math.MaxInt32),
			Category:     "traffic-control",
			Order:        20,
		},
		"enable-email": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Category:    "traffic-control",
			Order:       30,
		},
		"slack-webhook-url": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Category:     "traffic-control",
			ValidateFunc: validateSlackWebhookURL(),
			Order:        40,
		},
		"auto-traffic-shaping": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Category:    "traffic-control",
			Order:       50,
		},
	}
}

func mobileGatewayTrafficControlDeleteParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func mobileGatewayDNSUpdateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
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

func mobileGatewayLogsParam() map[string]*schema.Parameter {
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

func mobileGatewayStaticRouteInfoParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func mobileGatewayStaticRouteAddParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
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

func mobileGatewayStaticRouteUpdateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
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

func mobileGatewayStaticRouteDeleteParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
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

func mobileGatewaySIMInfoParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func mobileGatewaySIMAddParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"sim-id": {
			Type:         schema.TypeId,
			HandlerType:  schema.HandlerNoop,
			ValidateFunc: validateSakuraID(),
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

func mobileGatewaySIMUpdateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"sim-id": {
			Type:         schema.TypeId,
			HandlerType:  schema.HandlerNoop,
			ValidateFunc: validateSakuraID(),
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

func mobileGatewaySIMDeleteParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"sim-id": {
			Type:         schema.TypeId,
			HandlerType:  schema.HandlerNoop,
			ValidateFunc: validateSakuraID(),
			Required:     true,
			Category:     "interface",
			Order:        10,
		},
	}
}

func mobileGatewaySIMRouteInfoParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func mobileGatewaySIMRouteAddParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"prefix": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set prefix",
			Required:     true,
			ValidateFunc: validateIPv4AddressWithPrefix(),
			Category:     "SIM-Route",
			Order:        10,
		},
		"sim": {
			Type:         schema.TypeId,
			HandlerType:  schema.HandlerNoop,
			Description:  "set sim",
			Required:     true,
			ValidateFunc: validateSakuraID(),
			Category:     "SIM-Route",
			Order:        20,
		},
	}
}

func mobileGatewaySIMRouteUpdateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target sim-route",
			Required:    true,
			Category:    "SIM-Route",
			Order:       1,
		},
		"prefix": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set prefix",
			ValidateFunc: validateIPv4AddressWithPrefix(),
			Category:     "SIM-Route",
			Order:        10,
		},
		"sim": {
			Type:         schema.TypeId,
			HandlerType:  schema.HandlerNoop,
			Description:  "set sim",
			ValidateFunc: validateSakuraID(),
			Category:     "SIM-Route",
			Order:        20,
		},
	}
}

func mobileGatewaySIMRouteDeleteParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target sim-route",
			Required:    true,
			Category:    "SIM-Route",
			Order:       1,
		},
	}
}

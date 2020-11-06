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
	"strings"

	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/schema"
)

func VPCRouterResource() *schema.Resource {
	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find", "select"},
			Params:             vpcRouterListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: vpcRouterListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"create": {
			Type:             schema.CommandCreate,
			Params:           vpcRouterCreateParam(),
			ParamCategories:  vpcRouterCreateParamCategories,
			IncludeFields:    vpcRouterDetailIncludes(),
			ExcludeFields:    vpcRouterDetailExcludes(),
			Category:         "basics",
			Order:            20,
			UseCustomCommand: true,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        vpcRouterReadParam(),
			IncludeFields: vpcRouterDetailIncludes(),
			ExcludeFields: vpcRouterDetailExcludes(),
			Category:      "basics",
			Order:         30,
		},
		"update": {
			Type:             schema.CommandUpdate,
			Params:           vpcRouterUpdateParam(),
			ParamCategories:  vpcRouterUpdateParamCategories,
			IncludeFields:    vpcRouterDetailIncludes(),
			ExcludeFields:    vpcRouterDetailExcludes(),
			Category:         "basics",
			Order:            40,
			UseCustomCommand: true,
		},
		"delete": {
			Type:             schema.CommandDelete,
			Aliases:          []string{"rm"},
			Params:           vpcRouterDeleteParam(),
			ParamCategories:  vpcRouterDeleteParamCategories,
			IncludeFields:    vpcRouterDetailIncludes(),
			ExcludeFields:    vpcRouterDetailExcludes(),
			Category:         "basics",
			Order:            50,
			UseCustomCommand: true,
		},
		"boot": {
			Type:             schema.CommandManipulateMulti,
			Aliases:          []string{"power-on"},
			Params:           vpcRouterPowerOnParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            10,
			NoOutput:         true,
		},
		"shutdown": {
			Type:             schema.CommandManipulateMulti,
			Aliases:          []string{"power-off"},
			Params:           vpcRouterPowerOffParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            20,
			NoOutput:         true,
		},
		"shutdown-force": {
			Type:             schema.CommandManipulateMulti,
			Aliases:          []string{"stop"},
			Params:           vpcRouterPowerOffParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            25,
			NoOutput:         true,
		},
		"reset": {
			Type:             schema.CommandManipulateMulti,
			Params:           vpcRouterResetParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            30,
			NoOutput:         true,
		},
		"wait-for-boot": {
			Type:             schema.CommandManipulateMulti,
			Params:           vpcRouterWaitForParams(),
			Usage:            "Wait until boot is completed",
			UseCustomCommand: true,
			Category:         "power",
			Order:            40,
			NoOutput:         true,
			NeedlessConfirm:  true,
		},
		"wait-for-down": {
			Type:             schema.CommandManipulateMulti,
			Params:           vpcRouterWaitForParams(),
			Usage:            "Wait until shutdown is completed",
			UseCustomCommand: true,
			Category:         "power",
			Order:            50,
			NoOutput:         true,
			NeedlessConfirm:  true,
		},
		"enable-internet-connection": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterEnableInternetParam(),
			Usage:            "Enable internet connection from VPCRouter",
			UseCustomCommand: true,
			Category:         "nic",
			Order:            5,
			NoOutput:         true,
		},
		"disable-internet-connection": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterEnableInternetParam(),
			Usage:            "Enable internet connection from VPCRouter",
			UseCustomCommand: true,
			Category:         "nic",
			Order:            6,
			NoOutput:         true,
		},
		"interface-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             vpcRouterInterfaceInfoParam(),
			Aliases:            []string{"interface-list"},
			Usage:              "Show information of NIC(s) connected to vpc-router",
			TableType:          output.TableSimple,
			TableColumnDefines: vpcRouterInterfaceListColumns(),
			UseCustomCommand:   true,
			Category:           "nic",
			Order:              10,
			NeedlessConfirm:    true,
		},
		"interface-connect": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterInterfaceConnectParam(),
			Usage:            "Connected to switch",
			UseCustomCommand: true,
			Category:         "nic",
			Order:            20,
			NoOutput:         true,
		},
		"interface-update": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterInterfaceUpdateParam(),
			Usage:            "Update interface",
			UseCustomCommand: true,
			Category:         "nic",
			Order:            30,
			NoOutput:         true,
		},
		"interface-disconnect": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterInterfaceDisconnectParam(),
			Usage:            "Disconnected to switch",
			UseCustomCommand: true,
			Category:         "nic",
			Order:            40,
			NoOutput:         true,
		},
		"static-nat-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             vpcRouterStaticNATInfoParam(),
			Aliases:            []string{"static-nat-list"},
			Usage:              "Show information of static NAT settings",
			TableType:          output.TableSimple,
			TableColumnDefines: vpcRouterStaticNATListColumns(),
			UseCustomCommand:   true,
			Category:           "snat",
			Order:              10,
			NeedlessConfirm:    true,
		},
		"static-nat-add": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterStaticNATAddParam(),
			Usage:            "Add static NAT",
			UseCustomCommand: true,
			Category:         "snat",
			Order:            20,
			NoOutput:         true,
		},
		"static-nat-update": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterStaticNATUpdateParam(),
			Usage:            "Update static NAT",
			UseCustomCommand: true,
			Category:         "snat",
			Order:            30,
			NoOutput:         true,
		},
		"static-nat-delete": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterStaticNATDeleteParam(),
			Usage:            "Delete static NAT",
			UseCustomCommand: true,
			Category:         "snat",
			Order:            40,
			NoOutput:         true,
		},
		"port-forwarding-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             vpcRouterPortForwardingInfoParam(),
			Aliases:            []string{"port-forwarding-list"},
			Usage:              "Show information of port-forwarding settings",
			TableType:          output.TableSimple,
			TableColumnDefines: vpcRouterPortForwardingListColumns(),
			UseCustomCommand:   true,
			Category:           "rnat",
			Order:              10,
			NeedlessConfirm:    true,
		},
		"port-forwarding-add": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterPortForwardingAddParam(),
			Usage:            "Add port forwarding",
			UseCustomCommand: true,
			Category:         "rnat",
			Order:            20,
			NoOutput:         true,
		},
		"port-forwarding-update": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterPortForwardingUpdateParam(),
			Usage:            "Update port forwarding",
			UseCustomCommand: true,
			Category:         "rnat",
			Order:            30,
			NoOutput:         true,
		},
		"port-forwarding-delete": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterPortForwardingDeleteParam(),
			Usage:            "Delete port forwarding",
			UseCustomCommand: true,
			Category:         "rnat",
			Order:            40,
			NoOutput:         true,
		},
		"firewall-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             vpcRouterFirewallInfoParam(),
			Aliases:            []string{"firewall-list"},
			Usage:              "Show information of firewall rules",
			TableType:          output.TableSimple,
			TableColumnDefines: vpcRouterFirewallListColumns(),
			UseCustomCommand:   true,
			Category:           "fw",
			Order:              10,
			NeedlessConfirm:    true,
		},
		"firewall-add": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterFirewallAddParam(),
			Usage:            "Add firewall rule",
			UseCustomCommand: true,
			Category:         "fw",
			Order:            20,
			NoOutput:         true,
		},
		"firewall-update": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterFirewallUpdateParam(),
			Usage:            "Update firewall rule",
			UseCustomCommand: true,
			Category:         "fw",
			Order:            30,
			NoOutput:         true,
		},
		"firewall-delete": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterFirewallDeleteParam(),
			Usage:            "Delete firewall rule",
			UseCustomCommand: true,
			Category:         "fw",
			Order:            40,
			NoOutput:         true,
		},
		"dhcp-server-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             vpcRouterDHCPServerInfoParam(),
			Aliases:            []string{"dhcp-server-list"},
			Usage:              "Show information of DHCP servers",
			TableType:          output.TableSimple,
			TableColumnDefines: vpcRouterDHCPServerListColumns(),
			UseCustomCommand:   true,
			Category:           "dhcp-server",
			Order:              10,
			NeedlessConfirm:    true,
		},
		"dhcp-server-add": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterDHCPServerAddParam(),
			Usage:            "Add DHCP server",
			UseCustomCommand: true,
			Category:         "dhcp-server",
			Order:            20,
			NoOutput:         true,
		},
		"dhcp-server-update": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterDHCPServerUpdateParam(),
			Usage:            "Update DHCP server",
			UseCustomCommand: true,
			Category:         "dhcp-server",
			Order:            30,
			NoOutput:         true,
		},
		"dhcp-server-delete": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterDHCPServerDeleteParam(),
			Usage:            "Delete DHCP server",
			UseCustomCommand: true,
			Category:         "dhcp-server",
			Order:            40,
			NoOutput:         true,
		},
		"dhcp-static-mapping-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             vpcRouterDHCPStaticMappingInfoParam(),
			Aliases:            []string{"dhcp-static-mapping-list"},
			Usage:              "Show information of DHCP static mapping",
			TableType:          output.TableSimple,
			TableColumnDefines: vpcRouterDHCPStaticMappingListColumns(),
			UseCustomCommand:   true,
			Category:           "dhcp-static-mapping",
			Order:              10,
			NeedlessConfirm:    true,
		},
		"dhcp-static-mapping-add": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterDHCPStaticMappingAddParam(),
			Usage:            "Add DHCP static mapping",
			UseCustomCommand: true,
			Category:         "dhcp-static-mapping",
			Order:            20,
			NoOutput:         true,
		},
		"dhcp-static-mapping-update": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterDHCPStaticMappingUpdateParam(),
			Usage:            "Update DHCP static mapping",
			UseCustomCommand: true,
			Category:         "dhcp-static-mapping",
			Order:            30,
			NoOutput:         true,
		},
		"dhcp-static-mapping-delete": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterDHCPStaticMappingDeleteParam(),
			Usage:            "Delete DHCP static mapping",
			UseCustomCommand: true,
			Category:         "dhcp-static-mapping",
			Order:            40,
			NoOutput:         true,
		},
		"pptp-server-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             vpcRouterPPTPServerInfoParam(),
			Usage:              "Show information of PPTP server",
			TableType:          output.TableSimple,
			TableColumnDefines: vpcRouterPPTPServerListColumns(),
			UseCustomCommand:   true,
			Category:           "remote-access",
			Order:              10,
			NeedlessConfirm:    true,
		},
		"pptp-server-update": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterPPTPServerUpdateParam(),
			Usage:            "Update PPTP server setting",
			UseCustomCommand: true,
			Category:         "remote-access",
			Order:            15,
			NoOutput:         true,
		},
		"l2tp-server-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             vpcRouterL2TPServerInfoParam(),
			Usage:              "Show information of L2TP/IPSec server",
			TableType:          output.TableSimple,
			TableColumnDefines: vpcRouterL2TPServerListColumns(),
			UseCustomCommand:   true,
			Category:           "remote-access",
			Order:              20,
			NeedlessConfirm:    true,
		},
		"l2tp-server-update": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterL2TPServerUpdateParam(),
			Usage:            "Update L2TP/IPSec server setting",
			UseCustomCommand: true,
			Category:         "remote-access",
			Order:            25,
			NoOutput:         true,
		},
		"user-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             vpcRouterUserInfoParam(),
			Aliases:            []string{"user-list"},
			Usage:              "Show information of remote-access users",
			TableType:          output.TableSimple,
			TableColumnDefines: vpcRouterUserListColumns(),
			UseCustomCommand:   true,
			Category:           "remote-access",
			Order:              30,
			NeedlessConfirm:    true,
		},
		"user-add": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterUserAddParam(),
			Usage:            "Add remote-access user",
			UseCustomCommand: true,
			Category:         "remote-access",
			Order:            31,
			NoOutput:         true,
		},
		"user-update": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterUserUpdateParam(),
			Usage:            "Update remote-access user",
			UseCustomCommand: true,
			Category:         "remote-access",
			Order:            32,
			NoOutput:         true,
		},
		"user-delete": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterUserDeleteParam(),
			Usage:            "Delete remote-access user",
			UseCustomCommand: true,
			Category:         "remote-access",
			Order:            33,
			NoOutput:         true,
		},
		"site-to-site-vpn-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             vpcRouterS2SInfoParam(),
			Aliases:            []string{"site-to-site-vpn-list"},
			Usage:              "Show information of site-to-site IPSec VPN settings",
			TableType:          output.TableSimple,
			TableColumnDefines: vpcRouterS2SListColumns(),
			UseCustomCommand:   true,
			Category:           "s2s",
			Order:              10,
			NeedlessConfirm:    true,
		},
		"site-to-site-vpn-add": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterS2SAddParam(),
			Usage:            "Add site-to-site IPSec VPN setting",
			UseCustomCommand: true,
			Category:         "s2s",
			Order:            20,
			NoOutput:         true,
		},
		"site-to-site-vpn-update": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterS2SUpdateParam(),
			Usage:            "Update site-to-site IPSec VPN setting",
			UseCustomCommand: true,
			Category:         "s2s",
			Order:            30,
			NoOutput:         true,
		},
		"site-to-site-vpn-delete": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterS2SDeleteParam(),
			Usage:            "Delete site-to-site IPSec VPN setting",
			UseCustomCommand: true,
			Category:         "s2s",
			Order:            40,
			NoOutput:         true,
		},
		"site-to-site-vpn-peers": {
			Type:               schema.CommandManipulateSingle,
			Params:             vpcRouterS2SPeersParam(),
			Usage:              "Show status of site-to-site IPSec VPN peers",
			TableType:          output.TableSimple,
			TableColumnDefines: vpcRouterS2SPeersColumns(),
			UseCustomCommand:   true,
			Category:           "s2s",
			Order:              50,
			NeedlessConfirm:    true,
		},
		"static-route-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             vpcRouterStaticRouteInfoParam(),
			Aliases:            []string{"static-route-list"},
			Usage:              "Show information of static-routes",
			TableType:          output.TableSimple,
			TableColumnDefines: vpcRouterStaticRouteListColumns(),
			UseCustomCommand:   true,
			Category:           "routing",
			Order:              10,
			NeedlessConfirm:    true,
		},
		"static-route-add": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterStaticRouteAddParam(),
			Usage:            "Add static-route",
			UseCustomCommand: true,
			Category:         "routing",
			Order:            20,
			NoOutput:         true,
		},
		"static-route-update": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterStaticRouteUpdateParam(),
			Usage:            "Update static-route",
			UseCustomCommand: true,
			Category:         "routing",
			Order:            30,
			NoOutput:         true,
		},
		"static-route-delete": {
			Type:             schema.CommandManipulateSingle,
			Params:           vpcRouterStaticRouteDeleteParam(),
			Usage:            "Delete static-route",
			UseCustomCommand: true,
			Category:         "routing",
			Order:            40,
			NoOutput:         true,
		},
		"monitor": {
			Type:               schema.CommandRead,
			Params:             vpcRouterMonitorParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: vpcRouterMonitorColumns(),
			UseCustomCommand:   true,
			Order:              10,
			Category:           "monitor",
		},
		"logs": {
			Type:             schema.CommandRead,
			Params:           vpcRouterLogParam(),
			UseCustomCommand: true,
			Order:            20,
			Category:         "monitor",
			NoOutput:         true,
		},
	}

	return &schema.Resource{
		Commands:          commands,
		ResourceCategory:  CategoryAppliance,
		CommandCategories: VPCRouterCommandCategories,
	}
}

var VPCRouterCommandCategories = []schema.Category{
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
		Key:         "nic",
		DisplayName: "Network Interface Management",
		Order:       30,
	},
	{
		Key:         "snat",
		DisplayName: "Static NAT Setting Management",
		Order:       40,
	},
	{
		Key:         "rnat",
		DisplayName: "Port Forward Setting Management",
		Order:       45,
	},
	{
		Key:         "fw",
		DisplayName: "Firewall Setting Management",
		Order:       50,
	},
	{
		Key:         "dhcp-server",
		DisplayName: "DHCP Server Management",
		Order:       60,
	},
	{
		Key:         "dhcp-static-mapping",
		DisplayName: "DHCP Static Map Setting Management",
		Order:       65,
	},
	{
		Key:         "remote-access",
		DisplayName: "RemoteAccess(VPN) Setting Management",
		Order:       70,
	},
	{
		Key:         "s2s",
		DisplayName: "Site to Site IPSec VPN Management",
		Order:       80,
	},
	{
		Key:         "routing",
		DisplayName: "Static Route Management",
		Order:       90,
	},
	{
		Key:         "monitor",
		DisplayName: "Monitoring",
		Order:       500,
	},
	{
		Key:         "other",
		DisplayName: "Other",
		Order:       1000,
	},
}

var vpcRouterCreateParamCategories = []schema.Category{
	{
		Key:         "router",
		DisplayName: "VPCRouter options",
		Order:       10,
	},
	{
		Key:         "network",
		DisplayName: "Network options",
		Order:       20,
	},
	{
		Key:         "operation",
		DisplayName: "Operation options",
		Order:       30,
	},
	{
		Key:         "common",
		DisplayName: "Common options",
		Order:       40,
	},
}

var vpcRouterUpdateParamCategories = []schema.Category{
	{
		Key:         "router",
		DisplayName: "VPCRouter options",
		Order:       10,
	},
	{
		Key:         "common",
		DisplayName: "Common options",
		Order:       40,
	},
}

var vpcRouterDeleteParamCategories = []schema.Category{
	{
		Key:         "router",
		DisplayName: "VPCRouter options",
		Order:       10,
	},
}

func vpcRouterListParam() map[string]*schema.Parameter {
	return mergeParameterMap(CommonListParam, paramTagsCond)
}

func vpcRouterListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{
			Name:    "Power",
			Sources: []string{"Instance.Status"},
		},
		{
			Name:    "VRID",
			Sources: []string{"Settings.Router.VRID"},
		},
		{
			Name:    "Plan",
			Sources: []string{"Plan.ID"},
			ValueMapping: []map[string]string{
				{
					"1": "standard",
					"2": "premium",
					"3": "highspec1600",
					"4": "highspec4000",
				},
			},
		},
		{
			Name: "HA",
			FormatFunc: func(values map[string]string) string {
				if plan, ok := values["Plan.ID"]; ok && plan != "1" {
					return "true"
				}
				return "false"
			},
		},
		{
			Name: "Internet",
			FormatFunc: func(values map[string]string) string {
				if enabled, ok := values["Settings.Router.InternetConnection.Enabled"]; ok && enabled == "False" {
					return "false"
				}
				return "true"
			},
		},
		{
			Name: "IPAddress",
			FormatFunc: func(values map[string]string) string {
				if plan, ok := values["Plan.ID"]; ok {
					format := "%s/%s"
					switch plan {
					case "1": // standard plan
						return fmt.Sprintf(format,
							values["Interfaces.0.IPAddress"],
							values["Interfaces.0.Switch.Subnet.NetworkMaskLen"],
						)
					default: // other plan
						return fmt.Sprintf(format,
							values["Settings.Router.Interfaces.0.VirtualIPAddress"],
							values["Interfaces.0.Switch.Subnet.NetworkMaskLen"],
						)
					}
				}

				return ""
			},
		},
	}
}

func vpcRouterInterfaceListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Index"},
		{Name: "Type"},
		{Name: "Switch"},
		{
			Name:    "IPAddress(or VIP)",
			Sources: []string{"IPAddress"},
		},
		{
			Name:    "IPAddress(#1)",
			Sources: []string{"IPAddress1"},
		},
		{
			Name:    "IPAddress(#2)",
			Sources: []string{"IPAddress2"},
		},
		{Name: "Alias"},
		{Name: "NetworkMaskLen"},
	}
}

func vpcRouterStaticNATListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "__ORDER__"}, // magic column name(generated on demand)
		{Name: "GlobalAddress"},
		{Name: "PrivateAddress"},
		{Name: "Description"},
	}
}

func vpcRouterPortForwardingListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "__ORDER__"}, // magic column name(generated on demand)
		{Name: "Protocol"},
		{Name: "GlobalPort"},
		{Name: "PrivateAddress"},
		{Name: "PrivatePort"},
		{Name: "Description"},
	}
}

func vpcRouterFirewallListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{
			Name:    "NIC",
			Sources: []string{"Interface"},
			Format:  "eth%s",
		},
		{Name: "__ORDER__"}, // magic column name(generated on demand)
		{Name: "Protocol"},
		{Name: "SourceNetwork"},
		{Name: "SourcePort"},
		{Name: "DestinationNetwork"},
		{Name: "DestinationPort"},
		{Name: "Action"},
		{Name: "Logging"},
		{Name: "Description"},
	}
}

func vpcRouterDHCPServerListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{
			Name: "#",
			FormatFunc: func(values map[string]string) string {
				if nic, ok := values["Interface"]; ok {
					return strings.Replace(nic, "eth", "", -1)
				}
				return ""
			},
		},
		{Name: "Interface"},
		{Name: "RangeStart"},
		{Name: "RangeStop"},
		{Name: "DNSServerList"},
	}
}

func vpcRouterDHCPStaticMappingListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "__ORDER__"}, // magic column name(generated on demand)
		{Name: "MACAddress"},
		{Name: "IPAddress"},
	}
}

func vpcRouterPPTPServerListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Enabled"},
		{Name: "RangeStart"},
		{Name: "RangeStop"},
	}
}

func vpcRouterL2TPServerListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Enabled"},
		{Name: "RangeStart"},
		{Name: "RangeStop"},
		{Name: "PreSharedSecret"},
	}
}

func vpcRouterUserListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "__ORDER__"}, // magic column name(generated on demand)
		{Name: "UserName"},
		{Name: "Password"},
	}
}

func vpcRouterS2SListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "__ORDER__"}, // magic column name(generated on demand)
		{Name: "Peer"},
		{Name: "RemoteID"},
		{Name: "PreSharedSecret"},
		{
			Name:    "Routes",
			Sources: []string{"RoutesJoined"},
		},
		{
			Name:    "LocalPrefix",
			Sources: []string{"LocalPrefixJoined"},
		},
	}
}

func vpcRouterS2SPeersColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "__ORDER__"}, // magic column name(generated on demand)
		{Name: "Peer"},
		{Name: "Status"},
	}
}

func vpcRouterStaticRouteListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "__ORDER__"}, // magic column name(generated on demand)
		{Name: "Prefix"},
		{Name: "NextHop"},
	}
}

func vpcRouterMonitorColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Index"},
		{Name: "Type"},
		{Name: "Key"},
		{Name: "TimeStamp"},
		{Name: "UnixTime"},
		{Name: "Receive"},
		{Name: "Send"},
	}
}

func vpcRouterDetailIncludes() []string {
	return []string{}
}

func vpcRouterDetailExcludes() []string {
	return []string{}
}

var allowVPCRouterPlans = []string{"standard", "premium", "highspec", "highspec1600", "highspec4000"}

func vpcRouterCreateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"plan": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Required:     true,
			DefaultValue: "standard",
			Description:  fmt.Sprintf("set plan[%s]", strings.Join(allowVPCRouterPlans, "/")),
			ValidateFunc: validateInStrValues(allowVPCRouterPlans...),
			Category:     "router",
			Order:        10,
		},
		"switch-id": {
			Type:         schema.TypeId,
			HandlerType:  schema.HandlerNoop,
			Description:  "set connect switch ID",
			ValidateFunc: validateSakuraID(),
			Category:     "router",
			Order:        20,
		},
		"vrid": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"VRID"},
			Description:  "set VRID",
			DefaultValue: 1,
			Required:     true,
			Category:     "router",
			Order:        30,
		},
		"vip": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set virtual ipddress()",
			ValidateFunc: validateIPv4Address(),
			Category:     "network",
			Order:        10,
		},
		"ipaddress1": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"ip1"},
			Description:  "set ipaddress(#1)",
			ValidateFunc: validateIPv4Address(),
			Category:     "network",
			Order:        20,
		},
		"ipaddress2": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"ip2"},
			Description:  "set ipaddress(#2)",
			ValidateFunc: validateIPv4Address(),
			Category:     "network",
			Order:        30,
		},
		"disable-internet-connection": {
			Type:         schema.TypeBool,
			HandlerType:  schema.HandlerNoop,
			DefaultValue: false,
			Description:  "disable internet connection from VPCRouter",
			Category:     "network",
			Order:        35,
		},
		"boot-after-create": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "boot after create",
			Category:    "operation",
			Order:       10,
		},
		"name": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set resource display name",
			Required:     true,
			ValidateFunc: validateStrLen(1, 64),
			Category:     "common",
			Order:        510,
		},
		"description": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set resource description",
			Aliases:      []string{"desc"},
			ValidateFunc: validateStrLen(0, 254),
			Category:     "common",
			Order:        520,
		},
		"tags": {
			Type:         schema.TypeStringList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set resource tags",
			ValidateFunc: validateStringSlice(validateStrLen(1, 32)),
			Category:     "common",
			Order:        530,
		},
		"icon-id": {
			Type:         schema.TypeId,
			HandlerType:  schema.HandlerNoop,
			Description:  "set Icon ID",
			ValidateFunc: validateSakuraID(),
			Category:     "common",
			Order:        540,
		},
	}
}

func vpcRouterReadParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func vpcRouterUpdateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"syslog-host": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set syslog host IPAddress",
			ValidateFunc: validateIPv4Address(),
			Category:     "router",
			Order:        10,
		},
		"internet-connection": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set internet connection from VPCRouter",
			Category:    "router",
			Order:       20,
		},
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func vpcRouterDeleteParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"force": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"f"},
			Description: "forced-shutdown flag if server is running",
			Category:    "router",
			Order:       10,
		},
	}
}

func vpcRouterEnableInternetParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func vpcRouterPowerOnParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func vpcRouterPowerOffParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func vpcRouterResetParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func vpcRouterWaitForParams() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func vpcRouterInterfaceInfoParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func vpcRouterInterfaceConnectParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"interface": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "index of target private-interface",
			Required:     true,
			ValidateFunc: validateInStrValues("1", "2", "3", "4", "5", "6", "7"),
			Category:     "interface",
			Order:        10,
		},
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
			Aliases:      []string{"ip", "vip"},
			Description:  "set (virtual)ipaddress",
			ValidateFunc: validateIPv4Address(),
			Required:     true,
			Category:     "network",
			Order:        10,
		},
		"ipaddress1": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"ip1"},
			Description:  "set ipaddress(#1)",
			ValidateFunc: validateIPv4Address(),
			Category:     "network",
			Order:        20,
		},
		"ipaddress2": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"ip2"},
			Description:  "set ipaddress(#2)",
			ValidateFunc: validateIPv4Address(),
			Category:     "network",
			Order:        30,
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
		"with-reboot": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "reboot after connect",
			Category:    "operation",
			Order:       10,
		},
	}
}

func vpcRouterInterfaceUpdateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"interface": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "index of target interface",
			Required:     true,
			ValidateFunc: validateInStrValues("0", "1", "2", "3", "4", "5", "6", "7"),
			Category:     "interface",
			Order:        10,
		},
		"switch-id": {
			Type:         schema.TypeId,
			HandlerType:  schema.HandlerNoop,
			Description:  "set connect switch ID",
			ValidateFunc: validateSakuraID(),
			Category:     "interface",
			Order:        20,
		},
		"ipaddress": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"ip", "vip"},
			Description:  "set (virtual)ipaddress",
			ValidateFunc: validateIPv4Address(),
			Category:     "network",
			Order:        10,
		},
		"ipaddress1": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"ip1"},
			Description:  "set ipaddress(#1)",
			ValidateFunc: validateIPv4Address(),
			Category:     "network",
			Order:        20,
		},
		"ipaddress2": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"ip2"},
			Description:  "set ipaddress(#2)",
			ValidateFunc: validateIPv4Address(),
			Category:     "network",
			Order:        30,
		},
		"alias": {
			Type:         schema.TypeStringList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set ip aliases",
			ValidateFunc: validateStringSlice(validateIPv4Address()),
			Category:     "network",
			Order:        40,
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
		"with-reboot": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "reboot after connect",
			Category:    "operation",
			Order:       10,
		},
	}
}

func vpcRouterInterfaceDisconnectParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"interface": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "index of target private-interface",
			Required:     true,
			ValidateFunc: validateInStrValues("1", "2", "3", "4", "5", "6", "7"),
			Category:     "interface",
			Order:        10,
		},
		"with-reboot": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "reboot after connect",
			Category:    "operation",
			Order:       10,
		},
	}
}

func vpcRouterStaticNATInfoParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func vpcRouterStaticNATAddParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"global": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"global-ip", "global-address"},
			Description:  "set global ipaddress",
			ValidateFunc: validateIPv4Address(),
			Required:     true,
			Category:     "Static-NAT",
			Order:        10,
		},
		"private": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"private-ip", "private-address"},
			Description:  "set private ipaddress",
			ValidateFunc: validateIPv4Address(),
			Required:     true,
			Category:     "Static-NAT",
			Order:        20,
		},
		"description": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set description",
			Aliases:      []string{"desc"},
			ValidateFunc: validateStrLen(0, 512),
			Category:     "Static-NAT",
			Order:        30,
		},
	}
}

func vpcRouterStaticNATUpdateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target static NAT",
			Required:    true,
			Category:    "Static-NAT",
			Order:       1,
		},
		"global": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"global-ip", "global-address"},
			Description:  "set global ipaddress",
			ValidateFunc: validateIPv4Address(),
			Category:     "Static-NAT",
			Order:        10,
		},
		"private": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"private-ip", "private-address"},
			Description:  "set private ipaddress",
			ValidateFunc: validateIPv4Address(),
			Category:     "Static-NAT",
			Order:        20,
		},
		"description": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set description",
			Aliases:      []string{"desc"},
			ValidateFunc: validateStrLen(0, 512),
			Category:     "Static-NAT",
			Order:        30,
		},
	}
}

func vpcRouterStaticNATDeleteParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target static NAT",
			Required:    true,
			Category:    "Static-NAT",
			Order:       1,
		},
	}
}

func vpcRouterPortForwardingInfoParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func vpcRouterPortForwardingAddParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"protocol": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target protocol",
			ValidateFunc: validateInStrValues("tcp", "udp"),
			Required:     true,
			Category:     "Port-Forwarding",
			Order:        10,
		},
		"global-port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set global ipaddress",
			ValidateFunc: validateIntRange(1, 65535),
			Required:     true,
			Category:     "Port-Forwarding",
			Order:        20,
		},
		"private-ipaddress": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"private-ip", "private-address"},
			Description:  "set private ipaddress",
			ValidateFunc: validateIPv4Address(),
			Required:     true,
			Category:     "Port-Forwarding",
			Order:        30,
		},
		"private-port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set private ipaddress",
			ValidateFunc: validateIntRange(1, 65535),
			Required:     true,
			Category:     "Port-Forwarding",
			Order:        40,
		},
		"description": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set description",
			Aliases:      []string{"desc"},
			ValidateFunc: validateStrLen(0, 512),
			Category:     "Port-Forwarding",
			Order:        50,
		},
	}
}

func vpcRouterPortForwardingUpdateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target PortForward",
			Required:    true,
			Category:    "Port-Forwarding",
			Order:       1,
		},
		"protocol": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target protocol",
			ValidateFunc: validateInStrValues("tcp", "udp"),
			Category:     "Port-Forwarding",
			Order:        10,
		},
		"global-port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set global ipaddress",
			ValidateFunc: validateIntRange(1, 65535),
			Category:     "Port-Forwarding",
			Order:        20,
		},
		"private-ipaddress": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"private-ip", "private-address"},
			Description:  "set private ipaddress",
			ValidateFunc: validateIPv4Address(),
			Category:     "Port-Forwarding",
			Order:        30,
		},
		"private-port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set private ipaddress",
			ValidateFunc: validateIntRange(1, 65535),
			Category:     "Port-Forwarding",
			Order:        40,
		},
		"description": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set description",
			Aliases:      []string{"desc"},
			ValidateFunc: validateStrLen(0, 512),
			Category:     "Port-Forwarding",
			Order:        50,
		},
	}
}

func vpcRouterPortForwardingDeleteParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target PortForward",
			Required:    true,
			Category:    "Port-Forwarding",
			Order:       1,
		},
	}
}

func vpcRouterFirewallInfoParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"interface": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target NIC index",
			ValidateFunc: validateIntRange(0, 7),
			DefaultValue: 0,
			Category:     "Firewall",
			Order:        1,
		},
		"direction": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target direction[send/receive]",
			ValidateFunc: validateInStrValues("send", "receive"),
			Required:     true,
			DefaultValue: "receive",
			Category:     "Firewall",
			Order:        2,
		},
	}
}

func vpcRouterFirewallAddParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"interface": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target NIC index",
			ValidateFunc: validateIntRange(0, 7),
			DefaultValue: 0,
			Category:     "Firewall",
			Order:        1,
		},
		"direction": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target direction[send/receive]",
			ValidateFunc: validateInStrValues("send", "receive"),
			Required:     true,
			DefaultValue: "receive",
			Category:     "Firewall",
			Order:        2,
		},
		"protocol": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target protocol",
			ValidateFunc: validateInStrValues("tcp", "udp", "icmp", "ip"),
			Required:     true,
			Category:     "Firewall",
			Order:        10,
		},
		"source-network": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set source ipaddress or network address",
			ValidateFunc: validateIPv4AddressWithPrefixOption(),
			Category:     "Firewall",
			Order:        20,
		},
		"source-port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set source port",
			ValidateFunc: validateIntRange(1, 65535),
			Category:     "Firewall",
			Order:        30,
		},
		"destination-network": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"dest-network"},
			Description:  "set destination ipaddress or network address",
			ValidateFunc: validateIPv4AddressWithPrefixOption(),
			Category:     "Firewall",
			Order:        40,
		},
		"destination-port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"dest-port"},
			Description:  "set destination port",
			ValidateFunc: validateIntRange(1, 65535),
			Category:     "Firewall",
			Order:        50,
		},
		"action": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set action[allow/deny]",
			ValidateFunc: validateInStrValues("allow", "deny"),
			Required:     true,
			DefaultValue: "deny",
			Category:     "Firewall",
			Order:        60,
		},
		"enable-logging": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "enable logging",
			Category:    "Firewall",
			Order:       70,
		},
		"description": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set description",
			Aliases:      []string{"desc"},
			ValidateFunc: validateStrLen(0, 512),
			Category:     "Firewall",
			Order:        80,
		},
	}
}

func vpcRouterFirewallUpdateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"interface": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target NIC index",
			ValidateFunc: validateIntRange(0, 7),
			DefaultValue: 0,
			Category:     "Firewall",
			Order:        1,
		},
		"direction": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target direction[send/receive]",
			ValidateFunc: validateInStrValues("send", "receive"),
			Required:     true,
			DefaultValue: "receive",
			Category:     "Firewall",
			Order:        2,
		},
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target Firewall rule",
			Required:    true,
			Category:    "Firewall",
			Order:       3,
		},
		"protocol": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target protocol",
			ValidateFunc: validateInStrValues("tcp", "udp", "icmp", "ip"),
			Category:     "Firewall",
			Order:        10,
		},
		"source-network": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set source ipaddress or network address",
			ValidateFunc: validateIPv4AddressWithPrefixOption(),
			Category:     "Firewall",
			Order:        20,
		},
		"source-port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set source port",
			ValidateFunc: validateIntRange(1, 65535),
			Category:     "Firewall",
			Order:        30,
		},
		"destination-network": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"dest-network"},
			Description:  "set destination ipaddress or network address",
			ValidateFunc: validateIPv4AddressWithPrefixOption(),
			Category:     "Firewall",
			Order:        40,
		},
		"destination-port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"dest-port"},
			Description:  "set destination port",
			ValidateFunc: validateIntRange(1, 65535),
			Category:     "Firewall",
			Order:        50,
		},
		"action": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set action[allow/deny]",
			ValidateFunc: validateInStrValues("allow", "deny"),
			DefaultValue: "deny",
			Category:     "Firewall",
			Order:        60,
		},
		"enable-logging": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "enable logging",
			Category:    "Firewall",
			Order:       70,
		},
		"description": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set description",
			Aliases:      []string{"desc"},
			ValidateFunc: validateStrLen(0, 512),
			Category:     "Firewall",
			Order:        80,
		},
	}
}

func vpcRouterFirewallDeleteParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"interface": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target NIC index",
			ValidateFunc: validateIntRange(0, 7),
			DefaultValue: 0,
			Category:     "Firewall",
			Order:        1,
		},
		"direction": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target direction[send/receive]",
			ValidateFunc: validateInStrValues("send", "receive"),
			Required:     true,
			DefaultValue: "receive",
			Category:     "Firewall",
			Order:        2,
		},
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target Firewall rule",
			Required:    true,
			Category:    "Firewall",
			Order:       3,
		},
	}
}

func vpcRouterDHCPServerInfoParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func vpcRouterDHCPServerAddParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"interface": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target NIC(private NIC index)",
			ValidateFunc: validateIntRange(1, 7),
			Required:     true,
			Category:     "DHCP-Server",
			Order:        1,
		},
		"range-start": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set DHCP IPAddress Range(start)",
			ValidateFunc: validateIPv4Address(),
			Required:     true,
			Category:     "DHCP-Server",
			Order:        10,
		},
		"range-stop": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"range-end"},
			Description:  "set DHCP IPAddress Range(stop)",
			ValidateFunc: validateIPv4Address(),
			Required:     true,
			Category:     "DHCP-Server",
			Order:        20,
		},
		"dns-servers": {
			Type:         schema.TypeStringList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set DNS Server IPAddress",
			ValidateFunc: validateStringSlice(validateIPv4Address()),
			Category:     "DHCP-Server",
			Order:        30,
		},
	}
}

func vpcRouterDHCPServerUpdateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"interface": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target NIC(private NIC index)",
			ValidateFunc: validateIntRange(1, 7),
			Required:     true,
			Category:     "DHCP-Server",
			Order:        1,
		},
		"range-start": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set DHCP IPAddress Range(start)",
			ValidateFunc: validateIPv4Address(),
			Category:     "DHCP-Server",
			Order:        10,
		},
		"range-stop": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"range-end"},
			Description:  "set DHCP IPAddress Range(stop)",
			ValidateFunc: validateIPv4Address(),
			Category:     "DHCP-Server",
			Order:        20,
		},
		"dns-servers": {
			Type:         schema.TypeStringList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set DNS Server IPAddress",
			ValidateFunc: validateStringSlice(validateIPv4Address()),
			Category:     "DHCP-Server",
			Order:        30,
		},
	}
}

func vpcRouterDHCPServerDeleteParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"interface": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target NIC(private NIC index)",
			ValidateFunc: validateIntRange(1, 7),
			Required:     true,
			Category:     "DHCP-Server",
			Order:        1,
		},
	}
}

func vpcRouterDHCPStaticMappingInfoParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func vpcRouterDHCPStaticMappingAddParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"macaddress": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"mac"},
			Description:  "set mac address",
			ValidateFunc: validateMACAddress(),
			Required:     true,
			Category:     "DHCP-Static-Mapping",
			Order:        10,
		},
		"ipaddress": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"ip"},
			Description:  "set ipaddress",
			ValidateFunc: validateIPv4Address(),
			Required:     true,
			Category:     "DHCP-Static-Mapping",
			Order:        20,
		},
	}
}

func vpcRouterDHCPStaticMappingUpdateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target DHCP static mapping",
			Required:    true,
			Category:    "DHCP-Static-Mapping",
			Order:       1,
		},
		"macaddress": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"mac"},
			Description:  "set mac address",
			ValidateFunc: validateMACAddress(),
			Category:     "DHCP-Static-Mapping",
			Order:        10,
		},
		"ipaddress": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"ip"},
			Description:  "set ipaddress",
			ValidateFunc: validateIPv4Address(),
			Category:     "DHCP-Static-Mapping",
			Order:        20,
		},
	}
}

func vpcRouterDHCPStaticMappingDeleteParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target DHCP static mapping",
			Required:    true,
			Category:    "DHCP-Static-Mapping",
			Order:       1,
		},
	}
}

func vpcRouterPPTPServerInfoParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func vpcRouterPPTPServerUpdateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"disabled": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "enable/disable PPTP server",
			Category:    "PPTP",
			Order:       10,
		},
		"range-start": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set IPAddress Range(start)",
			ValidateFunc: validateIPv4Address(),
			Category:     "PPTP",
			Order:        20,
		},
		"range-stop": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"range-end"},
			Description:  "set IPAddress Range(stop)",
			ValidateFunc: validateIPv4Address(),
			Category:     "PPTP",
			Order:        30,
		},
	}
}

func vpcRouterL2TPServerInfoParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func vpcRouterL2TPServerUpdateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"disabled": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "enable/disable PPTP server",
			Category:    "L2TP-IPSec",
			Order:       10,
		},
		"range-start": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set IPAddress Range(start)",
			ValidateFunc: validateIPv4Address(),
			Category:     "L2TP-IPSec",
			Order:        20,
		},
		"range-stop": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"range-end"},
			Description:  "set IPAddress Range(stop)",
			ValidateFunc: validateIPv4Address(),
			Category:     "L2TP-IPSec",
			Order:        30,
		},
		"pre-shared-secret": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set PreSharedSecret",
			ValidateFunc: validateStrLen(0, 40),
			Category:     "L2TP-IPSec",
			Order:        40,
		},
	}
}

func vpcRouterUserInfoParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func vpcRouterUserAddParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"username": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set remote-access username",
			Aliases:      []string{"user"},
			Required:     true,
			ValidateFunc: validateStrLen(0, 20),
			Category:     "user",
			Order:        10,
		},
		"password": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set remote-access password",
			Aliases:      []string{"pass"},
			Required:     true,
			ValidateFunc: validateStrLen(0, 20),
			Category:     "user",
			Order:        20,
		},
	}
}

func vpcRouterUserUpdateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target remote-access user",
			Required:    true,
			Category:    "user",
			Order:       1,
		},
		"username": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set remote-access username",
			Aliases:      []string{"user"},
			ValidateFunc: validateStrLen(0, 20),
			Category:     "user",
			Order:        10,
		},
		"password": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set remote-access password",
			Aliases:      []string{"pass"},
			ValidateFunc: validateStrLen(0, 20),
			Category:     "user",
			Order:        20,
		},
	}
}

func vpcRouterUserDeleteParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target remote-access user",
			Required:    true,
			Category:    "user",
			Order:       1,
		},
	}
}

func vpcRouterS2SInfoParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func vpcRouterS2SAddParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"peer": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set peer address",
			Required:     true,
			ValidateFunc: validateIPv4Address(),
			Category:     "Site-To-Site IPSec VPN",
			Order:        10,
		},
		"remote-id": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set remote-id",
			Required:    true,
			Category:    "Site-To-Site IPSec VPN",
			Order:       20,
		},
		"pre-shared-secret": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set pre-shared-secret",
			Required:     true,
			ValidateFunc: validateStrLen(0, 40),
			Category:     "Site-To-Site IPSec VPN",
			Order:        30,
		},
		"routes": {
			Type:         schema.TypeStringList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set route list",
			Required:     true,
			ValidateFunc: validateStringSlice(validateIPv4AddressWithPrefix()),
			Category:     "Site-To-Site IPSec VPN",
			Order:        40,
		},
		"local-prefix": {
			Type:         schema.TypeStringList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set local prefix list",
			Required:     true,
			ValidateFunc: validateStringSlice(validateIPv4AddressWithPrefix()),
			Category:     "Site-To-Site IPSec VPN",
			Order:        50,
		},
	}
}

func vpcRouterS2SUpdateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target remote-access user",
			Required:    true,
			Category:    "Site-To-Site IPSec VPN",
			Order:       1,
		},
		"peer": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set peer address",
			ValidateFunc: validateIPv4Address(),
			Category:     "Site-To-Site IPSec VPN",
			Order:        10,
		},
		"remote-id": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set remote-id",
			Category:    "Site-To-Site IPSec VPN",
			Order:       20,
		},
		"pre-shared-secret": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set pre-shared-secret",
			ValidateFunc: validateStrLen(0, 40),
			Category:     "Site-To-Site IPSec VPN",
			Order:        30,
		},
		"routes": {
			Type:         schema.TypeStringList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set route list",
			ValidateFunc: validateStringSlice(validateIPv4AddressWithPrefix()),
			Category:     "Site-To-Site IPSec VPN",
			Order:        40,
		},
		"local-prefix": {
			Type:         schema.TypeStringList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set local prefix list",
			ValidateFunc: validateStringSlice(validateIPv4AddressWithPrefix()),
			Category:     "Site-To-Site IPSec VPN",
			Order:        50,
		},
	}
}

func vpcRouterS2SDeleteParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target remote-access user",
			Required:    true,
			Category:    "Site-To-Site IPSec VPN",
			Order:       1,
		},
	}
}

func vpcRouterS2SPeersParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func vpcRouterStaticRouteInfoParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func vpcRouterStaticRouteAddParam() map[string]*schema.Parameter {
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

func vpcRouterStaticRouteUpdateParam() map[string]*schema.Parameter {
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

func vpcRouterStaticRouteDeleteParam() map[string]*schema.Parameter {
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

func vpcRouterMonitorParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"interface": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "index of target interface",
			Required:     true,
			DefaultValue: "0",
			ValidateFunc: validateInStrValues("0", "1", "2", "3", "4", "5", "6", "7"),
			Category:     "monitor",
			Order:        10,
		},
		"start": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set start-time",
			ValidateFunc: validateDateTimeString(),
			Category:     "monitor",
			Order:        20,
		},
		"end": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set end-time",
			ValidateFunc: validateDateTimeString(),
			Category:     "monitor",
			Order:        30,
		},
		"key-format": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set monitoring value key-format",
			DefaultValue: "sakuracloud.vpcrouter.{{.ID}}.nic.{{.Index}}",
			Required:     true,
			Category:     "monitor",
			Order:        40,
		},
	}
}

var AllowVPCRouterLogNames = []string{"all", "vpn", "firewall-send", "firewall-receive"}

func vpcRouterLogParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"log-name": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"name"},
			Description:  "set target logfile name",
			DefaultValue: "all",
			ValidateFunc: validateInStrValues(AllowVPCRouterLogNames...),
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

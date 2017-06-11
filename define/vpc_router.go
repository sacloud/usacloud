package define

import (
	"fmt"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
	"strings"
)

func VPCRouterResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find"},
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
			Category:           "monitor",
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

func vpcRouterListParam() map[string]*schema.Schema {
	return mergeParameterMap(CommonListParam, paramTagsCond)
}

func vpcRouterListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
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
					"3": "highspec",
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

var allowVPCRouterPlans = []string{"standard", "premium", "highspec"}

func vpcRouterCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"plan": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Required:     true,
			DefaultValue: "standard",
			Description:  "set plan[standard/premium/highspec]",
			ValidateFunc: validateInStrValues(allowVPCRouterPlans...),
			CompleteFunc: completeInStrValues(allowVPCRouterPlans...),
			Category:     "router",
			Order:        10,
		},
		"switch-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set connect switch ID",
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeSwitchID(),
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
		"boot-after-create": {
			Type:         schema.TypeBool,
			HandlerType:  schema.HandlerNoop,
			Description:  "boot after create",
			DefaultValue: false,
			Category:     "operation",
			Order:        10,
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
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set Icon ID",
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeIconID(),
			Category:     "common",
			Order:        540,
		},
	}
}

func vpcRouterReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func vpcRouterUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"syslog-host": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set syslog host IPAddress",
			ValidateFunc: validateIPv4Address(),
			Category:     "router",
			Order:        10,
		},
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func vpcRouterDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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

func vpcRouterPowerOnParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func vpcRouterPowerOffParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func vpcRouterResetParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func vpcRouterWaitForParams() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func vpcRouterInterfaceInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func vpcRouterInterfaceConnectParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "index of target private-interface",
			Required:     true,
			ValidateFunc: validateInStrValues("1", "2", "3", "4", "5", "6", "7"),
			Category:     "interface",
			Order:        10,
		},
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

func vpcRouterInterfaceUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "index of target interface",
			Required:     true,
			ValidateFunc: validateInStrValues("0", "1", "2", "3", "4", "5", "6", "7"),
			Category:     "interface",
			Order:        10,
		},
		"switch-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set connect switch ID",
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeSwitchID(),
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

func vpcRouterInterfaceDisconnectParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
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

func vpcRouterStaticNATInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func vpcRouterStaticNATAddParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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

func vpcRouterStaticNATUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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

func vpcRouterStaticNATDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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

func vpcRouterPortForwardingInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func vpcRouterPortForwardingAddParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"protocol": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target protocol",
			ValidateFunc: validateInStrValues("tcp", "udp"),
			CompleteFunc: completeInStrValues("tcp", "udp"),
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

func vpcRouterPortForwardingUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target static NAT",
			Required:    true,
			Category:    "Port-Forwarding",
			Order:       1,
		},
		"protocol": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target protocol",
			ValidateFunc: validateInStrValues("tcp", "udp"),
			CompleteFunc: completeInStrValues("tcp", "udp"),
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

func vpcRouterPortForwardingDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target static NAT",
			Required:    true,
			Category:    "Port-Forwarding",
			Order:       1,
		},
	}
}

func vpcRouterFirewallInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"direction": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target direction[send/receive]",
			ValidateFunc: validateInStrValues("send", "receive"),
			CompleteFunc: completeInStrValues("send", "receive"),
			Required:     true,
			DefaultValue: "receive",
			Category:     "Firewall",
			Order:        1,
		},
	}
}

func vpcRouterFirewallAddParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"direction": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target direction[send/receive]",
			ValidateFunc: validateInStrValues("send", "receive"),
			CompleteFunc: completeInStrValues("send", "receive"),
			Required:     true,
			DefaultValue: "receive",
			Category:     "Firewall",
			Order:        1,
		},
		"protocol": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target protocol",
			ValidateFunc: validateInStrValues("tcp", "udp", "icmp", "ip"),
			CompleteFunc: completeInStrValues("tcp", "udp", "icmp", "ip"),
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
			CompleteFunc: completeInStrValues("allow", "deny"),
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

func vpcRouterFirewallUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"direction": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target direction[send/receive]",
			ValidateFunc: validateInStrValues("send", "receive"),
			CompleteFunc: completeInStrValues("send", "receive"),
			Required:     true,
			DefaultValue: "receive",
			Category:     "Firewall",
			Order:        1,
		},
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target static NAT",
			Required:    true,
			Category:    "Firewall",
			Order:       2,
		},
		"protocol": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target protocol",
			ValidateFunc: validateInStrValues("tcp", "udp", "icmp", "ip"),
			CompleteFunc: completeInStrValues("tcp", "udp", "icmp", "ip"),
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
			CompleteFunc: completeInStrValues("allow", "deny"),
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

func vpcRouterFirewallDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"direction": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target direction[send/receive]",
			ValidateFunc: validateInStrValues("send", "receive"),
			CompleteFunc: completeInStrValues("send", "receive"),
			Required:     true,
			DefaultValue: "receive",
			Category:     "Firewall",
			Order:        1,
		},
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target static NAT",
			Required:    true,
			Category:    "Firewall",
			Order:       2,
		},
	}
}

func vpcRouterDHCPServerInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func vpcRouterDHCPServerAddParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
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
	}
}

func vpcRouterDHCPServerUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
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
	}
}

func vpcRouterDHCPServerDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
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

func vpcRouterDHCPStaticMappingInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func vpcRouterDHCPStaticMappingAddParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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

func vpcRouterDHCPStaticMappingUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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

func vpcRouterDHCPStaticMappingDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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

func vpcRouterPPTPServerInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func vpcRouterPPTPServerUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enabled": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "enable/disable PPTP server",
			ValidateFunc: validateInStrValues("true", "false"),
			CompleteFunc: completeInStrValues("true", "false"),
			Required:     true,
			Category:     "PPTP",
			Order:        10,
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

func vpcRouterL2TPServerInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func vpcRouterL2TPServerUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enabled": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "enable/disable PPTP server",
			ValidateFunc: validateInStrValues("true", "false"),
			CompleteFunc: completeInStrValues("true", "false"),
			Required:     true,
			Category:     "L2TP-IPSec",
			Order:        10,
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

func vpcRouterUserInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func vpcRouterUserAddParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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

func vpcRouterUserUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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

func vpcRouterUserDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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

func vpcRouterS2SInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func vpcRouterS2SAddParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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

func vpcRouterS2SUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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

func vpcRouterS2SDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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

func vpcRouterStaticRouteInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func vpcRouterStaticRouteAddParam() map[string]*schema.Schema {
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

func vpcRouterStaticRouteUpdateParam() map[string]*schema.Schema {
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

func vpcRouterStaticRouteDeleteParam() map[string]*schema.Schema {
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

func vpcRouterMonitorParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
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
			DefaultValue: "sakuracloud.{{.ID}}.vpcrouter",
			Required:     true,
			Category:     "monitor",
			Order:        40,
		},
	}
}

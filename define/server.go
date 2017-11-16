package define

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/libsacloud/sacloud/ostype"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func ServerResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find", "selector"},
			Params:             serverListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: serverListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"build": {
			Type:             schema.CommandCreate,
			ParamCategories:  serverBuildParamCategories,
			Params:           serverBuildParam(),
			IncludeFields:    serverDetailIncludes(),
			ExcludeFields:    serverBuildResultExcludes(),
			UseCustomCommand: true,
			Category:         "basics",
			Order:            20,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        serverReadParam(),
			IncludeFields: serverDetailIncludes(),
			ExcludeFields: serverDetailExcludes(),
			Category:      "basics",
			Order:         30,
		},
		"update": {
			Type:             schema.CommandUpdate,
			Params:           serverUpdateParam(),
			IncludeFields:    serverDetailIncludes(),
			ExcludeFields:    serverDetailExcludes(),
			UseCustomCommand: true,
			Category:         "basics",
			Order:            40,
		},
		"delete": {
			Type:             schema.CommandDelete,
			Aliases:          []string{"rm"},
			Params:           serverDeleteParam(),
			IncludeFields:    serverDetailIncludes(),
			ExcludeFields:    serverDetailExcludes(),
			UseCustomCommand: true,
			Category:         "basics",
			Order:            50,
		},
		"plan-change": {
			Type:             schema.CommandManipulateMulti,
			Params:           serverPlanChangeParam(),
			Usage:            "Change server plan(core/memory)",
			IncludeFields:    serverDetailIncludes(),
			ExcludeFields:    serverDetailExcludes(),
			UseCustomCommand: true,
			Category:         "basics",
			Order:            60,
		},
		"boot": {
			Type:             schema.CommandManipulateMulti,
			Aliases:          []string{"power-on"},
			Params:           serverPowerOnParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            10,
			NoOutput:         true,
		},
		"shutdown": {
			Type:             schema.CommandManipulateMulti,
			Aliases:          []string{"power-off"},
			Params:           serverPowerOffParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            20,
			NoOutput:         true,
		},
		"shutdown-force": {
			Type:             schema.CommandManipulateMulti,
			Aliases:          []string{"stop"},
			Params:           serverPowerOffParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            25,
			NoOutput:         true,
		},
		"reset": {
			Type:             schema.CommandManipulateMulti,
			Params:           serverResetParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            30,
			NoOutput:         true,
		},
		"wait-for-boot": {
			Type:             schema.CommandManipulateMulti,
			Params:           serverWaitForParams(),
			Usage:            "Wait until boot is completed",
			UseCustomCommand: true,
			Category:         "power",
			Order:            40,
			NoOutput:         true,
			NeedlessConfirm:  true,
		},
		"wait-for-down": {
			Type:             schema.CommandManipulateMulti,
			Params:           serverWaitForParams(),
			Usage:            "Wait until shutdown is completed",
			UseCustomCommand: true,
			Category:         "power",
			Order:            50,
			NoOutput:         true,
			NeedlessConfirm:  true,
		},
		"ssh": {
			Type:             schema.CommandManipulateSingle,
			Params:           serverSSHParam(),
			Usage:            "Connect to server by SSH",
			UseCustomCommand: true,
			Category:         "connect",
			Order:            10,
			NoOutput:         true,
			NeedlessConfirm:  true,
		},
		"ssh-exec": {
			Type:                schema.CommandManipulateMulti,
			Params:              serverSSHParam(),
			Usage:               "Execute command on server connected by SSH",
			UseCustomCommand:    true,
			SkipAfterSecondArgs: true,
			NoSelector:          true,
			Category:            "connect",
			Order:               20,
			NoOutput:            true,
			NeedlessConfirm:     true,
		},
		"scp": {
			Type:             schema.CommandCustom,
			Params:           serverSCPParam(),
			Usage:            "Copy files/directories by SSH",
			ArgsUsage:        "[ServerID:]<FROM> [ServerID:]<TO>",
			UseCustomCommand: true,
			Category:         "connect",
			Order:            30,
			NoOutput:         true,
		},
		"vnc": {
			Type:             schema.CommandManipulateMulti,
			Params:           serverVNCParam(),
			Usage:            "Open VNC client using the OS's default application",
			UseCustomCommand: true,
			Category:         "connect",
			Order:            40,
			NoOutput:         true,
			ConfirmMessage:   "open VNC client",
		},
		"vnc-info": {
			Type:             schema.CommandRead,
			Params:           serverVNCParam(),
			Usage:            "Show VNC proxy information",
			UseCustomCommand: true,
			Category:         "connect",
			Order:            45,
			NeedlessConfirm:  true,
		},
		"vnc-send": {
			Type:                schema.CommandManipulateSingle,
			Params:              serverVNCSendParam(),
			Usage:               "Send keys over VNC connection",
			UseCustomCommand:    true,
			SkipAfterSecondArgs: true,
			Category:            "connect",
			Order:               46,
		},
		"vnc-snapshot": {
			Type:                schema.CommandManipulateSingle,
			Params:              serverVNCSnapshotParam(),
			Usage:               "Capture VNC snapshot",
			UseCustomCommand:    true,
			SkipAfterSecondArgs: true,
			Category:            "connect",
			Order:               47,
		},
		"disk-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             serverDiskInfoParam(),
			Aliases:            []string{"disk-list"},
			Usage:              "Show information of disk(s) connected to server",
			TableType:          output.TableSimple,
			TableColumnDefines: diskListColumns(),
			UseCustomCommand:   true,
			Category:           "disks",
			Order:              10,
			NeedlessConfirm:    true,
		},
		"disk-connect": {
			Type:               schema.CommandManipulateSingle,
			Params:             serverDiskConnectParam(),
			Usage:              "Connect disk to server",
			TableType:          output.TableSimple,
			TableColumnDefines: diskListColumns(),
			UseCustomCommand:   true,
			Category:           "disks",
			Order:              20,
			NoOutput:           true,
		},
		"disk-disconnect": {
			Type:               schema.CommandManipulateSingle,
			Params:             serverDiskDisconnectParam(),
			Usage:              "Disconnect disk from server",
			TableType:          output.TableSimple,
			TableColumnDefines: diskListColumns(),
			UseCustomCommand:   true,
			Category:           "disks",
			Order:              30,
			NoOutput:           true,
		},
		"interface-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             serverInterfaceInfoParam(),
			Aliases:            []string{"interface-list"},
			Usage:              "Show information of NIC(s) connected to server",
			TableType:          output.TableSimple,
			TableColumnDefines: serverInterfaceListColumns(),
			UseCustomCommand:   true,
			Category:           "network",
			Order:              10,
			NeedlessConfirm:    true,
		},
		"interface-add-for-internet": {
			Type:               schema.CommandManipulateSingle,
			Params:             serverInterfaceAddForInternetParam(),
			Usage:              "Create and connect NIC connected to the internet",
			TableType:          output.TableSimple,
			TableColumnDefines: serverInterfaceListColumns(),
			UseCustomCommand:   true,
			Category:           "network",
			Order:              20,
			NoOutput:           true,
		},
		"interface-add-for-router": {
			Type:               schema.CommandManipulateSingle,
			Params:             serverInterfaceAddForRouterParam(),
			Usage:              "Create and connect NIC connected to the router",
			TableType:          output.TableSimple,
			TableColumnDefines: serverInterfaceListColumns(),
			UseCustomCommand:   true,
			Category:           "network",
			Order:              30,
			NoOutput:           true,
		},
		"interface-add-for-switch": {
			Type:               schema.CommandManipulateSingle,
			Params:             serverInterfaceAddForSwitchParam(),
			Usage:              "Create and connect NIC connected to the switch",
			TableType:          output.TableSimple,
			TableColumnDefines: serverInterfaceListColumns(),
			UseCustomCommand:   true,
			Category:           "network",
			Order:              40,
			NoOutput:           true,
		},
		"interface-add-disconnected": {
			Type:               schema.CommandManipulateSingle,
			Params:             serverInterfaceAddDisconnectedParam(),
			Usage:              "Create and connect a disconnected NIC",
			TableType:          output.TableSimple,
			TableColumnDefines: serverInterfaceListColumns(),
			UseCustomCommand:   true,
			Category:           "network",
			Order:              50,
			NoOutput:           true,
		},
		"iso-info": {
			Type:             schema.CommandManipulateSingle,
			Params:           serverISOImageInfoParam(),
			Usage:            "Show information of ISO-Image inserted to server",
			UseCustomCommand: true,
			Category:         "iso",
			Order:            10,
			NeedlessConfirm:  true,
		},
		"iso-insert": {
			Type:             schema.CommandManipulateSingle,
			Params:           serverISOImageInsertParam(),
			Usage:            "Insert ISO-Image to server",
			UseCustomCommand: true,
			Category:         "iso",
			Order:            20,
			NoOutput:         true,
		},
		"iso-eject": {
			Type:             schema.CommandManipulateSingle,
			Params:           serverISOImageEjectParam(),
			Usage:            "Eject ISO-Image from server",
			UseCustomCommand: true,
			Category:         "iso",
			Order:            30,
			NoOutput:         true,
		},
		"monitor-cpu": {
			Type:               schema.CommandRead,
			Params:             serverMonitorCPUParam(),
			Usage:              "Collect CPU monitor values",
			TableType:          output.TableSimple,
			TableColumnDefines: serverMonitorCPUColumns(),
			UseCustomCommand:   true,
			Category:           "monitor",
			Order:              10,
		},
		"monitor-nic": {
			Type:               schema.CommandRead,
			Params:             serverMonitorNICParam(),
			Usage:              "Collect NIC(s) monitor values",
			TableType:          output.TableSimple,
			TableColumnDefines: serverMonitorNICColumns(),
			UseCustomCommand:   true,
			Category:           "monitor",
			Order:              20,
		},
		"monitor-disk": {
			Type:               schema.CommandRead,
			Params:             serverMonitorDiskParam(),
			Usage:              "Collect Disk(s) monitor values",
			TableType:          output.TableSimple,
			TableColumnDefines: serverMonitorDiskColumns(),
			UseCustomCommand:   true,
			Category:           "monitor",
			Order:              30,
		},
		"maintenance-info": {
			Type:               schema.CommandList,
			Params:             serverMaintenanceInfoParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: serverMaintenanceInfoColumns(),
			UseCustomCommand:   true,
			Category:           "other",
			Order:              10,
		},
	}

	return &schema.Resource{
		Commands:          commands,
		ResourceCategory:  CategoryComputing,
		CommandCategories: serverCommandCategories,
	}
}

func serverListParam() map[string]*schema.Schema {
	return mergeParameterMap(CommonListParam, paramTagsCond)
}

func serverListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{
			Name:    "CPU",
			Sources: []string{"ServerPlan.CPU"},
		},
		{
			Name:    "Memory",
			Sources: []string{"ServerPlan.MemoryMB"},
			Format:  "%sMB",
		},
		{
			Name: "IPAddress",
			FormatFunc: func(values map[string]string) string {
				if scope, ok := values["Interfaces.0.Switch.Scope"]; ok {
					format := "%s/%s"
					switch scope {
					case "shared":
						return fmt.Sprintf(format,
							values["Interfaces.0.IPAddress"],
							values["Interfaces.0.Switch.UserSubnet.NetworkMaskLen"],
						)
					case "user":
						return fmt.Sprintf(format,
							values["Interfaces.0.UserIPAddress"],
							values["Interfaces.0.Switch.UserSubnet.NetworkMaskLen"],
						)

					}

				}

				return ""
			},
		},
		{
			Name:    "Status",
			Sources: []string{"Instance.Status"},
		},
		{
			Name: "Host",
			FormatFunc: func(values map[string]string) string {
				pHost := ""
				if id, ok := values["PrivateHost.ID"]; ok {
					pHost = fmt.Sprintf("(PrivateHost:%s)", id)
				}
				return fmt.Sprintf("%s%s", values["Instance.Host.Name"], pHost)
			},
		},
	}
}

func serverInterfaceListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{
			Name:    "IPAddress",
			Sources: []string{"IPAddress", "UserIPAddress"},
			Format:  "%s%s",
		},
		{Name: "MACAddress"},
		{
			Name:    "Gateway",
			Sources: []string{"Switch.Subnet.DefaultRoute"},
		},
		{
			Name: "Network",
			Sources: []string{
				"Switch.Subnet.NetworkAddress",
				"Switch.Subnet.NetworkMaskLen",
			},
			Format: "%s/%s",
		},
		{
			Name:    "BandWidth",
			Sources: []string{"Switch.Subnet.Internet.BandWidthMbps"},
			Format:  "%sMbps",
		},
		{
			Name:    "SwitchID",
			Sources: []string{"Switch.ID"},
			Format:  "%s",
		},
		{
			Name:    "PacketFilter",
			Sources: []string{"PacketFilter.Name", "PacketFilter.ID"},
			Format:  "%s(%s)",
		},
	}
}

func serverMaintenanceInfoColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{
			Name:    "Host",
			Sources: []string{"Instance.Host.Name"},
			Format:  "%s",
		},
		{
			Name:    "Date",
			Sources: []string{"StartDate", "EndDate"},
			Format:  "%s 〜 %s",
		},
		{Name: "InfoURL"},
	}
}

func serverDetailIncludes() []string {
	return []string{}
}

func serverDetailExcludes() []string {
	return []string{
		"Instance.CDROMStorage",
		"Instance.CDROM.Storage",
		"ServerPlan.ID",
		"ServerPlan.Description",
		"ServerPlan.ServiceClass",
		"Zone.FTPServer",
		"Disks.0.Storage",
		"Disks.1.Storage",
		"Disks.2.Storage",
		"Disks.3.Storage",
	}
}

func serverBuildResultExcludes() []string {
	return []string{
		"Disks",
		"Server.Instance.CDROMStorage",
		"Server.Instance.CDROM.Storage",
		"Server.ServerPlan.ID",
		"Server.ServerPlan.Description",
		"Server.ServerPlan.ServiceClass",
		"Server.Zone.FTPServer",
		"Server.Disks.0.Storage",
		"Server.Disks.1.Storage",
		"Server.Disks.2.Storage",
		"Server.Disks.3.Storage",
	}
}

var serverCommandCategories = []schema.Category{
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
		Key:         "connect",
		DisplayName: "SSH/SCP/VNC",
		Order:       30,
	},
	{
		Key:         "disks",
		DisplayName: "Disk(s) Management",
		Order:       40,
	},
	{
		Key:         "network",
		DisplayName: "Network Management",
		Order:       50,
	},
	{
		Key:         "iso",
		DisplayName: "ISO Image Management",
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

var serverBuildParamCategories = []schema.Category{
	{
		Key:         "server-plan",
		DisplayName: "For server-plan options",
		Order:       10,
	},
	{
		Key:         "disk",
		DisplayName: "For disk options",
		Order:       20,
	},
	{
		Key:         "iso-image",
		DisplayName: "For ISO image options",
		Order:       25,
	},
	{
		Key:         "network",
		DisplayName: "For network options",
		Order:       30,
	},
	{
		Key:         "edit-disk",
		DisplayName: "For edit-disk options",
		Order:       40,
	},
	{
		Key:         "edit-disk-network",
		DisplayName: "For edit-disk(network settings) options",
		Order:       41,
	},
	{
		Key:         "edit-disk-startup-script",
		DisplayName: "For edit-disk(startup-script) options",
		Order:       42,
	},
	{
		Key:         "edit-disk-ssh-key",
		DisplayName: "For edit-disk(ssh-key) options",
		Order:       43,
	},
	{
		Key:         "server-info",
		DisplayName: "For server-info options",
		Order:       50,
	},
}

var AllowInterfaceDriver = []string{string(sacloud.InterfaceDriverVirtIO), string(sacloud.InterfaceDriverE1000)}

func serverBuildParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		/*
		 === server plan ===
		*/
		"core": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set CPU core count",
			DefaultValue: 1,
			CompleteFunc: completeServerCore(),
			Required:     true,
			Category:     "server-plan",
			Order:        10,
		},
		"memory": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set memory size(GB)",
			DefaultValue: 1,
			CompleteFunc: completeServerMemory(),
			Required:     true,
			Category:     "server-plan",
			Order:        20,
		},
		"private-host-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set private-host-id",
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completePrivateHostID(),
			Category:     "server-plan",
			Order:        30,
		},
		/*
		 === disk ===
		*/
		"disk-mode": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "disk create mode[create/connect/diskless]",
			DefaultValue: "create",
			Required:     true,
			ValidateFunc: validateInStrValues("create", "connect", "diskless"),
			CompleteFunc: completeInStrValues("create", "connect", "diskless"),
			Category:     "disk",
			Order:        10,
		},
		"os-type": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set source OS type",
			ValidateFunc: validateInStrValues(ostype.OSTypeShortNames...),
			CompleteFunc: completeInStrValues(ostype.OSTypeShortNames...),
			Category:     "disk",
			Order:        20,
		},
		"disk-plan": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			DefaultValue: "ssd",
			Description:  "set disk plan('hdd' or 'ssd')",
			ValidateFunc: validateInStrValues(allowDiskPlans...),
			CompleteFunc: completeInStrValues(allowDiskPlans...),
			Category:     "disk",
			Order:        30,
		},
		"disk-connection": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			DefaultValue: "virtio",
			Description:  "set disk connection('virtio' or 'ide')",
			ValidateFunc: validateInStrValues(allowDiskConnections...),
			CompleteFunc: completeInStrValues(allowDiskConnections...),
			Category:     "disk",
			Order:        40,
		},
		"disk-size": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set disk size(GB)",
			DefaultValue: 20,
			ValidateFunc: validateInIntValues(allowDiskSizes...),
			CompleteFunc: completeInIntValues(allowDiskSizes...),
			Category:     "disk",
			Order:        50,
		},

		"source-archive-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set source disk ID",
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeArchiveID(),
			Category:     "disk",
			Order:        60,
		},
		"source-disk-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set source disk ID",
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeDiskID(),
			Category:     "disk",
			Order:        70,
		},
		"distant-from": {
			Type:         schema.TypeIntList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set distant from disk IDs",
			ValidateFunc: validateIntSlice(validateSakuraID()),
			CompleteFunc: completeDiskID(),
			Category:     "disk",
			Order:        80,
		},
		"disk-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set connect disk ID",
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeDiskID(),
			Category:     "disk",
			Order:        90,
		},

		/*
		  === iso image ===
		*/
		"iso-image-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set iso-image ID",
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeISOImageID(),
			Category:     "iso-image",
			Order:        10,
		},

		/*
		  === network ===
		*/
		"network-mode": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "network connection mode[shared/switch/disconnect/none]",
			DefaultValue: "shared",
			Required:     true,
			ValidateFunc: validateInStrValues("shared", "switch", "disconnect", "none"),
			CompleteFunc: completeInStrValues("shared", "switch", "disconnect", "none"),
			Category:     "network",
			Order:        10,
		},
		"interface-driver": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set interface driver[virtio/e1000]",
			ValidateFunc: validateInStrValues(AllowInterfaceDriver...),
			CompleteFunc: completeInStrValues(AllowInterfaceDriver...),
			DefaultValue: string(sacloud.InterfaceDriverVirtIO),
			Category:     "network",
			Order:        20,
		},
		"packet-filter-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set packet filter ID",
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completePacketFilterID(),
			Category:     "network",
			Order:        30,
		},
		"switch-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set connect switch ID",
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeSwitchID(),
			Category:     "network",
			Order:        40,
		},
		/*
		  === disk edit ===
		*/
		"hostname": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set hostname",
			Category:    "edit-disk",
			Order:       10,
		},
		"password": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set password",
			EnvVars:     []string{"SAKURACLOUD_SERVER_PASSWORD"},
			Category:    "edit-disk",
			Order:       20,
		},

		"disable-password-auth": {
			Type:        schema.TypeBool,
			Aliases:     []string{"disable-pw-auth"},
			HandlerType: schema.HandlerNoop,
			Description: "disable password auth on SSH",
			Category:    "edit-disk",
			Order:       30,
		},
		"ipaddress": {
			Type:         schema.TypeString,
			Aliases:      []string{"ip"},
			HandlerType:  schema.HandlerNoop,
			Description:  "set ipaddress",
			ValidateFunc: validateIPv4Address(),
			Category:     "edit-disk-network",
			Order:        10,
		},
		"nw-masklen": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"network-masklen"},
			Description:  "set ipaddress  prefix",
			DefaultValue: 24,
			ValidateFunc: validateIntRange(8, 29),
			Category:     "edit-disk-network",
			Order:        20,
		},
		"default-route": {
			Type:         schema.TypeString,
			Aliases:      []string{"gateway"},
			HandlerType:  schema.HandlerNoop,
			Description:  "set default gateway",
			ValidateFunc: validateIPv4Address(),
			Category:     "edit-disk-network",
			Order:        30,
		},
		"startup-scripts": {
			Type:        schema.TypeStringList,
			Aliases:     []string{"notes"},
			HandlerType: schema.HandlerNoop,
			Description: "set startup script(s)",
			Category:    "edit-disk-startup-script",
			Order:       10,
		},
		"startup-script-ids": {
			Type:         schema.TypeIntList,
			Aliases:      []string{"note-ids"},
			HandlerType:  schema.HandlerNoop,
			Description:  "set startup script ID(s)",
			ValidateFunc: validateIntSlice(validateSakuraID()),
			CompleteFunc: completeNoteID(),
			Category:     "edit-disk-startup-script",
			Order:        20,
		},
		"startup-scripts-ephemeral": {
			Type:         schema.TypeBool,
			HandlerType:  schema.HandlerNoop,
			Description:  "set startup script persist mode",
			DefaultValue: true,
			Category:     "edit-disk-startup-script",
			Order:        30,
		},

		"ssh-key-mode": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "ssh-key mode[none/id/generate/upload]",
			ValidateFunc: validateInStrValues("none", "id", "generate", "upload"),
			CompleteFunc: completeInStrValues("none", "id", "generate", "upload"),
			Category:     "edit-disk-ssh-key",
			Order:        10,
		},
		"ssh-key-name": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set ssh-key name",
			Category:    "edit-disk-ssh-key",
			Order:       20,
		},
		"ssh-key-ids": {
			Type:         schema.TypeIntList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set ssh-key ID(s)",
			ValidateFunc: validateIntSlice(validateSakuraID()),
			CompleteFunc: completeSSHKeyID(),
			Category:     "edit-disk-ssh-key",
			Order:        30,
		},

		"ssh-key-pass-phrase": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set ssh-key pass phrase",
			ValidateFunc: validateStrLen(8, 64),
			Category:     "edit-disk-ssh-key",
			Order:        40,
		},
		"ssh-key-description": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set ssh-key description",
			Category:    "edit-disk-ssh-key",
			Order:       50,
		},
		"ssh-key-private-key-output": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set ssh-key privatekey output path",
			Category:    "edit-disk-ssh-key",
			Order:       60,
		},
		"ssh-key-public-keys": {
			Type:        schema.TypeStringList,
			HandlerType: schema.HandlerNoop,
			Description: "set ssh-key public key ",
			Category:    "edit-disk-ssh-key",
			Order:       70,
		},
		"ssh-key-public-key-files": {
			Type:         schema.TypeStringList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set ssh-key public key file",
			ValidateFunc: validateStringSlice(validateFileExists()),
			Category:     "edit-disk-ssh-key",
			Order:        80,
		},
		"ssh-key-ephemeral": {
			Type:         schema.TypeBool,
			HandlerType:  schema.HandlerNoop,
			Description:  "set ssh-key persist mode",
			DefaultValue: true,
			Category:     "edit-disk-ssh-key",
			Order:        90,
		},

		/*
		  === server info ===
		*/
		"name": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerPathThrough,
			Description:  "set resource display name",
			Required:     true,
			ValidateFunc: validateStrLen(1, 64),
			Category:     "server-info",
			Order:        10,
		},
		"description": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerPathThrough,
			Description:  "set resource description",
			Aliases:      []string{"desc"},
			ValidateFunc: validateStrLen(0, 254),
			Category:     "server-info",
			Order:        20,
		},
		"tags": {
			Type:         schema.TypeStringList,
			HandlerType:  schema.HandlerPathThrough,
			Description:  "set resource tags",
			ValidateFunc: validateStringSlice(validateStrLen(1, 32)),
			Category:     "server-info",
			Order:        30,
		},
		"icon-id": {
			Type:            schema.TypeInt64,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetIconByID",
			Description:     "set Icon ID",
			ValidateFunc:    validateSakuraID(),
			CompleteFunc:    completeIconID(),
			Category:        "server-info",
			Order:           40,
		},

		/*
		  === other options ===
		*/
		"us-keyboard": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "use us-keyboard",
			Order:       10,
		},
		"disable-boot-after-create": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "boot after create",
			Order:       20,
		},
	}
}

func serverReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func serverUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
		"interface-driver": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set interface driver[virtio/e1000]",
			ValidateFunc: validateInStrValues(AllowInterfaceDriver...),
			CompleteFunc: completeInStrValues(AllowInterfaceDriver...),
			DefaultValue: string(sacloud.InterfaceDriverVirtIO),
			Category:     "network",
			Order:        10,
		},
	}
}

func serverDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"force": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"f"},
			Description: "forced-shutdown flag if server is running",
			Category:    "operation",
			Order:       10,
		},
		"without-disk": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "don't delete connected disks with server",
			Category:    "operation",
			Order:       20,
		},
	}
}

func serverSSHParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"key": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"i"},
			Description:  "private-key file path",
			ValidateFunc: validateFileExists(),
			Category:     "auth",
			Order:        10,
		},
		"user": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"l"},
			Description: "user name",
			Category:    "auth",
			Order:       20,
		},
		"port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"p"},
			Description:  "port",
			Required:     true,
			DefaultValue: 22,
			Category:     "auth",
			Order:        30,
		},
		"password": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "password(or private-key pass phrase)",
			EnvVars:     []string{"SAKURACLOUD_SERVER_PASSWORD"},
			Category:    "auth",
			Order:       40,
		},
		"quiet": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"q"},
			Description: "disable information messages",
			Category:    "output",
			Order:       10,
		},
	}
}

func serverSCPParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"key": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"i"},
			Description:  "private-key file path",
			ValidateFunc: validateFileExists(),
			Category:     "auth",
			Order:        10,
		},
		"user": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"l"},
			Description: "user name",
			Category:    "auth",
			Order:       20,
		},
		"port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"p"},
			Description:  "port",
			Required:     true,
			DefaultValue: 22,
			Category:     "auth",
			Order:        30,
		},
		"password": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "password(or private-key pass phrase)",
			EnvVars:     []string{"SAKURACLOUD_SERVER_PASSWORD"},
			Category:    "auth",
			Order:       40,
		},
		"recursive": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"r"},
			Description: "set recursive copy flag",
			Category:    "operation",
			Order:       10,
		},
		"quiet": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"q"},
			Description: "disable information messages",
			Category:    "output",
			Order:       10,
		},
	}
}

func serverVNCParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"wait-for-boot": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "wait until the server starts up",
			Category:    "VNC",
		},
	}
}

func serverVNCSendParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"command": {
			Type:          schema.TypeString,
			HandlerType:   schema.HandlerNoop,
			Aliases:       []string{"c"},
			Description:   "command(compatible with HashiCorp Packer's boot_command)",
			ConflictsWith: []string{"command-file"},
			Category:      "VNC",
			Order:         10,
		},
		"command-file": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"f"},
			Description:  "command file(compatible with HashiCorp Packer's boot_command)",
			ValidateFunc: validateFileExists(),
			Category:     "VNC",
			Order:        20,
		},
		"use-us-keyboard": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "use US Keyboard",
			Category:    "VNC",
			Order:       30,
		},
		"debug": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"d"},
			Description: "write debug info",
			Category:    "VNC",
			Order:       40,
		},
		"wait-for-boot": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "wait until the server starts up",
			Category:    "VNC",
			Order:       50,
		},
	}
}

func serverVNCSnapshotParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"wait-for-boot": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "wait until the server starts up",
			Category:    "VNC",
			Order:       10,
		},
		"output-path": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"o"},
			Description: "snapshot output filepath",
			Category:    "VNC",
			Order:       20,
		},
	}
}

func serverPowerOnParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func serverPowerOffParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func serverResetParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func serverWaitForParams() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func serverPlanChangeParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		/*
		 === server plan ===
		*/
		"core": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set CPU core count",
			CompleteFunc: completeServerCore(),
			Required:     true,
			Category:     "plan",
			Order:        10,
		},
		"memory": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set memory size(GB)",
			CompleteFunc: completeServerMemory(),
			Required:     true,
			Category:     "plan",
			Order:        20,
		},
	}
}

func serverISOImageInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func serverISOImageInsertParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"iso-image-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set iso-image ID",
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeISOImageID(),
			Category:     "ISO-insert",
			Order:        10,
		},
		"size": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set iso size(GB)",
			DefaultValue: 5,
			ValidateFunc: validateInIntValues(5, 10),
			CompleteFunc: completeInIntValues(5, 10),
			Category:     "ISO-upload",
			Order:        10,
		},
		"iso-file": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set iso image file",
			Category:    "ISO-upload",
			Order:       20,
		},
		"name": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerPathThrough,
			Description:  "set resource display name",
			ValidateFunc: validateStrLen(1, 64),
			Category:     "ISO-upload",
			Order:        30,
		},
		"description": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerPathThrough,
			Description:  "set resource description",
			Aliases:      []string{"desc"},
			ValidateFunc: validateStrLen(0, 254),
			Category:     "ISO-upload",
			Order:        40,
		},
		"tags": {
			Type:         schema.TypeStringList,
			HandlerType:  schema.HandlerPathThrough,
			Description:  "set resource tags",
			ValidateFunc: validateStringSlice(validateStrLen(1, 32)),
			Category:     "ISO-upload",
			Order:        50,
		},
		"icon-id": {
			Type:            schema.TypeInt64,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetIconByID",
			Description:     "set Icon ID",
			ValidateFunc:    validateSakuraID(),
			CompleteFunc:    completeIconID(),
			Category:        "ISO-upload",
			Order:           60,
		},
	}
}

func serverISOImageEjectParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func serverDiskInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func serverDiskConnectParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"disk-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target disk ID",
			Required:     true,
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeDiskID(),
			Category:     "disk",
			Order:        10,
		},
	}
}

func serverDiskDisconnectParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"disk-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target disk ID",
			Required:     true,
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeDiskID(),
			Category:     "disk",
			Order:        10,
		},
	}
}

func serverInterfaceInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func serverInterfaceAddForInternetParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"without-disk-edit": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set skip edit-disk flag. if true, don't call DiskEdit API after interface added",
			Category:    "disk-edit",
			Order:       10,
		},
	}
}
func serverInterfaceAddForRouterParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"switch-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set connect switch(connected to router) ID",
			Required:     true,
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeSwitchID(),
			Category:     "connect",
			Order:        10,
		},
		"without-disk-edit": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set skip edit-disk flag. if true, don't call DiskEdit API after interface added",
			Category:    "disk-edit",
			Order:       10,
		},
		"ipaddress": {
			Type:         schema.TypeString,
			Aliases:      []string{"ip"},
			HandlerType:  schema.HandlerNoop,
			Description:  "set ipaddress",
			ValidateFunc: validateIPv4Address(),
			Category:     "disk-edit",
			Order:        20,
		},
		"default-route": {
			Type:         schema.TypeString,
			Aliases:      []string{"gateway"},
			HandlerType:  schema.HandlerNoop,
			Description:  "set default gateway",
			ValidateFunc: validateIPv4Address(),
			Category:     "disk-edit",
			Order:        30,
		},
		"nw-masklen": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"network-masklen"},
			Description:  "set ipaddress  prefix",
			DefaultValue: 24,
			ValidateFunc: validateIntRange(8, 29),
			Category:     "disk-edit",
			Order:        40,
		},
	}
}
func serverInterfaceAddForSwitchParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"switch-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set connect switch ID",
			Required:     true,
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeSwitchID(),
			Category:     "connect",
			Order:        10,
		},
		"without-disk-edit": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set skip edit-disk flag. if true, don't call DiskEdit API after interface added",
			Category:    "disk-edit",
			Order:       10,
		},
		"ipaddress": {
			Type:         schema.TypeString,
			Aliases:      []string{"ip"},
			HandlerType:  schema.HandlerNoop,
			Description:  "set ipaddress",
			ValidateFunc: validateIPv4Address(),
			Category:     "disk-edit",
			Order:        20,
		},
		"default-route": {
			Type:         schema.TypeString,
			Aliases:      []string{"gateway"},
			HandlerType:  schema.HandlerNoop,
			Description:  "set default gateway",
			ValidateFunc: validateIPv4Address(),
			Category:     "disk-edit",
			Order:        30,
		},
		"nw-masklen": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"network-masklen"},
			Description:  "set ipaddress  prefix",
			DefaultValue: 24,
			ValidateFunc: validateIntRange(8, 29),
			Category:     "disk-edit",
			Order:        40,
		},
	}
}
func serverInterfaceAddDisconnectedParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func serverMonitorCPUParam() map[string]*schema.Schema {
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
			DefaultValue: "sakuracloud.server.{{.ID}}.cpu",
			Required:     true,
			Category:     "monitor",
			Order:        30,
		},
	}
}

func serverMonitorCPUColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Key"},
		{Name: "TimeStamp"},
		{Name: "UnixTime"},
		{Name: "CPUTime"},
	}
}

func serverMonitorNICParam() map[string]*schema.Schema {
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
		"index": {
			Type:        schema.TypeIntList,
			HandlerType: schema.HandlerNoop,
			Description: "target index(es)",
			Category:    "monitor",
			Order:       30,
		},
		"key-format": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set monitoring value key-format",
			DefaultValue: "sakuracloud.server.{{.ID}}.nic.{{.Index}}",
			Required:     true,
			Category:     "monitor",
			Order:        40,
		},
	}
}

func serverMonitorNICColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Index"},
		{Name: "Key"},
		{
			Name:    "NICID",
			Sources: []string{"NicID"},
		},
		{Name: "TimeStamp"},
		{Name: "UnixTime"},
		{Name: "Send"},
		{Name: "Receive"},
	}
}
func serverMonitorDiskParam() map[string]*schema.Schema {
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
		"index": {
			Type:        schema.TypeIntList,
			HandlerType: schema.HandlerNoop,
			Description: "target index(es)",
			Category:    "monitor",
			Order:       30,
		},
		"key-format": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set monitoring value key-format",
			DefaultValue: "sakuracloud.server.{{.ID}}.disk.{{.Index}}",
			Required:     true,
			Category:     "monitor",
			Order:        40,
		},
	}
}

func serverMonitorDiskColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Index"},
		{Name: "Key"},
		{Name: "DiskID"},
		{Name: "TimeStamp"},
		{Name: "UnixTime"},
		{Name: "Read"},
		{Name: "Write"},
	}
}

func serverMaintenanceInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

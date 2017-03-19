package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func ServerResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"l", "ls", "find"},
			Params:             serverListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: serverListColumns(),
			Category:           "basic",
			Order:              10,
		},
		"build": {
			Type:             schema.CommandCreate,
			Aliases:          []string{"b"},
			ParamCategories:  serverBuildParamCategories,
			Params:           serverBuildParam(),
			IncludeFields:    serverDetailIncludes(),
			ExcludeFields:    serverBuildResultExcludes(),
			UseCustomCommand: true,
			Category:         "basic",
			Order:            20,
		},
		"read": {
			Type:          schema.CommandRead,
			Aliases:       []string{"r"},
			Params:        serverReadParam(),
			IncludeFields: serverDetailIncludes(),
			ExcludeFields: serverDetailExcludes(),
			Category:      "basic",
			Order:         30,
		},
		"update": {
			Type:          schema.CommandUpdate,
			Aliases:       []string{"u"},
			Params:        serverUpdateParam(),
			IncludeFields: serverDetailIncludes(),
			ExcludeFields: serverDetailExcludes(),
			Category:      "basic",
			Order:         40,
		},
		"delete": {
			Type:             schema.CommandDelete,
			Aliases:          []string{"d", "rm"},
			Params:           serverDeleteParam(),
			IncludeFields:    serverDetailIncludes(),
			ExcludeFields:    serverDetailExcludes(),
			UseCustomCommand: true,
			Category:         "basic",
			Order:            50,
			NeedConfirm:      true,
		},
		"plan-change": {
			Type:             schema.CommandManipulate,
			Params:           serverPlanChangeParam(),
			Usage:            "Change server plan(core/memory)",
			IncludeFields:    serverDetailIncludes(),
			ExcludeFields:    serverDetailExcludes(),
			UseCustomCommand: true,
			Category:         "basic",
			Order:            60,
		},
		"boot": {
			Type:             schema.CommandManipulate,
			Aliases:          []string{"power-on"},
			Params:           serverPowerOnParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            10,
			NoOutput:         true,
		},
		"shutdown": {
			Type:             schema.CommandManipulate,
			Aliases:          []string{"power-off"},
			Params:           serverPowerOffParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            20,
			NeedConfirm:      true,
			NoOutput:         true,
		},
		"reset": {
			Type:             schema.CommandManipulate,
			Params:           serverResetParam(),
			UseCustomCommand: true,
			Category:         "power",
			Order:            30,
			NeedConfirm:      true,
			NoOutput:         true,
		},
		"wait-for-boot": {
			Type:             schema.CommandManipulate,
			Params:           serverWaitForParams(),
			Usage:            "Wait until boot is completed",
			UseCustomCommand: true,
			Category:         "power",
			Order:            40,
			NoOutput:         true,
		},
		"wait-for-down": {
			Type:             schema.CommandManipulate,
			Params:           serverWaitForParams(),
			Usage:            "Wait until shutdown is completed",
			UseCustomCommand: true,
			Category:         "power",
			Order:            50,
			NoOutput:         true,
		},
		"ssh": {
			Type:             schema.CommandManipulate,
			Params:           serverSSHParam(),
			Usage:            "Connect to server by SSH",
			UseCustomCommand: true,
			Category:         "ssh",
			Order:            10,
			NoOutput:         true,
		},
		"ssh-exec": {
			Type:             schema.CommandManipulate,
			Params:           serverSSHParam(),
			Usage:            "Execute command on server connected by SSH",
			UseCustomCommand: true,
			Category:         "ssh",
			Order:            20,
			NoOutput:         true,
		},
		"scp": {
			Type:             schema.CommandManipulate,
			Params:           serverSCPParam(),
			Usage:            "Copy files/directories by SSH",
			ArgsUsage:        "[ServerID:]<FROM> [ServerID:]<TO>",
			UseCustomCommand: true,
			Category:         "ssh",
			Order:            30,
			NoOutput:         true,
		},
		"disk-info": {
			Type:               schema.CommandManipulate,
			Params:             serverDiskInfoParam(),
			Usage:              "Show information of disk(s) connected to server",
			TableType:          output.TableSimple,
			TableColumnDefines: diskListColumns(),
			UseCustomCommand:   true,
			Category:           "disks",
			Order:              10,
		},
		"disk-connect": {
			Type:               schema.CommandManipulate,
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
			Type:               schema.CommandManipulate,
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
			Type:               schema.CommandManipulate,
			Params:             serverInterfaceInfoParam(),
			Usage:              "Show information of NIC(s) connected to server",
			TableType:          output.TableSimple,
			TableColumnDefines: serverInterfaceListColumns(),
			UseCustomCommand:   true,
			Category:           "network",
			Order:              10,
		},
		"interface-add-for-internet": {
			Type:               schema.CommandManipulate,
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
			Type:               schema.CommandManipulate,
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
			Type:               schema.CommandManipulate,
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
			Type:               schema.CommandManipulate,
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
			Type:             schema.CommandManipulate,
			Params:           serverISOImageInfoParam(),
			Usage:            "Show information of ISO-Image inserted to server",
			UseCustomCommand: true,
			Category:         "iso",
			Order:            10,
		},
		"iso-insert": {
			Type:             schema.CommandManipulate,
			Params:           serverISOImageInsertParam(),
			Usage:            "Insert ISO-Image to server",
			UseCustomCommand: true,
			Category:         "iso",
			Order:            20,
			NoOutput:         true,
		},
		"iso-eject": {
			Type:             schema.CommandManipulate,
			Params:           serverISOImageEjectParam(),
			Usage:            "Eject ISO-Image from server",
			UseCustomCommand: true,
			Category:         "iso",
			Order:            30,
			NoOutput:         true,
		},
		"monitor-cpu": {
			Type:               schema.CommandManipulate,
			Params:             serverMonitorCPUParam(),
			Usage:              "Collect CPU monitor values",
			TableType:          output.TableSimple,
			TableColumnDefines: serverMonitorCPUColumns(),
			UseCustomCommand:   true,
			Category:           "monitor",
			Order:              10,
		},
		"monitor-nic": {
			Type:               schema.CommandManipulate,
			Params:             serverMonitorNICParam(),
			Usage:              "Collect NIC(s) monitor values",
			TableType:          output.TableSimple,
			TableColumnDefines: serverMonitorNICColumns(),
			UseCustomCommand:   true,
			Category:           "monitor",
			Order:              20,
		},
		"monitor-disk": {
			Type:               schema.CommandManipulate,
			Params:             serverMonitorDiskParam(),
			Usage:              "Collect Disk(s) monitor values",
			TableType:          output.TableSimple,
			TableColumnDefines: serverMonitorDiskColumns(),
			UseCustomCommand:   true,
			Category:           "monitor",
			Order:              30,
		},
	}

	return &schema.Resource{
		Commands:          commands,
		ResourceCategory:  CategoryComputing,
		CommandCategories: serverCommandCategories,
	}
}

func serverListParam() map[string]*schema.Schema {
	return CommonListParam
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
			Name:    "IPAddress",
			Sources: []string{"IPAddress", "UserIPAddress"},
			Format:  "%s%s",
		},
		{
			Name:    "Status",
			Sources: []string{"Instance.Status"},
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
			Name:    "Switch-ID",
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
		Key:         "ssh",
		DisplayName: "SSH/SCP",
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
			ValidateFunc: validateInStrValues(osTypeValues...),
			CompleteFunc: completeInStrValues(osTypeValues...),
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
		"use-nic-virtio": {
			Type:         schema.TypeBool,
			HandlerType:  schema.HandlerNoop,
			Description:  "use virtio on nic",
			DefaultValue: true,
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
			Type:         schema.TypeBool,
			HandlerType:  schema.HandlerNoop,
			Description:  "boot after create",
			DefaultValue: false,
			Order:        20,
		},
	}
}

var osTypeValues = []string{
	"centos", "ubuntu", "debian", "vyos", "coreos", "kusanagi", "site-guard", "freebsd",
	"windows2008", "windows2008-rds", "windows2008-rds-office",
	"windows2012", "windows2012-rds", "windows2012-rds-office",
	"windows2016", "windows2016-rds", "windows2016-rds-office",
}

func serverReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func serverUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id":          paramID,
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func serverDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"force": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"f"},
			Description: "force-shutdown flag if server is running",
		},
		"with-disk": {
			Type:         schema.TypeBool,
			HandlerType:  schema.HandlerNoop,
			Description:  "delete connected disks with server",
			DefaultValue: true,
		},
	}
}

func serverSSHParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"key": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"i"},
			Description:  "private-key file path",
			ValidateFunc: validateFileExists(),
		},
		"user": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"l"},
			Description: "user name",
		},
		"port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"p"},
			Description:  "port",
			Required:     true,
			DefaultValue: 22,
		},
		"password": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "password(or private-key pass phrase)",
			EnvVars:     []string{"SAKURACLOUD_SSH_PASSWORD"},
		},
		"quiet": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"q"},
			Description: "disable information messages",
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
		},
		"user": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"l"},
			Description: "user name",
		},
		"port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"p"},
			Description:  "port",
			Required:     true,
			DefaultValue: 22,
		},
		"password": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "password(or private-key pass phrase)",
			EnvVars:     []string{"SAKURACLOUD_SSH_PASSWORD"},
		},
		"recursive": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"r"},
			Description: "set recursive copy flag",
		},
		"quiet": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"q"},
			Description: "disable information messages",
		},
	}
}

func serverPowerOnParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func serverPowerOffParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"force": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"f"},
			Description: "force shutdown flag",
		},
	}
}

func serverResetParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func serverWaitForParams() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func serverPlanChangeParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		/*
		 === server plan ===
		*/
		"core": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set CPU core count",
			CompleteFunc: completeServerCore(),
			Required:     true,
		},
		"memory": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set memory size(GB)",
			CompleteFunc: completeServerMemory(),
			Required:     true,
		},
	}
}

func serverISOImageInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func serverISOImageInsertParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"iso-image-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set iso-image ID",
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeISOImageID(),
		},
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
		"size": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set iso size(GB)",
			DefaultValue: 5,
			ValidateFunc: validateInIntValues(5, 10),
			CompleteFunc: completeInIntValues(5, 10),
		},
		"iso-file": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set iso image file",
		},
	}
}

func serverISOImageEjectParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func serverDiskInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func serverDiskConnectParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"disk-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target disk ID",
			Required:     true,
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeDiskID(),
		},
	}
}

func serverDiskDisconnectParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"disk-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target disk ID",
			Required:     true,
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeDiskID(),
		},
	}
}

func serverInterfaceInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func serverInterfaceAddForInternetParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"without-disk-edit": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set skip edit-disk flag. if true, don't call DiskEdit API after interface added",
		},
	}
}
func serverInterfaceAddForRouterParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"without-disk-edit": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set skip edit-disk flag. if true, don't call DiskEdit API after interface added",
		},
		"switch-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set connect switch(connected to router) ID",
			Required:     true,
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeSwitchID(),
		},
		"ipaddress": {
			Type:         schema.TypeString,
			Aliases:      []string{"ip"},
			HandlerType:  schema.HandlerNoop,
			Description:  "set ipaddress",
			ValidateFunc: validateIPv4Address(),
		},
		"default-route": {
			Type:         schema.TypeString,
			Aliases:      []string{"gateway"},
			HandlerType:  schema.HandlerNoop,
			Description:  "set default gateway",
			ValidateFunc: validateIPv4Address(),
		},
		"nw-masklen": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"network-masklen"},
			Description:  "set ipaddress  prefix",
			DefaultValue: 24,
			ValidateFunc: validateIntRange(8, 29),
		},
	}
}
func serverInterfaceAddForSwitchParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"without-disk-edit": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set skip edit-disk flag. if true, don't call DiskEdit API after interface added",
		},
		"switch-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set connect switch ID",
			Required:     true,
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeSwitchID(),
		},
		"ipaddress": {
			Type:         schema.TypeString,
			Aliases:      []string{"ip"},
			HandlerType:  schema.HandlerNoop,
			Description:  "set ipaddress",
			ValidateFunc: validateIPv4Address(),
		},
		"default-route": {
			Type:         schema.TypeString,
			Aliases:      []string{"gateway"},
			HandlerType:  schema.HandlerNoop,
			Description:  "set default gateway",
			ValidateFunc: validateIPv4Address(),
		},
		"nw-masklen": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"network-masklen"},
			Description:  "set ipaddress  prefix",
			DefaultValue: 24,
			ValidateFunc: validateIntRange(8, 29),
		},
	}
}
func serverInterfaceAddDisconnectedParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func serverMonitorCPUParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
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
			DefaultValue: "sakuracloud.{{.ID}}.cpu",
			Required:     true,
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
		"id": paramID,
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
		"index": {
			Type:        schema.TypeIntList,
			HandlerType: schema.HandlerNoop,
			Description: "target index(es)",
		},
		"key-format": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set monitoring value key-format",
			DefaultValue: "sakuracloud.{{.ID}}.nic.{{.Index}}",
			Required:     true,
		},
	}
}

func serverMonitorNICColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Index"},
		{Name: "Key"},
		{
			Name:    "NIC-ID",
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
		"id": paramID,
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
		"index": {
			Type:        schema.TypeIntList,
			HandlerType: schema.HandlerNoop,
			Description: "target index(es)",
		},
		"key-format": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set monitoring value key-format",
			DefaultValue: "sakuracloud.{{.ID}}.disk.{{.Index}}",
			Required:     true,
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

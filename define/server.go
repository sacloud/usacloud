package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func ServerResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:                schema.CommandList,
			ListResultFieldName: "Servers",
			Aliases:             []string{"l", "ls", "find"},
			Params:              serverListParam(),
			TableType:           output.TableSimple,
			TableColumnDefines:  serverListColumns(),
		},
		"build": {
			Type:             schema.CommandCreate,
			Aliases:          []string{"b"},
			Params:           serverBuildParam(),
			IncludeFields:    serverDetailIncludes(),
			ExcludeFields:    serverBuildResultExcludes(),
			UseCustomCommand: true,
		},
		"read": {
			Type:          schema.CommandRead,
			Aliases:       []string{"r"},
			Params:        serverReadParam(),
			IncludeFields: serverDetailIncludes(),
			ExcludeFields: serverDetailExcludes(),
		},
		"update": {
			Type:          schema.CommandUpdate,
			Aliases:       []string{"u"},
			Params:        serverUpdateParam(),
			IncludeFields: serverDetailIncludes(),
			ExcludeFields: serverDetailExcludes(),
		},
		"delete": {
			Type:             schema.CommandDelete,
			Aliases:          []string{"d", "rm"},
			Params:           serverDeleteParam(),
			IncludeFields:    serverDetailIncludes(),
			ExcludeFields:    serverDetailExcludes(),
			UseCustomCommand: true,
		},
		"ssh": {
			Type:             schema.CommandManipulate,
			Params:           serverSSHParam(),
			UseCustomCommand: true,
		},
		"ssh-exec": {
			Type:             schema.CommandManipulate,
			Params:           serverSSHParam(),
			UseCustomCommand: true,
		},
		"boot": {
			Type:             schema.CommandManipulate,
			Aliases:          []string{"power-on"},
			Params:           serverPowerOnParam(),
			UseCustomCommand: true,
		},
		"shutdown": {
			Type:             schema.CommandManipulate,
			Aliases:          []string{"power-off"},
			Params:           serverPowerOffParam(),
			UseCustomCommand: true,
		},
		"reset": {
			Type:             schema.CommandManipulate,
			Params:           serverResetParam(),
			UseCustomCommand: true,
		},
		"plan-change": {
			Type:             schema.CommandManipulate,
			Params:           serverPlanChangeParam(),
			IncludeFields:    serverDetailIncludes(),
			ExcludeFields:    serverDetailExcludes(),
			UseCustomCommand: true,
		},
		"wait-for-boot": {
			Type:             schema.CommandManipulate,
			Params:           serverWaitForParams(),
			UseCustomCommand: true,
		},
		"wait-for-down": {
			Type:             schema.CommandManipulate,
			Params:           serverWaitForParams(),
			UseCustomCommand: true,
		},
		"iso-info": {
			Type:             schema.CommandManipulate,
			Params:           serverISOImageInfoParam(),
			UseCustomCommand: true,
		},
		"iso-insert": {
			Type:             schema.CommandManipulate,
			Params:           serverISOImageInsertParam(),
			UseCustomCommand: true,
		},
		"iso-eject": {
			Type:             schema.CommandManipulate,
			Params:           serverISOImageEjectParam(),
			UseCustomCommand: true,
		},
		"disk-info": {
			Type:               schema.CommandManipulate,
			Params:             serverDiskInfoParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: diskListColumns(),
			UseCustomCommand:   true,
		},
		"disk-connect": {
			Type:               schema.CommandManipulate,
			Params:             serverDiskConnectParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: diskListColumns(),
			UseCustomCommand:   true,
		},
		"disk-disconnect": {
			Type:               schema.CommandManipulate,
			Params:             serverDiskDisconnectParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: diskListColumns(),
			UseCustomCommand:   true,
		},
		"interface-info": {
			Type:               schema.CommandManipulate,
			Params:             serverInterfaceInfoParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: serverInterfaceListColumns(),
			UseCustomCommand:   true,
		},
		"interface-add-for-internet": {
			Type:               schema.CommandManipulate,
			Params:             serverInterfaceAddForInternetParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: serverInterfaceListColumns(),
			UseCustomCommand:   true,
		},
		"interface-add-for-router": {
			Type:               schema.CommandManipulate,
			Params:             serverInterfaceAddForRouterParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: serverInterfaceListColumns(),
			UseCustomCommand:   true,
		},
		"interface-add-for-switch": {
			Type:               schema.CommandManipulate,
			Params:             serverInterfaceAddForSwitchParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: serverInterfaceListColumns(),
			UseCustomCommand:   true,
		},
		"interface-add-disconnected": {
			Type:               schema.CommandManipulate,
			Params:             serverInterfaceAddDisconnectedParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: serverInterfaceListColumns(),
			UseCustomCommand:   true,
		},
	}

	return &schema.Resource{
		Commands: commands,
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
			Required:     true,
		},
		"memory": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set memory size(GB)",
			DefaultValue: 1,
			Required:     true,
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
		},
		"disk-plan": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			DefaultValue: "ssd",
			Description:  "set disk plan('hdd' or 'ssd')",
			ValidateFunc: validateInStrValues("ssd", "hdd"),
		},
		"disk-connection": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			DefaultValue: "virtio",
			Description:  "set disk connection('virtio' or 'ide')",
			ValidateFunc: validateInStrValues("virtio", "ide"),
		},
		"disk-size": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set disk size(GB)",
			DefaultValue: 20,
			ValidateFunc: validateInIntValues(20, 40, 60, 80, 100, 250, 500, 750, 1000, 2000, 4000),
		},
		"os-type": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set source OS type",
			ValidateFunc: validateInStrValues(osTypeValues...),
		},
		"source-archive-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set source disk ID",
			ValidateFunc: validateSakuraID(),
		},
		"source-disk-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set source disk ID",
			ValidateFunc: validateSakuraID(),
		},
		"distant-from": {
			Type:         schema.TypeIntList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set distant from disk IDs",
			ValidateFunc: validateIntSlice(validateSakuraID()),
		},
		"disk-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set connect disk ID",
			ValidateFunc: validateSakuraID(),
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
		},
		"use-nic-virtio": {
			Type:         schema.TypeBool,
			HandlerType:  schema.HandlerNoop,
			Description:  "use virtio on nic",
			DefaultValue: true,
		},
		"packet-filter-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set packet filter ID",
			ValidateFunc: validateSakuraID(),
		},
		"switch-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set connect switch ID",
			ValidateFunc: validateSakuraID(),
		},
		/*
		  === disk edit ===
		*/
		"hostname": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set hostname",
		},
		"password": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set password",
		},

		"disable-password-auth": {
			Type:        schema.TypeBool,
			Aliases:     []string{"disable-pw-auth"},
			HandlerType: schema.HandlerNoop,
			Description: "disable password auth on SSH",
		},
		"startup-script-ids": {
			Type:        schema.TypeIntList,
			Aliases:     []string{"note-ids"},
			HandlerType: schema.HandlerNoop,
			Description: "set startup script ID(s)",
		},
		"startup-scripts": {
			Type:        schema.TypeStringList,
			Aliases:     []string{"notes"},
			HandlerType: schema.HandlerNoop,
			Description: "set startup script(s)",
		},
		"startup-scripts-ephemeral": {
			Type:         schema.TypeBool,
			HandlerType:  schema.HandlerNoop,
			Description:  "set startup script persist mode(if true, script will delete after create server)",
			DefaultValue: true,
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
		"ssh-key-mode": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "ssh-key mode[none/id/generate/upload]",
			ValidateFunc: validateInStrValues("none", "id", "generate", "upload"),
		},
		"ssh-key-ids": {
			Type:         schema.TypeIntList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set ssh-key ID(s)",
			ValidateFunc: validateIntSlice(validateSakuraID()),
		},
		"ssh-key-name": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set ssh-key name",
		},
		"ssh-key-pass-phrase": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set ssh-key pass phrase",
			ValidateFunc: validateStrLen(8, 64),
		},
		"ssh-key-description": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set ssh-key description",
		},
		"ssh-key-private-key-output": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set ssh-key privatekey output path",
		},
		"ssh-key-public-keys": {
			Type:        schema.TypeStringList,
			HandlerType: schema.HandlerNoop,
			Description: "set ssh-key public key ",
		},
		"ssh-key-public-key-files": {
			Type:         schema.TypeStringList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set ssh-key public key file",
			ValidateFunc: validateStringSlice(validateFileExists()),
		},
		"ssh-key-ephemeral": {
			Type:         schema.TypeBool,
			HandlerType:  schema.HandlerNoop,
			Description:  "set ssh-key persist mode(if true, script will delete after create server)",
			DefaultValue: true,
		},
		/*
		  === iso image ===
		*/
		"iso-image-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set iso-image ID",
			ValidateFunc: validateSakuraID(),
		},

		/*
		  === server info ===
		*/
		"name":        paramRequiredName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     getParamSubResourceID("Icon"),
		"us-keyboard": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "use us-keyboard",
		},
		"disable-boot-after-create": {
			Type:         schema.TypeBool,
			HandlerType:  schema.HandlerNoop,
			Description:  "boot after create",
			DefaultValue: false,
		},
	}
}

var osTypeValues = []string{
	"centos", "ubuntu", "debian", "vyos", "coreos", "kusanagi", "site-guard", "freebsd",
	"windows2008", "windows2008-rds", "windows2008-rds-office",
	"windows2012", "windows2012-rds", "windows2012-rds-office",
	"windows2016",
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
		"icon-id":     getParamSubResourceID("Icon"),
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
		"proxy": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "proxy server",
		},
	}
}

func serverPowerOnParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"async": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set async flag(if true,return with non block)",
		},
	}
}

func serverPowerOffParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"force": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "force shutdown flag",
		},
		"async": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set async flag(if true,return with non block)",
		},
	}
}

func serverResetParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"async": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set async flag(if true,return with non block)",
		},
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
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "set CPU core count",
			Required:    true,
		},
		"memory": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "set memory size(GB)",
			Required:    true,
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
		},
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     getParamSubResourceID("Icon"),
		"size": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set iso size(GB)",
			DefaultValue: 5,
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

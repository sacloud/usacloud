package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func DiskResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"l", "ls", "find"},
			Params:             diskListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: diskListColumns(),
		},
		"create": {
			Type:             schema.CommandCreate,
			Aliases:          []string{"c"},
			Params:           diskCreateParam(),
			IncludeFields:    diskDetailIncludes(),
			ExcludeFields:    diskDetailExcludes(),
			UseCustomCommand: true,
		},
		"read": {
			Type:          schema.CommandRead,
			Aliases:       []string{"r"},
			Params:        diskReadParam(),
			IncludeFields: diskDetailIncludes(),
			ExcludeFields: diskDetailExcludes(),
		},
		"update": {
			Type:          schema.CommandUpdate,
			Aliases:       []string{"u"},
			Params:        diskUpdateParam(),
			IncludeFields: diskDetailIncludes(),
			ExcludeFields: diskDetailExcludes(),
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"d", "rm"},
			Params:        diskDeleteParam(),
			IncludeFields: diskDetailIncludes(),
			ExcludeFields: diskDetailExcludes(),
			NeedConfirm:   true,
		},
		"edit": {
			Type:             schema.CommandManipulate,
			Aliases:          []string{"config"},
			Params:           diskConfigParam(),
			IncludeFields:    diskDetailIncludes(),
			ExcludeFields:    diskDetailExcludes(),
			UseCustomCommand: true,
		},
		"wait-for-copy": {
			Type:             schema.CommandManipulate,
			Aliases:          []string{"wait"},
			Params:           diskWaitForCopyParam(),
			IncludeFields:    diskDetailIncludes(),
			ExcludeFields:    diskDetailExcludes(),
			UseCustomCommand: true,
			NoOutput:         true,
		},
		"reinstall-from-archive": {
			Type:             schema.CommandManipulate,
			Params:           diskReinstallFromArchiveParam(),
			IncludeFields:    diskDetailIncludes(),
			ExcludeFields:    diskDetailExcludes(),
			UseCustomCommand: true,
			NeedConfirm:      true,
			ConfirmMessage:   "re-install from archive",
			NoOutput:         true,
		},
		"reinstall-from-disk": {
			Type:             schema.CommandManipulate,
			Params:           diskReinstallFromDiskParam(),
			IncludeFields:    diskDetailIncludes(),
			ExcludeFields:    diskDetailExcludes(),
			UseCustomCommand: true,
			NeedConfirm:      true,
			ConfirmMessage:   "re-install from disk",
			NoOutput:         true,
		},
		"reinstall-to-blank": {
			Type:             schema.CommandManipulate,
			Params:           diskReinstallToBlankParam(),
			IncludeFields:    diskDetailIncludes(),
			ExcludeFields:    diskDetailExcludes(),
			UseCustomCommand: true,
			NeedConfirm:      true,
			ConfirmMessage:   "re-install to blank",
			NoOutput:         true,
		},
		"server-connect": {
			Type:             schema.CommandManipulate,
			Params:           diskServerConnectParam(),
			IncludeFields:    diskDetailIncludes(),
			ExcludeFields:    diskDetailExcludes(),
			UseCustomCommand: true,
			NoOutput:         true,
		},
		"server-disconnect": {
			Type:             schema.CommandManipulate,
			Params:           diskServerDisconnectParam(),
			IncludeFields:    diskDetailIncludes(),
			ExcludeFields:    diskDetailExcludes(),
			UseCustomCommand: true,
			NoOutput:         true,
		},
		"monitor": {
			Type:               schema.CommandManipulate,
			Params:             diskMonitorParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: diskMonitorColumns(),
			UseCustomCommand:   true,
		},
	}

	return &schema.Resource{
		Commands:         commands,
		ResourceCategory: CategoryStorage,
	}
}

func diskListParam() map[string]*schema.Schema {
	return mergeParameterMap(CommonListParam, paramScopeCond)
}

func diskListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{
			Name:    "Server",
			Sources: []string{"Server.ID", "Server.Name"},
			Format:  "%s(%s)",
		},
		{
			Name:    "Plan",
			Sources: []string{"Plan.ID"},
			ValueMapping: []map[string]string{
				{
					"4": "ssd",
					"2": "hdd",
				},
			},
		},
		{
			Name:    "Size",
			Sources: []string{"SizeMB"},
			Format:  "%sMB",
		},
		{Name: "Connection"},
	}
}

func diskDetailIncludes() []string {
	return []string{}
}

func diskDetailExcludes() []string {
	return []string{
		"SourceArchive.Storage.",
		"Storage.",
	}
}

var allowDiskPlans = []string{"ssd", "hdd"}
var allowDiskConnections = []string{"virtio", "ide"}
var allowDiskSizes = []int{20, 40, 60, 80, 100, 250, 500, 750, 1024, 2048, 4096}

func diskCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramRequiredName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
		"plan": {
			Type:            schema.TypeString,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetDiskPlan",
			Required:        true,
			DefaultValue:    "ssd",
			Description:     "set disk plan('hdd' or 'ssd')",
			ValidateFunc:    validateInStrValues(allowDiskPlans...),
			CompleteFunc:    completeInStrValues(allowDiskPlans...),
		},
		"connection": {
			Type:            schema.TypeString,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetDiskConnectionByStr",
			Required:        true,
			DefaultValue:    "virtio",
			Description:     "set disk connection('virtio' or 'ide')",
			ValidateFunc:    validateInStrValues(allowDiskConnections...),
			CompleteFunc:    completeInStrValues(allowDiskConnections...),
		},
		"size": {
			Type:            schema.TypeInt,
			HandlerType:     schema.HandlerPathThrough,
			Description:     "set disk size(GB)",
			DestinationProp: "SetSizeGB",
			DefaultValue:    20,
			Required:        true,
			ValidateFunc:    validateInIntValues(allowDiskSizes...),
			CompleteFunc:    completeInIntValues(allowDiskSizes...),
		},
		"source-archive-id": {
			Type:            schema.TypeInt64,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetSourceArchive",
			Description:     "set source disk ID",
			ValidateFunc:    validateSakuraID(),
			CompleteFunc:    completeArchiveID(),
			ConflictsWith:   []string{"source-disk-id"},
		},
		"source-disk-id": {
			Type:            schema.TypeInt64,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetSourceDisk",
			Description:     "set source disk ID",
			ValidateFunc:    validateSakuraID(),
			CompleteFunc:    completeDiskID(),
			ConflictsWith:   []string{"source-archive-id"},
		},
		"distant-from": {
			Type:         schema.TypeIntList,
			HandlerType:  schema.HandlerPathThrough,
			Description:  "set distant from disk IDs",
			ValidateFunc: validateIntSlice(validateSakuraID()),
			CompleteFunc: completeDiskID(),
		},
		"async": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set async flag(if true,return with non block)",
		},
	}
}

func diskReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func diskUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id":          paramID,
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
		"connection": {
			Type:            schema.TypeString,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetDiskConnectionByStr",
			Description:     "set disk connection('virtio' or 'ide')",
			ValidateFunc:    validateInStrValues(allowDiskConnections...),
			CompleteFunc:    completeInStrValues(allowDiskConnections...),
		},
	}
}

func diskDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func diskConfigParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"hostname": {
			Type:            schema.TypeString,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetHostName",
			Description:     "set hostname",
		},
		"password": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerPathThrough,
			Description: "set password",
		},
		"ssh-key-ids": {
			Type:            schema.TypeIntList,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetSSHKeys",
			Description:     "set ssh-key ID(s)",
			ValidateFunc:    validateIntSlice(validateSakuraID()),
			CompleteFunc:    completeSSHKeyID(),
		},
		"disable-password-auth": {
			Type:            schema.TypeBool,
			Aliases:         []string{"disable-pw-auth"},
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetDisablePWAuth",
			Description:     "disable password auth on SSH",
		},
		"startup-script-ids": {
			Type:            schema.TypeIntList,
			Aliases:         []string{"note-ids"},
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetNotes",
			Description:     "set startup-script ID(s)",
			ValidateFunc:    validateIntSlice(validateSakuraID()),
			CompleteFunc:    completeNoteID(),
		},
		"ipaddress": {
			Type:            schema.TypeString,
			Aliases:         []string{"ip"},
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetUserIPAddress",
			Description:     "set ipaddress",
		},
		"default-route": {
			Type:            schema.TypeString,
			Aliases:         []string{"gateway"},
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetDefaultRoute",
			Description:     "set default gateway",
		},
		"nw-masklen": {
			Type:            schema.TypeInt,
			HandlerType:     schema.HandlerPathThrough,
			Aliases:         []string{"network-masklen"},
			Description:     "set ipaddress  prefix",
			DestinationProp: "SetNetworkMaskLen",
			DefaultValue:    24,
			ValidateFunc:    validateIntRange(8, 29),
		},
	}
}

func diskWaitForCopyParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func diskReinstallFromArchiveParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"source-archive-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set source archive ID",
			Required:     true,
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeArchiveID(),
		},
		"distant-from": {
			Type:         schema.TypeIntList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set distant from disk IDs",
			ValidateFunc: validateIntSlice(validateSakuraID()),
			CompleteFunc: completeDiskID(),
		},
		"async": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set async flag(if true,return with non block)",
		},
	}
}

func diskReinstallFromDiskParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"source-disk-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set source disk ID",
			Required:     true,
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeDiskID(),
		},
		"distant-from": {
			Type:         schema.TypeIntList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set distant from disk IDs",
			ValidateFunc: validateIntSlice(validateSakuraID()),
			CompleteFunc: completeDiskID(),
		},
		"async": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set async flag(if true,return with non block)",
		},
	}
}

func diskReinstallToBlankParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"distant-from": {
			Type:         schema.TypeIntList,
			HandlerType:  schema.HandlerNoop,
			Description:  "set distant from disk IDs",
			ValidateFunc: validateIntSlice(validateSakuraID()),
			CompleteFunc: completeDiskID(),
		},
		"async": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set async flag(if true,return with non block)",
		},
	}
}

func diskServerConnectParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"server-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target server ID",
			Required:     true,
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeServerID(),
		},
	}
}

func diskServerDisconnectParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func diskMonitorParam() map[string]*schema.Schema {
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
			DefaultValue: "sakuracloud.{{.ID}}.disk",
			Required:     true,
		},
	}
}

func diskMonitorColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Key"},
		{Name: "TimeStamp"},
		{Name: "UnixTime"},
		{Name: "Read"},
		{Name: "Write"},
	}
}

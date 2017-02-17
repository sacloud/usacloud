package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func DiskResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:                schema.CommandList,
			ListResultFieldName: "Disks",
			Aliases:             []string{"l", "ls", "find"},
			Params:              diskListParam(),
			TableType:           output.TableSimple,
			TableColumnDefines:  diskListColumns(),
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
		},
	}

	return &schema.Resource{
		Commands: commands,
	}
}

func diskListParam() map[string]*schema.Schema {
	return mergeParameterMap(CommonListParam, paramScopeCond)
}

func diskListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{Name: "Scope"},
	}
}

func diskDetailIncludes() []string {
	return []string{}
}

func diskDetailExcludes() []string {
	return []string{}
}

func diskCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramRequiredName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon":        getParamSubResourceID("Icon"),
		"plan": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerPathThrough,
			Description: "set disk plan('hdd' or 'ssd')",
		},
		"size": {
			Type:            schema.TypeInt,
			HandlerType:     schema.HandlerPathThrough,
			Description:     "set disk size(GB)",
			DestinationProp: "SetSizeGB",
			DefaultValue:    20,
			ValidateFunc:    validateInIntValues(20, 40, 60, 80, 100, 250, 500, 750, 1000),
			ConflictsWith:   []string{"source-disk", "source-disk"},
		},
		"source-archive": {
			Type:            schema.TypeInt64,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetSourceArchive",
			Description:     "set source disk ID",
			ValidateFunc:    validateSakuraID(),
			ConflictsWith:   []string{"disk-file", "source-disk", "size"},
		},
		"source-disk": {
			Type:            schema.TypeInt64,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetSourceDisk",
			Description:     "set source disk ID",
			ValidateFunc:    validateSakuraID(),
			ConflictsWith:   []string{"disk-file", "source-disk", "size"},
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
		"icon":        getParamSubResourceID("Icon"),
	}
}

func diskDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func diskUploadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"disk-file": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set disk image file",
			Required:     true,
			ValidateFunc: validateFileExists(),
		},
	}
}

func diskDownloadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"file-destination": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set file destination path",
			Required:    true,
		},
	}
}

func diskOpenFTPParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func diskCloseFTPParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

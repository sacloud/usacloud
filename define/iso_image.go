package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func ISOImageResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"l", "ls", "find"},
			Params:             isoImageListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: isoImageListColumns(),
		},
		"create": {
			Type:             schema.CommandCreate,
			Aliases:          []string{"c"},
			Params:           isoImageCreateParam(),
			IncludeFields:    isoImageDetailIncludes(),
			ExcludeFields:    isoImageDetailExcludes(),
			UseCustomCommand: true,
		},
		"read": {
			Type:          schema.CommandRead,
			Aliases:       []string{"r"},
			Params:        isoImageReadParam(),
			IncludeFields: isoImageDetailIncludes(),
			ExcludeFields: isoImageDetailExcludes(),
		},
		"update": {
			Type:          schema.CommandUpdate,
			Aliases:       []string{"u"},
			Params:        isoImageUpdateParam(),
			IncludeFields: isoImageDetailIncludes(),
			ExcludeFields: isoImageDetailExcludes(),
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"d", "rm"},
			Params:        isoImageDeleteParam(),
			IncludeFields: isoImageDetailIncludes(),
			ExcludeFields: isoImageDetailExcludes(),
		},
		"upload": {
			Type:             schema.CommandManipulateSingle,
			Params:           isoImageUploadParam(),
			IncludeFields:    isoImageDetailIncludes(),
			ExcludeFields:    isoImageDetailExcludes(),
			UseCustomCommand: true,
		},
		"download": {
			Type:             schema.CommandManipulateSingle,
			Params:           isoImageDownloadParam(),
			IncludeFields:    isoImageDetailIncludes(),
			ExcludeFields:    isoImageDetailExcludes(),
			UseCustomCommand: true,
			NoOutput:         true,
		},
		"ftp-open": {
			Type:             schema.CommandManipulateMulti,
			Params:           isoImageOpenFTPParam(),
			IncludeFields:    isoImageDetailIncludes(),
			ExcludeFields:    isoImageDetailExcludes(),
			UseCustomCommand: true,
		},
		"ftp-close": {
			Type:             schema.CommandManipulateMulti,
			Params:           isoImageCloseFTPParam(),
			IncludeFields:    isoImageDetailIncludes(),
			ExcludeFields:    isoImageDetailExcludes(),
			UseCustomCommand: true,
			NoOutput:         true,
		},
	}

	return &schema.Resource{
		AltResource:         "CDROM",
		ListResultFieldName: "CDROMs",
		Commands:            commands,
		ResourceCategory:    CategoryStorage,
	}
}

func isoImageListParam() map[string]*schema.Schema {
	return mergeParameterMap(CommonListParam, paramScopeCond)
}

func isoImageListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{Name: "Scope"},
	}
}

func isoImageDetailIncludes() []string {
	return []string{}
}

func isoImageDetailExcludes() []string {
	return []string{
		"Storage.",
	}
}

func isoImageCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramRequiredName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
		"size": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set iso size(GB)",
			Required:     true,
			DefaultValue: 5,
			ValidateFunc: validateInIntValues(5, 10),
			CompleteFunc: completeInIntValues(5, 10),
		},
		"iso-file": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set iso image file",
			Required:     true,
			ValidateFunc: validateFileExists(),
		},
	}
}

func isoImageReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func isoImageUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func isoImageDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func isoImageUploadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"iso-file": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set iso image file",
			Required:     true,
			ValidateFunc: validateFileExists(),
		},
	}
}

func isoImageDownloadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"file-destination": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set file destination path",
			Required:    true,
		},
	}
}

func isoImageOpenFTPParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func isoImageCloseFTPParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

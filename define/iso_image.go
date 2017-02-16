package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func ISOImageResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:                schema.CommandList,
			ListResultFieldName: "CDROMs",
			Aliases:             []string{"l", "ls", "find"},
			Params:              isoImageListParam(),
			TableType:           output.TableSimple,
			TableColumnDefines:  isoImageListColumns(),
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
			Type:             schema.CommandManipulate,
			Params:           isoImageUploadParam(),
			IncludeFields:    isoImageDetailIncludes(),
			ExcludeFields:    isoImageDetailExcludes(),
			UseCustomCommand: true,
		},
		"download": {
			Type:             schema.CommandManipulate,
			Params:           isoImageDownloadParam(),
			IncludeFields:    isoImageDetailIncludes(),
			ExcludeFields:    isoImageDetailExcludes(),
			UseCustomCommand: true,
		},
		"ftp-open": {
			Type:             schema.CommandManipulate,
			Params:           isoImageOpenFTPParam(),
			IncludeFields:    isoImageDetailIncludes(),
			ExcludeFields:    isoImageDetailExcludes(),
			UseCustomCommand: true,
		},
		"ftp-close": {
			Type:             schema.CommandManipulate,
			Params:           isoImageCloseFTPParam(),
			IncludeFields:    isoImageDetailIncludes(),
			ExcludeFields:    isoImageDetailExcludes(),
			UseCustomCommand: true,
		},
	}

	return &schema.Resource{
		AltResource: "CDROM",
		Commands:    commands,
	}
}

func isoImageListParam() map[string]*schema.Schema {
	return CommonListParam
}

func isoImageListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
	}
}

func isoImageDetailIncludes() []string {
	return []string{}
}

func isoImageDetailExcludes() []string {
	return []string{}
}

var isoImageCommonParam = map[string]*schema.Schema{
	"name":        paramName,
	"description": paramDescription,
	"tags":        paramTags,
	"icon":        getParamSubResourceID("Icon"),
}

var isoImageFileParam = map[string]*schema.Schema{
	"iso-file": {
		Type:          schema.TypeString,
		HandlerType:   schema.HandlerCustomFunc,
		Description:   "set iso image file",
		Required:      true,
		ValidateFunc:  validateFileExists(),
		CustomHandler: iconSetImageContentUseBase64,
	},
}
var isoImageSizeParam = map[string]*schema.Schema{
	"size": {
		Type:            schema.TypeInt,
		HandlerType:     schema.HandlerPathThrough,
		Description:     "set iso size(GB)",
		DestinationProp: "SetSizeGB",
		Required:        true,
		DefaultValue:    5,
		ValidateFunc:    validateInIntValues(5, 10),
	},
}

func isoImageCreateParam() map[string]*schema.Schema {
	return mergeParameterMap(isoImageCommonParam, isoImageSizeParam, isoImageFileParam)
}

func isoImageReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func isoImageUpdateParam() map[string]*schema.Schema {
	updateParam := map[string]*schema.Schema{
		"id": paramID,
	}
	return mergeParameterMap(updateParam, isoImageCommonParam)
}

func isoImageDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func isoImageUploadParam() map[string]*schema.Schema {
	uploadParam := map[string]*schema.Schema{
		"id": paramID,
	}
	return mergeParameterMap(uploadParam, isoImageFileParam)
}

func isoImageDownloadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"iso-file": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set file destination path",
			Required:    true,
		},
	}
}

func isoImageOpenFTPParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func isoImageCloseFTPParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

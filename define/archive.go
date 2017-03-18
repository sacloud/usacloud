package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func ArchiveResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"l", "ls", "find"},
			Params:             archiveListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: archiveListColumns(),
		},
		"create": {
			Type:             schema.CommandCreate,
			Aliases:          []string{"c"},
			Params:           archiveCreateParam(),
			IncludeFields:    archiveDetailIncludes(),
			ExcludeFields:    archiveDetailExcludes(),
			UseCustomCommand: true,
		},
		"read": {
			Type:          schema.CommandRead,
			Aliases:       []string{"r"},
			Params:        archiveReadParam(),
			IncludeFields: archiveDetailIncludes(),
			ExcludeFields: archiveDetailExcludes(),
		},
		"update": {
			Type:          schema.CommandUpdate,
			Aliases:       []string{"u"},
			Params:        archiveUpdateParam(),
			IncludeFields: archiveDetailIncludes(),
			ExcludeFields: archiveDetailExcludes(),
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"d", "rm"},
			Params:        archiveDeleteParam(),
			IncludeFields: archiveDetailIncludes(),
			ExcludeFields: archiveDetailExcludes(),
			NeedConfirm:   true,
		},
		"upload": {
			Type:             schema.CommandManipulate,
			Params:           archiveUploadParam(),
			IncludeFields:    archiveDetailIncludes(),
			ExcludeFields:    archiveDetailExcludes(),
			UseCustomCommand: true,
		},
		"download": {
			Type:             schema.CommandManipulate,
			Params:           archiveDownloadParam(),
			IncludeFields:    archiveDetailIncludes(),
			ExcludeFields:    archiveDetailExcludes(),
			UseCustomCommand: true,
			NoOutput:         true,
		},
		"wait-for-copy": {
			Type:             schema.CommandManipulate,
			Params:           archiveWaitForCopyParam(),
			IncludeFields:    archiveDetailIncludes(),
			ExcludeFields:    archiveDetailExcludes(),
			UseCustomCommand: true,
			NoOutput:         true,
		},
		"ftp-open": {
			Type:             schema.CommandManipulate,
			Params:           archiveOpenFTPParam(),
			IncludeFields:    archiveDetailIncludes(),
			ExcludeFields:    archiveDetailExcludes(),
			UseCustomCommand: true,
		},
		"ftp-close": {
			Type:             schema.CommandManipulate,
			Params:           archiveCloseFTPParam(),
			IncludeFields:    archiveDetailIncludes(),
			ExcludeFields:    archiveDetailExcludes(),
			UseCustomCommand: true,
			NoOutput:         true,
		},
	}

	return &schema.Resource{
		Commands:         commands,
		ResourceCategory: CategoryStorage,
	}
}

func archiveListParam() map[string]*schema.Schema {
	return mergeParameterMap(CommonListParam, paramScopeCond)
}

func archiveListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{Name: "Scope"},
	}
}

func archiveDetailIncludes() []string {
	return []string{}
}

func archiveDetailExcludes() []string {
	return []string{
		"Storage.",
		"SourceArchive.Storage.",
	}
}

var allowSizes = []int{20, 40, 60, 80, 100, 250, 500, 750, 1024}

func archiveCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramRequiredName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
		"size": {
			Type:            schema.TypeInt,
			HandlerType:     schema.HandlerPathThrough,
			Description:     "set archive size(GB)",
			DestinationProp: "SetSizeGB",
			ValidateFunc:    validateInIntValues(allowSizes...),
			CompleteFunc:    completeInIntValues(allowSizes...),
			ConflictsWith:   []string{"source-archive-id", "source-disk-id"},
		},
		"archive-file": {
			Type:          schema.TypeString,
			HandlerType:   schema.HandlerNoop,
			Description:   "set archive image file",
			ValidateFunc:  validateFileExists(),
			ConflictsWith: []string{"source-archive-id", "source-disk-id"},
		},
		"source-disk-id": {
			Type:            schema.TypeInt64,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetSourceDisk",
			Description:     "set source disk ID",
			ValidateFunc:    validateSakuraID(),
			CompleteFunc:    completeDiskID(),
			ConflictsWith:   []string{"archive-file", "source-archive-id", "size"},
		},
		"source-archive-id": {
			Type:            schema.TypeInt64,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetSourceArchive",
			Description:     "set source archive ID",
			ValidateFunc:    validateSakuraID(),
			CompleteFunc:    completeArchiveID(),
			ConflictsWith:   []string{"archive-file", "source-disk-id", "size"},
		},
	}
}

func archiveReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func archiveUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id":          paramID,
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func archiveDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func archiveUploadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"archive-file": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set archive image file",
			Required:     true,
			ValidateFunc: validateFileExists(),
		},
	}
}

func archiveDownloadParam() map[string]*schema.Schema {
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

func archiveWaitForCopyParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func archiveOpenFTPParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func archiveCloseFTPParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func ArchiveResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find"},
			Params:             archiveListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: archiveListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"create": {
			Type:             schema.CommandCreate,
			Params:           archiveCreateParam(),
			IncludeFields:    archiveDetailIncludes(),
			ExcludeFields:    archiveDetailExcludes(),
			Category:         "basics",
			Order:            20,
			UseCustomCommand: true,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        archiveReadParam(),
			IncludeFields: archiveDetailIncludes(),
			ExcludeFields: archiveDetailExcludes(),
			Category:      "basics",
			Order:         30,
		},
		"update": {
			Type:          schema.CommandUpdate,
			Params:        archiveUpdateParam(),
			IncludeFields: archiveDetailIncludes(),
			ExcludeFields: archiveDetailExcludes(),
			Category:      "basics",
			Order:         40,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"rm"},
			Params:        archiveDeleteParam(),
			IncludeFields: archiveDetailIncludes(),
			ExcludeFields: archiveDetailExcludes(),
			Category:      "basics",
			Order:         50,
		},
		"upload": {
			Type:             schema.CommandManipulateSingle,
			Params:           archiveUploadParam(),
			IncludeFields:    archiveDetailIncludes(),
			ExcludeFields:    archiveDetailExcludes(),
			UseCustomCommand: true,
			Category:         "ftp",
			Order:            10,
		},
		"download": {
			Type:             schema.CommandManipulateSingle,
			Params:           archiveDownloadParam(),
			IncludeFields:    archiveDetailIncludes(),
			ExcludeFields:    archiveDetailExcludes(),
			UseCustomCommand: true,
			NoOutput:         true,
			Category:         "ftp",
			Order:            20,
		},
		"ftp-open": {
			Type:             schema.CommandManipulateMulti,
			Params:           archiveOpenFTPParam(),
			IncludeFields:    archiveDetailIncludes(),
			ExcludeFields:    archiveDetailExcludes(),
			UseCustomCommand: true,
			Category:         "ftp",
			Order:            30,
		},
		"ftp-close": {
			Type:             schema.CommandManipulateMulti,
			Params:           archiveCloseFTPParam(),
			IncludeFields:    archiveDetailIncludes(),
			ExcludeFields:    archiveDetailExcludes(),
			UseCustomCommand: true,
			NoOutput:         true,
			Category:         "ftp",
			Order:            40,
		},
		"wait-for-copy": {
			Type:             schema.CommandManipulateMulti,
			Params:           archiveWaitForCopyParam(),
			IncludeFields:    archiveDetailIncludes(),
			ExcludeFields:    archiveDetailExcludes(),
			UseCustomCommand: true,
			NoOutput:         true,
			NeedlessConfirm:  true,
			Category:         "other",
			Order:            10,
		},
	}

	return &schema.Resource{
		Commands:          commands,
		ResourceCategory:  CategoryStorage,
		CommandCategories: ArchiveCommandCategories,
	}
}

var ArchiveCommandCategories = []schema.Category{
	{
		Key:         "basics",
		DisplayName: "Basics",
		Order:       10,
	},
	{
		Key:         "ftp",
		DisplayName: "Upload/Download(SFTP)",
		Order:       20,
	},
	{
		Key:         "other",
		DisplayName: "Other",
		Order:       1000,
	},
}

func archiveListParam() map[string]*schema.Schema {
	return mergeParameterMap(
		CommonListParam,
		paramScopeCond,
		paramTagsCond,
		paramSourceArchiveIDCond,
		paramSourceDiskCond,
	)
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
		"source-disk-id": {
			Type:            schema.TypeInt64,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetSourceDisk",
			Description:     "set source disk ID",
			ValidateFunc:    validateSakuraID(),
			CompleteFunc:    completeDiskID(),
			ConflictsWith:   []string{"archive-file", "source-archive-id", "size"},
			Category:        "archive",
			Order:           10,
		},
		"source-archive-id": {
			Type:            schema.TypeInt64,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetSourceArchive",
			Description:     "set source archive ID",
			ValidateFunc:    validateSakuraID(),
			CompleteFunc:    completeArchiveID(),
			ConflictsWith:   []string{"archive-file", "source-disk-id", "size"},
			Category:        "archive",
			Order:           15,
		},
		"size": {
			Type:            schema.TypeInt,
			HandlerType:     schema.HandlerPathThrough,
			Description:     "set archive size(GB)",
			DestinationProp: "SetSizeGB",
			ValidateFunc:    validateInIntValues(allowSizes...),
			CompleteFunc:    completeInIntValues(allowSizes...),
			ConflictsWith:   []string{"source-archive-id", "source-disk-id"},
			Category:        "archive",
			Order:           20,
		},
		"archive-file": {
			Type:          schema.TypeString,
			HandlerType:   schema.HandlerNoop,
			Description:   "set archive image file",
			ValidateFunc:  validateFileExists(),
			ConflictsWith: []string{"source-archive-id", "source-disk-id"},
			Category:      "archive",
			Order:         30,
		},
	}
}

func archiveReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func archiveUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func archiveDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func archiveUploadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"archive-file": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set archive image file",
			Required:     true,
			ValidateFunc: validateFileExists(),
			Category:     "archive",
			Order:        10,
		},
	}
}

func archiveDownloadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"file-destination": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set file destination path",
			Required:    true,
			Category:    "archive",
			Order:       10,
		},
	}
}

func archiveWaitForCopyParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func archiveOpenFTPParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func archiveCloseFTPParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

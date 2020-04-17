// Copyright 2017-2020 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package define

import (
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/schema"
)

func ISOImageResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find", "select"},
			Params:             isoImageListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: isoImageListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"create": {
			Type:             schema.CommandCreate,
			Params:           isoImageCreateParam(),
			IncludeFields:    isoImageDetailIncludes(),
			ExcludeFields:    isoImageDetailExcludes(),
			UseCustomCommand: true,
			Category:         "basics",
			Order:            20,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        isoImageReadParam(),
			IncludeFields: isoImageDetailIncludes(),
			ExcludeFields: isoImageDetailExcludes(),
			Category:      "basics",
			Order:         30,
		},
		"update": {
			Type:          schema.CommandUpdate,
			Params:        isoImageUpdateParam(),
			IncludeFields: isoImageDetailIncludes(),
			ExcludeFields: isoImageDetailExcludes(),
			Category:      "basics",
			Order:         40,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"rm"},
			Params:        isoImageDeleteParam(),
			IncludeFields: isoImageDetailIncludes(),
			ExcludeFields: isoImageDetailExcludes(),
			Category:      "basics",
			Order:         50,
		},
		"upload": {
			Type:             schema.CommandManipulateSingle,
			Params:           isoImageUploadParam(),
			IncludeFields:    isoImageDetailIncludes(),
			ExcludeFields:    isoImageDetailExcludes(),
			UseCustomCommand: true,
			Category:         "ftp",
			Order:            10,
		},
		"download": {
			Type:             schema.CommandManipulateSingle,
			Params:           isoImageDownloadParam(),
			IncludeFields:    isoImageDetailIncludes(),
			ExcludeFields:    isoImageDetailExcludes(),
			UseCustomCommand: true,
			NoOutput:         true,
			Category:         "ftp",
			Order:            20,
		},
		"ftp-open": {
			Type:             schema.CommandManipulateMulti,
			Params:           isoImageOpenFTPParam(),
			IncludeFields:    isoImageDetailIncludes(),
			ExcludeFields:    isoImageDetailExcludes(),
			UseCustomCommand: true,
			Category:         "ftp",
			Order:            30,
		},
		"ftp-close": {
			Type:             schema.CommandManipulateMulti,
			Params:           isoImageCloseFTPParam(),
			IncludeFields:    isoImageDetailIncludes(),
			ExcludeFields:    isoImageDetailExcludes(),
			UseCustomCommand: true,
			NoOutput:         true,
			Category:         "ftp",
			Order:            40,
		},
	}

	return &schema.Resource{
		AltResource:         "CDROM",
		ListResultFieldName: "CDROMs",
		Commands:            commands,
		ResourceCategory:    CategoryStorage,
		CommandCategories:   isoImageCommandCategories,
	}
}

var isoImageCommandCategories = []schema.Category{
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

func isoImageListParam() map[string]*schema.Schema {
	return mergeParameterMap(CommonListParam, paramScopeCond, paramTagsCond)
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
		"size": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set iso size(GB)",
			Required:     true,
			DefaultValue: 5,
			ValidateFunc: validateInIntValues(5, 10),
			Category:     "ISO-Image",
			Order:        10,
		},
		"iso-file": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set iso image file",
			ValidateFunc: validateExistsFileOrStdIn(),
			Category:     "ISO-Image",
			Order:        20,
		},
		"name":        paramRequiredName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
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
			ValidateFunc: validateExistsFileOrStdIn(),
			Category:     "ISO-Image",
			Order:        10,
		},
	}
}

func isoImageDownloadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"file-destination": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set file destination path",
			Category:    "ISO-Image",
			Order:       10,
		},
	}
}

func isoImageOpenFTPParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func isoImageCloseFTPParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

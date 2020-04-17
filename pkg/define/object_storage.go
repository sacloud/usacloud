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

func ObjectStorageResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandCustom,
			Aliases:            []string{"ls"},
			Params:             objectStorageListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: objectStorageListColumns(),
			UseCustomCommand:   true,
			ArgsUsage:          "<remote path>",
			SkipAuth:           true,
			NeedlessConfirm:    true,
			Category:           "basics",
			Order:              10,
		},
		"put": {
			Type:             schema.CommandCustom,
			Params:           objectStoragePutParam(),
			UseCustomCommand: true,
			ArgsUsage:        "<local file/directory> <remote path>",
			SkipAuth:         true,
			NoOutput:         true,
			Category:         "basics",
			Order:            20,
		},
		"get": {
			Type:             schema.CommandCustom,
			Params:           objectStorageGetParam(),
			UseCustomCommand: true,
			ArgsUsage:        "<remote path> <local file/directory>",
			SkipAuth:         true,
			NoOutput:         true,
			NeedlessConfirm:  true,
			Category:         "basics",
			Order:            30,
		},
		"delete": {
			Type:             schema.CommandCustom,
			Aliases:          []string{"rm", "del"},
			Params:           objectStorageDelParam(),
			UseCustomCommand: true,
			ArgsUsage:        "<remote path>",
			SkipAuth:         true,
			NoOutput:         true,
			Category:         "basics",
			Order:            40,
		},
	}

	return &schema.Resource{
		Commands:         commands,
		Aliases:          []string{"ojs"},
		ResourceCategory: CategoryOther,
	}
}

func objectStorageListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Key"},
		{Name: "Size"},
		{Name: "ETag"},
	}
}

func objectStorageDetailIncludes() []string {
	return []string{}
}

func objectStorageDetailExcludes() []string {
	return []string{}
}

func objectStorageListParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access-key": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set access-key",
			EnvVars:     []string{"SACLOUD_OJS_ACCESS_KEY_ID", "AWS_ACCESS_KEY_ID"},
			Required:    true,
			Category:    "auth",
			Order:       10,
		},
		"secret-key": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set access-key",
			EnvVars:     []string{"SACLOUD_OJS_SECRET_ACCESS_KEY", "AWS_SECRET_ACCESS_KEY"},
			Required:    true,
			Category:    "auth",
			Order:       20,
		},
		"bucket": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set bucket",
			EnvVars:     []string{"SACLOUD_OJS_BUCKET_NAME"},
			Category:    "auth",
			Order:       30,
		},
	}
}

func objectStoragePutParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access-key": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set access-key",
			EnvVars:     []string{"SACLOUD_OJS_ACCESS_KEY_ID", "AWS_ACCESS_KEY_ID"},
			Required:    true,
			Category:    "auth",
			Order:       10,
		},
		"secret-key": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set access-key",
			EnvVars:     []string{"SACLOUD_OJS_SECRET_ACCESS_KEY", "AWS_SECRET_ACCESS_KEY"},
			Required:    true,
			Category:    "auth",
			Order:       20,
		},
		"bucket": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set bucket",
			EnvVars:     []string{"SACLOUD_OJS_BUCKET_NAME"},
			Category:    "auth",
			Order:       30,
		},
		"content-type": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set content-type",
			DefaultValue: "application/octet-stream",
			Category:     "operation",
			Order:        10,
		},
		"recursive": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"r"},
			Description: "put objects recursive",
			Category:    "operation",
			Order:       20,
		},
	}
}

func objectStorageGetParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access-key": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set access-key",
			EnvVars:     []string{"SACLOUD_OJS_ACCESS_KEY_ID", "AWS_ACCESS_KEY_ID"},
			Required:    true,
			Category:    "auth",
			Order:       10,
		},
		"secret-key": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set access-key",
			EnvVars:     []string{"SACLOUD_OJS_SECRET_ACCESS_KEY", "AWS_SECRET_ACCESS_KEY"},
			Required:    true,
			Category:    "auth",
			Order:       20,
		},
		"bucket": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set bucket",
			EnvVars:     []string{"SACLOUD_OJS_BUCKET_NAME"},
			Category:    "auth",
			Order:       30,
		},
		"recursive": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"r"},
			Description: "get objects recursive",
			Category:    "operation",
			Order:       10,
		},
	}
}

func objectStorageDelParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access-key": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set access-key",
			EnvVars:     []string{"SACLOUD_OJS_ACCESS_KEY_ID", "AWS_ACCESS_KEY_ID"},
			Required:    true,
			Category:    "auth",
			Order:       10,
		},
		"secret-key": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set access-key",
			EnvVars:     []string{"SACLOUD_OJS_SECRET_ACCESS_KEY", "AWS_SECRET_ACCESS_KEY"},
			Required:    true,
			Category:    "auth",
			Order:       20,
		},
		"bucket": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set bucket",
			EnvVars:     []string{"SACLOUD_OJS_BUCKET_NAME"},
			Category:    "auth",
			Order:       30,
		},
		"recursive": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"r"},
			Description: "delete objects recursive",
			Category:    "operation",
			Order:       10,
		},
	}
}

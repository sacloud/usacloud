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

func LicenseResource() *schema.Resource {
	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find"},
			Params:             licenseListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: licenseListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"create": {
			Type:          schema.CommandCreate,
			Params:        licenseCreateParam(),
			IncludeFields: licenseDetailIncludes(),
			ExcludeFields: licenseDetailExcludes(),
			Category:      "basics",
			Order:         20,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        licenseReadParam(),
			IncludeFields: licenseDetailIncludes(),
			ExcludeFields: licenseDetailExcludes(),
			NoSelector:    true,
			Category:      "basics",
			Order:         30,
		},
		"update": {
			Type:          schema.CommandUpdate,
			Params:        licenseUpdateParam(),
			IncludeFields: licenseDetailIncludes(),
			ExcludeFields: licenseDetailExcludes(),
			NoSelector:    true,
			Category:      "basics",
			Order:         40,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"rm"},
			Params:        licenseDeleteParam(),
			IncludeFields: licenseDetailIncludes(),
			ExcludeFields: licenseDetailExcludes(),
			NoSelector:    true,
			Category:      "basics",
			Order:         50,
		},
	}

	return &schema.Resource{
		Commands:            commands,
		ResourceCategory:    CategoryCommonItem,
		ListResultFieldName: "Licenses",
		IsGlobal:            true,
	}
}

func licenseListParam() map[string]*schema.Schema {
	return CommonListParam
}

func licenseListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{
			Name:    "LicenseInfo:ID",
			Sources: []string{"LicenseInfo.ID"},
		},
		{
			Name:    "LicenseInfo:Name",
			Sources: []string{"LicenseInfo.Name"},
		},
	}
}

func licenseDetailIncludes() []string {
	return []string{}
}

func licenseDetailExcludes() []string {
	return []string{}
}

func licenseCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": paramRequiredName,
		"license-info-id": {
			Type:            schema.TypeId,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "LicenseInfoID",
			Description:     "set LicenseInfo ID",
			Category:        "license",
			Order:           10,
		},
	}
}

func licenseReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func licenseUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": paramName,
	}
}

func licenseDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

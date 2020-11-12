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

package disk

import (
	"reflect"

	"github.com/sacloud/usacloud/pkg/output"

	"github.com/sacloud/libsacloud/v2/helper/service/disk"
	"github.com/sacloud/usacloud/pkg/cmd/core"
)

var Resource = &core.Resource{
	Name:        "disk",
	ServiceType: reflect.TypeOf(&disk.Service{}),
	Category:    core.ResourceCategoryStorage,
	CommandCategories: []core.Category{
		{
			Key:         "basics",
			DisplayName: "Basics",
			Order:       10,
		},
		{
			Key:         "edit",
			DisplayName: "Disk Edit",
			Order:       20,
		},
		{
			Key:         "re-install",
			DisplayName: "Re-Install",
			Order:       25,
		},
		{
			Key:         "server",
			DisplayName: "Server Connection Management",
			Order:       30,
		},
		{
			Key:         "monitor",
			DisplayName: "Monitoring",
			Order:       40,
		},
		{
			Key:         "other",
			DisplayName: "Other",
			Order:       1000,
		},
	},
}

var defaultColumnDefs = []output.ColumnDef{
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
	{
		Name:    "Storage",
		Sources: []string{"Storage.Name"},
	},
}

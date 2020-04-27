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
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/schema"
)

func AutoBackupResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find", "select"},
			Params:             autoBackupListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: autoBackupListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"create": {
			Type:          schema.CommandCreate,
			Params:        autoBackupCreateParam(),
			IncludeFields: autoBackupDetailIncludes(),
			ExcludeFields: autoBackupDetailExcludes(),
			Category:      "basics",
			Order:         20,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        autoBackupReadParam(),
			IncludeFields: autoBackupDetailIncludes(),
			ExcludeFields: autoBackupDetailExcludes(),
			Category:      "basics",
			Order:         30,
		},
		"update": {
			Type:          schema.CommandUpdate,
			Params:        autoBackupUpdateParam(),
			IncludeFields: autoBackupDetailIncludes(),
			ExcludeFields: autoBackupDetailExcludes(),
			Category:      "basics",
			Order:         40,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"rm"},
			Params:        autoBackupDeleteParam(),
			IncludeFields: autoBackupDetailIncludes(),
			ExcludeFields: autoBackupDetailExcludes(),
			Category:      "basics",
			Order:         50,
		},
	}

	return &schema.Resource{
		Commands:            commands,
		ResourceCategory:    CategoryStorage,
		ListResultFieldName: "AutoBackups",
	}
}

func autoBackupListParam() map[string]*schema.Schema {
	return mergeParameterMap(CommonListParam, paramTagsCond)
}

func autoBackupListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{
			Name:    "DiskID",
			Sources: []string{"Status.DiskId"},
		},
		{
			Name:    "StartHour",
			Sources: []string{"Settings.Autobackup.BackupHour"},
		},
		{
			Name:    "Generation",
			Sources: []string{"Settings.Autobackup.MaximumNumberOfArchives"},
		},
		{
			Name: "Weekdays",
			Sources: []string{
				"Settings.Autobackup.BackupSpanWeekdays.0",
				"Settings.Autobackup.BackupSpanWeekdays.1",
				"Settings.Autobackup.BackupSpanWeekdays.2",
				"Settings.Autobackup.BackupSpanWeekdays.3",
				"Settings.Autobackup.BackupSpanWeekdays.4",
				"Settings.Autobackup.BackupSpanWeekdays.5",
				"Settings.Autobackup.BackupSpanWeekdays.6",
			},
			Format: "%s %s %s %s %s %s %s",
		},
	}
}

func autoBackupDetailIncludes() []string {
	return []string{}
}

func autoBackupDetailExcludes() []string {
	return []string{}
}

func autoBackupCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramRequiredName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
		"disk-id": {
			Type:         schema.TypeId,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target diskID ",
			Required:     true,
			ValidateFunc: validateSakuraID(),
			Category:     "backup",
			Order:        10,
		},
		"weekdays": {
			Type:               schema.TypeStringList,
			HandlerType:        schema.HandlerPathThrough,
			Description:        "set backup target weekdays[all or mon/tue/wed/thu/fri/sat/sun]",
			DestinationProp:    "BackupSpanWeekdays",
			DestinationWrapper: "weekday.FromStrings",
			ValidateFunc: validateStringSlice(
				validateInStrValues(append(types.BackupWeekdayStrings, "all")...),
			),
			DefaultValue: []string{"all"},
			Required:     true,
			Category:     "backup",
			Order:        20,
		},
		"generation": {
			Type:            schema.TypeInt,
			HandlerType:     schema.HandlerPathThrough,
			Description:     "set backup generation[1-10]",
			DestinationProp: "MaximumNumberOfArchives",
			ValidateFunc:    validateIntRange(1, 10),
			DefaultValue:    1,
			Required:        true,
			Category:        "backup",
			Order:           30,
		},
	}
}

func autoBackupReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func autoBackupUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
		"weekdays": {
			Type:               schema.TypeStringList,
			HandlerType:        schema.HandlerPathThrough,
			Description:        "set backup target weekdays[all or mon/tue/wed/thu/fri/sat/sun]",
			DestinationProp:    "BackupSpanWeekdays",
			DestinationWrapper: "weekday.FromStrings",
			ValidateFunc: validateStringSlice(
				validateInStrValues(append(types.BackupWeekdayStrings, "all")...),
			),
			Category: "backup",
			Order:    20,
		},
		"generation": {
			Type:            schema.TypeInt,
			HandlerType:     schema.HandlerPathThrough,
			Description:     "set backup generation[1-10]",
			DestinationProp: "MaximumNumberOfArchives",
			ValidateFunc:    validateIntRange(1, 10),
			Category:        "backup",
			Order:           30,
		},
	}
}

func autoBackupDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

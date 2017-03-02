package define

import (
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func AutoBackupResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"l", "ls", "find"},
			Params:             autoBackupListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: autoBackupListColumns(),
		},
		"create": {
			Type:             schema.CommandCreate,
			Aliases:          []string{"c"},
			Params:           autoBackupCreateParam(),
			IncludeFields:    autoBackupDetailIncludes(),
			ExcludeFields:    autoBackupDetailExcludes(),
			UseCustomCommand: true,
		},
		"read": {
			Type:          schema.CommandRead,
			Aliases:       []string{"r"},
			Params:        autoBackupReadParam(),
			IncludeFields: autoBackupDetailIncludes(),
			ExcludeFields: autoBackupDetailExcludes(),
		},
		"update": {
			Type:             schema.CommandUpdate,
			Aliases:          []string{"u"},
			Params:           autoBackupUpdateParam(),
			IncludeFields:    autoBackupDetailIncludes(),
			ExcludeFields:    autoBackupDetailExcludes(),
			UseCustomCommand: true,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"d", "rm"},
			Params:        autoBackupDeleteParam(),
			IncludeFields: autoBackupDetailIncludes(),
			ExcludeFields: autoBackupDetailExcludes(),
		},
	}

	return &schema.Resource{
		Commands:            commands,
		ResourceCategory:    CategoryStorage,
		ListResultFieldName: "CommonServiceAutoBackupItems",
	}
}

func autoBackupListParam() map[string]*schema.Schema {
	return CommonListParam
}

func autoBackupListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{
			Name:    "Disk-ID",
			Sources: []string{"Status.DiskId"},
		},
		{
			Name:    "Start-Hour",
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
		"generation": {
			Type:            schema.TypeInt,
			HandlerType:     schema.HandlerPathThrough,
			Description:     "set backup generation[1-10]",
			DestinationProp: "SetBackupMaximumNumberOfArchives",
			ValidateFunc:    validateIntRange(1, 10),
			DefaultValue:    1,
			Required:        true,
		},
		"start-hour": {
			Type:            schema.TypeInt,
			HandlerType:     schema.HandlerPathThrough,
			Description:     "set backup start hour[0/6/12/18]",
			DestinationProp: "SetBackupHour",
			ValidateFunc:    validateInIntValues(sacloud.AllowAutoBackupHour()...),
			CompleteFunc:    completeInIntValues(sacloud.AllowAutoBackupHour()...),
			DefaultValue:    0,
			Required:        true,
		},
		"weekdays": {
			Type:            schema.TypeStringList,
			HandlerType:     schema.HandlerPathThrough,
			Description:     "set backup target weekdays[all or mon/tue/wed/thu/fri/sat/sun]",
			DestinationProp: "SetBackupSpanWeekdays",
			ValidateFunc: validateStringSlice(
				validateInStrValues(append(sacloud.AllowAutoBackupWeekdays(), "all")...),
			),
			CompleteFunc: completeInStrValues(append(sacloud.AllowAutoBackupWeekdays(), "all")...),
			DefaultValue: []string{"all"},
			Required:     true,
		},
		"disk-id": {
			Type:         schema.TypeInt64,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target diskID ",
			Required:     true,
			ValidateFunc: validateSakuraID(),
			CompleteFunc: completeDiskID(),
		},
	}
}

func autoBackupReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func autoBackupUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id":          paramID,
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
		"generation": {
			Type:            schema.TypeInt,
			HandlerType:     schema.HandlerPathThrough,
			Description:     "set backup generation[1-10]",
			DestinationProp: "SetBackupMaximumNumberOfArchives",
			ValidateFunc:    validateIntRange(1, 10),
		},
		"start-hour": {
			Type:            schema.TypeInt,
			HandlerType:     schema.HandlerPathThrough,
			Description:     "set backup start hour[0/6/12/18]",
			DestinationProp: "SetBackupHour",
			ValidateFunc:    validateInIntValues(sacloud.AllowAutoBackupHour()...),
			CompleteFunc:    completeInIntValues(sacloud.AllowAutoBackupHour()...),
		},
		"weekdays": {
			Type:            schema.TypeStringList,
			HandlerType:     schema.HandlerPathThrough,
			Description:     "set backup target weekdays[all or mon/tue/wed/thu/fri/sat/sun]",
			DestinationProp: "SetBackupSpanWeekdays",
			ValidateFunc: validateStringSlice(
				validateInStrValues(append(sacloud.AllowAutoBackupWeekdays(), "all")...),
			),
			CompleteFunc: completeInStrValues(append(sacloud.AllowAutoBackupWeekdays(), "all")...),
		},
	}
}

func autoBackupDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

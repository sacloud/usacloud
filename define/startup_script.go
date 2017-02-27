package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func StartupScriptResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:                schema.CommandList,
			ListResultFieldName: "Notes",
			Aliases:             []string{"l", "ls", "find"},
			Params:              startupScriptListParam(),
			TableType:           output.TableSimple,
			TableColumnDefines:  startupScriptListColumns(),
		},
		"create": {
			Type:             schema.CommandCreate,
			Aliases:          []string{"c"},
			Params:           startupScriptCreateParam(),
			IncludeFields:    startupScriptDetailIncludes(),
			ExcludeFields:    startupScriptDetailExcludes(),
			UseCustomCommand: true,
		},
		"read": {
			Type:          schema.CommandRead,
			Aliases:       []string{"r"},
			Params:        startupScriptReadParam(),
			IncludeFields: startupScriptDetailIncludes(),
			ExcludeFields: startupScriptDetailExcludes(),
		},
		"update": {
			Type:             schema.CommandUpdate,
			Aliases:          []string{"u"},
			Params:           startupScriptUpdateParam(),
			IncludeFields:    startupScriptDetailIncludes(),
			ExcludeFields:    startupScriptDetailExcludes(),
			UseCustomCommand: true,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"d", "rm"},
			Params:        startupScriptDeleteParam(),
			IncludeFields: startupScriptDetailIncludes(),
			ExcludeFields: startupScriptDetailExcludes(),
		},
	}

	return &schema.Resource{
		Commands:         commands,
		Aliases:          []string{"note"},
		AltResource:      "Note",
		ResourceCategory: CategoryCommonItem,
	}
}

func startupScriptListParam() map[string]*schema.Schema {
	return mergeParameterMap(CommonListParam, paramScopeCond)
}

func startupScriptListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{Name: "Scope"},
	}
}

func startupScriptDetailIncludes() []string {
	return []string{}
}

func startupScriptDetailExcludes() []string {
	return []string{}
}

func startupScriptCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":    paramRequiredName,
		"tags":    paramTags,
		"icon-id": getParamSubResourceID("Icon"),
		"script-content": {
			Type:          schema.TypeString,
			HandlerType:   schema.HandlerNoop,
			Aliases:       []string{"note-content"},
			Description:   "set script content",
			ConflictsWith: []string{"script"},
		},
		"script": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"note"},
			Description:  "set script from file",
			ValidateFunc: validateFileExists(),
		},
	}
}

func startupScriptReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func startupScriptUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id":      paramID,
		"name":    paramRequiredName,
		"tags":    paramTags,
		"icon-id": getParamSubResourceID("Icon"),
		"script-content": {
			Type:          schema.TypeString,
			HandlerType:   schema.HandlerNoop,
			Aliases:       []string{"note-content"},
			Description:   "set script content",
			ConflictsWith: []string{"script"},
		},
		"script": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Aliases:      []string{"note"},
			Description:  "set script from file",
			ValidateFunc: validateFileExists(),
		},
	}
}

func startupScriptDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

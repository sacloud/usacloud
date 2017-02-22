package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func WebAccelResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"delete-cache": {
			Type:               schema.CommandManipulate,
			Aliases:            []string{"purge"},
			Params:             webAccelDeleteCacheParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: webAccelListColumns(),
			UseCustomCommand:   true,
		},
	}

	return &schema.Resource{
		Commands: commands,
	}
}
func webAccelListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Result"},
		{Name: "Status"},
		{Name: "URL"},
	}
}

func webAccelDeleteCacheParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"url": {
			Type:        schema.TypeStringList,
			HandlerType: schema.HandlerNoop,
			Required:    true,
			Description: "set delete-cache(purge) targets",
		},
	}
}

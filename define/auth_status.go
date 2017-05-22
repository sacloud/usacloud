package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func AuthStatusResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"show": {
			Type:               schema.CommandCustom,
			Params:             authShowParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: authShowColumns(),
			UseCustomCommand:   true,
			NeedlessConfirm:    true,
		},
	}

	return &schema.Resource{
		Commands:         commands,
		DefaultCommand:   "show",
		ResourceCategory: CategoryAuth,
	}
}

func authShowParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func authShowColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{
			Name:    "AccountID",
			Sources: []string{"Account.ID"},
		},
		{
			Name:    "AccountCode",
			Sources: []string{"Account.Code"},
		},
		{
			Name:    "AccountName",
			Sources: []string{"Account.Name"},
		},
		{
			Name:    "MemberCode",
			Sources: []string{"Member.Code"},
		},
		{Name: "Permission"},
		{Name: "ExternalPermission"},
	}
}

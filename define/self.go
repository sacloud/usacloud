package define

import (
	"github.com/sacloud/usacloud/schema"
)

func SelfResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"info": {
			Type:             schema.CommandCustom,
			Params:           selfInfoParam(),
			UseCustomCommand: true,
			NoOutput:         true,
			NeedlessConfirm:  true,
			SkipAuth:         true,
		},
	}

	return &schema.Resource{
		Commands:         commands,
		DefaultCommand:   "info",
		ResourceCategory: CategoryOther,
		Usage:            "Show self info",
	}
}

func selfInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

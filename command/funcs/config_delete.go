package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/command/profile"
)

func ConfigDelete(ctx command.Context, params *params.DeleteConfigParam) error {
	if ctx.NArgs() == 0 || ctx.Args()[0] == "" {
		return fmt.Errorf("Profile name is required")
	}

	profileName := ctx.Args()[0]
	return profile.RemoveConfigFile(profileName)
}

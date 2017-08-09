package completion

import (
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ConfigShowCompleteArgs(ctx command.Context, params *params.ShowConfigParam, cur, prev, commandName string) {
	writeAllProfileName(command.GlobalOption.Out)
}

package completion

import (
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ConfigEditCompleteArgs(ctx command.Context, params *params.EditConfigParam, cur, prev, commandName string) {
	writeAllProfileName(command.GlobalOption.Out)
}

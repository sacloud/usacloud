package completion

import (
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ConfigDeleteCompleteArgs(ctx command.Context, params *params.DeleteConfigParam, cur, prev, commandName string) {
	writeAllProfileName(command.GlobalOption.Out)
}

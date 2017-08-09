package completion

import (
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ConfigUseCompleteArgs(ctx command.Context, params *params.UseConfigParam, cur, prev, commandName string) {
	writeAllProfileName(command.GlobalOption.Out)
}

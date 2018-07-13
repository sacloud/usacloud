package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/command/profile"
)

func ConfigList(ctx command.Context, params *params.ListConfigParam) error {
	profiles, err := profile.List()
	if err != nil {
		return err
	}

	for _, p := range profiles {
		fmt.Fprintln(command.GlobalOption.Out, p)
	}
	return nil
}

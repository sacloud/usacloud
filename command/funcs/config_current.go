package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/command/profile"
)

func ConfigCurrent(ctx command.Context, params *params.CurrentConfigParam) error {

	p, err := profile.GetCurrentName()
	if err != nil {
		return fmt.Errorf("ConfigCurrent is failed: %s", err)
	}

	fmt.Fprintf(command.GlobalOption.Out, "%s\n", p)
	return nil
}

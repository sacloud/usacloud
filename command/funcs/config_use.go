package funcs

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/command/profile"
	"github.com/sacloud/usacloud/helper/printer"
)

func ConfigUse(ctx command.Context, params *params.UseConfigParam) error {

	if ctx.NArgs() == 0 || ctx.Args()[0] == "" {
		return fmt.Errorf("Profile name is required")
	}

	profileName := ctx.Args()[0]
	if err := profile.SetCurrentName(profileName); err != nil {
		return err
	}

	printer.Fprintf(command.GlobalOption.Out, color.New(color.FgHiGreen), "\nCurrent profile: %q\n", profileName)
	return nil
}

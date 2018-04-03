package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/command/profile"
)

func ConfigShow(ctx command.Context, params *params.ShowConfigParam) error {

	profileName := ""
	if ctx.NArgs() == 0 {
		n, err := profile.GetCurrentName()
		if err != nil {
			return err
		}
		profileName = n
	} else {
		profileName = ctx.Args()[0]
	}

	conf, err := profile.LoadConfigFile(profileName)
	if err != nil {
		return err
	}

	out := command.GlobalOption.Out
	fmt.Fprintf(out, "\n")
	fmt.Fprintf(out, "token               = %s\n", conf.AccessToken)
	fmt.Fprintf(out, "secret              = %s\n", conf.AccessTokenSecret)
	fmt.Fprintf(out, "zone                = %s\n", conf.Zone)
	fmt.Fprintf(out, "default-output-type = %s\n", conf.DefaultOutputType)
	fmt.Fprintf(out, "\n")
	return nil
}

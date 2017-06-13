package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterUserUpdate(ctx command.Context, params *params.UserUpdateVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterUserUpdate is failed: %s", e)
	}

	if !p.HasRemoteAccessUsers() {
		return fmt.Errorf("VPCRouter[%d] don't have any remote-access users", params.Id)
	}

	// validate
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.Router.RemoteAccessUsers.Config) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	cnf := p.Settings.Router.RemoteAccessUsers.Config[params.Index-1]
	if ctx.IsSet("username") {
		cnf.UserName = params.Username
	}
	if ctx.IsSet("password") {
		cnf.Password = params.Password
	}

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterUserUpdate is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterUserUpdate is failed: %s", err)
	}

	return nil

}

package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterUserDelete(ctx command.Context, params *params.UserDeleteVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterUserDelete is failed: %s", e)
	}

	if !p.HasRemoteAccessUsers() {
		return fmt.Errorf("VPCRouter[%d] don't have any remote-access users", params.Id)
	}

	// validate
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.Router.RemoteAccessUsers.Config) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	p.Settings.Router.RemoveRemoteAccessUserAt(params.Index - 1)

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterPortForwardingDelete is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterPortForwardingDelete is failed: %s", err)
	}

	return nil
}

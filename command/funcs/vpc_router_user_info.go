package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterUserInfo(ctx command.Context, params *params.UserInfoVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterUserInfo is failed: %s", e)
	}

	if !p.HasRemoteAccessUsers() {
		fmt.Fprintf(command.GlobalOption.Err, "VPCRouter[%d] don't have any remote-access users\n", params.Id)
		return nil
	}

	confList := p.Settings.Router.RemoteAccessUsers.Config
	// build parameters to display table
	list := []interface{}{}
	for i := range confList {
		list = append(list, &confList[i])
	}

	return ctx.GetOutput().Print(list...)
}

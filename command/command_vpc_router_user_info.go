package command

import (
	"fmt"
)

func VPCRouterUserInfo(ctx Context, params *UserInfoVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterUserInfo is failed: %s", e)
	}

	if !p.HasRemoteAccessUsers() {
		fmt.Fprintf(GlobalOption.Err, "VPCRouter[%d] don't have any remote-access users\n", params.Id)
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

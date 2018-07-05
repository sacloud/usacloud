package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterEnableInternetConnection(ctx command.Context, params *params.EnableInternetConnectionVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterEnableInternetConnection is failed: %s", e)
	}

	if !p.HasSetting() {
		p.InitVPCRouterSetting()
	}

	if p.Settings.Router.InternetConnection == nil {
		p.Settings.Router.InternetConnection = &sacloud.VPCRouterInternetConnection{
			Enabled: "True",
		}
	}
	p.Settings.Router.InternetConnection.Enabled = "True"

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterEnableInternetConnection is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterEnableInternetConnection is failed: %s", err)
	}
	return nil

}

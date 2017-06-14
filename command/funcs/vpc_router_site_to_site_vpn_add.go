package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterSiteToSiteVpnAdd(ctx command.Context, params *params.SiteToSiteVpnAddVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterSiteToSiteVpnAdd is failed: %s", e)
	}

	if !p.HasSetting() {
		p.InitVPCRouterSetting()
	}

	p.Settings.Router.AddSiteToSiteIPsecVPN(
		params.LocalPrefix,
		params.Peer,
		params.PreSharedSecret,
		params.RemoteId,
		params.Routes,
	)

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterSiteToSiteVpnAdd is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterSiteToSiteVpnAdd is failed: %s", err)
	}
	return nil

}

package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterSiteToSiteVpnUpdate(ctx command.Context, params *params.SiteToSiteVpnUpdateVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterSiteToSiteVpnUpdate is failed: %s", e)
	}

	if !p.HasSiteToSiteIPsecVPN() {
		return fmt.Errorf("VPCRouter[%d] don't have any site-to-site IPSec VPN settings", params.Id)
	}

	// validate
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.Router.SiteToSiteIPsecVPN.Config) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	cnf := p.Settings.Router.SiteToSiteIPsecVPN.Config[params.Index-1]
	if ctx.IsSet("peer") {
		cnf.Peer = params.Peer
	}
	if ctx.IsSet("remote-id") {
		cnf.RemoteID = params.RemoteId
	}
	if ctx.IsSet("pre-shared-secret") {
		cnf.PreSharedSecret = params.PreSharedSecret
	}
	if ctx.IsSet("routes") {
		cnf.Routes = params.Routes
	}
	if ctx.IsSet("local-prefix") {
		cnf.LocalPrefix = params.LocalPrefix
	}

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterSiteToSiteVpnUpdate is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterSiteToSiteVpnUpdate is failed: %s", err)
	}

	return nil

}

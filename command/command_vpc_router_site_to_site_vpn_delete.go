package command

import (
	"fmt"
)

func VPCRouterSiteToSiteVpnDelete(ctx Context, params *SiteToSiteVpnDeleteVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterSiteToSiteVpnDelete is failed: %s", e)
	}

	if !p.HasSiteToSiteIPsecVPN() {
		return fmt.Errorf("VPCRouter[%d] don't have any site-to-site IPSec VPN settings", params.Id)
	}

	// validate
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.Router.SiteToSiteIPsecVPN.Config) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	p.Settings.Router.RemoveSiteToSiteIPsecVPNAt(params.Index - 1)

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterSiteToSiteVpnDelete is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterSiteToSiteVpnDelete is failed: %s", err)
	}

	return nil

}

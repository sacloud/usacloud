package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterSiteToSiteVpnPeers(ctx command.Context, params *params.SiteToSiteVpnPeersVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterSiteToSiteVpnPeers is failed: %s", e)
	}

	if !p.HasSiteToSiteIPsecVPN() {
		fmt.Fprintf(command.GlobalOption.Err, "VPCRouter[%d] don't have any site-to-site IPSec VPN settings\n", params.Id)
		return nil
	}

	status, err := api.Status(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterSiteToSiteVpnPeers is failed: %s", err)
	}

	// build parameters to display table
	list := []interface{}{}
	for i := range status.SiteToSiteIPsecVPNPeers {
		list = append(list, &status.SiteToSiteIPsecVPNPeers[i])
	}

	return ctx.GetOutput().Print(list...)
}

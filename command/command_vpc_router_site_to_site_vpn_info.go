package command

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
	"strings"
)

func VPCRouterSiteToSiteVpnInfo(ctx Context, params *SiteToSiteVpnInfoVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterSiteToSiteVpnInfo is failed: %s", e)
	}

	if !p.HasSiteToSiteIPsecVPN() {
		fmt.Fprintf(GlobalOption.Err, "VPCRouter[%d] don't have any site-to-site IPSec VPN settings\n", params.Id)
		return nil
	}

	confList := p.Settings.Router.SiteToSiteIPsecVPN.Config

	type s2sSetting struct {
		*sacloud.VPCRouterSiteToSiteIPsecVPNConfig
		RoutesJoined      string
		LocalPrefixJoined string
	}

	// build parameters to display table
	list := []interface{}{}
	for i := range confList {
		s := &s2sSetting{
			VPCRouterSiteToSiteIPsecVPNConfig: confList[i],
			RoutesJoined:                      strings.Join(confList[i].Routes, ","),
			LocalPrefixJoined:                 strings.Join(confList[i].LocalPrefix, ","),
		}
		list = append(list, s)
	}

	return ctx.GetOutput().Print(list...)

}

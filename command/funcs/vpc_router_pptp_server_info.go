package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterPptpServerInfo(ctx command.Context, params *params.PptpServerInfoVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterPptpServerInfo is failed: %s", e)
	}

	type pptpConf struct {
		*sacloud.VPCRouterPPTPServerConfig
		Enabled string
	}

	var cnf *pptpConf
	if p.HasPPTPServer() {
		cnf = &pptpConf{
			VPCRouterPPTPServerConfig: p.Settings.Router.PPTPServer.Config,
			Enabled:                   p.Settings.Router.PPTPServer.Enabled,
		}
	} else {
		cnf = &pptpConf{
			Enabled: "False",
		}
	}

	return ctx.GetOutput().Print(cnf)

}

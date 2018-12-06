package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterL2tpServerInfo(ctx command.Context, params *params.L2tpServerInfoVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterL2tpServerInfo is failed: %s", e)
	}

	type l2tpConf struct {
		*sacloud.VPCRouterL2TPIPsecServerConfig
		Enabled string
	}

	var cnf *l2tpConf
	if p.HasL2TPIPsecServer() {
		cnf = &l2tpConf{
			VPCRouterL2TPIPsecServerConfig: p.Settings.Router.L2TPIPsecServer.Config,
			Enabled:                        p.Settings.Router.L2TPIPsecServer.Enabled,
		}
	} else {
		cnf = &l2tpConf{
			Enabled: "False",
		}
	}

	return ctx.GetOutput().Print(cnf)
}

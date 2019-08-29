package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ProxyLBCreate(ctx command.Context, params *params.CreateProxyLBParam) error {

	// Validate params
	if ctx.IsSet("sorry-server-ipaddress") || ctx.IsSet("sorry-server-port") {
		if params.SorryServerIpaddress == "" || params.SorryServerPort == 0 {
			return fmt.Errorf("both of sorry-server-ipaddress and sorry-server-port are required")
		}
	}

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()
	p := api.New(params.Name)

	// set params
	p.SetPlan(sacloud.ProxyLBPlan(params.Plan))
	p.SetDescription(params.Description)
	p.SetTags(params.Tags)
	p.SetIconByID(params.IconId)

	protocol := params.Protocol
	switch protocol {
	case "http":
		p.SetHTTPHealthCheck(params.HostHeader, params.Path, params.DelayLoop)
	case "tcp":
		p.SetTCPHealthCheck(params.DelayLoop)
	default:
		return fmt.Errorf("invalid protocol: %s", protocol)
	}

	if params.StickySession {
		p.Settings.ProxyLB.StickySession = sacloud.ProxyLBSessionSetting{
			Enabled: true,
			Method:  sacloud.ProxyLBStickySessionDefaultMethod,
		}
	}

	p.SetSorryServer(params.SorryServerIpaddress, params.SorryServerPort)

	p.Settings.ProxyLB.Timeout = &sacloud.ProxyLBTimeout{
		InactiveSec: params.Timeout,
	}

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("ProxyLBCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}

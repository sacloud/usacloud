package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ProxyLBServerUpdate(ctx command.Context, params *params.ServerUpdateProxyLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ProxyLBServerUpdate is failed: %s", e)
	}

	if len(p.Settings.ProxyLB.Servers) == 0 {
		return fmt.Errorf("ProxyLB don't have any servers")
	}

	// validate index
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.ProxyLB.Servers) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	server := &p.Settings.ProxyLB.Servers[params.Index-1]

	if ctx.IsSet("ipaddress") {
		server.IPAddress = params.Ipaddress
	}
	if ctx.IsSet("port") {
		server.Port = params.Port
	}
	if ctx.IsSet("disabled") {
		server.Enabled = !params.Disabled
	}

	p, e = api.UpdateSetting(params.Id, p)
	if e != nil {
		return fmt.Errorf("ProxyLBServerUpdate is failed: %s", e)
	}

	var list []interface{}
	for i := range p.Settings.ProxyLB.Servers {
		list = append(list, &p.Settings.ProxyLB.Servers[i])
	}
	return ctx.GetOutput().Print(list...)

}

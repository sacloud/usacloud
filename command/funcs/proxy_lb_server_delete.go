package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ProxyLBServerDelete(ctx command.Context, params *params.ServerDeleteProxyLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ProxyLBServerDelete is failed: %s", e)
	}

	if len(p.Settings.ProxyLB.Servers) == 0 {
		return fmt.Errorf("ProxyLB don't have any servers")
	}

	// validate index
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.ProxyLB.Servers) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	serverIndex := params.Index - 1
	var updServers []sacloud.ProxyLBServer
	for i := range p.Settings.ProxyLB.Servers {
		if i != serverIndex {
			updServers = append(updServers, p.Settings.ProxyLB.Servers[i])
		}
	}

	p.Settings.ProxyLB.Servers = updServers

	p, e = api.UpdateSetting(params.Id, p)
	if e != nil {
		return fmt.Errorf("ProxyLBServerDelete is failed: %s", e)
	}

	var list []interface{}
	for i := range p.Settings.ProxyLB.Servers {
		list = append(list, &p.Settings.ProxyLB.Servers[i])
	}
	return ctx.GetOutput().Print(list...)
}

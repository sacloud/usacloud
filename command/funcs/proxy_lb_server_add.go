package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ProxyLBServerAdd(ctx command.Context, params *params.ServerAddProxyLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ProxyLBServerAdd is failed: %s", e)
	}

	// check duplicate
	for _, server := range p.Settings.ProxyLB.Servers {
		if server.IPAddress == params.Ipaddress && server.Port == params.Port {
			return fmt.Errorf("server %s/%d is already exists", server.IPAddress, server.Port)
		}
	}

	p.AddServer(params.Ipaddress, params.Port, !params.Disabled)

	p, e = api.UpdateSetting(params.Id, p)
	if e != nil {
		return fmt.Errorf("ProxyLBServerAdd is failed: %s", e)
	}

	var list []interface{}
	for i := range p.Settings.ProxyLB.Servers {
		list = append(list, &p.Settings.ProxyLB.Servers[i])
	}
	return ctx.GetOutput().Print(list...)
}

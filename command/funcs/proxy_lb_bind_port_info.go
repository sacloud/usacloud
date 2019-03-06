package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ProxyLBBindPortInfo(ctx command.Context, params *params.BindPortInfoProxyLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ProxyLBBindPortInfo is failed: %s", e)
	}

	var list []interface{}
	for i := range p.Settings.ProxyLB.BindPorts {
		list = append(list, &p.Settings.ProxyLB.BindPorts[i])
	}
	return ctx.GetOutput().Print(list...)
}

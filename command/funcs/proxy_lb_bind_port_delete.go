package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ProxyLBBindPortDelete(ctx command.Context, params *params.BindPortDeleteProxyLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ProxyLBBindPortDelete is failed: %s", e)
	}
	if len(p.Settings.ProxyLB.BindPorts) == 0 {
		return fmt.Errorf("ProxyLB don't have any bind-ports")
	}

	// validate index
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.ProxyLB.BindPorts) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	index := params.Index - 1
	var updPorts []*sacloud.ProxyLBBindPorts
	for i := range p.Settings.ProxyLB.BindPorts {
		if i != index {
			updPorts = append(updPorts, p.Settings.ProxyLB.BindPorts[i])
		}
	}

	p.Settings.ProxyLB.BindPorts = updPorts

	p, e = api.UpdateSetting(params.Id, p)
	if e != nil {
		return fmt.Errorf("ProxyLBBindPortDelete is failed: %s", e)
	}

	var list []interface{}
	for i := range p.Settings.ProxyLB.BindPorts {
		list = append(list, &p.Settings.ProxyLB.BindPorts[i])
	}
	return ctx.GetOutput().Print(list...)

}

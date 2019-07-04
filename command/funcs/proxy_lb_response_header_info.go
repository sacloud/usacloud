package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ProxyLBResponseHeaderInfo(ctx command.Context, params *params.ResponseHeaderInfoProxyLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ProxyLBResponseHeaderInfo is failed: %s", e)
	}
	if len(p.Settings.ProxyLB.BindPorts) == 0 {
		return fmt.Errorf("ProxyLB don't have any bind-ports")
	}

	// validate index
	if params.PortIndex <= 0 || params.PortIndex-1 >= len(p.Settings.ProxyLB.BindPorts) {
		return fmt.Errorf("port-index(%d) is out of range", params.PortIndex)
	}

	bindPort := p.Settings.ProxyLB.BindPorts[params.PortIndex-1]

	if len(bindPort.AddResponseHeader) == 0 {
		return fmt.Errorf("Port %s:%d don't have any additional response headers", bindPort.ProxyMode, bindPort.Port)
	}

	var list []interface{}
	for i := range bindPort.AddResponseHeader {
		list = append(list, &bindPort.AddResponseHeader[i])
	}
	return ctx.GetOutput().Print(list...)

}

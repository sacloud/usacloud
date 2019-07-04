package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ProxyLBResponseHeaderDelete(ctx command.Context, params *params.ResponseHeaderDeleteProxyLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ProxyLBResponseHeaderDelete is failed: %s", e)
	}
	if len(p.Settings.ProxyLB.BindPorts) == 0 {
		return fmt.Errorf("ProxyLB don't have any bind-ports")
	}

	// validate index
	if params.PortIndex <= 0 || params.PortIndex-1 >= len(p.Settings.ProxyLB.BindPorts) {
		return fmt.Errorf("port-index(%d) is out of range", params.PortIndex)
	}

	bindPort := p.Settings.ProxyLB.BindPorts[params.PortIndex-1]

	// validate header index
	if params.Index <= 0 || params.Index-1 >= len(bindPort.AddResponseHeader) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	newHeaders := []*sacloud.ProxyLBResponseHeader{}
	for i, header := range bindPort.AddResponseHeader {
		if i != params.Index-1 {
			newHeaders = append(newHeaders, header)
		}
	}
	bindPort.AddResponseHeader = newHeaders

	p, e = api.UpdateSetting(params.Id, p)
	if e != nil {
		return fmt.Errorf("ProxyLBResponseHeaderDelete is failed: %s", e)
	}

	var list []interface{}
	for i := range bindPort.AddResponseHeader {
		list = append(list, &bindPort.AddResponseHeader[i])
	}
	return ctx.GetOutput().Print(list...)

}

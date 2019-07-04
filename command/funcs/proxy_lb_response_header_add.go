package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ProxyLBResponseHeaderAdd(ctx command.Context, params *params.ResponseHeaderAddProxyLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ProxyLBResponseHeaderAdd is failed: %s", e)
	}
	if len(p.Settings.ProxyLB.BindPorts) == 0 {
		return fmt.Errorf("ProxyLB don't have any bind-ports")
	}

	// validate index
	if params.PortIndex <= 0 || params.PortIndex-1 >= len(p.Settings.ProxyLB.BindPorts) {
		return fmt.Errorf("port-index(%d) is out of range", params.PortIndex)
	}

	bindPort := p.Settings.ProxyLB.BindPorts[params.PortIndex-1]

	// validate duplicate
	for _, header := range bindPort.AddResponseHeader {
		if header.Header == params.Header {
			return fmt.Errorf("Header %q already exists", params.Header)
		}
	}

	bindPort.AddResponseHeader = append(bindPort.AddResponseHeader, &sacloud.ProxyLBResponseHeader{
		Header: params.Header,
		Value:  params.Value,
	})

	p, e = api.UpdateSetting(params.Id, p)
	if e != nil {
		return fmt.Errorf("ProxyLBResponseHeaderAdd is failed: %s", e)
	}

	var list []interface{}
	for i := range bindPort.AddResponseHeader {
		list = append(list, &bindPort.AddResponseHeader[i])
	}
	return ctx.GetOutput().Print(list...)
}

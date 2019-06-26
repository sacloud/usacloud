package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ProxyLBAcmeInfo(ctx command.Context, params *params.AcmeInfoProxyLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ProxyLBAcmeInfo is failed: %s", e)
	}

	return ctx.GetOutput().Print(&p.Settings.ProxyLB.LetsEncrypt)
}

package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func GSLBServerUpdate(ctx command.Context, params *params.ServerUpdateGSLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetGSLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("GSLBServerUpdate is failed: %s", e)
	}

	if len(p.Settings.GSLB.Servers) == 0 {
		return fmt.Errorf("GSLB don't have any servers")
	}

	// validate index
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.GSLB.Servers) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	// validate duplicate
	if ctx.IsSet("ipaddress") {
		for i, s := range p.Settings.GSLB.Servers {
			if i != params.Index && s.IPAddress == params.Ipaddress {
				return fmt.Errorf("GSLB already have server(%s)", params.Ipaddress)
			}
		}
	}

	server := &p.Settings.GSLB.Servers[params.Index-1]

	if ctx.IsSet("ipaddress") {
		server.IPAddress = params.Ipaddress
	}

	if ctx.IsSet("disalbed") {
		// update
		enabled := "True"
		if params.Disabled {
			enabled = "False"
		}
		server.Enabled = enabled

	}

	if ctx.IsSet("weight") {
		if params.Weight != 0 {
			server.Weight = fmt.Sprintf("%d", params.Weight)
		}
	}

	p, e = api.Update(params.Id, p)
	if e != nil {
		return fmt.Errorf("GSLBServerUpdate is failed: %s", e)
	}

	list := []interface{}{}
	for i := range p.Settings.GSLB.Servers {
		list = append(list, p.Settings.GSLB.Servers[i])
	}

	return ctx.GetOutput().Print(list...)

}

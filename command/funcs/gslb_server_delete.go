package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func GSLBServerDelete(ctx command.Context, params *params.ServerDeleteGSLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetGSLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("GSLBServerDelete is failed: %s", e)
	}

	if len(p.Settings.GSLB.Servers) == 0 {
		return fmt.Errorf("GSLB don't have any servers")
	}

	// validate index
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.GSLB.Servers) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	// delete by ipaddress
	p.Settings.GSLB.DeleteServer(p.Settings.GSLB.Servers[params.Index-1].IPAddress)

	p, e = api.Update(params.Id, p)
	if e != nil {
		return fmt.Errorf("GSLBServerDelete is failed: %s", e)
	}

	list := []interface{}{}
	for i := range p.Settings.GSLB.Servers {
		list = append(list, p.Settings.GSLB.Servers[i])
	}

	return ctx.GetOutput().Print(list...)

}

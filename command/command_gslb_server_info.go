package command

import (
	"fmt"
)

func GSLBServerInfo(ctx Context, params *ServerInfoGSLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetGSLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("GSLBServerList is failed: %s", e)
	}

	if len(p.Settings.GSLB.Servers) == 0 {
		fmt.Fprintf(GlobalOption.Err, "GSLB don't have any servers\n")
		return nil
	}

	list := []interface{}{}
	for i := range p.Settings.GSLB.Servers {
		list = append(list, p.Settings.GSLB.Servers[i])
	}

	return ctx.GetOutput().Print(list...)

}

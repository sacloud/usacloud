package funcs

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func GSLBServerAdd(ctx command.Context, params *params.ServerAddGSLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetGSLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("GSLBServerAdd is failed: %s", e)
	}

	// validate maxlen
	if len(p.Settings.GSLB.Servers) == 6 {
		return fmt.Errorf("GSLB already have max(6) servers")
	}

	// validate duplicate
	for _, s := range p.Settings.GSLB.Servers {
		if s.IPAddress == params.Ipaddress {
			return fmt.Errorf("GSLB already have server(%s)", params.Ipaddress)
		}
	}

	// add
	enabled := "True"
	if params.Disabled {
		enabled = "False"
	}

	server := &sacloud.GSLBServer{
		IPAddress: params.Ipaddress,
		Enabled:   enabled,
	}

	if params.Weight != 0 {
		server.Weight = fmt.Sprintf("%d", params.Weight)
	}
	p.AddGSLBServer(server)

	p, e = api.Update(params.Id, p)
	if e != nil {
		return fmt.Errorf("GSLBServerAdd is failed: %s", e)
	}

	list := []interface{}{}
	for i := range p.Settings.GSLB.Servers {
		list = append(list, p.Settings.GSLB.Servers[i])
	}

	return ctx.GetOutput().Print(list...)

}

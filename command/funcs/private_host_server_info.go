package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func PrivateHostServerInfo(ctx command.Context, params *params.ServerInfoPrivateHostParam) error {

	client := ctx.GetAPIClient()
	api := client.GetPrivateHostAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("PrivateHostServerInfo is failed: %s", e)
	}

	res, err := client.GetServerAPI().Find()
	if err != nil {
		return fmt.Errorf("PrivateHostServerInfo is failed: %s", err)
	}

	list := []interface{}{}
	for i, s := range res.Servers {

		if s.PrivateHost != nil && s.PrivateHost.ID == p.ID {
			list = append(list, &res.Servers[i])
		}
	}
	return ctx.GetOutput().Print(list...)
}

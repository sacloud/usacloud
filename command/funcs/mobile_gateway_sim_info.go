package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func MobileGatewaySimInfo(ctx command.Context, params *params.SimInfoMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("MobileGatewaySimInfo is failed: %s", e)
	}

	sims, err := api.ListSIM(p.ID, nil)
	if err != nil {
		return fmt.Errorf("MobileGatewaySimInfo is failed: %s", e)
	}

	list := []interface{}{}
	for i := range sims {
		list = append(list, &sims[i])
	}
	return ctx.GetOutput().Print(list...)
}

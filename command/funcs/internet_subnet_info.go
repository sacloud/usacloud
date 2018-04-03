package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func InternetSubnetInfo(ctx command.Context, params *params.SubnetInfoInternetParam) error {

	client := ctx.GetAPIClient()
	api := client.GetInternetAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("InternetSubnetInfo is failed: %s", e)
	}

	list := []interface{}{}
	for _, s := range p.Switch.Subnets {

		sn, err := getSubnetByID(ctx, s.ID)
		if err != nil {
			return fmt.Errorf("InternetSubnetInfo is failed: %s", err)
		}

		list = append(list, sn)
	}
	return ctx.GetOutput().Print(list...)

}

package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func MobileGatewayInterfaceInfo(ctx command.Context, params *params.InterfaceInfoMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("MobileGatewayInterfaceInfo is failed: %s", e)
	}

	res := []interface{}{
		map[string]interface{}{
			"Type":           "Public",
			"Switch":         "shared",
			"IPAddress":      p.Interfaces[0].IPAddress,
			"NetworkMaskLen": p.Interfaces[0].Switch.Subnet.NetworkMaskLen,
		},
	}

	if len(p.Settings.MobileGateway.Interfaces) > 1 && p.Settings.MobileGateway.Interfaces[1] != nil {
		res = append(res, map[string]interface{}{
			"Type":           "Private",
			"Switch":         p.Interfaces[1].Switch.GetStrID(),
			"IPAddress":      p.Settings.MobileGateway.Interfaces[1].IPAddress[0],
			"NetworkMaskLen": p.Settings.MobileGateway.Interfaces[1].NetworkMaskLen,
		})
	}

	return ctx.GetOutput().Print(res...)
}

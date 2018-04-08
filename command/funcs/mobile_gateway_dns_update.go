package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func MobileGatewayDnsUpdate(ctx command.Context, params *params.DnsUpdateMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("MobileGatewayDnsUpdate is failed: %s", e)
	}

	_, err := api.SetDNS(params.Id, sacloud.NewMobileGatewayResolver(params.Dns1, params.Dns2))
	if err != nil {
		return fmt.Errorf("MobileGatewayDnsUpdate is failed: %s", err)
	}
	return nil
}

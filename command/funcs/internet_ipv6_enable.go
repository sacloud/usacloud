package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func InternetIpv6Enable(ctx command.Context, params *params.Ipv6EnableInternetParam) error {

	client := ctx.GetAPIClient()
	api := client.GetInternetAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("InternetIpv6Enable is failed: %s", e)
	}

	// has switch?
	sw := p.GetSwitch()
	if sw == nil {
		return fmt.Errorf("InternetIpv6Enable is failed: %s", "Invalid state: missing Switch resource")
	}

	if len(sw.IPv6Nets) > 0 {
		fmt.Fprintln(command.GlobalOption.Err, "IPv6 is already enabled on this resource")
		return nil
	}

	ipv6net, err := api.EnableIPv6(params.Id)
	if err != nil {
		return fmt.Errorf("InternetIpv6Enable is failed: %s", err)
	}

	return ctx.GetOutput().Print(&ipv6net)

}

package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func InternetIpv6Disable(ctx command.Context, params *params.Ipv6DisableInternetParam) error {

	client := ctx.GetAPIClient()
	api := client.GetInternetAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("InternetIpv6Disable is failed: %s", e)
	}

	// has switch?
	sw := p.GetSwitch()
	if sw == nil {
		return fmt.Errorf("InternetIpv6Disable is failed: %s", "Invalid state: missing Switch resource")
	}

	if len(sw.IPv6Nets) == 0 {
		fmt.Fprintln(command.GlobalOption.Err, "IPv6 is already disabled on this resource")
		return nil
	}

	_, err := api.DisableIPv6(params.Id, sw.IPv6Nets[0].ID)
	if err != nil {
		return fmt.Errorf("InternetIpv6Disable is failed: %s", err)
	}

	return nil

}

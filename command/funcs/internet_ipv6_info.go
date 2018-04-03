package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func InternetIpv6Info(ctx command.Context, params *params.Ipv6InfoInternetParam) error {

	client := ctx.GetAPIClient()
	api := client.GetInternetAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("InternetIpv6NetInfo is failed: %s", e)
	}

	// has switch?
	sw := p.GetSwitch()
	if sw == nil {
		return fmt.Errorf("InternetIpv6NetInfo is failed: %s", "Invalid state: missing Switch resource")
	}

	if len(sw.IPv6Nets) == 0 {
		fmt.Fprintln(command.GlobalOption.Err, "IPv6 is disabled on this resource")
		return nil
	}

	res := []interface{}{}
	for i := range sw.IPv6Nets {
		res = append(res, &sw.IPv6Nets[i])
	}
	return ctx.GetOutput().Print(res...)
}

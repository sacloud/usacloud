package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func Ipv6PtrUpdate(ctx command.Context, params *params.PtrUpdateIpv6Param) error {
	client := ctx.GetAPIClient()
	api := client.GetIPv6AddrAPI()

	targetIP, err := getIPv6AddrFromArgs(ctx.Args())
	if err != nil {
		return err
	}

	ip, err := api.Read(targetIP)
	if err != nil {
		return fmt.Errorf("IPv6PtrUpdate is failed: %s", err)
	}

	ip, err = api.Update(targetIP, params.Hostname)
	if err != nil {
		return fmt.Errorf("IPv6PtrUpdate is failed: %s", err)
	}

	ipv6net, err := client.GetIPv6NetAPI().Read(ip.IPv6Net.ID)
	if err != nil {
		return fmt.Errorf("IPv6PtrUpdate is failed: %s", err)
	}
	ip.IPv6Net = ipv6net

	return ctx.GetOutput().Print(ip)
}

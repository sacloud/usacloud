package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func Ipv4PtrUpdate(ctx command.Context, params *params.PtrUpdateIpv4Param) error {
	client := ctx.GetAPIClient()
	api := client.GetIPAddressAPI()

	targetIP, err := getIPv4AddrFromArgs(ctx.Args())
	if err != nil {
		return err
	}

	ip, err := api.Read(targetIP)
	if err != nil {
		return fmt.Errorf("IPv4PtrUpdate is failed: %s", err)
	}

	ip, err = api.Update(targetIP, params.Hostname)
	if err != nil {
		return fmt.Errorf("IPv4PtrUpdate is failed: %s", err)
	}

	return ctx.GetOutput().Print(ip)
}

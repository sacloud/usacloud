package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func Ipv4PtrDelete(ctx command.Context, params *params.PtrDeleteIpv4Param) error {
	client := ctx.GetAPIClient()
	api := client.GetIPAddressAPI()

	targetIP, err := getIPv4AddrFromArgs(ctx.Args())
	if err != nil {
		return err
	}

	ip, err := api.Read(targetIP)
	if err != nil {
		return fmt.Errorf("IPv4PtrDelete is failed: %s", err)
	}

	if ip.HostName == "" {
		return fmt.Errorf("PTR record has not been set for IPAddress %q", targetIP)
	}

	ip, err = api.Update(targetIP, "")
	if err != nil {
		return fmt.Errorf("IPv4PtrDelete is failed: %s", err)
	}

	return ctx.GetOutput().Print(ip)
}

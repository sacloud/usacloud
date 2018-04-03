package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/helper/rdp"
)

func ServerRemoteDesktopInfo(ctx command.Context, params *params.RemoteDesktopInfoServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerRdpInfo is failed: %s", e)
	}

	if !p.IsUp() {
		return fmt.Errorf("ServerRdpInfo is failed: %s", "server is not running")
	}

	// collect server IPAddress
	ip := p.Interfaces[0].IPAddress
	if ip == "" {
		ip = p.Interfaces[0].UserIPAddress
	}
	if ip == "" {
		return fmt.Errorf("ServerRdpInfo is failed: collecting IPAddress from server is failed: %#v", p)
	}

	rdpClient := &rdp.Opener{
		User:      params.User,
		Port:      params.Port,
		IPAddress: ip,
	}

	fmt.Fprint(command.GlobalOption.Out, rdpClient.RDPFileContent())
	return nil

}

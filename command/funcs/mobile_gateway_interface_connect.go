package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func MobileGatewayInterfaceConnect(ctx command.Context, params *params.InterfaceConnectMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("MobileGatewayInterfaceConnect is failed: %s", e)
	}

	if len(p.Interfaces) > 1 && p.Interfaces[1].GetSwitch() != nil {
		return fmt.Errorf("Interface[%s] is already connected to switch", "eth1")
	}
	if p.IsUp() {
		return fmt.Errorf("MobileGateway(%d) is still running", params.Id)
	}

	var err error

	_, err = client.Switch.Read(params.SwitchId)
	if err != nil {
		return fmt.Errorf("Switch[%d] is not found", params.SwitchId)
	}

	//connect
	_, err = api.ConnectToSwitch(p.ID, params.SwitchId)
	if err != nil {
		return fmt.Errorf("MobileGatewayInterfaceConnect is failed: %s", err)
	}

	p.SetPrivateInterface(params.Ipaddress, params.NwMasklen)
	_, err = api.Update(p.ID, p)
	if err != nil {
		return fmt.Errorf("MobileGatewayInterfaceConnect is failed: %s", err)
	}

	_, err = api.Config(p.ID)
	if err != nil {
		return fmt.Errorf("MobileGatewayInterfaceConnect is failed: %s", err)
	}

	return nil
}

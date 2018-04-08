package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func MobileGatewayInterfaceDisconnect(ctx command.Context, params *params.InterfaceDisconnectMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("MobileGatewayInterfaceDisconnect is failed: %s", e)
	}

	if len(p.Interfaces) < 2 || p.Interfaces[1].GetSwitch() == nil {
		return fmt.Errorf("Interface[%s] is already disconnected from switch", "eth1")
	}
	if p.IsUp() {
		return fmt.Errorf("MobileGateway(%d) is still running", params.Id)
	}

	var err error
	// disconnect
	_, err = api.DisconnectFromSwitch(p.ID)
	if err != nil {
		return fmt.Errorf("MobileGatewayInterfaceDisconnect is failed: %s", err)
	}

	p.ClearPrivateInterface()
	_, err = api.Update(p.ID, p)
	if err != nil {
		return fmt.Errorf("MobileGatewayInterfaceDisconnect is failed: %s", err)
	}

	_, err = api.Config(p.ID)
	if err != nil {
		return fmt.Errorf("MobileGatewayInterfaceDisconnect is failed: %s", err)
	}

	return nil
}

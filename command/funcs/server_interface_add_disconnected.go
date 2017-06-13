package funcs

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ServerInterfaceAddDisconnected(ctx command.Context, params *params.InterfaceAddDisconnectedServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerInterfaceAddDisconnected is failed: %s", e)
	}

	if len(p.GetInterfaces()) >= sacloud.ServerMaxInterfaceLen {
		return fmt.Errorf("Server already connected maximum count of interfaces")
	}

	if p.IsUp() {
		return fmt.Errorf("ServerInterfaceAddDisconnected is failed: %s", "server is running")
	}

	// call manipurate functions
	interfaceAPI := client.GetInterfaceAPI()
	_, err := interfaceAPI.CreateAndConnectToServer(params.Id)
	if err != nil {
		return fmt.Errorf("ServerInterfaceAddDisconnected is failed: %s", err)
	}

	return nil

}

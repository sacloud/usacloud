package command

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
)

func ServerInterfaceAddDisconnected(ctx Context, params *InterfaceAddDisconnectedServerParam) error {

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

	// read again
	p, e = api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerInterfaceAddDisconnected is failed: %s", e)
	}

	interfaces := p.GetInterfaces()

	// build parameters to display table
	list := []interface{}{}
	for i := range interfaces {
		list = append(list, &interfaces[i])
	}

	return ctx.GetOutput().Print(list...)

}

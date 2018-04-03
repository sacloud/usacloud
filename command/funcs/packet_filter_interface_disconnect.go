package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func PacketFilterInterfaceDisconnect(ctx command.Context, params *params.InterfaceDisconnectPacketFilterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetPacketFilterAPI()
	interfaceAPI := client.GetInterfaceAPI()

	// read packet filter
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("PacketFilterInterfaceDisconnect is failed: %s", e)
	}

	// read interface
	nic, e := interfaceAPI.Read(params.InterfaceId)
	if e != nil {
		return fmt.Errorf("PacketFilterInterfaceDisconnect is failed: %s", e)
	}

	if nic.PacketFilter == nil || nic.PacketFilter.ID != params.Id {
		return fmt.Errorf("interface(%d) is not connected packet filter", nic.ID)
	}

	// call manipurate functions
	_, err := interfaceAPI.DisconnectFromPacketFilter(params.InterfaceId)
	if err != nil {
		return fmt.Errorf("PacketFilterInterfaceDisconnect is failed: %s", err)
	}
	return nil

}

package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func InterfacePacketFilterDisconnect(ctx command.Context, params *params.PacketFilterDisconnectInterfaceParam) error {

	client := ctx.GetAPIClient()
	api := client.GetInterfaceAPI()
	packetFilterAPI := client.GetPacketFilterAPI()

	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("InterfacePacketFilterDisconnect is failed: %s", e)
	}

	// read packet filter
	_, e = packetFilterAPI.Read(params.PacketFilterId)
	if e != nil {
		return fmt.Errorf("InterfacePacketFilterDisconnect is failed: %s", e)
	}

	if p.PacketFilter == nil || p.PacketFilter.ID != params.PacketFilterId {
		return fmt.Errorf("interface is not connected packet filter(%d)", params.PacketFilterId)
	}

	// call manipurate functions
	_, err := api.DisconnectFromPacketFilter(params.Id)
	if err != nil {
		return fmt.Errorf("InterfacePacketFilterDisconnect is failed: %s", err)
	}
	return nil

}

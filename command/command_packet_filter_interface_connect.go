package command

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
)

func PacketFilterInterfaceConnect(ctx Context, params *InterfaceConnectPacketFilterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetPacketFilterAPI()
	interfaceAPI := client.GetInterfaceAPI()

	// read packet filter
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("PacketFilterInterfaceConnect is failed: %s", e)
	}

	// read interface
	nic, e := interfaceAPI.Read(params.InterfaceId)
	if e != nil {
		return fmt.Errorf("PacketFilterInterfaceConnect is failed: %s", e)
	}

	if nic.PacketFilter != nil && nic.PacketFilter.ID > sacloud.EmptyID {
		return fmt.Errorf("interface(%d) is already connected packet filter", nic.ID)
	}

	// call manipurate functions
	_, err := interfaceAPI.ConnectToPacketFilter(params.InterfaceId, params.Id)
	if err != nil {
		return fmt.Errorf("PacketFilterInterfaceConnect is failed: %s", err)
	}
	return nil

}

package command

import (
	"fmt"
)

func SwitchBridgeConnect(ctx Context, params *BridgeConnectSwitchParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSwitchAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("SwitchBridgeConnect is failed: %s", e)
	}

	if p.Bridge != nil {
		return fmt.Errorf("SwitchBridgeConnect is failed: Bridge is already connected")
	}

	// call manipurate functions
	_, err := api.ConnectToBridge(params.Id, params.BridgeId)
	if err != nil {
		return fmt.Errorf("SwitchBridgeConnect is failed: %s", err)
	}

	// read again
	p, e = api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("SwitchBridgeConnect is failed: %s", e)
	}
	return ctx.GetOutput().Print(p)

}

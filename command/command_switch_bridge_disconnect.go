package command

import (
	"fmt"
)

func SwitchBridgeDisconnect(ctx Context, params *BridgeDisconnectSwitchParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSwitchAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("SwitchBridgeDisconnect is failed: %s", e)
	}

	if p.Bridge == nil {
		return fmt.Errorf("SwitchBridgeDisconnect is failed: Bridge is already disconnected")
	}

	// call manipurate functions
	_, err := api.DisconnectFromBridge(params.Id)
	if err != nil {
		return fmt.Errorf("SwitchBridgeDisconnect is failed: %s", err)
	}

	return nil

}

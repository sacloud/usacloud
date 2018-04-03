package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func SwitchBridgeConnect(ctx command.Context, params *params.BridgeConnectSwitchParam) error {

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

	return nil

}

package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func DiskServerDisconnect(ctx command.Context, params *params.ServerDisconnectDiskParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDiskAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DiskServerDisconnect is failed: %s", e)
	}

	// disk is connected to server?
	if p.Server == nil {
		return fmt.Errorf("DiskServerDisconnect is failed: %s", "Disk needs to be connected to server")
	}

	// server is stopped?
	if !p.Server.IsDown() {
		return fmt.Errorf("DiskServerDisconnect is failed: %s", "Server needs to be stopped")
	}

	// call manipurate functions
	_, err := api.DisconnectFromServer(params.Id)
	if err != nil {
		return fmt.Errorf("DiskServerDisconnect is failed: %s", err)
	}

	return nil
}

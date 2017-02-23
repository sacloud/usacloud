package command

import (
	"fmt"
)

func DiskServerDisconnect(ctx Context, params *ServerDisconnectDiskParam) error {

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

	// read again
	p, e = api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DiskServerDisconnect is failed: %s", e)
	}

	return ctx.GetOutput().Print(p)
}

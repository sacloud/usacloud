package command

import (
	"fmt"
)

func DiskServerConnect(ctx Context, params *ServerConnectDiskParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDiskAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DiskServerConnect is failed: %s", e)
	}

	// disk is disconnected from server?
	if p.Server != nil {
		return fmt.Errorf("DiskServerConnect is failed: %s", "Disk is already connected to server")
	}

	// server is exists?
	server, e := client.GetServerAPI().Read(params.ServerId)
	if e != nil || server == nil {
		return fmt.Errorf("DiskServerConnect is failed: Server(ID:%d) isnot exists: %s", params.Id, e)
	}

	// server is stopped?
	if !server.IsDown() {
		return fmt.Errorf("DiskServerConnect is failed: %s", "Server needs to be stopped")
	}

	// call manipurate functions
	_, err := api.ConnectToServer(params.Id, params.ServerId)
	if err != nil {
		return fmt.Errorf("DiskServerConnect is failed: %s", err)
	}

	// read again
	p, e = api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DiskServerConnect is failed: %s", e)
	}

	return ctx.GetOutput().Print(p)

}

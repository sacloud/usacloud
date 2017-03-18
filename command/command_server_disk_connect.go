package command

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
)

func ServerDiskConnect(ctx Context, params *DiskConnectServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerDiskConnect is failed: %s", e)
	}

	if len(p.Disks) >= sacloud.ServerMaxDiskLen {
		return fmt.Errorf("Server already connected maximum count of disks")
	}

	if p.IsUp() {
		return fmt.Errorf("ServerDiskConnect is failed: %s", "server is running")
	}

	// call manipurate functions
	_, err := client.GetDiskAPI().ConnectToServer(params.DiskId, params.Id)
	if err != nil {
		return fmt.Errorf("ServerDiskConnect is failed: %s", err)
	}

	return nil

}

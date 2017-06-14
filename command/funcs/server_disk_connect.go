package funcs

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ServerDiskConnect(ctx command.Context, params *params.DiskConnectServerParam) error {

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

package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ServerDiskDisconnect(ctx command.Context, params *params.DiskDisconnectServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerDiskDisconnect is failed: %s", e)
	}
	disks := p.GetDisks()
	if len(disks) == 0 {
		return fmt.Errorf("Server don't have any disks")
	}
	if p.IsUp() {
		return fmt.Errorf("ServerDiskDisconnect is failed: %s", "server is running")
	}

	exists := false
	for _, disk := range disks {
		if disk.ID == params.DiskId {
			exists = true
			break
		}
	}
	if !exists {
		return fmt.Errorf("Server don't have disk(%d)", params.DiskId)
	}

	// call manipurate functions
	_, err := client.GetDiskAPI().DisconnectFromServer(params.DiskId)
	if err != nil {
		return fmt.Errorf("ServerDiskDisconnect is failed: %s", err)
	}

	return nil

}

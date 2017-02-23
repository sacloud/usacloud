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

	// read again
	p, e = api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerDiskConnect is failed: %s", e)
	}
	disks := p.GetDisks()

	// collect disk info by DiskAPI
	diskAPI := client.GetDiskAPI()
	for _, disk := range disks {
		diskAPI.FilterMultiBy("ID", disk.ID)
	}
	res, err := diskAPI.Find()
	if err != nil {
		if e != nil {
			return fmt.Errorf("ServerDiskConnect is failed: %s", e)
		}
	}

	// build parameters to display table
	list := []interface{}{}
	for i := range res.Disks {
		list = append(list, &res.Disks[i])
	}

	return ctx.GetOutput().Print(list...)

}

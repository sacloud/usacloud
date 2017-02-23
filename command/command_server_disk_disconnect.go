package command

import (
	"fmt"
)

func ServerDiskDisconnect(ctx Context, params *DiskDisconnectServerParam) error {

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

	// read again
	p, e = api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerDiskDisconnect is failed: %s", e)
	}

	disks = p.GetDisks()
	if len(disks) == 0 {
		return nil
	}

	// collect disk info by DiskAPI
	diskAPI := client.GetDiskAPI()
	for _, disk := range disks {
		diskAPI.FilterMultiBy("ID", disk.ID)
	}
	res, err := diskAPI.Find()
	if err != nil {
		if e != nil {
			return fmt.Errorf("ServerDiskDisconnect is failed: %s", e)
		}
	}

	// build parameters to display table
	list := []interface{}{}
	for i := range res.Disks {
		list = append(list, &res.Disks[i])
	}

	return ctx.GetOutput().Print(list...)

}

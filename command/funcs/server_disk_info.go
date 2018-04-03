package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ServerDiskInfo(ctx command.Context, params *params.DiskInfoServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerDiskInfo is failed: %s", e)
	}

	disks := p.GetDisks()
	if len(disks) == 0 {
		fmt.Fprintf(command.GlobalOption.Err, "Server don't have any disks\n")
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
			return fmt.Errorf("ServerDiskInfo is failed: %s", e)
		}
	}

	// build parameters to display table
	list := []interface{}{}
	for i := range res.Disks {
		list = append(list, &res.Disks[i])
	}

	return ctx.GetOutput().Print(list...)

}

package command

import (
	"fmt"
)

func DiskReinstallFromArchive(ctx Context, params *ReinstallFromArchiveDiskParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDiskAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DiskReinstallFromArchive is failed: %s", e)
	}

	// call manipurate functions
	_, err := api.ReinstallFromArchive(params.Id, params.SourceArchive, params.DistantFrom...)
	if err != nil {
		return fmt.Errorf("DiskReinstallFromArchive is failed: %s", err)
	}

	if !params.Async {
		err = api.SleepWhileCopying(params.Id, client.DefaultTimeoutDuration)
		if err != nil {
			return fmt.Errorf("DiskReinstallFromArchive is failed: %s", err)
		}
		p, err = api.Read(params.Id)
		if err != nil {
			return fmt.Errorf("DiskReinstallFromArchive is failed: %s", err)
		}
	}

	return ctx.GetOutput().Print(p)

}

package command

import (
	"fmt"
)

func DiskReinstallFromArchive(ctx Context, params *ReinstallFromArchiveDiskParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDiskAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DiskReinstallFromArchive is failed: %s", e)
	}

	// call manipurate functions
	_, err := api.ReinstallFromArchive(params.Id, params.SourceArchiveId, params.DistantFrom...)
	if err != nil {
		return fmt.Errorf("DiskReinstallFromArchive is failed: %s", err)
	}

	if !params.Async {
		err = api.SleepWhileCopying(params.Id, client.DefaultTimeoutDuration)
		if err != nil {
			return fmt.Errorf("DiskReinstallFromArchive is failed: %s", err)
		}
	}

	return nil

}

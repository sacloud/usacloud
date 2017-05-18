package command

import (
	"fmt"
	"github.com/sacloud/usacloud/command/internal"
)

func DiskReinstallFromArchive(ctx Context, params *ReinstallFromArchiveDiskParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDiskAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DiskReinstallFromArchive is failed: %s", e)
	}

	err := internal.ExecWithProgress(
		fmt.Sprintf("Still installing[ID:%d]...", params.Id),
		fmt.Sprintf("Reinstall disk[ID:%d]", params.Id),
		GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {
			_, err := api.ReinstallFromArchive(params.Id, params.SourceArchiveId, params.DistantFrom...)
			if err != nil {
				errChan <- err
				return
			}
			err = api.SleepWhileCopying(params.Id, client.DefaultTimeoutDuration)
			if err != nil {
				errChan <- err
				return
			}
			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("DiskReinstallFromArchive is failed: %s", err)
	}

	return nil
}

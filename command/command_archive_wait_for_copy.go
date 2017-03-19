package command

import (
	"fmt"
	"github.com/sacloud/usacloud/command/internal"
)

func ArchiveWaitForCopy(ctx Context, params *WaitForCopyArchiveParam) error {

	client := ctx.GetAPIClient()
	api := client.GetArchiveAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ArchiveWaitForCopy is failed: %s", e)
	}

	// wait for copy with progress
	spinner := internal.NewSpinner(
		"Coping...",
		"Copy archive is complete.\n",
		internal.CharSetProgress,
		GlobalOption.Progress)
	spinner.Start()
	compChan, progChan, errChan := api.AsyncSleepWhileCopying(p.ID, client.DefaultTimeoutDuration)
copy:
	for {
		select {
		case <-compChan:
			spinner.Stop()
			break copy
		case <-progChan:
		// noop
		case err := <-errChan:
			return fmt.Errorf("ArchiveWaitForCopy is failed: %s", err)
		}
	}

	return nil

}

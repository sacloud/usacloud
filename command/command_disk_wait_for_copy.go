package command

import (
	"fmt"
	"github.com/sacloud/usacloud/command/internal"
)

func DiskWaitForCopy(ctx Context, params *WaitForCopyDiskParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDiskAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DiskWaitForCopy is failed: %s", e)
	}

	// wait for copy with progress
	spinner := internal.NewSpinner(
		"Coping...",
		"Copy disk is complete.\n",
		internal.CharSetProgress,
		GlobalOption.Progress)
	spinner.Start()
	compChan, progChan, errChan := api.AsyncSleepWhileCopying(p.ID, client.DefaultTimeoutDuration)
copy:
	for {
		select {
		case r := <-compChan:
			p = r
			spinner.Stop()
			break copy
		case <-progChan:
		// noop
		case err := <-errChan:
			return fmt.Errorf("DiskWaitForCopy is failed: %s", err)
		}
	}

	return ctx.GetOutput().Print(p)

}

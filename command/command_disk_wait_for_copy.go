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
	err := internal.ExecWithProgress(
		fmt.Sprintf("Still coping[ID:%d]...", params.Id),
		fmt.Sprintf("Copy disk[ID:%d]", params.Id),
		GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {
			err := api.SleepWhileCopying(p.ID, client.DefaultTimeoutDuration)
			if err != nil {
				errChan <- err
			}
			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("DiskWaitForCopy is failed: %s", err)
	}

	return ctx.GetOutput().Print(p)
}

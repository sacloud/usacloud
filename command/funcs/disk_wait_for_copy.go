package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
)

func DiskWaitForCopy(ctx command.Context, params *params.WaitForCopyDiskParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDiskAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DiskWaitForCopy is failed: %s", e)
	}

	// wait for copy with progress
	err := internal.ExecWithProgress(
		fmt.Sprintf("Still copying[ID:%d]...", params.Id),
		fmt.Sprintf("Copy disk[ID:%d]", params.Id),
		command.GlobalOption.Progress,
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

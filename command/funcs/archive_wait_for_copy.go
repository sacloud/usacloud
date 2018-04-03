package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
)

func ArchiveWaitForCopy(ctx command.Context, params *params.WaitForCopyArchiveParam) error {

	client := ctx.GetAPIClient()
	api := client.GetArchiveAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ArchiveWaitForCopy is failed: %s", e)
	}

	// wait for copy with progress
	err := internal.ExecWithProgress(
		fmt.Sprintf("Still copying[ID:%d]...", params.Id),
		fmt.Sprintf("Copy archive[ID:%d]", params.Id),
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
		return fmt.Errorf("ArchiveWaitForCopy is failed: %s", err)
	}

	return nil
}

package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
)

func DiskReinstallToBlank(ctx command.Context, params *params.ReinstallToBlankDiskParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDiskAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DiskReinstallToBlank is failed: %s", e)
	}

	err := internal.ExecWithProgress(
		fmt.Sprintf("Still installing[ID:%d]...", params.Id),
		fmt.Sprintf("Reinstall disk[ID:%d]", params.Id),
		command.GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {
			_, err := api.ReinstallFromBlank(params.Id, p.SizeMB)
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
		return fmt.Errorf("DiskReinstallToBlank is failed: %s", err)
	}

	return nil
}

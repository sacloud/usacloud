package command

import (
	"fmt"
	"github.com/sacloud/usacloud/command/internal"
)

func DiskReinstallToBlank(ctx Context, params *ReinstallToBlankDiskParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDiskAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DiskReinstallToBlank is failed: %s", e)
	}

	compChan := make(chan bool)
	errChan := make(chan error)
	spinner := internal.NewProgress(
		fmt.Sprintf("Still installing[ID:%d]...", params.Id),
		fmt.Sprintf("Reinstall disk[ID:%d]", params.Id),
		GlobalOption.Progress)
	spinner.Start()

	go func() {
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
	}()

copy:
	for {
		select {
		case <-compChan:
			spinner.Stop()
			break copy
		case err := <-errChan:
			return fmt.Errorf("DiskReinstallToBlank is failed: %s", err)
		}
	}

	return nil

}

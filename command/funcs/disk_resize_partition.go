package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
)

func DiskResizePartition(ctx command.Context, params *params.ResizePartitionDiskParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDiskAPI()

	// wait for copy with progress
	err := internal.ExecWithProgress(
		fmt.Sprintf("Still resizing[ID:%d]...", params.Id),
		fmt.Sprintf("Resize-Partition disk[ID:%d]", params.Id),
		command.GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {
			// call manipurate functions
			_, err := api.ResizePartitionBackground(params.Id)
			if err != nil {
				errChan <- err
				return
			}
			if err := api.SleepWhileCopying(params.Id, client.DefaultTimeoutDuration); err != nil {
				errChan <- err
				return
			}

			// check result
			disk, err := api.Read(params.Id)
			if err != nil {
				errChan <- err
				return
			}

			if disk.JobStatus != nil && disk.JobStatus.Status == "failed" {
				msg := ""
				if disk.JobStatus.ConfigError != nil {
					ce := disk.JobStatus.ConfigError
					msg = fmt.Sprintf("%s: %s", ce.ErrorCode, ce.ErrorMsg)
				}
				errChan <- fmt.Errorf("DiskResize job is failed: %s", msg)
				return
			}

			compChan <- true
		},
	)

	if err != nil {
		return fmt.Errorf("DiskResizePartition is failed: %s", err)
	}

	// read
	res, err := api.Read(params.Id)
	if err != nil {
		return fmt.Errorf("DiskResizePartition is failed: %s", err)
	}
	return ctx.GetOutput().Print(res)
}

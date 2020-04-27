// Copyright 2017-2020 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
	"github.com/sacloud/usacloud/pkg/progress"
)

func DiskResizePartition(ctx cli.Context, params *params.ResizePartitionDiskParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDiskAPI()

	// wait for copy with progress
	err := progress.ExecWithProgress(
		fmt.Sprintf("Still resizing[ID:%d]...", params.Id),
		fmt.Sprintf("Resize-Partition disk[ID:%d]", params.Id),
		ctx.IO().Progress(),
		ctx.Option().NoColor,
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

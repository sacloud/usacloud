// Copyright 2017-2019 The Usacloud Authors
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
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
)

func ServerVncSnapshot(ctx command.Context, params *params.VncSnapshotServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerVncSnapshot is failed: %s", e)
	}
	if !p.IsUp() {
		return fmt.Errorf("ServerVncSnapshot is failed: %s", "server is not running")
	}

	if !p.IsUp() && params.WaitForBoot {

		err := internal.ExecWithProgress(
			fmt.Sprintf("Still booting[ID:%d]...", params.Id),
			fmt.Sprintf("Connect to server[ID:%d]", params.Id),
			command.GlobalOption.Progress,
			func(compChan chan bool, errChan chan error) {
				// call manipurate functions
				err := api.SleepUntilUp(params.Id, client.DefaultTimeoutDuration)
				if err != nil {
					errChan <- err
					return
				}
				compChan <- true
			},
		)
		if err != nil {
			return fmt.Errorf("ServerVncSnapshot is failed: %s", e)
		}
	}

	snapshotReq := api.NewVNCSnapshotRequest()
	vncSnapshotResponse, err := api.GetVNCSnapshot(params.Id, snapshotReq)
	if err != nil {
		return fmt.Errorf("ServerVncSnapshot is failed: %s", err)
	}
	vncImage, err := base64.StdEncoding.DecodeString(vncSnapshotResponse.Image)
	if err != nil {
		return fmt.Errorf("ServerVncSnapshot is failed: %s", err)
	}

	var filename string
	if len(params.OutputPath) > 0 {
		filename = params.OutputPath
	} else {
		filename = fmt.Sprintf("%d_%s.gif", params.Id, time.Now().Format("20060102-150405"))
	}
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("ServerVncSnapshot is failed: %s", err)
	}
	defer file.Close()

	_, err = file.Write(vncImage)
	if err != nil {
		return fmt.Errorf("ServerVncSnapshot is failed: %s", err)
	}

	out := command.GlobalOption.Err
	fmt.Fprintln(out, "Snapshot created:", filename)

	return nil
}

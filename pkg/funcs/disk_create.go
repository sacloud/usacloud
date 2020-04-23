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

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
	"github.com/sacloud/usacloud/pkg/util/progress"
)

func DiskCreate(ctx cli.Context, params *params.CreateDiskParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDiskAPI()
	p := api.New()

	// set params
	var distantFrom []sacloud.ID
	for _, id := range params.DistantFrom {
		distantFrom = append(distantFrom, sacloud.ID(id))
	}

	p.SetDescription(params.Description)
	p.SetIconByID(params.IconId)
	p.SetDiskPlan(params.Plan)
	p.SetSizeGB(params.Size)
	p.SetDistantFrom(distantFrom)
	p.SetName(params.Name)
	p.SetTags(params.Tags)
	p.SetDiskConnection(sacloud.EDiskConnection(params.Connection))
	p.SetSourceArchive(params.SourceArchiveId)
	p.SetSourceDisk(params.SourceDiskId)

	// wait for copy with progress
	var res *sacloud.Disk
	var err error
	err = progress.ExecWithProgress(
		"Still creating...",
		"Create disk",
		ctx.IO().Progress(),
		ctx.Option().NoColor,
		func(compChan chan bool, errChan chan error) {
			// call Create(id)
			res, err = api.Create(p)
			if err != nil {
				errChan <- err
				return
			}
			err = api.SleepWhileCopying(res.ID, client.DefaultTimeoutDuration)
			if err != nil {
				errChan <- err
				return
			}
			res, err = api.Read(res.ID)
			if err != nil {
				errChan <- err
				return
			}
			compChan <- true
		},
	)

	if err != nil {
		return fmt.Errorf("DiskCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}

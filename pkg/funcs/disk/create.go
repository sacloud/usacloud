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

package disk

import (
	"github.com/sacloud/libsacloud/v2/pkg/size"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func Create(ctx cli.Context, params *params.CreateDiskParam) error {
	client := sacloud.NewDiskOp(ctx.Client())

	// wait for copy with progress
	var res *sacloud.Disk
	err := ctx.ExecWithProgress(func() error {
		disk, err := client.Create(ctx, ctx.Zone(), &sacloud.DiskCreateRequest{
			DiskPlanID:      types.DiskPlanIDMap[params.Plan],
			Connection:      types.DiskConnectionMap[params.Connection],
			SourceDiskID:    params.SourceDiskId,
			SourceArchiveID: params.SourceArchiveId,
			SizeMB:          params.Size * size.GiB,
			Name:            params.Name,
			Description:     params.Description,
			Tags:            params.Tags,
			IconID:          params.IconId,
		}, params.DistantFrom)
		if err != nil {
			return err
		}

		raw, err := sacloud.WaiterForReady(func() (interface{}, error) {
			return client.Read(ctx, ctx.Zone(), disk.ID)
		}).WaitForState(ctx)
		if err != nil {
			return err
		}

		res = raw.(*sacloud.Disk)
		return err
	},
	)

	if err != nil {
		return err
	}
	return ctx.Output().Print(res)
}

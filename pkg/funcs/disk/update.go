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
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func Update(ctx cli.Context, params *params.UpdateDiskParam) error {
	client := sacloud.NewDiskOp(ctx.Client())
	data, err := client.Read(ctx, ctx.Zone(), params.Id)
	if err != nil {
		return err
	}

	p := &sacloud.DiskUpdateRequest{
		Connection:  data.Connection,
		Name:        data.Name,
		Description: data.Description,
		Tags:        data.Tags,
		IconID:      data.IconID,
	}
	if params.Changed("Connection") {
		p.Connection = types.EDiskConnection(params.Connection)
	}
	if params.Changed("Name") {
		p.Name = params.Name
	}
	if params.Changed("Description") {
		p.Description = params.Description
	}
	if params.Changed("Tags") {
		p.Tags = params.Tags
	}
	if params.Changed("IconId") {
		p.IconID = params.IconId
	}
	// call Update(id)
	res, err := client.Update(ctx, ctx.Zone(), params.Id, p)
	if err != nil {
		return err
	}
	return ctx.Output().Print(res)
}

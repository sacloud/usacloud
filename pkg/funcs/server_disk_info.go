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
)

func ServerDiskInfo(ctx cli.Context, params *params.DiskInfoServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerDiskInfo is failed: %s", e)
	}

	disks := p.GetDisks()
	if len(disks) == 0 {
		fmt.Fprintf(ctx.IO().Err(), "Server don't have any disks\n")
		return nil
	}

	// collect disk info by DiskAPI
	diskAPI := client.GetDiskAPI()
	for _, disk := range disks {
		diskAPI.FilterMultiBy("ID", disk.ID)
	}
	res, err := diskAPI.Find()
	if err != nil {
		if e != nil {
			return fmt.Errorf("ServerDiskInfo is failed: %s", e)
		}
	}

	// build parameters to display table
	list := []interface{}{}
	for i := range res.Disks {
		list = append(list, &res.Disks[i])
	}

	return ctx.GetOutput().Print(list...)

}

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

package database

import (
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func Reset(ctx cli.Context, params *params.ResetDatabaseParam) error {
	client := sacloud.NewDatabaseOp(ctx.Client())
	// set params
	err := ctx.ExecWithProgress(func() error {
		if err := client.Reset(ctx, ctx.Zone(), params.Id); err != nil {
			return err
		}
		_, err := sacloud.WaiterForApplianceUp(func() (interface{}, error) {
			return client.Read(ctx, ctx.Zone(), params.Id)
		}, 30).WaitForState(ctx)
		return err
	})
	if err != nil {
		return fmt.Errorf("DatabaseReset is failed: %s", err)
	}

	return nil
}

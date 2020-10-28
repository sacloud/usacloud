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

func ShutdownForce(ctx cli.Context, params *params.ShutdownForceDatabaseParam) error {
	client := sacloud.NewDatabaseOp(ctx.Client())
	db, err := client.Read(ctx, ctx.Zone(), params.Id)
	if err != nil {
		return fmt.Errorf("DatabaseShutdown is failed: %s", err)
	}

	if db.InstanceStatus.IsDown() {
		return nil // already downed.
	}

	err = ctx.ExecWithProgress(func() error {
		if err := client.Shutdown(ctx, ctx.Zone(), params.Id, &sacloud.ShutdownOption{Force: true}); err != nil {
			return err
		}
		_, err := sacloud.WaiterForDown(func() (interface{}, error) {
			return client.Read(ctx, ctx.Zone(), params.Id)
		}).WaitForState(ctx)
		return err
	})
	if err != nil {
		return fmt.Errorf("DatabaseShutdown is failed: %s", err)
	}

	return nil
}

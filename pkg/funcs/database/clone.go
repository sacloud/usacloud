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
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func Clone(ctx cli.Context, params *params.CloneDatabaseParam) error {
	client := sacloud.NewDatabaseOp(ctx.Client())
	db, err := client.Read(ctx, ctx.Zone(), params.Id)
	if err != nil {
		return fmt.Errorf("DatabaseClone is failed: %s", err)
	}

	req := &sacloud.DatabaseCreateRequest{
		PlanID:             db.PlanID,
		SwitchID:           db.SwitchID,
		IPAddresses:        db.IPAddresses,
		NetworkMaskLen:     db.NetworkMaskLen,
		DefaultRoute:       db.DefaultRoute,
		Conf:               db.Conf,
		SourceID:           db.ID,
		CommonSetting:      db.CommonSetting,
		BackupSetting:      db.BackupSetting,
		ReplicationSetting: db.ReplicationSetting,
		Name:               db.Name,
		Description:        db.Description,
		Tags:               db.Tags,
		IconID:             db.IconID,
	}

	if !params.SwitchId.IsEmpty() {
		req.SwitchID = params.SwitchId
	}
	if params.NwMaskLen != 0 {
		req.NetworkMaskLen = params.NwMaskLen
	}
	if params.DefaultRoute != "" {
		req.Description = params.Description
	}
	if params.Port != 0 {
		req.CommonSetting.ServicePort = params.Port
	}
	if !params.Changed("plan") {
		req.PlanID = types.ID(params.Plan)
	}

	res, err := client.Create(ctx, ctx.Zone(), req)
	if err != nil {
		return fmt.Errorf("DatabaseClone is failed: %s", err)
	}

	// wait for boot
	err = ctx.ExecWithProgress(func() error {
		_, err := sacloud.WaiterForApplianceUp(func() (interface{}, error) {
			return client.Read(ctx, ctx.Zone(), res.ID)
		}, 10).WaitForState(ctx)
		return err
	})
	if err != nil {
		return fmt.Errorf("DatabaseClone is failed: %s", err)
	}
	return ctx.Output().Print(res)
}

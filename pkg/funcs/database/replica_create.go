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

	"github.com/sacloud/libsacloud/v2/sacloud/types"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func ReplicaCreate(ctx cli.Context, params *params.ReplicaCreateDatabaseParam) error {
	client := sacloud.NewDatabaseOp(ctx.Client())
	db, err := client.Read(ctx, ctx.Zone(), params.Id)
	if err != nil {
		return fmt.Errorf("DatabaseReplicaCreate is failed: %s", err)
	}

	// Validate master instance
	if db.ReplicationSetting == nil || db.ReplicationSetting.Model != types.DatabaseReplicationModels.MasterSlave {
		return fmt.Errorf("database[%q] is not setted as replication master", db.Name)
	}

	if params.SwitchId.IsEmpty() {
		params.SwitchId = db.SwitchID
	}
	if params.NwMaskLen == 0 {
		params.NwMaskLen = db.NetworkMaskLen
	}
	if params.DefaultRoute == "" {
		params.DefaultRoute = db.DefaultRoute
	}

	res, err := client.Create(ctx, ctx.Zone(), &sacloud.DatabaseCreateRequest{
		PlanID:         types.ID(db.PlanID.Int64() + 1),
		SwitchID:       params.SwitchId,
		IPAddresses:    []string{params.Ipaddress1},
		NetworkMaskLen: params.NwMaskLen,
		DefaultRoute:   params.DefaultRoute,
		Conf:           db.Conf,
		CommonSetting: &sacloud.DatabaseSettingCommon{
			ServicePort: db.CommonSetting.ServicePort,
			// SourceNetwork: params.SourceNetwork,
		},
		BackupSetting: db.BackupSetting,
		ReplicationSetting: &sacloud.DatabaseReplicationSetting{
			Model:       types.DatabaseReplicationModels.AsyncReplica,
			IPAddress:   db.IPAddresses[0],
			Port:        db.CommonSetting.ServicePort,
			User:        db.CommonSetting.ReplicaUser,
			Password:    db.CommonSetting.ReplicaPassword,
			ApplianceID: db.ID,
		},
		Name:        params.Name,
		Description: params.Description,
		Tags:        params.Tags,
		IconID:      params.IconId,
	})
	if err != nil {
		return fmt.Errorf("DatabaseReplicaCreate is failed: %s", err)
	}

	// wait for boot
	err = ctx.ExecWithProgress(func() error {
		_, err := sacloud.WaiterForApplianceUp(func() (interface{}, error) {
			return client.Read(ctx, ctx.Zone(), res.ID)
		}, 10).WaitForState(ctx)
		return err
	})
	if err != nil {
		return fmt.Errorf("DatabaseCreate is failed: %s", err)
	}

	return ctx.Output().Print(res)
}

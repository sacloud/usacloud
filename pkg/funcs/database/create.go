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

	databaseBuilder "github.com/sacloud/libsacloud/v2/helper/builder/database"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
	"github.com/sacloud/usacloud/pkg/weekday"
)

func Create(ctx cli.Context, params *params.CreateDatabaseParam) error {
	swClient := sacloud.NewSwitchOp(ctx.Client())
	_, err := swClient.Read(ctx, ctx.Zone(), params.SwitchId)
	if err != nil {
		return fmt.Errorf("Switch(%d) not found", params.SwitchId)
	}

	// set params
	dbTypeInfo := types.RDBMSVersions[types.RDBMSTypeFromString(params.Database)]
	builder := &databaseBuilder.Builder{
		PlanID:         types.ID(params.Plan),
		SwitchID:       params.SwitchId,
		IPAddresses:    []string{params.Ipaddress1},
		NetworkMaskLen: params.NwMaskLen,
		DefaultRoute:   params.DefaultRoute,
		Conf: &sacloud.DatabaseRemarkDBConfCommon{
			DatabaseName:     dbTypeInfo.Name,
			DatabaseVersion:  dbTypeInfo.Version,
			DatabaseRevision: dbTypeInfo.Revision,
			DefaultUser:      params.Username,
			UserPassword:     params.Password,
		},
		SourceID: 0,
		CommonSetting: &sacloud.DatabaseSettingCommon{
			WebUI:           types.WebUI(fmt.Sprintf("%t", params.EnableWebUi)),
			ServicePort:     params.Port,
			SourceNetwork:   params.SourceNetworks,
			DefaultUser:     params.Username,
			UserPassword:    params.Password,
			ReplicaPassword: params.ReplicaUserPassword,
		},
		BackupSetting: &sacloud.DatabaseSettingBackup{
			Time:      params.BackupTime,
			DayOfWeek: weekday.FromStrings(params.BackupWeekdays),
		},
		ReplicationSetting: nil,
		//ReplicationSetting: &sacloud.DatabaseReplicationSetting{
		//	Model:       "",
		//	IPAddress:   "",
		//	Port:        0,
		//	User:        "",
		//	Password:    "",
		//	ApplianceID: 0,
		//},
		Name:         params.Name,
		Description:  params.Description,
		Tags:         params.Tags,
		IconID:       params.IconId,
		SetupOptions: nil,
		Client:       databaseBuilder.NewAPIClient(ctx.Client()),
	}
	if err := builder.Validate(ctx, ctx.Zone()); err != nil {
		return err
	}

	var db *sacloud.Database
	err = ctx.ExecWithProgress(func() error {
		created, err := builder.Build(ctx, ctx.Zone())
		if err != nil {
			return err
		}
		db = created
		return nil
	})
	if err != nil {
		return fmt.Errorf("DatabaseCreate is failed: %s", err)
	}

	return ctx.Output().Print(db)
}

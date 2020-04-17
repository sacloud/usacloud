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
	"github.com/sacloud/usacloud/pkg/internal"
	"github.com/sacloud/usacloud/pkg/params"
)

func DatabaseClone(ctx cli.Context, params *params.CloneDatabaseParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDatabaseAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DatabaseClone is failed: %s", e)
	}

	// set params
	servicePort := p.Settings.DBConf.Common.ServicePort
	port, err := servicePort.Int64()
	if servicePort.String() != "" && err != nil {
		return fmt.Errorf("DatabaseClone is failed: %s", err)
	}

	if params.SwitchId == 0 {
		params.SwitchId = p.Switch.ID
	}
	if params.NwMaskLen == 0 {
		params.NwMaskLen = p.Remark.Network.NetworkMaskLen
	}
	if params.DefaultRoute == "" {
		params.DefaultRoute = p.Remark.Network.DefaultRoute
	}
	if params.Port == 0 {
		params.Port = int(port)
	}
	if !ctx.IsSet("plan") {
		params.Plan = int(p.Plan.ID)
	}

	cloneParam := &sacloud.CreateDatabaseValue{
		Plan: sacloud.DatabasePlan(params.Plan),
		// DefaultUser:     p.Settings.DBConf.Common.DefaultUser,
		// UserPassword:    p.Settings.DBConf.Common.UserPassword,
		SourceNetwork:   params.SourceNetworks,
		ServicePort:     params.Port,
		EnableBackup:    params.EnableBackup,
		BackupTime:      params.BackupTime,
		BackupDayOfWeek: params.BackupWeekdays,
		WebUI:           params.EnableWebUi,
		SwitchID:        params.SwitchId,
		IPAddress1:      params.Ipaddress1,
		MaskLen:         params.NwMaskLen,
		DefaultRoute:    params.DefaultRoute,
		Name:            params.Name,
		Description:     params.Description,
		Tags:            params.Tags,
		Icon:            sacloud.NewResource(params.IconId),
		DatabaseName:    p.Remark.DBConf.Common.DatabaseName,
		DatabaseVersion: p.Remark.DBConf.Common.DatabaseVersion,
		ReplicaPassword: params.ReplicaUserPassword,
		SourceAppliance: p.Resource,
	}

	clone := sacloud.CreateNewDatabase(cloneParam)

	// call manipurate functions
	res, err := api.Create(clone)
	if err != nil {
		return fmt.Errorf("DatabaseClone is failed: %s", err)
	}

	// wait for boot
	err = internal.ExecWithProgress(
		fmt.Sprintf("Still cloning[ID:%d]...", res.ID),
		fmt.Sprintf("Clone database[ID:%d]", res.ID),
		ctx.IO().Progress(),
		func(compChan chan bool, errChan chan error) {
			// call manipurate functions
			err := api.SleepWhileCopying(res.ID, client.DefaultTimeoutDuration, 20)
			if err != nil {
				errChan <- err
				return
			}
			err = api.SleepUntilDatabaseRunning(res.ID, client.DefaultTimeoutDuration, 30)
			if err != nil {
				errChan <- err
				return
			}
			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("DatabaseClone is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}

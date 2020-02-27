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
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
)

func DatabaseCreate(ctx command.Context, params *params.CreateDatabaseParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDatabaseAPI()

	// validation
	sw, err := client.GetSwitchAPI().Read(params.SwitchId)
	if err != nil {
		return fmt.Errorf("Switch(%d) is not found", params.SwitchId)
	}

	// set params
	var p *sacloud.CreateDatabaseValue
	switch params.Database {
	case "postgresql":
		p = sacloud.NewCreatePostgreSQLDatabaseValue()
	case "mariadb":
		p = sacloud.NewCreateMariaDBDatabaseValue()
	}

	p.SwitchID = sw.ID
	p.Plan = sacloud.DatabasePlan(params.Plan)
	p.DefaultUser = params.Username
	p.UserPassword = params.Password
	p.ReplicaPassword = params.ReplicaUserPassword

	if ctx.IsSet("port") {
		p.ServicePort = params.Port
	}
	p.IPAddress1 = params.Ipaddress1
	p.MaskLen = params.NwMaskLen
	p.DefaultRoute = params.DefaultRoute
	p.SourceNetwork = params.SourceNetworks
	if params.EnableWebUi {
		p.WebUI = params.EnableWebUi
	}
	p.EnableBackup = params.EnableBackup

	exists := false
	for _, v := range params.BackupWeekdays {
		if v == "all" {
			exists = true
			break
		}
	}
	if exists {
		p.BackupDayOfWeek = sacloud.AllowDatabaseBackupWeekdays()
	} else {
		p.BackupDayOfWeek = params.BackupWeekdays
	}

	p.BackupTime = params.BackupTime

	p.Name = params.Name
	p.Tags = params.Tags
	p.Description = params.Description
	p.Icon = sacloud.NewResource(params.IconId)

	// call Create(id)
	dbParam := api.New(p)
	res, err := api.Create(dbParam)
	if err != nil {
		return fmt.Errorf("DatabaseCreate is failed: %s", err)
	}

	// wait for boot
	err = internal.ExecWithProgress(
		fmt.Sprintf("Still creating[ID:%d]...", res.ID),
		fmt.Sprintf("Create database[ID:%d]", res.ID),
		command.GlobalOption.Progress,
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
		return fmt.Errorf("DatabaseCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)
}

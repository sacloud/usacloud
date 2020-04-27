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

package vpcrouter

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func L2TPServerInfo(ctx cli.Context, params *params.L2TPServerInfoVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterL2TPServerInfo is failed: %s", e)
	}

	type l2tpConf struct {
		*sacloud.VPCRouterL2TPIPsecServerConfig
		Enabled string
	}

	var cnf *l2tpConf
	if p.HasL2TPIPsecServer() {
		cnf = &l2tpConf{
			VPCRouterL2TPIPsecServerConfig: p.Settings.Router.L2TPIPsecServer.Config,
			Enabled:                        p.Settings.Router.L2TPIPsecServer.Enabled,
		}
	} else {
		cnf = &l2tpConf{
			Enabled: "False",
		}
	}

	return ctx.Output().Print(cnf)
}

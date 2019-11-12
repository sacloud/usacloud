// Copyright 2017-2019 The Usacloud Authors
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
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterL2tpServerInfo(ctx command.Context, params *params.L2tpServerInfoVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterL2tpServerInfo is failed: %s", e)
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

	return ctx.GetOutput().Print(cnf)
}

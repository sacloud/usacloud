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
	"github.com/sacloud/usacloud/pkg/params"
)

func VPCRouterPptpServerInfo(ctx cli.Context, params *params.PptpServerInfoVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterPptpServerInfo is failed: %s", e)
	}

	type pptpConf struct {
		*sacloud.VPCRouterPPTPServerConfig
		Enabled string
	}

	var cnf *pptpConf
	if p.HasPPTPServer() {
		cnf = &pptpConf{
			VPCRouterPPTPServerConfig: p.Settings.Router.PPTPServer.Config,
			Enabled:                   p.Settings.Router.PPTPServer.Enabled,
		}
	} else {
		cnf = &pptpConf{
			Enabled: "False",
		}
	}

	return ctx.GetOutput().Print(cnf)

}

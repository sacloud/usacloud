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

package privatehost

import (
	"fmt"

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func ServerInfo(ctx cli.Context, params *params.ServerInfoPrivateHostParam) error {

	client := ctx.GetAPIClient()
	api := client.GetPrivateHostAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("PrivateHostServerInfo is failed: %s", e)
	}

	res, err := client.GetServerAPI().Find()
	if err != nil {
		return fmt.Errorf("PrivateHostServerInfo is failed: %s", err)
	}

	list := []interface{}{}
	for i, s := range res.Servers {

		if s.PrivateHost != nil && s.PrivateHost.ID == p.ID {
			list = append(list, &res.Servers[i])
		}
	}
	return ctx.Output().Print(list...)
}

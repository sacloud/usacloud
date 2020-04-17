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

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func VPCRouterUserInfo(ctx cli.Context, params *params.UserInfoVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterUserInfo is failed: %s", e)
	}

	if !p.HasRemoteAccessUsers() {
		fmt.Fprintf(ctx.IO().Err(), "VPCRouter[%d] don't have any remote-access users\n", params.Id)
		return nil
	}

	confList := p.Settings.Router.RemoteAccessUsers.Config
	// build parameters to display table
	list := []interface{}{}
	for i := range confList {
		list = append(list, &confList[i])
	}

	return ctx.GetOutput().Print(list...)
}

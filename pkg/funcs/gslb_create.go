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

func GSLBCreate(ctx cli.Context, params *params.CreateGSLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetGSLBAPI()
	p := api.New(params.Name)

	// set params
	p.SetTags(params.Tags)
	p.SetIconByID(params.IconId)
	p.SetName(params.Name)
	p.SetDescription(params.Description)

	// set health check
	switch params.Protocol {
	case "http":
		if params.Path == "" || params.ResponseCode == 0 {
			return fmt.Errorf("path and response-code is required when protocol is http")
		}
		p.SetHTTPHealthCheck(params.HostHeader, params.Path, params.ResponseCode)
	case "https":
		if params.Path == "" || params.ResponseCode == 0 {
			return fmt.Errorf("path and response-code is required when protocol is https")
		}
		p.SetHTTPSHealthCheck(params.HostHeader, params.Path, params.ResponseCode)
	case "ping":
		p.SetPingHealthCheck()
	case "tcp":
		if params.Port == 0 {
			return fmt.Errorf("port is required when protocol is tcp")
		}
		p.SetTCPHealthCheck(params.Port)
	}
	p.SetSorryServer(params.SorryServer)
	p.SetDelayLoop(params.DelayLoop)
	p.SetWeightedEnable(params.Weighted)

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("GSLBCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}

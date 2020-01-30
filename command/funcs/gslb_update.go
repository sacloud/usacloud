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
	"strconv"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func GSLBUpdate(ctx command.Context, params *params.UpdateGSLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetGSLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("GSLBUpdate is failed: %s", e)
	}

	// set params

	if ctx.IsSet("icon-id") {
		p.SetIconByID(params.IconId)
	}
	if ctx.IsSet("name") {
		p.SetName(params.Name)
	}
	if ctx.IsSet("description") {
		p.SetDescription(params.Description)
	}
	if ctx.IsSet("tags") {
		p.SetTags(params.Tags)
	}

	if ctx.IsSet("protocol") {

		switch params.Protocol {
		case "http":

			if p.Settings.GSLB.HealthCheck.Protocol != "http" && (params.Path == "" || params.ResponseCode == 0) {
				return fmt.Errorf("path and response-code is required when protocol is http")
			}

			hostHeader := p.Settings.GSLB.HealthCheck.Host
			path := p.Settings.GSLB.HealthCheck.Path
			responseCode := p.Settings.GSLB.HealthCheck.Status

			if ctx.IsSet("host-header") {
				hostHeader = params.HostHeader
			}
			if ctx.IsSet("path") {
				path = params.Path
			}
			if ctx.IsSet("response-code") {
				responseCode = fmt.Sprintf("%d", params.ResponseCode)
			}

			code, err := strconv.Atoi(responseCode)
			if err != nil {
				return fmt.Errorf("GSLBUpdate is failed: %s", e)
			}
			p.SetHTTPHealthCheck(hostHeader, path, code)
		case "https":

			if p.Settings.GSLB.HealthCheck.Protocol != "https" && (params.Path == "" || params.ResponseCode == 0) {
				return fmt.Errorf("path and response-code is required when protocol is http")
			}

			hostHeader := p.Settings.GSLB.HealthCheck.Host
			path := p.Settings.GSLB.HealthCheck.Path
			responseCode := p.Settings.GSLB.HealthCheck.Status

			if ctx.IsSet("host-header") {
				hostHeader = params.HostHeader
			}
			if ctx.IsSet("path") {
				path = params.Path
			}
			if ctx.IsSet("response-code") {
				responseCode = fmt.Sprintf("%d", params.ResponseCode)
			}

			code, err := strconv.Atoi(responseCode)
			if err != nil {
				return fmt.Errorf("GSLBUpdate is failed: %s", e)
			}
			p.SetHTTPSHealthCheck(hostHeader, path, code)
		case "ping":
			p.SetPingHealthCheck()
		case "tcp":
			if p.Settings.GSLB.HealthCheck.Protocol != "tcp" && params.Port == 0 {
				return fmt.Errorf("port is required when protocol is tcp")
			}
			port := p.Settings.GSLB.HealthCheck.Port
			if ctx.IsSet("port") {
				port = fmt.Sprintf("%d", params.Port)
			}
			intPort, err := strconv.Atoi(port)
			if err != nil {
				return fmt.Errorf("GSLBUpdate is failed: %s", e)
			}
			p.SetTCPHealthCheck(intPort)
		}

	}

	if ctx.IsSet("sorry-server") {
		p.SetSorryServer(params.SorryServer)
	}
	if ctx.IsSet("weighted") {
		p.SetWeightedEnable(params.Weighted)
	}
	if ctx.IsSet("delay-loop") {
		p.SetDelayLoop(params.DelayLoop)
	}

	// call Update(id)
	res, err := api.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("GSLBUpdate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}

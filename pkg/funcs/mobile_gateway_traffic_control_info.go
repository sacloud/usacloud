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
	"net/http"

	sacloudapi "github.com/sacloud/libsacloud/api"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func MobileGatewayTrafficControlInfo(ctx cli.Context, params *params.TrafficControlInfoMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()

	// set params

	// call Read(id)
	conf, err := api.GetTrafficMonitoringConfig(params.Id)
	if err != nil {
		if e, ok := err.(sacloudapi.Error); ok && e.ResponseCode() == http.StatusNotFound {
			return fmt.Errorf("MobileGatewayTrafficControlInfo is failed: Traffic Monitoring is disabled")
		}
		return fmt.Errorf("MobileGatewayTrafficControlInfo is failed: %s", err)
	}

	status, err := api.GetTrafficStatus(params.Id)
	if err != nil {
		if e, ok := err.(sacloudapi.Error); ok && e.ResponseCode() == http.StatusNotFound {
			return fmt.Errorf("MobileGatewayTrafficControlInfo is failed: Traffic Status Not Found")
		}
		return fmt.Errorf("MobileGatewayTrafficControlInfo is failed: %s", err)
	}

	info := &struct {
		*sacloud.TrafficMonitoringConfig `json:"config,omitempty"`
		*sacloud.TrafficStatus           `json:"status,omitempty"`
	}{
		TrafficMonitoringConfig: conf,
		TrafficStatus:           status,
	}

	return ctx.GetOutput().Print(info)
}

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

func MobileGatewayTrafficControlEnable(ctx cli.Context, params *params.TrafficControlEnableMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("MobileGatewayTrafficControlEnable is failed: %s", e)
	}

	// set params
	config := &sacloud.TrafficMonitoringConfig{}
	config.TrafficQuotaInMB = params.Quota
	config.BandWidthLimitInKbps = params.BandWidthLimit
	config.AutoTrafficShaping = params.AutoTrafficShaping

	config.EMailConfig = &sacloud.TrafficMonitoringNotifyEmail{
		Enabled: params.EnableEmail,
	}

	config.SlackConfig = &sacloud.TrafficMonitoringNotifySlack{
		Enabled: false,
	}
	if params.SlackWebhookUrl != "" {
		config.SlackConfig = &sacloud.TrafficMonitoringNotifySlack{
			Enabled:             true,
			IncomingWebhooksURL: params.SlackWebhookUrl,
		}
	}

	// call Update(id)
	if _, err := api.SetTrafficMonitoringConfig(params.Id, config); err != nil {
		return fmt.Errorf("MobileGatewayTrafficControlEnable is failed: %s", err)
	}

	return nil
}

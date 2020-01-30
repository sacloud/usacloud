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
	"github.com/sacloud/usacloud/command/params"
)

func MobileGatewayTrafficControlUpdate(ctx command.Context, params *params.TrafficControlUpdateMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("MobileGatewayTrafficControlUpdate is failed: %s", e)
	}

	// set params
	config, err := api.GetTrafficMonitoringConfig(params.Id)
	if err != nil {
		return fmt.Errorf("MobileGatewayTrafficControlUpdate is failed: %s", err)
	}

	if ctx.IsSet("quota") {
		config.TrafficQuotaInMB = params.Quota
	}
	if ctx.IsSet("band-width-limit") {
		config.BandWidthLimitInKbps = params.BandWidthLimit
	}
	if ctx.IsSet("auto-traffic-shaping") {
		config.AutoTrafficShaping = params.AutoTrafficShaping
	}

	if ctx.IsSet("enable-email") {
		config.EMailConfig = &sacloud.TrafficMonitoringNotifyEmail{
			Enabled: params.EnableEmail,
		}
	}

	if ctx.IsSet("slack-webhook-url") {
		if params.SlackWebhookUrl == "" {
			config.SlackConfig = &sacloud.TrafficMonitoringNotifySlack{
				Enabled: false,
			}
		} else {
			config.SlackConfig = &sacloud.TrafficMonitoringNotifySlack{
				Enabled:             true,
				IncomingWebhooksURL: params.SlackWebhookUrl,
			}
		}
	}

	// call Update(id)
	if _, err = api.SetTrafficMonitoringConfig(params.Id, config); err != nil {
		return fmt.Errorf("MobileGatewayTrafficControlEnable is failed: %s", err)
	}

	return nil
}

package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func MobileGatewayTrafficControlEnable(ctx command.Context, params *params.TrafficControlEnableMobileGatewayParam) error {

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

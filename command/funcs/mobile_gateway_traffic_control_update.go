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

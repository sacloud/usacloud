package funcs

import (
	"fmt"
	"net/http"

	sacloudapi "github.com/sacloud/libsacloud/api"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func MobileGatewayTrafficControlInfo(ctx command.Context, params *params.TrafficControlInfoMobileGatewayParam) error {

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

package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
)

func MobileGatewayCreate(ctx command.Context, params *params.CreateMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()

	createMGWValues := &sacloud.CreateMobileGatewayValue{
		Name:        params.Name,
		Description: params.Description,
		Tags:        params.Tags,
	}
	mgwSetting := &sacloud.MobileGatewaySetting{
		InternetConnection: &sacloud.MGWInternetConnection{
			Enabled: "False",
		},
		Interfaces: []*sacloud.MGWInterface{
			nil,
		},
	}
	if params.InternetConnection {
		mgwSetting.InternetConnection.Enabled = "True"
	}

	p, err := sacloud.CreateNewMobileGateway(createMGWValues, mgwSetting)
	if err != nil {
		return fmt.Errorf("MobileGatewayCreate is failed: %s", err)
	}

	p.SetIconByID(params.IconId)

	var res *sacloud.MobileGateway
	err = internal.ExecWithProgress(
		fmt.Sprintf("Still creating..."),
		fmt.Sprintf("Create mobile-gateway"),
		command.GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {
			// call manipurate functions
			// call Create(id)
			res, err = api.Create(p)
			if err != nil {
				errChan <- fmt.Errorf("MobileGatewayCreate is failed: %s", err)
				return
			}
			err = api.SleepWhileCopying(res.ID, client.DefaultTimeoutDuration, 20)
			if err != nil {
				errChan <- err
				return
			}

			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("MobileGatewayBoot is failed: %s", err)
	}

	res, err = api.Read(res.ID)
	if err != nil {
		return fmt.Errorf("MobileGatewayBoot is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)
}

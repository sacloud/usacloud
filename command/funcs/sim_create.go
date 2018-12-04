package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/define"
)

func SIMCreate(ctx command.Context, params *params.CreateSIMParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSIMAPI()
	p := api.New(params.Name, params.Iccid, params.Passcode)

	// set params
	p.SetDescription(params.Description)
	p.SetTags(params.Tags)
	p.SetIconByID(params.IconId)

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("SIMCreate is failed: %s", err)
	}

	var careers []*sacloud.SIMNetworkOperatorConfig
	for _, career := range params.Career {
		careers = append(careers, &sacloud.SIMNetworkOperatorConfig{
			Allow: true,
			Name:  define.SIMCareers[career],
		})
	}
	if _, err := api.SetNetworkOperator(res.ID, careers...); err != nil {
		return fmt.Errorf("SIMCreate is failed: %s", err)
	}

	if !params.Disabled {
		// activate sim
		if _, err := api.Activate(res.ID); err != nil {
			return fmt.Errorf("SIMCreate is failed: %s", err)
		}
	}

	if params.Imei != "" {
		// set imei lock
		if _, err := api.IMEILock(res.ID, params.Imei); err != nil {
			return fmt.Errorf("SIMCreate is failed: %s", err)
		}
	}

	return ctx.GetOutput().Print(res)

}

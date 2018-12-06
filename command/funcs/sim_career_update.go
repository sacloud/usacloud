package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/define"
)

func SIMCareerUpdate(ctx command.Context, params *params.CareerUpdateSIMParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSIMAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("SIMCareerUpdate is failed: %s", e)
	}

	// set params
	var careers []*sacloud.SIMNetworkOperatorConfig
	for _, career := range params.Career {
		careers = append(careers, &sacloud.SIMNetworkOperatorConfig{
			Allow: true,
			Name:  define.SIMCareers[career],
		})
	}
	if _, err := api.SetNetworkOperator(params.Id, careers...); err != nil {
		return fmt.Errorf("SIMCareerUpdate is failed: %s", err)
	}

	return nil
}

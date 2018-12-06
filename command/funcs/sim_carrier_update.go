package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/define"
)

func SIMCarrierUpdate(ctx command.Context, params *params.CarrierUpdateSIMParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSIMAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("SIMCarrierUpdate is failed: %s", e)
	}

	// set params
	var carriers []*sacloud.SIMNetworkOperatorConfig
	for _, carrier := range params.Carrier {
		carriers = append(carriers, &sacloud.SIMNetworkOperatorConfig{
			Allow: true,
			Name:  define.SIMCarrier[carrier],
		})
	}
	if _, err := api.SetNetworkOperator(params.Id, carriers...); err != nil {
		return fmt.Errorf("SIMCarrierUpdate is failed: %s", err)
	}

	return nil
}

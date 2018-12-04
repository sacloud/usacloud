package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func SIMDelete(ctx command.Context, params *params.DeleteSIMParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSIMAPI()

	// call
	res, err := api.Reset().Include("*").Include("Status.sim").Find()
	if err != nil {
		return fmt.Errorf("SIMDelete is failed: %s", err)
	}
	var sim *sacloud.SIM
	for _, s := range res.CommonServiceSIMItems {
		if s.ID == params.Id {
			sim = &s
			break
		}
	}

	if sim == nil {
		return fmt.Errorf("SIMDelete is failed: SIM[%d] is not found", params.Id)
	}

	if params.Force && sim.Status.SIMInfo.Activated {
		_, err := api.Deactivate(params.Id)
		if err != nil {
			return fmt.Errorf("SIMDelete is failed: %s", err)
		}
	}

	// call Delete(id)
	if _, err := api.Delete(params.Id); err != nil {
		return fmt.Errorf("SIMDelete is failed: %s", err)
	}
	return nil
}

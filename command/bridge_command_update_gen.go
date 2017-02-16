package command

import (
	"fmt"
)

func BridgeUpdate(ctx Context, params *UpdateBridgeParam) error {

	client := ctx.GetAPIClient()
	api := client.GetBridgeAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("BridgeUpdate is failed: %s", e)
	}

	// set params

	if ctx.IsSet("name") {
		p.SetName(params.Name)
	}
	if ctx.IsSet("description") {
		p.SetDescription(params.Description)
	}

	// call Update(id)
	res, err := api.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("BridgeUpdate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}

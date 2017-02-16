package command

import (
	"fmt"
)

func BridgeCreate(ctx Context, params *CreateBridgeParam) error {

	client := ctx.GetAPIClient()
	api := client.GetBridgeAPI()
	p := api.New()

	// set params

	p.SetDescription(params.Description)

	p.SetName(params.Name)

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("BridgeCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}

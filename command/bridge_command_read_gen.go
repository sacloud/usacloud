package command

import (
	"fmt"
)

func BridgeRead(ctx Context, params *ReadBridgeParam) error {

	client := ctx.GetAPIClient()
	api := client.GetBridgeAPI()

	// set params

	// call Read(id)
	res, err := api.Read(params.Id)
	if err != nil {
		return fmt.Errorf("BridgeRead is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}

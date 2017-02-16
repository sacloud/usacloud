package command

import (
	"fmt"
)

func BridgeDelete(ctx Context, params *DeleteBridgeParam) error {

	client := ctx.GetAPIClient()
	api := client.GetBridgeAPI()

	// set params

	// call Delete(id)
	res, err := api.Delete(params.Id)
	if err != nil {
		return fmt.Errorf("BridgeDelete is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}

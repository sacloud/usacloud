package command

import (
	"fmt"
)

func SwitchDelete(ctx Context, params *DeleteSwitchParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSwitchAPI()

	// set params

	// call Delete(id)
	res, err := api.Delete(params.Id)
	if err != nil {
		return fmt.Errorf("SwitchDelete is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}

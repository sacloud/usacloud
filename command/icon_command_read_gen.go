package command

import (
	"fmt"
)

func IconRead(ctx Context, params *ReadIconParam) error {

	client := ctx.GetAPIClient()
	api := client.GetIconAPI()

	// set params

	// call Read(id)
	res, err := api.Read(params.Id)
	if err != nil {
		return fmt.Errorf("IconRead is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}

package command

import (
	"fmt"
)

func IconDelete(ctx Context, params *DeleteIconParam) error {

	client := ctx.GetAPIClient()
	api := client.GetIconAPI()

	// set params

	// call Delete(id)
	res, err := api.Delete(params.Id)
	if err != nil {
		return fmt.Errorf("IconDelete is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}

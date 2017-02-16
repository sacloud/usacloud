package command

import (
	"fmt"
)

func ISOImageDelete(ctx Context, params *DeleteISOImageParam) error {

	client := ctx.GetAPIClient()
	api := client.GetCDROMAPI()

	// set params

	// call Delete(id)
	res, err := api.Delete(params.Id)
	if err != nil {
		return fmt.Errorf("ISOImageDelete is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}

package command

import (
	"fmt"
)

func ISOImageRead(ctx Context, params *ReadISOImageParam) error {

	client := ctx.GetAPIClient()
	api := client.GetCDROMAPI()

	// set params

	// call Read(id)
	res, err := api.Read(params.Id)
	if err != nil {
		return fmt.Errorf("ISOImageRead is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}

package command

import (
	"fmt"
)

func InternetRead(ctx Context, params *ReadInternetParam) error {

	client := ctx.GetAPIClient()
	api := client.GetInternetAPI()

	// set params

	// call Read(id)
	res, err := api.Read(params.Id)
	if err != nil {
		return fmt.Errorf("InternetRead is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}

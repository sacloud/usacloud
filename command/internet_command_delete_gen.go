package command

import (
	"fmt"
)

func InternetDelete(ctx Context, params *DeleteInternetParam) error {

	client := ctx.GetAPIClient()
	api := client.GetInternetAPI()

	// set params

	// call Delete(id)
	res, err := api.Delete(params.Id)
	if err != nil {
		return fmt.Errorf("InternetDelete is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}

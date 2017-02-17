package command

import (
	"fmt"
)

func InternetUpdateBandwidth(ctx Context, params *UpdateBandwidthInternetParam) error {

	client := ctx.GetAPIClient()
	api := client.GetInternetAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("InternetUpdateBandwidth is failed: %s", e)
	}

	// set params

	if ctx.IsSet("band-width") {
		p.SetBandWidthMbps(params.BandWidth)
	}

	// call manipurate functions
	res, err := api.UpdateBandWidth(params.Id, params.BandWidth)
	if err != nil {
		return fmt.Errorf("InternetUpdateBandwidth is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}

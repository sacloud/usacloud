package command

import (
	"fmt"
)

func ISOImageFtpOpen(ctx Context, params *FtpOpenISOImageParam) error {

	client := ctx.GetAPIClient()
	api := client.GetCDROMAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ISOImageFtpOpen is failed: %s", e)
	}

	// check scope(download can user scope only)
	if !p.IsUserScope() {
		return fmt.Errorf("ISOImageFtpOpen is failed: %s", "Only the ISO Image of scope=`user` can be downloaded")
	}

	// call manipurate functions
	res, err := api.OpenFTP(p.ID, true)
	if err != nil {
		return fmt.Errorf("ISOImageFtpOpen is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}

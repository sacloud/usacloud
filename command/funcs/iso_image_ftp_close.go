package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ISOImageFtpClose(ctx command.Context, params *params.FtpCloseISOImageParam) error {

	client := ctx.GetAPIClient()
	api := client.GetCDROMAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ISOImageFtpClose is failed: %s", e)
	}

	// check scope(download can user scope only)
	if !p.IsUserScope() {
		return fmt.Errorf("ISOImageFtpClose is failed: %s", "Only the ISO Image of scope=`user` can be downloaded")
	}

	// call manipurate functions
	// close
	_, err := api.CloseFTP(p.ID)
	if err != nil {
		return fmt.Errorf("ISOImageFtpClose is failed: %s", err)
	}

	return nil

}

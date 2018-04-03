package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ArchiveFtpClose(ctx command.Context, params *params.FtpCloseArchiveParam) error {

	client := ctx.GetAPIClient()
	api := client.GetArchiveAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ArchiveFtpClose is failed: %s", e)
	}

	// check scope(download can user scope only)
	if !p.IsUserScope() {
		return fmt.Errorf("ArchiveFtpClose is failed: %s", "Only the ISO Image of scope=`user` can be downloaded")
	}

	// call manipurate functions
	// close
	_, err := api.CloseFTP(p.ID)
	if err != nil {
		return fmt.Errorf("ArchiveFtpClose is failed: %s", err)
	}

	return nil
}

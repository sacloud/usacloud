package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/ftp"
)

func ArchiveDownload(ctx command.Context, params *params.DownloadArchiveParam) error {

	client := ctx.GetAPIClient()
	api := client.GetArchiveAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ArchiveDownload is failed: %s", e)
	}

	// check scope(download can user scope only)
	if !p.IsUserScope() {
		return fmt.Errorf("ArchiveDownload is failed: %s", "Only the ISO Image of scope=`user` can be downloaded")
	}

	// call manipurate functions
	res, err := api.OpenFTP(p.ID)
	if err != nil {
		return fmt.Errorf("ArchiveDownload is failed: %s", err)
	}

	// download
	ftpsClient := ftp.NewClient(res.User, res.Password, res.HostName)

	err = internal.ExecWithProgress(
		fmt.Sprintf("Still downloading[ID:%d]...", params.Id),
		fmt.Sprintf("Download archive[ID:%d]", params.Id),
		command.GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {
			err = ftpsClient.Download(params.FileDestination)
			if err != nil {
				errChan <- err
				return
			}
			compChan <- true

		},
	)

	if err != nil {
		return fmt.Errorf("ArchiveDownload is failed: %s", err)
	}

	// close
	_, err = api.CloseFTP(p.ID)
	if err != nil {
		return fmt.Errorf("ArchiveDownload is failed: %s", err)
	}

	return nil
}

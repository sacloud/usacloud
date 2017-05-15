package command

import (
	"fmt"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/ftp"
)

func ISOImageDownload(ctx Context, params *DownloadISOImageParam) error {

	client := ctx.GetAPIClient()
	api := client.GetCDROMAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ISOImageDownload is failed: %s", e)
	}

	// check scope(download can user scope only)
	if !p.IsUserScope() {
		return fmt.Errorf("ISOImageDownload is failed: %s", "Only the ISO Image of scope=`user` can be downloaded")
	}

	// call manipurate functions
	res, err := api.OpenFTP(p.ID, false)
	if err != nil {
		return fmt.Errorf("ISOImageDownload is failed: %s", err)
	}

	// download
	ftpsClient := ftp.NewClient(res.User, res.Password, res.HostName)
	compChan := make(chan bool)
	errChan := make(chan error)

	spinner := internal.NewProgress(
		fmt.Sprintf("Still downloading[ID:%d]...", params.Id),
		fmt.Sprintf("Download ISOImage[ID:%d]", params.Id),
		GlobalOption.Progress)
	go func() {
		spinner.Start()
		err = ftpsClient.Download(params.FileDestination)
		if err != nil {
			errChan <- err
			return
		}
		compChan <- true
	}()

download:
	for {
		select {
		case <-compChan:
			spinner.Stop()
			break download
		case err := <-errChan:
			return fmt.Errorf("ISOImageDownload is failed: %s", err)
		}
	}

	// close
	_, err = api.CloseFTP(p.ID)
	if err != nil {
		return fmt.Errorf("ISOImageDownload is failed: %s", err)
	}

	return nil
}

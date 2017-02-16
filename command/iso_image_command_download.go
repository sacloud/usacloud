package command

import (
	"fmt"
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
	ftpsClient := ftp.NewFTPClient(res.User, res.Password, res.HostName)
	err = ftpsClient.Download(params.IsoFile)
	if err != nil {
		return fmt.Errorf("ISOImageDownload is failed: %s", err)
	}

	// close
	_, err = api.CloseFTP(p.ID)
	if err != nil {
		return fmt.Errorf("ISOImageDownload is failed: %s", err)
	}

	p, e = api.Read(p.ID)
	if err != nil {
		return fmt.Errorf("ISOImageDownload is failed: %s", err)
	}

	return ctx.GetOutput().Print()

}

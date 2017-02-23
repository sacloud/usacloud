package command

import (
	"fmt"
	"github.com/sacloud/usacloud/ftp"
)

func ArchiveDownload(ctx Context, params *DownloadArchiveParam) error {

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
	ftpsClient := ftp.NewFTPClient(res.User, res.Password, res.HostName)
	err = ftpsClient.Download(params.FileDestination)
	if err != nil {
		return fmt.Errorf("ArchiveDownload is failed: %s", err)
	}

	// close
	_, err = api.CloseFTP(p.ID)
	if err != nil {
		return fmt.Errorf("ArchiveDownload is failed: %s", err)
	}

	p, e = api.Read(p.ID)
	if err != nil {
		return fmt.Errorf("ArchiveDownload is failed: %s", err)
	}

	return ctx.GetOutput().Print()
}

package command

import (
	"fmt"
	"github.com/sacloud/usacloud/ftp"
)

func ISOImageUpload(ctx Context, params *UploadISOImageParam) error {

	client := ctx.GetAPIClient()
	api := client.GetCDROMAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ISOImageUpload is failed: %s", e)
	}

	// set params

	params.getCommandDef().Params["iso-file"].CustomHandler("IsoFile", params, p)

	// open FTP
	ftpServer, err := api.OpenFTP(p.ID, false)
	if err != nil {
		return fmt.Errorf("ISOImageUpload is failed: %s", err)
	}

	// upload
	ftpsClient := ftp.NewFTPClient(ftpServer.User, ftpServer.Password, ftpServer.HostName)
	err = ftpsClient.Upload(params.GetIsoFile())
	if err != nil {
		return fmt.Errorf("ISOImageUpload is failed: %s", err)
	}

	// close FTP
	_, err = api.CloseFTP(p.ID)
	if err != nil {
		return fmt.Errorf("ISOImageUpload is failed: %s", err)
	}

	// read again
	p, err = api.Read(p.ID)
	if err != nil {
		return fmt.Errorf("ISOImageUpload is failed: %s", err)
	}
	return ctx.GetOutput().Print(p)
}

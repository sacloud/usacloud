package funcs

import (
	"fmt"

	"github.com/sacloud/ftps"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
)

func ISOImageUpload(ctx command.Context, params *params.UploadISOImageParam) error {

	client := ctx.GetAPIClient()
	api := client.GetCDROMAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ISOImageUpload is failed: %s", e)
	}

	// open FTP
	ftpServer, err := api.OpenFTP(p.ID, false)
	if err != nil {
		return fmt.Errorf("ISOImageUpload is failed: %s", err)
	}

	// upload
	ftpsClient := ftps.NewClient(ftpServer.User, ftpServer.Password, ftpServer.HostName)
	err = internal.ExecWithProgress(
		fmt.Sprintf("Still uploading[ID:%d]...", params.Id),
		fmt.Sprintf("Upload iso-image[ID:%d]", params.Id),
		command.GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {

			file, df, err := fileOrStdin(params.GetIsoFile())
			if err != nil {
				errChan <- err
				return
			}
			defer df()

			err = ftpsClient.UploadFile("upload.iso", file)
			if err != nil {
				errChan <- err
				return
			}
			compChan <- true
		},
	)
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

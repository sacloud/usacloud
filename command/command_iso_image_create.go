package command

import (
	"fmt"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/ftp"
)

func ISOImageCreate(ctx Context, params *CreateISOImageParam) error {

	client := ctx.GetAPIClient()
	api := client.GetCDROMAPI()
	p := api.New()

	// set params

	p.SetTags(params.Tags)
	p.SetIconByID(params.IconId)
	p.SetSizeGB(params.Size)
	p.SetName(params.Name)
	p.SetDescription(params.Description)

	// call Create(id)
	res, ftpServer, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("ISOImageCreate is failed: %s", err)
	}

	// upload
	ftpsClient := ftp.NewClient(ftpServer.User, ftpServer.Password, ftpServer.HostName)
	err = internal.ExecWithProgress(
		fmt.Sprintf("Still uploading[ID:%d]...", res.ID),
		fmt.Sprintf("Upload iso-image[ID:%d]", res.ID),
		GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {
			err = ftpsClient.Upload(params.GetIsoFile())
			if err != nil {
				errChan <- err
			}
			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("ISOImageCreate is failed: %s", err)
	}

	// close FTP
	_, err = api.CloseFTP(res.ID)
	if err != nil {
		return fmt.Errorf("ISOImageCreate is failed: %s", err)
	}

	// read again
	res, err = api.Read(res.ID)
	if err != nil {
		return fmt.Errorf("ISOImageCreate is failed: %s", err)
	}
	return ctx.GetOutput().Print(res)

}

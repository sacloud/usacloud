package command

import (
	"fmt"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/ftp"
)

func ArchiveUpload(ctx Context, params *UploadArchiveParam) error {

	client := ctx.GetAPIClient()
	api := client.GetArchiveAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ArchiveUpload is failed: %s", e)
	}

	// open FTP
	ftpServer, err := api.OpenFTP(p.ID)
	if err != nil {
		return fmt.Errorf("ArchiveUpload is failed: %s", err)
	}

	// upload
	ftpsClient := ftp.NewClient(ftpServer.User, ftpServer.Password, ftpServer.HostName)
	compChan := make(chan bool)
	errChan := make(chan error)

	spinner := internal.NewProgress(
		fmt.Sprintf("Still uploading[ID:%d]...", params.Id),
		fmt.Sprintf("Upload archive[ID:%d]", params.Id),
		GlobalOption.Progress)
	go func() {
		spinner.Start()
		err = ftpsClient.Upload(params.GetArchiveFile())
		if err != nil {
			errChan <- err
		}
		compChan <- true
	}()
upload:
	for {
		select {
		case <-compChan:
			spinner.Stop()
			break upload
		case err := <-errChan:
			return fmt.Errorf("ArchiveUpload is failed: %s", err)
		}
	}

	// close FTP
	_, err = api.CloseFTP(p.ID)
	if err != nil {
		return fmt.Errorf("ArchiveUpload is failed: %s", err)
	}

	// read again
	p, err = api.Read(p.ID)
	if err != nil {
		return fmt.Errorf("ArchiveUpload is failed: %s", err)
	}
	return ctx.GetOutput().Print(p)

}

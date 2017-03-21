package command

import (
	"fmt"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/ftp"
)

func ArchiveCreate(ctx Context, params *CreateArchiveParam) error {

	client := ctx.GetAPIClient()
	api := client.GetArchiveAPI()
	p := api.New()

	// set params
	p.SetDescription(params.Description)
	p.SetTags(params.Tags)
	p.SetIconByID(params.IconId)
	p.SetName(params.Name)

	needUpload := params.SourceDiskId == 0 && params.SourceArchiveId == 0

	if params.SourceDiskId != 0 {
		p.SetSourceDisk(params.SourceDiskId)
	} else if params.SourceArchiveId != 0 {
		p.SetSourceArchive(params.SourceArchiveId)
	} else {
		p.SetSizeGB(params.Size)
	}
	// manual validation
	if needUpload {
		if params.ArchiveFile == "" || params.Size == 0 {
			return fmt.Errorf("ArchiveCreate is required both of %q and %q if when source disk/archive is missing", "--archive-file", "--size")
		}
	}

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("ArchiveCreate is failed: %s", err)
	}

	if needUpload {
		ftpServer, err := api.OpenFTP(res.ID)
		if err != nil {
			return fmt.Errorf("ArchiveCreate is failed: %s", err)
		}

		// upload
		ftpsClient := ftp.NewFTPClient(ftpServer.User, ftpServer.Password, ftpServer.HostName)
		compChan := make(chan bool)
		errChan := make(chan error)

		spinner := internal.NewProgress(
			fmt.Sprintf("Still uploading[ID:%d]...", res.ID),
			fmt.Sprintf("Upload archive[ID:%d]", res.ID),
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
				return fmt.Errorf("ArchiveCreate is failed: %s", err)
			}
		}

		// close FTP
		_, err = api.CloseFTP(res.ID)
		if err != nil {
			return fmt.Errorf("ArchiveCreate is failed: %s", err)
		}

		// read again
		res, err = api.Read(res.ID)
		if err != nil {
			return fmt.Errorf("ArchiveCreate is failed: %s", err)
		}
	} else {
		// wait for copy with progress
		spinner := internal.NewProgress(
			fmt.Sprintf("Still coping[ID:%d]...", res.ID),
			fmt.Sprintf("Copy archive[ID:%d]", res.ID),
			GlobalOption.Progress)
		spinner.Start()
		compChan, progChan, errChan := api.AsyncSleepWhileCopying(res.ID, client.DefaultTimeoutDuration)
	copy:
		for {
			select {
			case r := <-compChan:
				res = r
				spinner.Stop()
				break copy
			case <-progChan:
			// noop
			case err := <-errChan:
				return fmt.Errorf("ArchiveCreate is failed: %s", err)
			}
		}

	}

	return ctx.GetOutput().Print(res)

}

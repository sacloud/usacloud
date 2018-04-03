package funcs

import (
	"fmt"

	"github.com/sacloud/ftps"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
)

func ArchiveCreate(ctx command.Context, params *params.CreateArchiveParam) error {

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
		if params.Size == 0 {
			return fmt.Errorf("ArchiveCreate is required %q if when source disk/archive is missing", "--size")
		}
		errs := validateExistsFileOrStdIn("archive-file", params.ArchiveFile)
		if len(errs) > 0 {
			return errs[0]
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
		ftpsClient := ftps.NewClient(ftpServer.User, ftpServer.Password, ftpServer.HostName)

		err = internal.ExecWithProgress(
			fmt.Sprintf("Still uploading[ID:%d]...", res.ID),
			fmt.Sprintf("Upload archive[ID:%d]", res.ID),
			command.GlobalOption.Progress,
			func(compChan chan bool, errChan chan error) {

				file, df, err := fileOrStdin(params.GetArchiveFile())
				if err != nil {
					errChan <- err
					return
				}
				defer df()

				err = ftpsClient.UploadFile("upload.raw", file)
				if err != nil {
					errChan <- err
					return
				}
				compChan <- true
			},
		)
		if err != nil {
			return fmt.Errorf("ArchiveCreate is failed: %s", err)
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
		err := internal.ExecWithProgress(
			fmt.Sprintf("Still copying[ID:%d]...", res.ID),
			fmt.Sprintf("Copy archive[ID:%d]", res.ID),
			command.GlobalOption.Progress,
			func(compChan chan bool, errChan chan error) {
				err = api.SleepWhileCopying(res.ID, client.DefaultTimeoutDuration)
				if err != nil {
					errChan <- err
				}
				compChan <- true
			},
		)
		if err != nil {
			return fmt.Errorf("ArchiveCreate is failed: %s", err)
		}

	}

	return ctx.GetOutput().Print(res)

}

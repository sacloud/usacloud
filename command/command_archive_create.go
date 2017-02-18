package command

import (
	"fmt"
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
			return fmt.Errorf("ArchiveCreate is required %q and %q if when source disk/archive is missing", "archive-file", "size")
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
		err = ftpsClient.Upload(params.GetArchiveFile())
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
	}

	return ctx.GetOutput().Print(res)

}

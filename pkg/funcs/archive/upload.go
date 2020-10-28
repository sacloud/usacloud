// Copyright 2017-2020 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package archive

import (
	"fmt"

	"github.com/sacloud/usacloud/pkg/term"

	"github.com/sacloud/ftps"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func Upload(ctx cli.Context, params *params.UploadArchiveParam) error {
	client := sacloud.NewArchiveOp(ctx.Client())
	archive, err := client.Read(ctx, ctx.Zone(), params.Id)
	if err != nil {
		return fmt.Errorf("ArchiveUpload is failed: %s", err)
	}

	// open FTP
	ftpServer, err := client.OpenFTP(ctx, ctx.Zone(), archive.ID, &sacloud.OpenFTPRequest{ChangePassword: true})
	if err != nil {
		return fmt.Errorf("ArchiveUpload is failed: %s", err)
	}

	// upload
	ftpsClient := ftps.NewClient(ftpServer.User, ftpServer.Password, ftpServer.HostName)
	err = ctx.ExecWithProgress(func() error {
		file, df, err := term.FileOrStdin(params.GetArchiveFile())
		if err != nil {
			return err
		}
		defer df()

		return ftpsClient.UploadFile("upload.raw", file)
	})
	if err != nil {
		return fmt.Errorf("ArchiveUpload is failed: %s", err)
	}

	// close FTP
	if err := client.CloseFTP(ctx, ctx.Zone(), archive.ID); err != nil {
		return fmt.Errorf("ArchiveUpload is failed: %s", err)
	}

	// read again
	archive, err = client.Read(ctx, ctx.Zone(), archive.ID)
	if err != nil {
		return fmt.Errorf("ArchiveUpload is failed: %s", err)
	}
	return ctx.Output().Print(archive)
}

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

package isoimage

import (
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/usacloud/pkg/term"

	"github.com/sacloud/ftps"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
	"github.com/sacloud/usacloud/pkg/progress"
)

func Upload(ctx cli.Context, params *params.UploadISOImageParam) error {
	client := sacloud.NewCDROMOp(ctx.Client())
	if _, err := client.Read(ctx, ctx.Zone(), params.Id); err != nil {
		return fmt.Errorf("Upload is failed: %s", err)
	}

	// open FTP
	ftpServer, err := client.OpenFTP(ctx, ctx.Zone(), params.Id, &sacloud.OpenFTPRequest{ChangePassword: true})
	if err != nil {
		return fmt.Errorf("ISOImageUpload is failed: %s", err)
	}

	// upload
	ftpsClient := ftps.NewClient(ftpServer.User, ftpServer.Password, ftpServer.HostName)
	err = progress.ExecWithProgress(
		fmt.Sprintf("Still uploading[ID:%d]...", params.Id),
		fmt.Sprintf("Upload iso-image[ID:%d]", params.Id),
		ctx.IO().Progress(),
		ctx.Option().NoColor,
		func(compChan chan bool, errChan chan error) {
			file, df, err := term.FileOrStdin(params.GetISOFile())
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
	if err := client.CloseFTP(ctx, ctx.Zone(), params.Id); err != nil {
		return fmt.Errorf("ISOImageUpload is failed: %s", err)
	}
	return nil
}

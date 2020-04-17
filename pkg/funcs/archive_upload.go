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

package funcs

import (
	"fmt"

	"github.com/sacloud/ftps"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/internal"
	"github.com/sacloud/usacloud/pkg/params"
)

func ArchiveUpload(ctx cli.Context, params *params.UploadArchiveParam) error {

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
	ftpsClient := ftps.NewClient(ftpServer.User, ftpServer.Password, ftpServer.HostName)

	err = internal.ExecWithProgress(
		fmt.Sprintf("Still uploading[ID:%d]...", params.Id),
		fmt.Sprintf("Upload archive[ID:%d]", params.Id),
		ctx.IO().Progress(),
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
		return fmt.Errorf("ArchiveUpload is failed: %s", err)
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

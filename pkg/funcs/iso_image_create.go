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
	"github.com/sacloud/usacloud/pkg/params"
	"github.com/sacloud/usacloud/pkg/progress"
)

func ISOImageCreate(ctx cli.Context, params *params.CreateISOImageParam) error {

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
	ftpsClient := ftps.NewClient(ftpServer.User, ftpServer.Password, ftpServer.HostName)
	err = progress.ExecWithProgress(
		fmt.Sprintf("Still uploading[ID:%d]...", res.ID),
		fmt.Sprintf("Upload iso-image[ID:%d]", res.ID),
		ctx.IO().Progress(),
		ctx.Option().NoColor,
		func(compChan chan bool, errChan chan error) {

			file, df, err := fileOrStdin(params.GetISOFile())
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

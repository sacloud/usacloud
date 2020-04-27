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

package server

import (
	"fmt"

	"github.com/sacloud/ftps"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	usacloud_params "github.com/sacloud/usacloud/pkg/params"
	"github.com/sacloud/usacloud/pkg/progress"
)

func ISOInsert(ctx cli.Context, params *usacloud_params.ISOInsertServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerISOInsert is failed: %s", e)
	}

	if p.Instance.CDROM != nil && p.Instance.CDROM.ID != sacloud.EmptyID {
		fmt.Fprint(ctx.IO().Err(), "ISOImage is already inserted to server\n")
		return nil
	}

	imageID := params.ISOImageId
	if imageID == sacloud.EmptyID {

		// Upload iso image
		api := client.GetCDROMAPI()
		iso := api.New()

		// set params
		iso.SetTags(params.Tags)
		iso.SetIconByID(params.IconId)
		iso.SetSizeGB(params.Size)
		iso.SetName(params.Name)
		iso.SetDescription(params.Description)

		// call Create(id)
		res, ftpServer, err := api.Create(iso)
		if err != nil {
			return fmt.Errorf("ServerISOInsert is failed: %s", err)
		}

		// upload
		ftpsClient := ftps.NewClient(ftpServer.User, ftpServer.Password, ftpServer.HostName)
		err = progress.ExecWithProgress(
			fmt.Sprintf("Still uploading[ID:%d]...", params.Id),
			fmt.Sprintf("Upload iso-image[ID:%d]", params.Id),
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
			return fmt.Errorf("ServerISOInsert is failed: %s", err)
		}

		// close FTP
		_, err = api.CloseFTP(res.ID)
		if err != nil {
			return fmt.Errorf("ServerISOInsert is failed: %s", err)
		}

		imageID = res.ID
	}

	// call manipurate functions
	_, err := api.InsertCDROM(params.Id, imageID)
	if err != nil {
		return fmt.Errorf("ServerISOInsert is failed: %s", err)
	}

	return nil
	// return ctx.Output().Print(res)

}

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
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	usacloud_params "github.com/sacloud/usacloud/command/params"
)

func ServerIsoInsert(ctx command.Context, params *usacloud_params.IsoInsertServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerIsoInsert is failed: %s", e)
	}

	if p.Instance.CDROM != nil && p.Instance.CDROM.ID != sacloud.EmptyID {
		fmt.Fprint(command.GlobalOption.Err, "ISOImage is already inserted to server\n")
		return nil
	}

	imageID := params.IsoImageId
	if imageID == sacloud.EmptyID {

		//validate
		isoParams := &usacloud_params.CreateISOImageParam{
			Tags:        params.Tags,
			IconId:      params.IconId,
			Size:        params.Size,
			Name:        params.Name,
			Description: params.Description,
			IsoFile:     params.IsoFile,
		}
		if errs := isoParams.Validate(); len(errs) > 0 {
			return command.FlattenErrors(errs)
		}

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
			return fmt.Errorf("ServerIsoInsert is failed: %s", err)
		}

		// upload
		ftpsClient := ftps.NewClient(ftpServer.User, ftpServer.Password, ftpServer.HostName)
		err = internal.ExecWithProgress(
			fmt.Sprintf("Still uploading[ID:%d]...", params.Id),
			fmt.Sprintf("Upload iso-image[ID:%d]", params.Id),
			command.GlobalOption.Progress,
			func(compChan chan bool, errChan chan error) {

				file, df, err := fileOrStdin(params.GetIsoFile())
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
			return fmt.Errorf("ServerIsoInsert is failed: %s", err)
		}

		// close FTP
		_, err = api.CloseFTP(res.ID)
		if err != nil {
			return fmt.Errorf("ServerIsoInsert is failed: %s", err)
		}

		imageID = res.ID
	}

	// call manipurate functions
	_, err := api.InsertCDROM(params.Id, imageID)
	if err != nil {
		return fmt.Errorf("ServerIsoInsert is failed: %s", err)
	}

	return nil
	// return ctx.GetOutput().Print(res)

}

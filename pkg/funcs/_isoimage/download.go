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
	"os"

	"github.com/sacloud/ftps"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
	"github.com/sacloud/usacloud/pkg/progress"
)

func Download(ctx cli.Context, params *params.DownloadISOImageParam) error {
	client := sacloud.NewCDROMOp(ctx.Client())
	cdrom, err := client.Read(ctx, ctx.Zone(), params.Id)
	if err != nil {
		return fmt.Errorf("ISOImageDownload is failed: %s", err)
	}

	// check scope(download can user scope only)
	if cdrom.Scope != types.Scopes.User {
		return fmt.Errorf("ISOImageDownload is failed: %s", "Only the ISO Image of scope=`user` can be downloaded")
	}

	if params.FileDestination == "" || params.FileDestination == "-" {
		if !params.Assumeyes {
			return fmt.Errorf("To output to STDOUT, specify the --assumeyes(-y) option")
		}
	}

	ftpServer, err := client.OpenFTP(ctx, ctx.Zone(), params.Id, &sacloud.OpenFTPRequest{ChangePassword: true})
	if err != nil {
		return fmt.Errorf("ISOImageDownload is failed: %s", err)
	}

	// download
	ftpsClient := ftps.NewClient(ftpServer.User, ftpServer.Password, ftpServer.HostName)

	err = progress.ExecWithProgress(
		fmt.Sprintf("Still downloading[ID:%d]...", params.Id),
		fmt.Sprintf("Download iso-image[ID:%d]", params.Id),
		ctx.IO().Progress(),
		ctx.Option().NoColor,
		func(compChan chan bool, errChan chan error) {
			path := params.FileDestination
			if path == "" || path == "-" {
				err = ftpsClient.DownloadFile(os.Stdout)
			} else {
				err = ftpsClient.Download(path)
			}
			if err != nil {
				errChan <- err
				return
			}
			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("ISOImageDownload is failed: %s", err)
	}

	// close
	if err := client.CloseFTP(ctx, ctx.Zone(), params.Id); err != nil {
		return fmt.Errorf("ISOImageDownload is failed: %s", err)
	}

	return nil
}

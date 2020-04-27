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
	"errors"
	"fmt"
	"os"

	"github.com/sacloud/ftps"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func Download(ctx cli.Context, params *params.DownloadArchiveParam) error {
	client := sacloud.NewArchiveOp(ctx.Client())
	archive, err := client.Read(ctx, ctx.Zone(), params.Id)
	if err != nil {
		return fmt.Errorf("ArchiveDownload is failed: %s", err)
	}

	// check scope(download can user scope only)
	if archive.Scope != types.Scopes.User {
		return fmt.Errorf("ArchiveDownload is failed: %s", "download not allowed")
	}

	if params.FileDestination == "" || params.FileDestination == "-" {
		if !params.Assumeyes {
			return errors.New("please specify the --assumeyes(-y) option when using STDOUT")
		}
	}

	// call manipurate functions
	res, err := client.OpenFTP(ctx, ctx.Zone(), params.Id, &sacloud.OpenFTPRequest{ChangePassword: true})
	if err != nil {
		return fmt.Errorf("ArchiveDownload is failed: %s", err)
	}

	// download
	ftpsClient := ftps.NewClient(res.User, res.Password, res.HostName)

	err = ctx.ExecWithProgress(func() error {
		path := params.FileDestination
		if path == "" || path == "-" {
			return ftpsClient.DownloadFile(os.Stdout)
		}
		return ftpsClient.Download(path)
	},
	)

	if err != nil {
		return fmt.Errorf("ArchiveDownload is failed: %s", err)
	}

	// close
	if err = client.CloseFTP(ctx, ctx.Zone(), archive.ID); err != nil {
		return fmt.Errorf("ArchiveDownload is failed: %s", err)
	}

	return nil
}

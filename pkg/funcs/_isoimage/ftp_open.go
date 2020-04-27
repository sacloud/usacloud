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
	"github.com/sacloud/libsacloud/v2/sacloud/types"

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func FTPOpen(ctx cli.Context, params *params.FTPOpenISOImageParam) error {
	client := sacloud.NewCDROMOp(ctx.Client())
	cdrom, err := client.Read(ctx, ctx.Zone(), params.Id)
	if err != nil {
		return fmt.Errorf("FTPOpen is failed: %s", err)
	}

	// check scope(download can user scope only)
	if cdrom.Scope != types.Scopes.User {
		return fmt.Errorf("ISOImageFTPOpen is failed: %s", "Only the ISO Image of scope=`user` can be downloaded")
	}

	// call manipurate functions
	ftpServer, err := client.OpenFTP(ctx, ctx.Zone(), params.Id, &sacloud.OpenFTPRequest{ChangePassword: true})
	if err != nil {
		return fmt.Errorf("ISOImageFTPOpen is failed: %s", err)
	}

	return ctx.Output().Print(ftpServer)
}

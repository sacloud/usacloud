// Copyright 2016-2020 The Libsacloud Authors
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

package cdrom

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/sacloud/ftps"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func (s *Service) Download(req *DownloadRequest) error {
	return s.DownloadWithContext(context.Background(), req)
}

func (s *Service) DownloadWithContext(ctx context.Context, req *DownloadRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	client := sacloud.NewCDROMOp(s.caller)
	resource, err := client.Read(ctx, req.Zone, req.ID)
	if err != nil {
		return fmt.Errorf("reading CDROM[%s] failed: %s", req.ID, err)
	}

	if resource.Scope != types.Scopes.User {
		return fmt.Errorf("CDROM[%s] is not allowed to download", req.ID)
	}

	ftpServer, err := client.OpenFTP(ctx, req.Zone, req.ID, &sacloud.OpenFTPRequest{ChangePassword: true})
	if err != nil {
		return fmt.Errorf("requesting FTP server information failed: %s", err)
	}

	ftpsClient := ftps.NewClient(ftpServer.User, ftpServer.Password, ftpServer.HostName)
	switch req.Path {
	case "":
		var out io.Writer = os.Stdout
		if req.Writer != nil {
			out = req.Writer
		}
		if err := ftpsClient.DownloadWriter(out); err != nil {
			return fmt.Errorf("downloading via FTP failed: %s", err)
		}
	default:
		if err := ftpsClient.Download(req.Path); err != nil {
			return fmt.Errorf("downloading via FTP failed: %s", err)
		}
	}

	// close
	if err := client.CloseFTP(ctx, req.Zone, req.ID); err != nil {
		return fmt.Errorf("closing FTP server failed: %s", err)
	}
	return nil
}

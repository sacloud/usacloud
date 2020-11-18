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

package archive

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/sacloud/ftps"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func (s *Service) Upload(req *UploadRequest) error {
	return s.UploadWithContext(context.Background(), req)
}

func (s *Service) UploadWithContext(ctx context.Context, req *UploadRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	client := sacloud.NewArchiveOp(s.caller)
	resource, err := client.Read(ctx, req.Zone, req.ID)
	if err != nil {
		return fmt.Errorf("reading Archive[%s] failed: %s", req.ID, err)
	}

	if resource.Scope != types.Scopes.User {
		return fmt.Errorf("Archive[%s] is not allowed to download", req.ID)
	}

	ftpServer, err := client.OpenFTP(ctx, req.Zone, req.ID, &sacloud.OpenFTPRequest{ChangePassword: true})
	if err != nil {
		return fmt.Errorf("requesting FTP server information failed: %s", err)
	}

	ftpsClient := ftps.NewClient(ftpServer.User, ftpServer.Password, ftpServer.HostName)
	var reader io.Reader
	switch req.Path {
	case "":
		reader = os.Stdin
		if req.Reader != nil {
			reader = req.Reader
		}
	default:
		f, err := os.Open(req.Path)
		if err != nil {
			return fmt.Errorf("opening upload file failed: %s", err)
		}
		defer f.Close()
		reader = f
	}

	if err := ftpsClient.UploadReader("upload.raw", reader); err != nil {
		return fmt.Errorf("uploading file failed: %s", err)
	}

	// close FTP
	if err := client.CloseFTP(ctx, req.Zone, resource.ID); err != nil {
		return fmt.Errorf("closing FTP server failed: %s", err)
	}
	return nil
}

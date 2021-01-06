// Copyright 2016-2021 The Libsacloud Authors
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
	"github.com/sacloud/libsacloud/v2/pkg/size"

	"github.com/sacloud/libsacloud/v2/sacloud"
)

func (s *Service) Create(req *CreateRequest) (*sacloud.CDROM, error) {
	return s.CreateWithContext(context.Background(), req)
}

func (s *Service) CreateWithContext(ctx context.Context, req *CreateRequest) (*sacloud.CDROM, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	var reader io.Reader
	switch req.SourcePath {
	case "":
		reader = req.SourceReader
	default:
		file, err := os.Open(req.SourcePath)
		if err != nil {
			return nil, fmt.Errorf("reading source file[%s] failed: %s", req.SourcePath, err)
		}
		defer file.Close() // nolint
		reader = file
	}

	client := sacloud.NewCDROMOp(s.caller)
	cdrom, ftpServer, err := client.Create(ctx, req.Zone, &sacloud.CDROMCreateRequest{
		SizeMB:      req.SizeGB * size.GiB,
		Name:        req.Name,
		Description: req.Description,
		Tags:        req.Tags,
		IconID:      req.IconID,
	})
	if err != nil {
		return nil, err
	}

	ftpsClient := ftps.NewClient(ftpServer.User, ftpServer.Password, ftpServer.HostName)
	if err := ftpsClient.UploadReader("data.iso", reader); err != nil {
		return nil, err
	}

	if err := client.CloseFTP(ctx, req.Zone, cdrom.ID); err != nil {
		return nil, err
	}

	// refresh
	return client.Read(ctx, req.Zone, cdrom.ID)
}

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

	archiveBuilder "github.com/sacloud/libsacloud/v2/helper/builder/archive"
	"github.com/sacloud/libsacloud/v2/sacloud"
)

func (s *Service) Create(req *CreateRequest) (*sacloud.Archive, error) {
	return s.CreateWithContext(context.Background(), req)
}

func (s *Service) CreateWithContext(ctx context.Context, req *CreateRequest) (*sacloud.Archive, error) {
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

	builder := (&archiveBuilder.Director{
		Name:              req.Name,
		Description:       req.Description,
		Tags:              req.Tags,
		IconID:            req.IconID,
		SizeGB:            req.SizeGB,
		SourceReader:      reader,
		SourceDiskID:      req.SourceDiskID,
		SourceArchiveID:   req.SourceArchiveID,
		SourceArchiveZone: req.SourceArchiveZone,
		SourceSharedKey:   "",
		Client:            archiveBuilder.NewAPIClient(s.caller),
	}).Builder()
	if err := builder.Validate(ctx, req.Zone); err != nil {
		return nil, err
	}
	return builder.Build(ctx, req.Zone)
}

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
	"fmt"
	"io"
	"os"

	archiveBuilder "github.com/sacloud/libsacloud/v2/helper/builder/archive"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func Create(ctx cli.Context, params *params.CreateArchiveParam) error {
	var uploadSource io.Reader
	if params.SourceArchiveId.IsEmpty() && params.SourceDiskId.IsEmpty() {
		// validate manually
		if params.Size == 0 {
			return fmt.Errorf("ArchiveCreate is required %q if when source disk/archive is missing", "--size")
		}
		errs := cli.ValidateExistsFileOrStdIn("archive-file", params.ArchiveFile)
		if len(errs) > 0 {
			return errs[0]
		}
		in, err := os.Open(params.ArchiveFile)
		if err != nil {
			return fmt.Errorf("ArchiveCreate is failed: %s", err)
		}
		defer in.Close() // nolint
		uploadSource = in
	}

	director := &archiveBuilder.Director{
		Name:            params.Name,
		Description:     params.Description,
		Tags:            params.Tags,
		IconID:          params.IconId,
		SizeGB:          params.Size,
		SourceReader:    uploadSource,
		SourceDiskID:    params.SourceDiskId,
		SourceArchiveID: params.SourceArchiveId,
		Client:          archiveBuilder.NewAPIClient(ctx.Client()),
	}
	if err := director.Builder().Validate(ctx, ctx.Zone()); err != nil {
		return fmt.Errorf("ArchiveCreate is failed: %s", err)
	}

	// create and wait
	var archive *sacloud.Archive
	err := ctx.ExecWithProgress(func() error {
		created, err := director.Builder().Build(ctx, ctx.Zone())
		if err != nil {
			return err
		}
		archive = created
		return nil
	})

	if err != nil {
		return fmt.Errorf("ArchiveCreate is failed: %s", err)
	}
	return ctx.Output().Print(archive)
}

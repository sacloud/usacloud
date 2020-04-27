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

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func WaitForCopy(ctx cli.Context, params *params.WaitForCopyArchiveParam) error {
	client := sacloud.NewArchiveOp(ctx.Client())
	archive, err := client.Read(ctx, ctx.Zone(), params.Id)
	if err != nil {
		return fmt.Errorf("ArchiveWaitForCopy is failed: %s", err)
	}

	// wait for copy with progress
	err = ctx.ExecWithProgress(func() error {
		_, err := sacloud.WaiterForReady(func() (interface{}, error) {
			return client.Read(ctx, ctx.Zone(), archive.ID)
		}).WaitForState(ctx)
		return err
	})
	if err != nil {
		return fmt.Errorf("ArchiveWaitForCopy is failed: %s", err)
	}
	return nil
}

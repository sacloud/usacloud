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

package disk

import (
	"context"

	"github.com/sacloud/libsacloud/v2/helper/wait"

	"github.com/sacloud/libsacloud/v2/sacloud"
)

func (s *Service) ReinstallFromArchive(req *ReinstallFromArchiveRequest) (*sacloud.Disk, error) {
	return s.ReinstallFromArchiveWithContext(context.Background(), req)
}

func (s *Service) ReinstallFromArchiveWithContext(ctx context.Context, req *ReinstallFromArchiveRequest) (*sacloud.Disk, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	params, err := req.toRequestParameter(ctx, s.caller)
	if err != nil {
		return nil, err
	}

	client := sacloud.NewDiskOp(s.caller)
	disk, err := client.Install(ctx, req.Zone, req.ID, params, req.DistantFrom)
	if err != nil {
		return nil, err
	}
	// wait for ready
	disk, err = wait.UntilDiskIsReady(ctx, client, req.Zone, disk.ID)
	if err != nil {
		return disk, err
	}

	// edit disk
	if req.EditParameter != nil {
		if err := client.Config(ctx, req.Zone, disk.ID, req.EditParameter); err != nil {
			return nil, err
		}
		disk, err = wait.UntilDiskIsReady(ctx, client, req.Zone, disk.ID)
		if err != nil {
			return disk, err
		}
	}

	return disk, nil
}

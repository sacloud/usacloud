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

package internet

import (
	"context"
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"
)

func (s *Service) Update(req *UpdateRequest) (*sacloud.Internet, error) {
	return s.UpdateWithContext(context.Background(), req)
}

func (s *Service) UpdateWithContext(ctx context.Context, req *UpdateRequest) (*sacloud.Internet, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	client := sacloud.NewInternetOp(s.caller)
	current, err := client.Read(ctx, req.Zone, req.ID)
	if err != nil {
		return nil, fmt.Errorf("reading Internet[%s] failed: %s", req.ID, err)
	}

	var ret *sacloud.Internet
	if req.BasicParameterChanged() {
		params, err := req.ToRequestParameter(current)
		if err != nil {
			return nil, fmt.Errorf("processing request parameter failed: %s", err)
		}

		updated, err := client.Update(ctx, req.Zone, req.ID, params)
		if err != nil {
			return nil, err
		}
		ret = updated
	}

	if req.BandWidthChanged() {
		updated, err := client.UpdateBandWidth(ctx, req.Zone, req.ID, &sacloud.InternetUpdateBandWidthRequest{BandWidthMbps: *req.BandWidthMbps})
		if err != nil {
			return nil, err
		}
		ret = updated
	}
	return ret, nil
}

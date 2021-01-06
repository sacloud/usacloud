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

package ipv6addr

import (
	"context"
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"
)

func (s *Service) Update(req *UpdateRequest) (*sacloud.IPv6Addr, error) {
	return s.UpdateWithContext(context.Background(), req)
}

func (s *Service) UpdateWithContext(ctx context.Context, req *UpdateRequest) (*sacloud.IPv6Addr, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	client := sacloud.NewIPv6AddrOp(s.caller)
	_, err := client.Read(ctx, req.Zone, req.IPv6Addr)
	if err != nil {
		return nil, fmt.Errorf("reading IPv6Addr[%s] failed: %s", req.IPv6Addr, err)
	}

	return client.Update(ctx, req.Zone, req.IPv6Addr, &sacloud.IPv6AddrUpdateRequest{HostName: req.HostName})
}

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

func (s *Service) DisableIPv6(req *DisableIPv6Request) error {
	return s.DisableIPv6WithContext(context.Background(), req)
}

func (s *Service) DisableIPv6WithContext(ctx context.Context, req *DisableIPv6Request) error {
	if err := req.Validate(); err != nil {
		return err
	}

	internetOp := sacloud.NewInternetOp(s.caller)
	current, err := internetOp.Read(ctx, req.Zone, req.ID)
	if err != nil {
		return fmt.Errorf("reading the internet resource[%s] failed: %s", req.ID, err)
	}

	if len(current.Switch.IPv6Nets) == 0 {
		return nil // noop if not exists
	}

	return internetOp.DisableIPv6(ctx, req.Zone, req.ID, current.Switch.IPv6Nets[0].ID)
}

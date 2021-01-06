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

package localrouter

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
)

func (s *Service) Health(req *HealthRequest) (*sacloud.LocalRouterHealth, error) {
	return s.HealthWithContext(context.Background(), req)
}

func (s *Service) HealthWithContext(ctx context.Context, req *HealthRequest) (*sacloud.LocalRouterHealth, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	client := sacloud.NewLocalRouterOp(s.caller)
	return client.HealthStatus(ctx, req.ID)
}

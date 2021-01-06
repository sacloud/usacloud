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

package esme

import (
	"context"
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"
)

func (s *Service) Logs(req *LogsRequest) ([]*sacloud.ESMELogs, error) {
	return s.LogsWithContext(context.Background(), req)
}

func (s *Service) LogsWithContext(ctx context.Context, req *LogsRequest) ([]*sacloud.ESMELogs, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	client := sacloud.NewESMEOp(s.caller)
	_, err := client.Read(ctx, req.ID)
	if err != nil {
		return nil, fmt.Errorf("reading ESME[%s] failed: %s", req.ID, err)
	}

	return client.Logs(ctx, req.ID)
}

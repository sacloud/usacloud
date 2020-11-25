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

package ipaddress

import (
	"context"
	"fmt"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
)

const (
	defaultRetryMax      = 30
	defaultRetryInterval = 10
)

func (s *Service) UpdateHostName(req *UpdateHostNameRequest) error {
	return s.UpdateHostNameWithContext(context.Background(), req)
}

func (s *Service) UpdateHostNameWithContext(ctx context.Context, req *UpdateHostNameRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	client := sacloud.NewIPAddressOp(s.caller)
	_, err := client.Read(ctx, req.Zone, req.IPAddress)
	if err != nil {
		return err
	}

	if req.RetryMax == 0 {
		req.RetryMax = defaultRetryMax
	}
	if req.RetryInterval == 0 {
		req.RetryInterval = defaultRetryInterval
	}

	i := 0
	success := false
	for i < req.RetryMax {
		if err := ctx.Err(); err != nil {
			return fmt.Errorf("updating the HostName for %s failed: %s", req.IPAddress, err)
		}
		if _, err = client.UpdateHostName(ctx, req.Zone, req.IPAddress, req.HostName); err == nil {
			success = true
			break
		}
		time.Sleep(time.Duration(req.RetryInterval) * time.Second)
		i++
	}

	if !success {
		return fmt.Errorf("updating the HostName for %s failed: giving up on setting the HostName: max attempts exceeded", req.IPAddress)
	}
	return nil
}

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

package server

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
)

func (s *Service) EjectCDROM(req *EjectCDROMRequest) error {
	return s.EjectCDROMWithContext(context.Background(), req)
}

func (s *Service) EjectCDROMWithContext(ctx context.Context, req *EjectCDROMRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	client := sacloud.NewServerOp(s.caller)
	server, err := client.Read(ctx, req.Zone, req.ID)
	if err != nil {
		return err
	}
	if server.CDROMID.IsEmpty() {
		return nil // noop
	}

	return client.EjectCDROM(ctx, req.Zone, req.ID, &sacloud.EjectCDROMRequest{ID: server.CDROMID})
}

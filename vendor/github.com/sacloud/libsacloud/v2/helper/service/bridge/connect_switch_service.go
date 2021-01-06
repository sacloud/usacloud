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

package bridge

import (
	"context"
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"
)

func (s *Service) ConnectSwitch(req *ConnectSwitchRequest) error {
	return s.ConnectSwitchWithContext(context.Background(), req)
}

func (s *Service) ConnectSwitchWithContext(ctx context.Context, req *ConnectSwitchRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	bridgeOp := sacloud.NewBridgeOp(s.caller)
	switchOp := sacloud.NewSwitchOp(s.caller)

	bridge, err := bridgeOp.Read(ctx, req.Zone, req.ID)
	if err != nil {
		return err
	}
	if bridge.SwitchInZone != nil {
		return fmt.Errorf("target bridge[%s] is already connected to the switch[%s]", req.ID, bridge.SwitchInZone.ID)
	}

	return switchOp.ConnectToBridge(ctx, req.Zone, req.SwitchID, req.ID)
}

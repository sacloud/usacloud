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

package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func SwitchBridgeDisconnect(ctx command.Context, params *params.BridgeDisconnectSwitchParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSwitchAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("SwitchBridgeDisconnect is failed: %s", e)
	}

	if p.Bridge == nil {
		return fmt.Errorf("SwitchBridgeDisconnect is failed: Bridge is already disconnected")
	}

	// call manipurate functions
	_, err := api.DisconnectFromBridge(params.Id)
	if err != nil {
		return fmt.Errorf("SwitchBridgeDisconnect is failed: %s", err)
	}

	return nil

}

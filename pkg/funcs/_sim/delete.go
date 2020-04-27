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

package sim

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func Delete(ctx cli.Context, params *params.DeleteSIMParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSIMAPI()

	// call
	res, err := api.Reset().Include("*").Include("Status.sim").Find()
	if err != nil {
		return fmt.Errorf("SIMDelete is failed: %s", err)
	}
	var sim *sacloud.SIM
	for _, s := range res.CommonServiceSIMItems {
		if s.ID == params.Id {
			sim = &s
			break
		}
	}

	if sim == nil {
		return fmt.Errorf("SIMDelete is failed: SIM[%d] is not found", params.Id)
	}

	if params.Force && sim.Status.SIMInfo.Activated {
		_, err := api.Deactivate(params.Id)
		if err != nil {
			return fmt.Errorf("SIMDelete is failed: %s", err)
		}
	}

	// call Delete(id)
	if _, err := api.Delete(params.Id); err != nil {
		return fmt.Errorf("SIMDelete is failed: %s", err)
	}
	return nil
}

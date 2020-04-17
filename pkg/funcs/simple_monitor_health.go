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

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func SimpleMonitorHealth(ctx cli.Context, params *params.HealthSimpleMonitorParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSimpleMonitorAPI()

	res, err := api.Health(params.Id)
	if err != nil {
		return fmt.Errorf("SimpleMonitorHealth is failed: %s", err)
	}

	output := struct {
		ID sacloud.ID
		*sacloud.SimpleMonitorHealthCheckStatus
	}{
		ID:                             params.Id,
		SimpleMonitorHealthCheckStatus: res,
	}

	return ctx.GetOutput().Print(output)

}

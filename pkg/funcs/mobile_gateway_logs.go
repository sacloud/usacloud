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
	"encoding/json"
	"fmt"
	"time"

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/internal"
	"github.com/sacloud/usacloud/pkg/params"
)

func MobileGatewayLogs(ctx cli.Context, params *params.LogsMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()

	// call Read(id)
	_, err := api.Read(params.Id)
	if err != nil {
		return fmt.Errorf("MobileGatewayLogs is failed: %s", err)
	}

	logBuf := internal.NewHashQueue(500)
	out := ctx.IO().Out()

	if params.Follow {
		for {
			// call Read(id)
			logs, err := api.Logs(params.Id, nil)
			if err != nil {
				return fmt.Errorf("MobileGatewayLogs is failed: %s", err)
			}

			for _, log := range logs {
				data, err := json.MarshalIndent(log, "", "\t")
				if err != nil {
					return fmt.Errorf("MobileGatewayLogs is failed: %s", err)
				}
				line := string(data)
				if logBuf.PutIfAbsent(line) {
					fmt.Fprintf(out, "%s\n", line)
				}
			}
			time.Sleep(time.Duration(params.RefreshInterval) * time.Second)
		}
	} else {
		// call Read(id)
		logs, err := api.Logs(params.Id, nil)
		if err != nil {
			return fmt.Errorf("MobileGatewayLogs is failed: %s", err)
		}

		if len(logs) == 0 {
			fmt.Fprintf(ctx.IO().Err(), "Result is empty\n")
			return nil
		}

		data, err := json.MarshalIndent(logs, "", "\t")
		if err != nil {
			return fmt.Errorf("MobileGatewayLogs is failed: %s", err)
		}
		fmt.Fprintf(out, "%s\n", string(data))
	}

	return nil

}

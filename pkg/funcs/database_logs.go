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
	"time"

	"github.com/fatih/color"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/helper/printer"
	"github.com/sacloud/usacloud/pkg/internal"
	"github.com/sacloud/usacloud/pkg/params"
)

func DatabaseLogs(ctx cli.Context, params *params.LogsDatabaseParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDatabaseAPI()

	// validate params
	if params.Follow {
		if params.LogName == "all" {
			fmt.Fprintf(ctx.IO().Err(), "[WARN] -f/--follow option can not use with --log-name=%q, ignored", "all")
			params.Follow = false
		}
		if params.ListLogNames {
			params.Follow = false
		}
	}
	if ctx.IsSet("refresh-interval") && params.LogName == "all" {
		fmt.Fprintf(ctx.IO().Err(), "[WARN] --refresh-interval option can not use with --log-name=%q, ignored", "all")
	}

	logBuf := internal.NewHashQueue(500)
	out := ctx.IO().Out()
	printer := printer.Printer{NoColor: ctx.Option().NoColor}

	for {
		// call Read(id)
		res, err := api.Status(params.Id)
		if err != nil {
			return fmt.Errorf("DatabaseLogs is failed: %s", err)
		}

		if !res.IsUp() {
			fmt.Fprintf(ctx.IO().Progress(), "Database[%d] is not running\n", params.Id)
		} else {

			if params.ListLogNames {
				fmt.Fprintln(out, "all")
				for _, name := range res.DBConf.Log {
					fmt.Fprintln(out, name.Name)
				}
				return nil
			}

			logs := map[string][]string{}
			for _, l := range res.DBConf.Log {
				if params.LogName == "all" || params.LogName == l.Name {
					logs[l.Name] = l.Logs()
				}
			}
			if len(logs) == 0 {
				return fmt.Errorf("log-name[%q] is not found", params.LogName)
			}

			for key, lines := range logs {
				if params.LogName == "all" {
					printer.Fprintf(out, color.New(color.FgHiGreen), "\n==> [%s]:start\n", key)
				}
				for _, line := range lines {
					if logBuf.PutIfAbsent(line) {
						fmt.Fprintf(out, "%s\n", line)
					}
				}
				if params.LogName == "all" {
					printer.Fprintf(out, color.New(color.FgHiGreen), "\n<== [%s]:end\n", key)
				}
			}
		}
		if !params.Follow {
			break
		}
		time.Sleep(time.Duration(params.RefreshInterval) * time.Second)
	}
	return nil

}

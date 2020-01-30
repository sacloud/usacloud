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
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ServerMaintenanceInfo(ctx command.Context, params *params.MaintenanceInfoServerParam) error {

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()

	type mainteInfo struct {
		*sacloud.Server
		MaintenanceInfo *sacloud.NewsFeed
		StartDate       string
		EndDate         string
		InfoURL         string
	}

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return fmt.Errorf("ServerMaintenanceInfo is failed: %s", err)
	}

	list := []interface{}{}
	timeLayout := "01/02 15:04"
	for i, s := range res.Servers {
		if s.MaintenanceScheduled() {

			info, err := client.NewsFeed.GetFeedByURL(s.GetMaintenanceInfoURL())
			if err != nil {
				return fmt.Errorf("GetFeedByURL(%s) is failed: %s", s.GetMaintenanceInfoURL(), err)
			}
			if info == nil {
				continue
			}

			v := &mainteInfo{
				Server:          &res.Servers[i],
				MaintenanceInfo: info,
				StartDate:       info.EventStart().Format(timeLayout),
				EndDate:         info.EventEnd().Format(timeLayout),
				InfoURL:         s.GetMaintenanceInfoURL(),
			}

			list = append(list, v)
		}
	}
	return ctx.GetOutput().Print(list...)

}

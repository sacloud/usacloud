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
	"bytes"
	"fmt"
	"sort"
	"text/template"
	"time"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ServerMonitorNic(ctx command.Context, params *params.MonitorNicServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()

	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerMonitorNic is failed: %s", e)
	}

	end := parseDateTimeString(params.End)
	start := end.Add(-1 * time.Hour)
	if params.Start != "" {
		start = parseDateTimeString(params.Start)
	}

	// validate start <= end
	if !(start.Unix() <= end.Unix()) {
		return fmt.Errorf("Invalid Parameter : start(%s) or end(%s) is invalid", start, end)
	}

	// validate nic indexes
	for _, v := range params.Index {
		if v > int64(len(p.Interfaces))-1 {
			return fmt.Errorf("Invalid NIC index: NIC(#%d) is not exists", v)
		}
	}
	if len(params.Index) == 0 {
		for i := range p.Interfaces {
			params.Index = append(params.Index, int64(i))
		}
	}

	req := sacloud.NewResourceMonitorRequest(&start, &end)

	list := MonitorValues{}
	for _, nicIndex := range params.Index {

		res, err := client.Interface.Monitor(p.Interfaces[nicIndex].ID, req)
		if err != nil {
			return fmt.Errorf("ServerMonitorNic is failed: %s", err)
		}

		// collect values
		receiveValues, err := res.FlattenPacketReceiveValue()
		if err != nil {
			return fmt.Errorf("ServerMonitorNic is failed: %s", err)
		}
		sendValues, err := res.FlattenPacketSendValue()
		if err != nil {
			return fmt.Errorf("ServerMonitorNic is failed: %s", err)
		}

		// sort
		sort.Slice(receiveValues, func(i, j int) bool { return receiveValues[i].Time.Before(receiveValues[j].Time) })
		sort.Slice(sendValues, func(i, j int) bool { return sendValues[i].Time.Before(sendValues[j].Time) })

		// build key string
		key := ""
		buf := bytes.NewBufferString("")
		t := template.New("t")
		template.Must(t.Parse(params.KeyFormat))
		err = t.Execute(buf, map[string]interface{}{
			"ID":    params.Id,
			"Index": nicIndex,
		})
		if err != nil {
			return fmt.Errorf("ServerMonitorNic is failed: %s", err)
		}
		key = buf.String()

		// build sortable struct
		for i := range receiveValues {
			list = append(list, MonitorValue{
				"Index":     fmt.Sprintf("%d", nicIndex),
				"Key":       key,
				"NicID":     fmt.Sprintf("%d", p.Interfaces[nicIndex].ID),
				"TimeStamp": fmt.Sprintf("%s", receiveValues[i].Time),
				"UnixTime":  fmt.Sprintf("%d", receiveValues[i].Time.Unix()),
				"Send":      fmt.Sprintf("%f", receiveValues[i].Value),
				"Receive":   fmt.Sprintf("%f", sendValues[i].Value),
			})
		}

	}
	sort.Sort(list)

	values := []interface{}{}
	for _, v := range list {
		values = append(values, v)
	}

	return ctx.GetOutput().Print(values...)

}

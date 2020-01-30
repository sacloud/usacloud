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

func DatabaseMonitorNic(ctx command.Context, params *params.MonitorNicDatabaseParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDatabaseAPI()

	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DatabaseMonitorNic is failed: %s", e)
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

	req := sacloud.NewResourceMonitorRequest(&start, &end)

	list := MonitorValues{}

	res, err := api.MonitorInterface(params.Id, req)
	if err != nil {
		return fmt.Errorf("DatabaseMonitorNic is failed: %s", err)
	}

	// collect values
	receiveValues, err := res.FlattenPacketReceiveValue()
	if err != nil {
		return fmt.Errorf("DatabaseMonitorNic is failed: %s", err)
	}
	sendValues, err := res.FlattenPacketSendValue()
	if err != nil {
		return fmt.Errorf("DatabaseMonitorNic is failed: %s", err)
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
		"ID": params.Id,
	})
	if err != nil {
		return fmt.Errorf("DatabaseMonitorNic is failed: %s", err)
	}
	key = buf.String()

	// build sortable struct
	for i := range receiveValues {
		list = append(list, MonitorValue{
			"Key":       key,
			"TimeStamp": fmt.Sprintf("%s", receiveValues[i].Time),
			"UnixTime":  fmt.Sprintf("%d", receiveValues[i].Time.Unix()),
			"Send":      fmt.Sprintf("%f", receiveValues[i].Value),
			"Receive":   fmt.Sprintf("%f", sendValues[i].Value),
		})
	}
	sort.Sort(list)

	values := []interface{}{}
	for _, v := range list {
		values = append(values, v)
	}

	return ctx.GetOutput().Print(values...)

}

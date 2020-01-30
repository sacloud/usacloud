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

func DatabaseMonitorBackupDisk(ctx command.Context, params *params.MonitorBackupDiskDatabaseParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDatabaseAPI()

	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DatabaseMonitorSystemDisk is failed: %s", e)
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
	res, err := api.MonitorBackupDisk(params.Id, req)
	if err != nil {
		return fmt.Errorf("DatabaseMonitorSystemDisk is failed: %s", err)
	}

	// collect values
	readValues, err := res.FlattenDiskReadValue()
	if err != nil {
		return fmt.Errorf("DatabaseMonitorSystemDisk is failed: %s", err)
	}
	writeValues, err := res.FlattenDiskWriteValue()
	if err != nil {
		return fmt.Errorf("DatabaseMonitorSystemDisk is failed: %s", err)
	}

	// sort
	sort.Slice(readValues, func(i, j int) bool { return readValues[i].Time.Before(readValues[j].Time) })
	sort.Slice(writeValues, func(i, j int) bool { return writeValues[i].Time.Before(writeValues[j].Time) })

	// build key string
	key := ""
	buf := bytes.NewBufferString("")
	t := template.New("t")
	template.Must(t.Parse(params.KeyFormat))
	err = t.Execute(buf, map[string]interface{}{
		"ID": params.Id,
	})
	if err != nil {
		return fmt.Errorf("DatabaseMonitorSystemDisk is failed: %s", err)
	}
	key = buf.String()

	// build sortable struct
	for i := range readValues {
		list = append(list, MonitorValue{
			"Key":       key,
			"TimeStamp": fmt.Sprintf("%s", readValues[i].Time),
			"UnixTime":  fmt.Sprintf("%d", readValues[i].Time.Unix()),
			"Read":      fmt.Sprintf("%f", readValues[i].Value),
			"Write":     fmt.Sprintf("%f", writeValues[i].Value),
		})
	}
	sort.Sort(list)

	values := []interface{}{}
	for _, v := range list {
		values = append(values, v)
	}

	return ctx.GetOutput().Print(values...)

}

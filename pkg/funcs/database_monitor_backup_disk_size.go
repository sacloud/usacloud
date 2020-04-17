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
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func DatabaseMonitorBackupDiskSize(ctx cli.Context, params *params.MonitorBackupDiskSizeDatabaseParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDatabaseAPI()

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

	res, err := api.MonitorDatabase(params.Id, req)
	if err != nil {
		return fmt.Errorf("DatabaseMonitorBackupDiskSize is failed: %s", err)
	}

	// collect values
	totalValues, err := res.FlattenTotalDisk2SizeValue()
	if err != nil {
		return fmt.Errorf("DatabaseMonitorBackupDiskSize is failed: %s", err)
	}
	usedValues, err := res.FlattenUsedDisk2SizeValue()
	if err != nil {
		return fmt.Errorf("DatabaseMonitorBackupDiskSize is failed: %s", err)
	}

	// sort
	sort.Slice(totalValues, func(i, j int) bool { return totalValues[i].Time.Before(totalValues[j].Time) })
	sort.Slice(usedValues, func(i, j int) bool { return usedValues[i].Time.Before(usedValues[j].Time) })

	// build key string
	key := ""
	buf := bytes.NewBufferString("")
	t := template.New("t")
	template.Must(t.Parse(params.KeyFormat))
	err = t.Execute(buf, map[string]interface{}{
		"ID": params.Id,
	})
	if err != nil {
		return fmt.Errorf("DatabaseMonitorBackupDiskSize is failed: %s", err)
	}
	key = buf.String()

	// build sortable struct
	list := MonitorValues{}
	for i := range totalValues {
		list = append(list, MonitorValue{
			"Key":       key,
			"TimeStamp": fmt.Sprintf("%s", totalValues[i].Time),
			"UnixTime":  fmt.Sprintf("%d", totalValues[i].Time.Unix()),
			"Used":      fmt.Sprintf("%f", usedValues[i].Value),
			"Total":     fmt.Sprintf("%f", totalValues[i].Value),
		})
	}
	sort.Sort(list)

	values := []interface{}{}
	for _, v := range list {
		values = append(values, v)
	}

	return ctx.GetOutput().Print(values...)

}

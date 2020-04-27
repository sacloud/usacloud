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

package internet

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

func Monitor(ctx cli.Context, params *params.MonitorInternetParam) error {

	client := ctx.GetAPIClient()
	api := client.GetInternetAPI()

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

	res, err := api.Monitor(params.Id, req)
	if err != nil {
		return fmt.Errorf("InternetMonitor is failed: %s", err)
	}

	// collect values
	inValues, err := res.FlattenInternetInValue()
	if err != nil {
		return fmt.Errorf("InternetMonitor is failed: %s", err)
	}
	outValues, err := res.FlattenInternetOutValue()
	if err != nil {
		return fmt.Errorf("InternetMonitor is failed: %s", err)
	}

	// sort
	sort.Slice(inValues, func(i, j int) bool { return inValues[i].Time.Before(inValues[j].Time) })
	sort.Slice(outValues, func(i, j int) bool { return outValues[i].Time.Before(outValues[j].Time) })

	// build key string
	key := ""
	buf := bytes.NewBufferString("")
	t := template.New("t")
	template.Must(t.Parse(params.KeyFormat))
	err = t.Execute(buf, map[string]interface{}{
		"ID": params.Id,
	})
	if err != nil {
		return fmt.Errorf("DiskMonitor is failed: %s", err)
	}
	key = buf.String()

	// build sortable struct
	list := MonitorValues{}
	for i := range inValues {
		list = append(list, MonitorValue{
			"Key":       key,
			"TimeStamp": fmt.Sprintf("%s", inValues[i].Time),
			"UnixTime":  fmt.Sprintf("%d", inValues[i].Time.Unix()),
			"In":        fmt.Sprintf("%f", inValues[i].Value),
			"Out":       fmt.Sprintf("%f", outValues[i].Value),
		})
	}
	sort.Sort(list)

	values := []interface{}{}
	for _, v := range list {
		values = append(values, v)
	}

	return ctx.Output().Print(values...)
}

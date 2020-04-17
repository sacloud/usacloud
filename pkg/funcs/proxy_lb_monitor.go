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

func ProxyLBMonitor(ctx cli.Context, params *params.MonitorProxyLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()

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
		return fmt.Errorf("ProxyLB is failed: %s", err)
	}

	// collect values
	activeConnValues, err := res.FlattenActiveConnections()
	if err != nil {
		return fmt.Errorf("ProxyLB is failed: %s", err)
	}
	connPerSecValues, err := res.FlattenConnectionsPerSec()
	if err != nil {
		return fmt.Errorf("ProxyLB is failed: %s", err)
	}

	// sort
	sort.Slice(activeConnValues, func(i, j int) bool { return activeConnValues[i].Time.Before(activeConnValues[j].Time) })
	sort.Slice(connPerSecValues, func(i, j int) bool { return connPerSecValues[i].Time.Before(connPerSecValues[j].Time) })

	// build key string
	key := ""
	buf := bytes.NewBufferString("")
	t := template.New("t")
	template.Must(t.Parse(params.KeyFormat))
	err = t.Execute(buf, map[string]interface{}{
		"ID": params.Id,
	})
	if err != nil {
		return fmt.Errorf("LoadBalancerMonitor is failed: %s", err)
	}
	key = buf.String()

	// build sortable struct
	list := MonitorValues{}
	for i := range activeConnValues {
		list = append(list, MonitorValue{
			"Key":       key,
			"TimeStamp": fmt.Sprintf("%s", activeConnValues[i].Time),
			"UnixTime":  fmt.Sprintf("%d", activeConnValues[i].Time.Unix()),
			"Receive":   fmt.Sprintf("%f", activeConnValues[i].Value),
			"Send":      fmt.Sprintf("%f", connPerSecValues[i].Value),
		})
	}
	sort.Sort(list)

	var values []interface{}
	for _, v := range list {
		values = append(values, v)
	}
	return ctx.GetOutput().Print(values...)
}

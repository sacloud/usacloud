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
	"sync"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func SimpleMonitorList(ctx command.Context, params *params.ListSimpleMonitorParam) error {

	client := ctx.GetAPIClient()
	finder := client.GetSimpleMonitorAPI()

	finder.SetEmpty()

	if !command.IsEmpty(params.Name) {
		for _, v := range params.Name {
			finder.SetFilterBy("Name", v)
		}
	}
	if !command.IsEmpty(params.Id) {
		for _, v := range params.Id {
			finder.SetFilterMultiBy("ID", v)
		}
	}
	if !command.IsEmpty(params.From) {
		finder.SetOffset(params.From)
	}
	if !command.IsEmpty(params.Max) {
		finder.SetLimit(params.Max)
	}
	if !command.IsEmpty(params.Sort) {
		for _, v := range params.Sort {
			setSortBy(finder, v)
		}
	}

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return fmt.Errorf("SimpleMonitorList is failed: %s", err)
	}

	var results []*simpleMonitorFindResult
	for i := range res.SimpleMonitors {
		results = append(results, &simpleMonitorFindResult{
			SimpleMonitor: &res.SimpleMonitors[i],
		})
	}
	if len(results) > 0 {
		var wg sync.WaitGroup
		wg.Add(len(results))
		for i := range res.SimpleMonitors {
			go func(s *simpleMonitorFindResult) {
				health, _ := finder.Health(s.SimpleMonitor.ID)
				s.HealthCheck = health
				wg.Done()
			}(results[i])
		}
		wg.Wait()
	}

	var list []interface{}
	for i := range results {

		if !params.GetCommandDef().Params["tags"].FilterFunc(list, results[i], params.Tags) {
			continue
		}

		if !params.GetCommandDef().Params["health"].FilterFunc(list, results[i], params.Health) {
			continue
		}

		list = append(list, results[i])
	}
	return ctx.GetOutput().Print(list...)

}

type simpleMonitorFindResult struct {
	SimpleMonitor *sacloud.SimpleMonitor
	HealthCheck   *sacloud.SimpleMonitorHealthCheckStatus
}

func (s *simpleMonitorFindResult) HasTag(target string) bool {
	return s.SimpleMonitor.HasTag(target)
}

func (s *simpleMonitorFindResult) HealthCheckResult() *sacloud.SimpleMonitorHealthCheckStatus {
	return s.HealthCheck
}

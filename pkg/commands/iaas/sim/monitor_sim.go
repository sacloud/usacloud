// Copyright 2017-2022 The sacloud/usacloud Authors
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

package sim

import (
	"github.com/sacloud/usacloud/pkg/ccol"
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/output"
)

var monitorSIMCommand = &core.Command{
	Name:       "monitor-sim",
	Aliases:    []string{"monitor"},
	Category:   "monitor",
	Order:      10,
	NoProgress: true,

	ColumnDefs: []output.ColumnDef{
		ccol.ID,
		{Name: "Time"},
		{Name: "UplinkBPS"},
		{Name: "DownlinkBPS"},
	},

	SelectorType: core.SelectorTypeRequireMulti,

	ParameterInitializer: func() interface{} {
		return newMonitorSIMParameter()
	},
}

type monitorSIMParameter struct {
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.MonitorParameter `cli:",squash" mapconv:",squash"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`
}

func newMonitorSIMParameter() *monitorSIMParameter {
	return &monitorSIMParameter{}
}

func init() {
	Resource.AddCommand(monitorSIMCommand)
}

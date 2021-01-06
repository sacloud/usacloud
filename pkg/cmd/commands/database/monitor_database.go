// Copyright 2017-2021 The Usacloud Authors
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

package database

import (
	"github.com/sacloud/usacloud/pkg/cmd/ccol"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/output"
)

var monitorDatabaseCommand = &core.Command{
	Name:       "monitor-database",
	Category:   "monitor",
	Order:      20,
	NoProgress: true,

	ColumnDefs: []output.ColumnDef{
		ccol.Zone,
		ccol.ID,
		{Name: "Time"},
		{Name: "DatabaseTime"},
		{
			Name:     "TotalMemory",
			Template: "{{ .TotalMemorySize }}"},
		{
			Name:     "UsedMemory",
			Template: "{{ .UsedMemorySize }}"},
		{
			Name:     "TotalDisk[System]",
			Template: "{{ .TotalDisk1Size }}"},
		{
			Name:     "UsedDisk[System]",
			Template: "{{ .UsedDisk1Size }}"},
		{
			Name:     "TotalDisk[Backup]",
			Template: "{{ .TotalDisk2Size }}"},
		{
			Name:     "UsedDisk[Backup]",
			Template: "{{ .UsedDisk2Size }}"},
		{
			Name:     "Binlog",
			Template: "{{ .BinlogUsedSizeKiB }}"},
		{
			Name:     "DelayTime(sec)",
			Template: "{{ .DelayTimeSec }}"},
	},

	SelectorType: core.SelectorTypeRequireMulti,

	ParameterInitializer: func() interface{} {
		return newMonitorDatabaseParameter()
	},
}

type monitorDatabaseParameter struct {
	cflag.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.MonitorParameter `cli:",squash" mapconv:",squash"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`
}

func newMonitorDatabaseParameter() *monitorDatabaseParameter {
	return &monitorDatabaseParameter{}
}

func init() {
	Resource.AddCommand(monitorDatabaseCommand)
}

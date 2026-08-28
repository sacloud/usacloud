// Copyright 2017-2025 The sacloud/usacloud Authors
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

package schedule

import (
	"github.com/sacloud/usacloud/pkg/ccol"
	"github.com/sacloud/usacloud/pkg/output"
)

var defaultColumnDefs = []output.ColumnDef{
	ccol.ID,
	ccol.Name,
	ccol.Tags,
	{Name: "ProcessConfigurationID", Template: `{{ .Settings.ScheduleSettings.ProcessConfigurationID }}`},
	{Name: "RecurringStep", Template: `{{ .Settings.ScheduleSettings.RecurringStep.Value }}`},
	{Name: "RecurringUnit", Template: `{{ .Settings.ScheduleSettings.RecurringUnit.Value }}`},
	{Name: "Crontab", Template: `{{ .Settings.ScheduleSettings.Crontab.Value }}`},
	{Name: "StartsAt", Template: `{{ if .Settings.ScheduleSettings.StartsAt.IsString }}{{ .Settings.ScheduleSettings.StartsAt.String }}{{ else }}{{ .Settings.ScheduleSettings.StartsAt.Int64 }}{{ end }}`},
}

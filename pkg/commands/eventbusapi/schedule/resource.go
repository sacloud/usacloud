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

import "github.com/sacloud/usacloud/pkg/core"

var Resource = &core.Resource{
	Name:  "schedule",
	Usage: "Run a process configuration at specified times",
	LongUsage: `A schedule configures when EventBus calls an existing process configuration.
It supports an execution interval or a five-field crontab expression. An interval specifies the
repeat interval, its unit (day, hour, or minute), and the execution start time.

Create the process configuration first. The minimum interval is one minute. Job execution is best
effort and does not guarantee strict real-time execution.

Documentation: https://manual.sakura.ad.jp/cloud/appliance/eventbus/control_panel.html`,
	Example: `  usacloud eventbus-api schedule create --help
  usacloud eventbus-api schedule list`,
}

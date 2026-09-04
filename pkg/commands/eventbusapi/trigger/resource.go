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

package trigger

import "github.com/sacloud/usacloud/pkg/core"

var Resource = &core.Resource{
	Name:  "trigger",
	Usage: "Run a process configuration when an event is detected",
	LongUsage: `A trigger detects an event and runs a job by calling an existing process configuration.
It configures an event source, event types, and conditions. Supported sources include Monitoring
Suite alerts and the event log.

Create the process configuration first. Documented use cases include adding a server to a load
balancer when a server-start event is detected and running AutoScale when an alert is detected.

Documentation:
  https://manual.sakura.ad.jp/cloud/appliance/eventbus/control_panel.html
  https://manual.sakura.ad.jp/cloud/appliance/eventbus/events.html`,
	Example: `  usacloud eventbus-api trigger create --help
  usacloud eventbus-api trigger list`,
}

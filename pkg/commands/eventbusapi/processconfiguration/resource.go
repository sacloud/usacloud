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

package processconfiguration

import "github.com/sacloud/usacloud/pkg/core"

var Resource = &core.Resource{
	Name:  "process-configuration",
	Usage: "Define the job that EventBus runs",
	LongUsage: `A process configuration configures which service EventBus calls as a job.
The supported destination services are Simple Notification, SimpleMQ, and AutoScale. Create the
destination service before registering it in a process configuration.

Create a process configuration before creating a schedule or trigger. Secret settings for an
existing process configuration can be changed with update-secret.

Documentation: https://manual.sakura.ad.jp/cloud/appliance/eventbus/control_panel.html`,
	Example: `  usacloud eventbus-api process-configuration create --help
  usacloud eventbus-api process-configuration list`,
}

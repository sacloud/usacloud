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

package eventbusapi

import (
	"github.com/sacloud/usacloud/pkg/commands/eventbusapi/processconfiguration"
	"github.com/sacloud/usacloud/pkg/commands/eventbusapi/schedule"
	"github.com/sacloud/usacloud/pkg/commands/eventbusapi/trigger"
	"github.com/sacloud/usacloud/pkg/commands/iaas/category"
	"github.com/sacloud/usacloud/pkg/core"
)

var Resource = &core.Resource{
	Name:  "eventbus-api",
	Usage: "Run jobs on schedules or in response to events",
	LongUsage: `EventBus is a managed service that combines event detection and job scheduling.
It can run jobs from schedule triggers or when a change in an event source is detected. Supported
jobs call Simple Notification, SimpleMQ, or AutoScale.

Create a process configuration before creating a schedule or trigger. These resources are global and
shared by all SAKURA Cloud network zones. Job execution is best effort and does not guarantee strict
real-time execution.

Documentation: https://manual.sakura.ad.jp/cloud/appliance/eventbus/about.html`,
	Example: `  usacloud eventbus-api process-configuration create --help
  usacloud eventbus-api schedule create --help
  usacloud eventbus-api trigger create --help`,
	Experimental:     true,
	IsGlobalResource: true,
	Category:         category.ResourceCategoryAPI,
}

func init() {
	Resource.AddChild(schedule.Resource)
	Resource.AddChild(trigger.Resource)
	Resource.AddChild(processconfiguration.Resource)
}

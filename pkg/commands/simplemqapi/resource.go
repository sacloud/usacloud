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

package simplemqapi

import (
	"github.com/sacloud/usacloud/pkg/commands/iaas/category"
	"github.com/sacloud/usacloud/pkg/commands/simplemqapi/message"
	"github.com/sacloud/usacloud/pkg/commands/simplemqapi/queue"
	"github.com/sacloud/usacloud/pkg/core"
)

var Resource = &core.Resource{
	Name:  "simplemq-api",
	Usage: "Exchange messages asynchronously between software components",
	LongUsage: `SimpleMQ is a managed message queue for asynchronous communication between software
components. It uses pull-based, at-least-once delivery, so a message can be delivered more than once.

Create a queue before using message commands. Each queue has its own API key, which message commands
read from a file or standard input. Queues are global and shared by all SAKURA Cloud network zones.

Documentation: https://manual.sakura.ad.jp/cloud/appliance/simplemq/about.html`,
	Example: `  usacloud simplemq-api queue create --help
  usacloud simplemq-api message send --help`,
	Experimental:     true,
	IsGlobalResource: true,
	Category:         category.ResourceCategoryAPI,
}

func init() {
	Resource.AddChild(queue.Resource)
	Resource.AddChild(message.Resource)
}

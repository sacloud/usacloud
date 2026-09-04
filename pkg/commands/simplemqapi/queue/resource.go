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

package queue

import "github.com/sacloud/usacloud/pkg/core"

var Resource = &core.Resource{
	Name:  "queue",
	Usage: "Manage queues that store SimpleMQ messages",
	LongUsage: `A queue stores messages exchanged through SimpleMQ. Create a queue before using the
message resource. Queue names are required, immutable, and unique within a project.

Each queue has its own Message API key. The key is shown once when issued, and rotating it invalidates
the old key. Queue management requires a cloud API key with create and delete access.

Documentation: https://manual.sakura.ad.jp/cloud/appliance/simplemq/control_panel.html`,
	Example: `  usacloud simplemq-api queue create --help
  usacloud simplemq-api queue list`,
}

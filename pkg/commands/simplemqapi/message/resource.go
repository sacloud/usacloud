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

package message

import "github.com/sacloud/usacloud/pkg/core"

var Resource = &core.Resource{
	Name:  "message",
	Usage: "Send, receive, extend, and delete queue messages",
	LongUsage: `Message commands transfer messages through an existing queue using that queue's API key.
Create the queue first. Provide its name and read the API key from a file or standard input.

Delivery is pull-based and at least once. Receiving hides a message until its visibility timeout.
Delete a message after successful processing, or extend its timeout when processing takes longer.
Message content is Base64 and is limited to 256 KB.

Documentation: https://manual.sakura.ad.jp/cloud/appliance/simplemq/about.html`,
	Example: `  usacloud simplemq-api message send --help
  usacloud simplemq-api message receive --help`,
}

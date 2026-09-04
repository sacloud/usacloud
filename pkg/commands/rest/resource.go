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

package rest

import (
	"github.com/sacloud/usacloud/pkg/commands/iaas/category"
	"github.com/sacloud/usacloud/pkg/core"
)

var Resource = &core.Resource{
	Name:  "rest",
	Usage: "Send an authenticated request directly to the SAKURA Cloud API",
	LongUsage: `Send an authenticated request directly to the SAKURA Cloud API.

Specify either a full HTTPS URL or an API path such as /server. For a path,
usacloud builds the API URL from --zone (or the zone in the active profile).

Examples:
  # List servers in the configured zone
  usacloud rest /server

  # Request an absolute URL (the --zone option is not used)
  usacloud rest https://secure.sakura.ad.jp/cloud/zone/is1a/api/cloud/1.1/server

  # Send JSON in a request body
  usacloud rest -X POST -d '{"Server":{"Name":"example"}}' /server

  # Query a response with JMESPath
  usacloud rest --query 'Servers[].Name' /server`,
	ArgsUsage:          "URL",
	Category:           category.ResourceCategoryOther,
	IsGlobalResource:   true,
	DefaultCommandName: "request",
}

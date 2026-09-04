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

import "github.com/sacloud/usacloud/pkg/output"

var defaultColumnDefs = []output.ColumnDef{
	{Name: "ID", Template: `{{ if .ID.IsString }}{{ .ID.String }}{{ else }}{{ .ID.Int }}{{ end }}`},
	{Name: "Name", Template: `{{ .Status.QueueName }}`},
	{Name: "Description", Template: `{{ .Description.Value }}`},
	{Name: "VisibilityTimeoutSeconds", Template: `{{ .Settings.VisibilityTimeoutSeconds }}`},
	{Name: "ExpireSeconds", Template: `{{ .Settings.ExpireSeconds }}`},
	{Name: "Tags", Template: `{{ .Tags }}`},
	{Name: "Availability", Template: `{{ .Availability }}`},
}

var countMessagesColumnDefs = []output.ColumnDef{
	{Name: "MessageCount", Template: `{{ .MessageCount }}`},
}

var rotateAPIKeyColumnDefs = []output.ColumnDef{
	{Name: "APIKey", Template: `{{ .APIKey }}`},
}

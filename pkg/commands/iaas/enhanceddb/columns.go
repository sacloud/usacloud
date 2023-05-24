// Copyright 2017-2023 The sacloud/usacloud Authors
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

package enhanceddb

import (
	"github.com/sacloud/usacloud/pkg/output"
)

var defaultColumnDefs = []output.ColumnDef{
	{Name: "ID", Template: "{{ if .EnhancedDB }}{{ .EnhancedDB.ID }}{{ else }}{{ .ID }}{{ end }}"},
	{Name: "Name", Template: "{{ if .EnhancedDB }}{{ .EnhancedDB.Name }}{{ else }}{{ .Name }}{{ end }}"},
	{Name: "Tags", Template: "{{ if .EnhancedDB }}{{ .EnhancedDB.Tags}}{{ else }}{{ .Tags }}{{ end }}"},
	{Name: "Region", Template: "{{ if .EnhancedDB }}{{ .EnhancedDB.Region }}{{ else }}{{ .Region }}{{ end }}"},
	{Name: "HostName", Template: "{{ if .EnhancedDB }}{{ .EnhancedDB.HostName }}{{ else }}{{ .HostName }}{{ end }}"},
	{Name: "Port", Template: "{{ if .EnhancedDB }}{{ .EnhancedDB.Port }}{{ else }}{{ .Port }}{{ end }}"},
	{Name: "DatabaseName", Template: "{{ if .EnhancedDB }}{{ .EnhancedDB.DatabaseName }}{{ else }}{{ .DatabaseName }}{{ end }}"},
	{Name: "Description", Template: "{{ if .EnhancedDB }}{{ .EnhancedDB.Description | ellipsis 20 | to_single_line}}{{ else }}{{ .Description | ellipsis 20 | to_single_line}}{{ end }}"},
}

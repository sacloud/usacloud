// Copyright 2017-2021 The Usacloud Authors
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

package ccol

import "github.com/sacloud/usacloud/pkg/output"

var (
	Zone           = output.ColumnDef{Name: "Zone"}
	ID             = output.ColumnDef{Name: "ID"}
	Name           = output.ColumnDef{Name: "Name", Template: "{{ .Name | ellipsis 30 }}"}
	Tags           = output.ColumnDef{Name: "Tags", Template: "{{ if .Tags }}{{ .Tags | ellipsis 20 }}{{ end }}"}
	Description    = output.ColumnDef{Name: "Description", Template: "{{ .Description | ellipsis 20 | to_single_line}}"}
	Size           = output.ColumnDef{Name: "Size", Template: "{{ .SizeMB | mib_to_gib }}GB"}
	Scope          = output.ColumnDef{Name: "Scope", Template: "{{ .Scope | scope_to_key }}"}
	InstanceStatus = output.ColumnDef{Name: "InstanceStatus", Template: `{{ .InstanceStatus }}{{ if ne .Availability "available" }}({{ .Availability }}){{ end }}`}
)

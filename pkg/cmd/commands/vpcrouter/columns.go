// Copyright 2017-2020 The Usacloud Authors
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

package vpcrouter

import (
	"github.com/sacloud/usacloud/pkg/cmd/ccol"
	"github.com/sacloud/usacloud/pkg/output"
)

var defaultColumnDefs = []output.ColumnDef{
	ccol.Zone,
	ccol.ID,
	ccol.Name,
	ccol.Tags,
	ccol.Description,
	{Name: "Plan", Template: `{{ .PlanID | vpc_router_plan_to_key }}`},
	{Name: "VRID", Template: `{{ if .Settings }}{{ if gt .Settings.VRID 0 }}{{ .Settings.VRID }}{{ end }}{{ end }}`},
	{Name: "Upstream", Template: `{{ with index .OriginalValue.Interfaces 0 }}{{ .UpstreamType }}{{ end }}`},
	{
		Name: "IPAddress",
		Template: `
{{- if eq .PlanID 1 }}
{{- with index .Interfaces 0 }}{{ .IPAddress }}{{ end -}}
{{- else }}
{{- with index .Settings.Interfaces 0 }}{{ .VirtualIPAddress }}{{ end -}}
{{- end -}}
`,
	},
	ccol.InstanceStatus,
}

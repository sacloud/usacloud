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

package server

import (
	"github.com/sacloud/usacloud/pkg/cmd/ccol"
	"github.com/sacloud/usacloud/pkg/output"
)

var defaultColumnDefs = []output.ColumnDef{
	{Name: "Zone", Template: `{{ .Zone.Name }}`}, // Note: Serverは元々Zoneフィールドを持つためccol.Zoneとバッティングする。そのためここでは個別定義している
	ccol.ID,
	ccol.Name,
	ccol.Tags,
	{
		Name: "CPU",
	},
	{
		Name:     "Memory",
		Template: `{{ .MemoryMB | mib_to_gib }}`,
	},

	{
		Name:     "IPAddress",
		Template: "{{ if gt (len .Interfaces) 0 }}{{ with index .Interfaces 0 }}{{ if .IPAddress }}{{ .IPAddress }}/{{ .SubnetNetworkMaskLen }}{{ else }}{{ .UserIPAddress }}/{{ .UserSubnetNetworkMaskLen }}{{ end }}{{ end }}{{ end }}",
	},
	{
		Name:     "Upstream(Mbps)",
		Template: `{{ if gt (len .OriginalValue.Interfaces) 0 }}{{ with index .OriginalValue.Interfaces 0 }}{{ .UpstreamType }}{{ end }}/{{ .OriginalValue.BandWidthAt 0 }}{{ end }}`,
	},
	ccol.InstanceStatus,
	{Name: "InstanceHost", Template: "{{ .InstanceHostName }}"},
}

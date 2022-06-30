// Copyright 2017-2022 The sacloud/usacloud Authors
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

package internet

import (
	"github.com/sacloud/usacloud/pkg/ccol"
	"github.com/sacloud/usacloud/pkg/output"
)

var defaultColumnDefs = []output.ColumnDef{
	ccol.Zone,
	ccol.ID,
	ccol.Name,
	ccol.Tags,
	ccol.Description,
	{Name: "BandWidthMbps"},
	{Name: "IPAddresses", Template: "{{ range $i, $v := .Switch.Subnets }}{{ if gt $i 0 }}\n{{ end }}{{ $v.NetworkAddress }}/{{ $v.NetworkMaskLen }}{{ end }}"},
	{Name: "IPv6Net", Template: "{{ if .Switch.IPv6Nets }}{{ with index .Switch.IPv6Nets 0 }}{{ .IPv6Prefix }}/{{ .IPv6PrefixLen }} {{ end }}{{ end }}"},
}

var subnetColumnDefs = []output.ColumnDef{
	ccol.Zone,
	ccol.ID,
	{Name: "SwitchID"},
	{Name: "InternetID"},
	{Name: "NextHop"},
	{Name: "IPAddresses", Template: "{{ .NetworkAddress }}/{{ .NetworkMaskLen }}"},
}

// Copyright 2017-2022 The Usacloud Authors
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

package webaccelerator

import (
	"github.com/sacloud/usacloud/pkg/cmd/ccol"
	"github.com/sacloud/usacloud/pkg/output"
)

var defaultColumnDefs = []output.ColumnDef{
	ccol.ID,
	ccol.Name,
	{Name: "DomainType"},
	{Name: "Domain"},
	{Name: "Origin"},
	{Name: "Status"},
	{Name: "GibSentInLastWeek"},
}

var certificateColumnDefs = []output.ColumnDef{
	ccol.ID,
	{Name: "CertID", Template: `{{ .Current.ID }}`},
	{Name: "DNSNames", Template: `{{ .Current.DNSNames | join "\n" }}`},
	{Name: "SerialNumber", Template: `{{ .Current.SerialNumber }}`},
	{Name: "Validity", Template: `{{ .Current.NotBefore | unix_time_to_date }} ~ {{ .Current.NotAfter | unix_time_to_date }}`},
}

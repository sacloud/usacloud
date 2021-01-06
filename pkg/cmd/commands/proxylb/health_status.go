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

package proxylb

import (
	"github.com/sacloud/usacloud/pkg/cmd/ccol"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/output"
)

var healthStatusCommand = &core.Command{
	Name:       "health-status",
	Aliases:    []string{"health"},
	Category:   "operation",
	Order:      10,
	NoProgress: true,

	ColumnDefs: []output.ColumnDef{
		ccol.ID,
		{Name: "ActiveConn"},
		{Name: "CPS"},
		{Name: "CurrentVIP"},
		{Name: "Servers",
			Template: "{{ range $i, $v := .Servers }}{{ if gt $i 0 }}\n{{ end }}{{ $v.IPAddress }}:{{ $v.Port }} => {{ $v.Status }} (CPS: {{ $v.CPS.Int64 }} / ActiveConn: {{ $v.ActiveConn.Int64 }}){{ end }}",
		},
	},

	SelectorType: core.SelectorTypeRequireMulti,

	ParameterInitializer: func() interface{} {
		return newHealthStatusParameter()
	},
}

type healthStatusParameter struct {
	cflag.IDParameter     `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter `cli:",squash" mapconv:"-"`
}

func newHealthStatusParameter() *healthStatusParameter {
	return &healthStatusParameter{}
}

func init() {
	Resource.AddCommand(healthStatusCommand)
}

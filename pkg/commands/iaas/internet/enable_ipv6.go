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

package internet

import (
	"github.com/sacloud/usacloud/pkg/ccol"
	cflag2 "github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/output"
)

var enableIPv6Command = &core.Command{
	Name:         "enable-ipv6",
	Aliases:      []string{"ipv6-enable"},
	Category:     "ipv6",
	Order:        20,
	SelectorType: core.SelectorTypeRequireMulti,

	ColumnDefs: []output.ColumnDef{
		ccol.Zone,
		ccol.ID,
		{Name: "IPAddresses", Template: "{{.IPv6Prefix}}/{{.IPv6PrefixLen}}"},
	},

	ParameterInitializer: func() interface{} {
		return newEnableIPv6Parameter()
	},
}

type enableIPv6Parameter struct {
	cflag2.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag2.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag2.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag2.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag2.OutputParameter  `cli:",squash" mapconv:"-"`
}

func newEnableIPv6Parameter() *enableIPv6Parameter {
	return &enableIPv6Parameter{}
}

func init() {
	Resource.AddCommand(enableIPv6Command)
}

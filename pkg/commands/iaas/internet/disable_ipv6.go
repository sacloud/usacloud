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

package internet

import (
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/core"
)

var disableIPv6Command = &core.Command{
	Name:         "disable-ipv6",
	Aliases:      []string{"ipv6-disable"},
	Category:     "ipv6",
	Order:        30,
	SelectorType: core.SelectorTypeRequireMulti,

	ColumnDefs: defaultColumnDefs,

	ParameterInitializer: func() interface{} {
		return newDisableIPv6Parameter()
	},
}

type disableIPv6Parameter struct {
	cflag.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
}

func newDisableIPv6Parameter() *disableIPv6Parameter {
	return &disableIPv6Parameter{}
}

func init() {
	Resource.AddCommand(disableIPv6Command)
}

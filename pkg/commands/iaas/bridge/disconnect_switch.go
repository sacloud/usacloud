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

package bridge

import (
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/core"
)

var disconnectSwitchCommand = &core.Command{
	Name:     "disconnect-switch",
	Aliases:  []string{"switch-disconnect"},
	Category: "operation",
	Order:    10,

	ColumnDefs:   defaultColumnDefs,
	SelectorType: core.SelectorTypeRequireSingle,

	ParameterInitializer: func() interface{} {
		return newDisconnectSwitchParameter()
	},
}

type disconnectSwitchParameter struct {
	cflag.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
}

func newDisconnectSwitchParameter() *disconnectSwitchParameter {
	return &disconnectSwitchParameter{}
}

func init() {
	Resource.AddCommand(disconnectSwitchCommand)
}

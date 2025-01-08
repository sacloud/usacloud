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

package disk

import (
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/core"
)

var disconnectFromServerCommand = &core.Command{
	Name:         "disconnect-from-server",
	Aliases:      []string{"server-disconnect"}, // v0との互換用
	Category:     "operation",
	Order:        20,
	SelectorType: core.SelectorTypeRequireSingle,

	ParameterInitializer: func() interface{} {
		return newDisconnectFromServerParameter()
	},
}

type disconnectFromServerParameter struct {
	cflag.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
}

func newDisconnectFromServerParameter() *disconnectFromServerParameter {
	return &disconnectFromServerParameter{}
}

func init() {
	Resource.AddCommand(disconnectFromServerCommand)
}

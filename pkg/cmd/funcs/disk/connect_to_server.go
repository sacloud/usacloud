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

package disk

import (
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cmd/core"
)

var connectToServerCommand = &core.Command{
	Name:         "connect-to-server",
	Aliases:      []string{"server-connect"}, // v0との互換用
	Category:     "server",
	Order:        10,
	SelectorType: core.SelectorTypeRequireSingle,

	ParameterInitializer: func() interface{} {
		return newConnectToServerParameter()
	},
}

type connectToServerParameter struct {
	core.ZoneParameter `cli:",squash" mapconv:",squash"`
	core.IDParameter   `cli:",squash" mapconv:",squash"`

	ServerID types.ID `cli:",category=disk" validate:"required"`

	core.ConfirmParameter `cli:",squash" mapconv:"-"`
}

func newConnectToServerParameter() *connectToServerParameter {
	return &connectToServerParameter{}
}

func init() {
	Resource.AddCommand(connectToServerCommand)
}

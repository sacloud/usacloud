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

package vpcrouter

import (
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
)

var shutdownCommand = &core.Command{
	Name:     "shutdown",
	Aliases:  []string{"power-off"},
	Category: "power",
	Order:    20,

	ColumnDefs: defaultColumnDefs,

	SelectorType: core.SelectorTypeRequireMulti,

	ParameterInitializer: func() interface{} {
		return newShutdownParameter()
	},
}

type shutdownParameter struct {
	cflag.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`

	cflag.NoWaitParameter `cli:",squash" mapconv:",squash"`
	ForceShutdown         bool `cli:",short=f,aliases=force"`
}

func newShutdownParameter() *shutdownParameter {
	return &shutdownParameter{}
}

func init() {
	Resource.AddCommand(shutdownCommand)
}

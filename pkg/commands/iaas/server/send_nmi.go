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

package server

import (
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/core"
)

var sendNMICommand = &core.Command{
	Name:     "send-nmi",
	Category: "power",
	Order:    30,

	ColumnDefs: defaultColumnDefs,

	SelectorType: core.SelectorTypeRequireMulti,

	ParameterInitializer: func() interface{} {
		return newSendNMIParameter()
	},
}

type sendNMIParameter struct {
	cflag.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
}

func newSendNMIParameter() *sendNMIParameter {
	return &sendNMIParameter{}
}

func init() {
	Resource.AddCommand(sendNMICommand)
}

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
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/core"
)

var addSubnetCommand = &core.Command{
	Name:         "add-subnet",
	Aliases:      []string{"subnet-add"},
	Category:     "subnet",
	Order:        20,
	SelectorType: core.SelectorTypeRequireMulti,

	ColumnDefs: subnetColumnDefs,

	ParameterInitializer: func() interface{} {
		return newAddSubnetParameter()
	},
}

type addSubnetParameter struct {
	cflag.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	NetworkMaskLen int    `cli:"netmask,aliases=network-mask-len,options=internet_network_mask_len" validate:"required,internet_network_mask_len"`
	NextHop        string `validate:"required,ipv4"`
}

func newAddSubnetParameter() *addSubnetParameter {
	return &addSubnetParameter{
		NetworkMaskLen: 28,
	}
}

func init() {
	Resource.AddCommand(addSubnetCommand)
}

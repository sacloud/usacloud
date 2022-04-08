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
	"github.com/sacloud/iaas-api-go/types"
	cflag2 "github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/core"
)

var deleteSubnetCommand = &core.Command{
	Name:         "delete-subnet",
	Aliases:      []string{"subnet-delete"},
	Category:     "subnet",
	Order:        30,
	SelectorType: core.SelectorTypeRequireSingle, // --subnet-idを指定する関係上、Singleしかありえない

	ParameterInitializer: func() interface{} {
		return newDeleteSubnetParameter()
	},
}

type deleteSubnetParameter struct {
	cflag2.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag2.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag2.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag2.ConfirmParameter `cli:",squash" mapconv:"-"`

	cflag2.FailIfNotFoundParameter `cli:",squash" mapconv:",squash"`
	SubnetID                       types.ID `validate:"required"`
}

func newDeleteSubnetParameter() *deleteSubnetParameter {
	return &deleteSubnetParameter{}
}

func init() {
	Resource.AddCommand(deleteSubnetCommand)
}

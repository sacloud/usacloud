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

package nfs

import (
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
)

var createCommand = &core.Command{
	Name:     "create",
	Category: "basic",
	Order:    20,

	ColumnDefs: defaultColumnDefs,

	ParameterInitializer: func() interface{} {
		return newCreateParameter()
	},
}

type createParameter struct {
	cflag.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag.InputParameter   `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameParameter   `cli:",squash" mapconv:",squash"`
	cflag.DescParameter   `cli:",squash" mapconv:",squash"`
	cflag.TagsParameter   `cli:",squash" mapconv:",squash"`
	cflag.IconIDParameter `cli:",squash" mapconv:",squash"`

	Plan string `cli:"plan,options=nfs_plan,category=plan,order=10" mapconv:",filters=nfs_plan_to_value" validate:"required,nfs_plan"`
	Size int    `cli:",category=plan,order=20" validate:"required"`

	SwitchID       types.ID `cli:",category=network,order=10" validate:"required"`
	IPAddresses    []string `cli:"ip-address,aliases=ipaddress,category=network,order=20" validate:"required,min=1,max=2,dive,ipv4"`
	NetworkMaskLen int      `cli:",category=network,order=30" validate:"required,min=1,max=32"`
	DefaultRoute   string   `cli:",category=network,order=40" validate:"omitempty,ipv4"`

	cflag.NoWaitParameter `cli:",squash" mapconv:",squash"`
}

func newCreateParameter() *createParameter {
	return &createParameter{
		Plan: "ssd",
	}
}

func init() {
	Resource.AddCommand(createCommand)
}

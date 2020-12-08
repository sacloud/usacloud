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

package vpcrouter

import (
	"github.com/sacloud/libsacloud/v2/helper/service/vpcrouter"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/util"
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
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	Name        string   `validate:"required"`
	Description string   `validate:"description"`
	Tags        []string `validate:"tags"`
	IconID      types.ID

	Plan string `cli:"plan,options=vpc_router_plan_premium" mapconv:"PlanID,filters=vpc_router_plan_premium_to_value" validate:"required,vpc_router_plan_premium"`

	PublicNetworkInterface vpcrouter.PremiumNICSetting `mapconv:"NICSetting,omitempty"`

	PrivateNetworkInterfacesData string                                   `cli:"private-network-interfaces" mapconv:"-"`
	PrivateNetworkInterfaces     []*vpcrouter.AdditionalPremiumNICSetting `cli:"-" mapconv:"AdditionalNICSettings"`

	RouterSetting routerSetting `cli:",squash" mapconv:",recursive"`

	BootAfterCreate bool
	NoWait          bool
}

func newCreateParameter() *createParameter {
	return &createParameter{
		Plan: "premium",
	}
}

func init() {
	Resource.AddCommand(createCommand)
}

// Customize パラメータ変換処理
func (p *createParameter) Customize(ctx cli.Context) error {
	if p.PrivateNetworkInterfacesData != "" {
		var nics []*vpcrouter.AdditionalPremiumNICSetting
		if err := util.MarshalJSONFromPathOrContent(p.PrivateNetworkInterfacesData, &nics); err != nil {
			return err
		}
		p.PrivateNetworkInterfaces = append(p.PrivateNetworkInterfaces, nics...)
	}
	return p.RouterSetting.Customize(ctx)
}

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

var updateStandardCommand = &core.Command{
	Name:         "update-standard",
	Category:     "basic",
	Order:        45,
	SelectorType: core.SelectorTypeRequireMulti,

	ColumnDefs: defaultColumnDefs,

	ParameterInitializer: func() interface{} {
		return newUpdateStandardParameter()
	},
}

type updateStandardParameter struct {
	cflag.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	Name        *string   `validate:"omitempty,min=1"`
	Description *string   `validate:"omitempty,description"`
	Tags        *[]string `validate:"omitempty,tags"`
	IconID      *types.ID

	PrivateNetworkInterfacesData string                                           `cli:"private-network-interfaces" mapconv:"-"`
	PrivateNetworkInterfaces     *[]*vpcrouter.AdditionalStandardNICSettingUpdate `cli:"-" mapconv:"AdditionalNICSettings"`

	RouterSetting routerSettingUpdate `cli:",squash" mapconv:",omitempty,recursive"`

	SettingsHash string
	NoWait       bool
}

func newUpdateStandardParameter() *updateStandardParameter {
	return &updateStandardParameter{
		// TODO デフォルト値はここで設定する
	}
}

func init() {
	Resource.AddCommand(updateStandardCommand)
}

// Customize パラメータ変換処理
func (p *updateStandardParameter) Customize(ctx cli.Context) error {
	if p.PrivateNetworkInterfacesData != "" {
		var nics []*vpcrouter.AdditionalStandardNICSettingUpdate
		if err := util.MarshalJSONFromPathOrContent(p.PrivateNetworkInterfacesData, &nics); err != nil {
			return err
		}
		if p.PrivateNetworkInterfaces == nil {
			p.PrivateNetworkInterfaces = &[]*vpcrouter.AdditionalStandardNICSettingUpdate{}
		}
		*p.PrivateNetworkInterfaces = append(*p.PrivateNetworkInterfaces, nics...)
	}
	return p.RouterSetting.Customize(ctx)
}

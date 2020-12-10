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

package mobilegateway

import (
	"github.com/sacloud/libsacloud/v2/helper/service/mobilegateway"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/util"
)

var updateCommand = &core.Command{
	Name:         "update",
	Category:     "basic",
	Order:        40,
	SelectorType: core.SelectorTypeRequireMulti,

	ColumnDefs: defaultColumnDefs,

	ParameterInitializer: func() interface{} {
		return newUpdateParameter()
	},
}

type updateParameter struct {
	cflag.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.InputParameter   `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.DescUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.TagsUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.IconIDUpdateParameter `cli:",squash" mapconv:",omitempty,squash"`

	PrivateInterface mobilegateway.PrivateInterfaceSettingUpdate `mapconv:",omitempty" validate:"omitempty"`

	InternetConnectionEnabled       *bool
	InterDeviceCommunicationEnabled *bool

	SIMsData *string                      `cli:"sims" mapconv:"-"`
	SIMs     *[]*mobilegateway.SIMSetting `cli:"-"`

	SIMRoutesData *string                           `cli:"sim-routes" mapconv:"-"`
	SIMRoutes     *[]*mobilegateway.SIMRouteSetting `cli:"-"`

	StaticRoutesData *string                              `cli:"static-routes" mapconv:"-"`
	StaticRoutes     *[]*sacloud.MobileGatewayStaticRoute `cli:"-"`

	DNS           mobilegateway.DNSSettingUpdate    `cli:",squash" mapconv:",omitempty"`
	TrafficConfig mobilegateway.TrafficConfigUpdate `mapconv:",omitempty"`

	cflag.NoWaitParameter `cli:",squash" mapconv:",squash"`
}

func newUpdateParameter() *updateParameter {
	return &updateParameter{}
}

func init() {
	Resource.AddCommand(updateCommand)
}

// Customize パラメータ変換処理
func (p *updateParameter) Customize(_ cli.Context) error {
	if p.SIMsData != nil && *p.SIMsData != "" {
		var sims []*mobilegateway.SIMSetting
		if err := util.MarshalJSONFromPathOrContent(*p.SIMsData, &sims); err != nil {
			return err
		}
		if p.SIMs == nil {
			p.SIMs = &[]*mobilegateway.SIMSetting{}
		}
		*p.SIMs = append(*p.SIMs, sims...)
	}

	if p.SIMRoutesData != nil && *p.SIMRoutesData != "" {
		var simRoutes []*mobilegateway.SIMRouteSetting
		if err := util.MarshalJSONFromPathOrContent(*p.SIMRoutesData, &simRoutes); err != nil {
			return err
		}
		if p.SIMRoutes == nil {
			p.SIMRoutes = &[]*mobilegateway.SIMRouteSetting{}
		}
		*p.SIMRoutes = append(*p.SIMRoutes, simRoutes...)
	}

	if p.StaticRoutesData != nil && *p.StaticRoutesData != "" {
		var staticRoutes []*sacloud.MobileGatewayStaticRoute
		if err := util.MarshalJSONFromPathOrContent(*p.StaticRoutesData, &staticRoutes); err != nil {
			return err
		}
		if p.StaticRoutes == nil {
			p.StaticRoutes = &[]*sacloud.MobileGatewayStaticRoute{}
		}
		*p.StaticRoutes = append(*p.StaticRoutes, staticRoutes...)
	}

	return nil
}

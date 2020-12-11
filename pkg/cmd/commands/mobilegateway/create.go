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

	InternetConnectionEnabled       bool
	InterDeviceCommunicationEnabled bool

	SIMsData string                      `cli:"sims" mapconv:"-"`
	SIMs     []*mobilegateway.SIMSetting `cli:"-"`

	SIMRoutesData string                           `cli:"sim-routes" mapconv:"-"`
	SIMRoutes     []*mobilegateway.SIMRouteSetting `cli:"-"`

	StaticRoutesData string                              `cli:"static-routes" mapconv:"-"`
	StaticRoutes     []*sacloud.MobileGatewayStaticRoute `cli:"-"`

	PrivateInterface mobilegateway.PrivateInterfaceSetting `cli:",category=network" mapconv:",omitempty" validate:"omitempty"`
	DNS              sacloud.MobileGatewayDNSSetting       `cli:",squash" mapconv:",omitempty" validate:"omitempty"`
	TrafficConfig    sacloud.MobileGatewayTrafficControl   `mapconv:",omitempty"`

	cflag.NoWaitParameter `cli:",squash" mapconv:",squash"`
	BootAfterCreate       bool
}

func newCreateParameter() *createParameter {
	return &createParameter{}
}

func init() {
	Resource.AddCommand(createCommand)
}

// Customize パラメータ変換処理
func (p *createParameter) Customize(_ cli.Context) error {
	if p.SIMsData != "" {
		var sims []*mobilegateway.SIMSetting
		if err := util.MarshalJSONFromPathOrContent(p.SIMsData, &sims); err != nil {
			return err
		}
		p.SIMs = append(p.SIMs, sims...)
	}

	if p.SIMRoutesData != "" {
		var simRoutes []*mobilegateway.SIMRouteSetting
		if err := util.MarshalJSONFromPathOrContent(p.SIMRoutesData, &simRoutes); err != nil {
			return err
		}
		p.SIMRoutes = append(p.SIMRoutes, simRoutes...)
	}

	if p.StaticRoutesData != "" {
		var staticRoutes []*sacloud.MobileGatewayStaticRoute
		if err := util.MarshalJSONFromPathOrContent(p.StaticRoutesData, &staticRoutes); err != nil {
			return err
		}
		p.StaticRoutes = append(p.StaticRoutes, staticRoutes...)
	}

	return nil
}

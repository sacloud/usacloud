// Copyright 2017-2025 The sacloud/usacloud Authors
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
	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-service-go/mobilegateway"
	"github.com/sacloud/packages-go/pointer"
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/examples"
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
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.DescUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.TagsUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.IconIDUpdateParameter `cli:",squash" mapconv:",omitempty,squash"`

	PrivateInterface mobilegateway.PrivateInterfaceSettingUpdate `cli:",category=network" mapconv:",omitempty" validate:"omitempty"`

	InternetConnectionEnabled       *bool
	InterDeviceCommunicationEnabled *bool

	SIMsData *string                      `cli:"sims" mapconv:"-" json:"-"`
	SIMs     *[]*mobilegateway.SIMSetting `cli:"-"`

	SIMRoutesData *string                           `cli:"sim-routes" mapconv:"-" json:"-"`
	SIMRoutes     *[]*mobilegateway.SIMRouteSetting `cli:"-"`

	StaticRoutesData *string                           `cli:"static-routes" mapconv:"-" json:"-"`
	StaticRoutes     *[]*iaas.MobileGatewayStaticRoute `cli:"-"`

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
		var staticRoutes []*iaas.MobileGatewayStaticRoute
		if err := util.MarshalJSONFromPathOrContent(*p.StaticRoutesData, &staticRoutes); err != nil {
			return err
		}
		if p.StaticRoutes == nil {
			p.StaticRoutes = &[]*iaas.MobileGatewayStaticRoute{}
		}
		*p.StaticRoutes = append(*p.StaticRoutes, staticRoutes...)
	}

	return nil
}

func (p *updateParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &updateParameter{
		ZoneParameter:                   examples.Zones(ctx.Option().Zones),
		NameUpdateParameter:             examples.NameUpdate,
		DescUpdateParameter:             examples.DescriptionUpdate,
		TagsUpdateParameter:             examples.TagsUpdate,
		IconIDUpdateParameter:           examples.IconIDUpdate,
		InternetConnectionEnabled:       pointer.NewBool(true),
		InterDeviceCommunicationEnabled: pointer.NewBool(true),
		SIMs: &[]*mobilegateway.SIMSetting{
			{
				SIMID:     examples.ID,
				IPAddress: examples.IPAddress,
			},
		},
		SIMRoutes: &[]*mobilegateway.SIMRouteSetting{
			{
				SIMID:  examples.ID,
				Prefix: "192.0.2.0/24",
			},
		},
		StaticRoutes: &[]*iaas.MobileGatewayStaticRoute{
			{
				NextHop: "192.0.2.2",
				Prefix:  "192.0.2.0/24",
			},
		},
		PrivateInterface: mobilegateway.PrivateInterfaceSettingUpdate{
			SwitchID:       &examples.ID,
			IPAddress:      &examples.IPAddress,
			NetworkMaskLen: &examples.NetworkMaskLen,
		},
		DNS: mobilegateway.DNSSettingUpdate{
			DNS1: pointer.NewString("133.242.0.3 | 210.188.224.10 | n.n.n.n"),
			DNS2: pointer.NewString("133.242.0.4 | 210.188.224.11 | n.n.n.n"),
		},
		TrafficConfig: mobilegateway.TrafficConfigUpdate{
			TrafficQuotaInMB:       pointer.NewInt(10),
			BandWidthLimitInKbps:   pointer.NewInt(128),
			EmailNotifyEnabled:     pointer.NewBool(true),
			SlackNotifyEnabled:     pointer.NewBool(true),
			SlackNotifyWebhooksURL: &examples.SlackNotifyWebhooksURL,
			AutoTrafficShaping:     pointer.NewBool(true),
		},
		NoWaitParameter: cflag.NoWaitParameter{
			NoWait: false,
		},
	}
}

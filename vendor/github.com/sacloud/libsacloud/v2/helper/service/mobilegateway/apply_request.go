// Copyright 2016-2020 The Libsacloud Authors
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
	"github.com/sacloud/libsacloud/v2/helper/builder"
	mobileGatewayBuilder "github.com/sacloud/libsacloud/v2/helper/builder/mobilegateway"
	"github.com/sacloud/libsacloud/v2/helper/service"
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type ApplyRequest struct {
	Zone string `request:"-" validate:"required"`

	ID                              types.ID `request:"-"`
	Name                            string   `validate:"required"`
	Description                     string   `validate:"min=0,max=512"`
	Tags                            types.Tags
	IconID                          types.ID
	PrivateInterface                *PrivateInterfaceSetting `validate:"omitempty"`
	StaticRoutes                    []*sacloud.MobileGatewayStaticRoute
	SIMRoutes                       []*SIMRouteSetting
	InternetConnectionEnabled       bool
	InterDeviceCommunicationEnabled bool
	DNS                             *DNSSetting
	SIMs                            []*SIMSetting
	TrafficConfig                   *TrafficConfig

	SettingsHash    string
	BootAfterCreate bool
	NoWait          bool
}

// PrivateInterfaceSetting represents API parameter/response structure
type PrivateInterfaceSetting struct {
	SwitchID       types.ID `request:",omitempty"`
	IPAddress      string   `request:",omitempty" validate:"required,ipv4"`
	NetworkMaskLen int      `request:",omitempty"`
}

// SIMRouteSetting represents API parameter/response structure
type SIMRouteSetting struct {
	SIMID  types.ID
	Prefix string `validate:"required"`
}

// SIMSetting represents API parameter/response structure
type SIMSetting struct {
	SIMID     types.ID
	IPAddress string `validate:"required,ipv4"`
}

type DNSSetting struct {
	DNS1 string `request:",omitempty" validate:"required_with=DNS2,omitempty,ipv4"`
	DNS2 string `request:",omitempty" validate:"required_with=DNS1,omitempty,ipv4"`
}

type TrafficConfig struct {
	TrafficQuotaInMB       int    `request:",omitempty"`
	BandWidthLimitInKbps   int    `request:",omitempty"`
	EmailNotifyEnabled     bool   `request:",omitempty"`
	SlackNotifyEnabled     bool   `request:",omitempty"`
	SlackNotifyWebhooksURL string `request:",omitempty"`
	AutoTrafficShaping     bool   `request:",omitempty"`
}

func (req *ApplyRequest) Validate() error {
	return validate.Struct(req)
}

func (req *ApplyRequest) Builder(caller sacloud.APICaller) (*mobileGatewayBuilder.Builder, error) {
	var privateInterface *mobileGatewayBuilder.PrivateInterfaceSetting
	if req.PrivateInterface != nil {
		privateInterface = &mobileGatewayBuilder.PrivateInterfaceSetting{
			SwitchID:       req.PrivateInterface.SwitchID,
			IPAddress:      req.PrivateInterface.IPAddress,
			NetworkMaskLen: req.PrivateInterface.NetworkMaskLen,
		}
	}

	var simRoutes []*mobileGatewayBuilder.SIMRouteSetting
	for _, sr := range req.SIMRoutes {
		simRoutes = append(simRoutes, &mobileGatewayBuilder.SIMRouteSetting{
			SIMID:  sr.SIMID,
			Prefix: sr.Prefix,
		})
	}

	var sims []*mobileGatewayBuilder.SIMSetting
	for _, s := range req.SIMs {
		sims = append(sims, &mobileGatewayBuilder.SIMSetting{
			SIMID:     s.SIMID,
			IPAddress: s.IPAddress,
		})
	}

	var dns *sacloud.MobileGatewayDNSSetting
	if req.DNS != nil {
		dns = &sacloud.MobileGatewayDNSSetting{
			DNS1: req.DNS.DNS1,
			DNS2: req.DNS.DNS2,
		}
	}

	var trafficConfig *sacloud.MobileGatewayTrafficControl
	if req.TrafficConfig != nil {
		trafficConfig = &sacloud.MobileGatewayTrafficControl{}
		if err := service.RequestConvertTo(req.TrafficConfig, trafficConfig); err != nil {
			return nil, err
		}
	}

	return &mobileGatewayBuilder.Builder{
		Name:                            req.Name,
		Description:                     req.Description,
		Tags:                            req.Tags,
		IconID:                          req.IconID,
		PrivateInterface:                privateInterface,
		StaticRoutes:                    req.StaticRoutes,
		SIMRoutes:                       simRoutes,
		InternetConnectionEnabled:       req.InternetConnectionEnabled,
		InterDeviceCommunicationEnabled: req.InterDeviceCommunicationEnabled,
		DNS:                             dns,
		SIMs:                            sims,
		TrafficConfig:                   trafficConfig,
		SettingsHash:                    req.SettingsHash,
		NoWait:                          req.NoWait,
		SetupOptions:                    &builder.RetryableSetupParameter{BootAfterBuild: req.BootAfterCreate},
		Client:                          mobileGatewayBuilder.NewAPIClient(caller),
	}, nil
}

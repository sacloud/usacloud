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
	DNS                             *sacloud.MobileGatewayDNSSetting
	SIMs                            []*SIMSetting
	TrafficConfig                   *sacloud.MobileGatewayTrafficControl

	SettingsHash    string
	BootAfterCreate bool
	NoWait          bool
}

// PrivateInterfaceSetting represents API parameter/response structure
type PrivateInterfaceSetting struct {
	SwitchID       types.ID
	IPAddress      string `validate:"ipv4"`
	NetworkMaskLen int
}

// SIMRouteSetting represents API parameter/response structure
type SIMRouteSetting struct {
	SIMID  types.ID
	Prefix string
}

// SIMSetting represents API parameter/response structure
type SIMSetting struct {
	SIMID     types.ID
	IPAddress string `validate:"ipv4"`
}

func (req *ApplyRequest) Validate() error {
	return validate.Struct(req)
}

func (req *ApplyRequest) Builder(caller sacloud.APICaller) *mobileGatewayBuilder.Builder {
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
		DNS:                             req.DNS,
		SIMs:                            sims,
		TrafficConfig:                   req.TrafficConfig,
		SettingsHash:                    req.SettingsHash,
		NoWait:                          req.NoWait,
		SetupOptions:                    &builder.RetryableSetupParameter{BootAfterBuild: req.BootAfterCreate},
		Client:                          mobileGatewayBuilder.NewAPIClient(caller),
	}
}

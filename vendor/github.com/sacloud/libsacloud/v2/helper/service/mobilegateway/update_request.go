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
	"context"

	"github.com/sacloud/libsacloud/v2/helper/service"
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type UpdateRequest struct {
	Zone string   `validate:"required"`
	ID   types.ID `validate:"required"`

	Name                            *string                              `request:",omitempty"`
	Description                     *string                              `request:",omitempty" validate:"min=0,max=512"`
	Tags                            *types.Tags                          `request:",omitempty"`
	IconID                          *types.ID                            `request:",omitempty"`
	PrivateInterface                *PrivateInterfaceSetting             `request:",omitempty"`
	StaticRoutes                    *[]*sacloud.MobileGatewayStaticRoute `request:",omitempty"`
	SIMRoutes                       *[]*SIMRouteSetting                  `request:",omitempty"`
	InternetConnectionEnabled       *bool                                `request:",omitempty"`
	InterDeviceCommunicationEnabled *bool                                `request:",omitempty"`
	DNS                             *sacloud.MobileGatewayDNSSetting     `request:",omitempty"`
	SIMs                            *[]*SIMSetting                       `request:",omitempty"`
	TrafficConfig                   *sacloud.MobileGatewayTrafficControl `request:",omitempty"`

	SettingsHash string
	NoWait       bool
}

func (req *UpdateRequest) Validate() error {
	return validate.Struct(req)
}

func (req *UpdateRequest) ApplyRequest(ctx context.Context, caller sacloud.APICaller) (*ApplyRequest, error) {
	mgwOp := sacloud.NewMobileGatewayOp(caller)
	current, err := mgwOp.Read(ctx, req.Zone, req.ID)
	if err != nil {
		return nil, err
	}

	var privateInterface *PrivateInterfaceSetting
	for i, nic := range current.InterfaceSettings {
		if nic.Index == 1 {
			privateInterface = &PrivateInterfaceSetting{
				SwitchID:       current.Interfaces[i].SwitchID,
				IPAddress:      nic.IPAddress[0],
				NetworkMaskLen: nic.NetworkMaskLen,
			}
		}
	}

	simRoutes, err := mgwOp.GetSIMRoutes(ctx, req.Zone, req.ID)
	if err != nil {
		return nil, err
	}
	var simRouteSettings []*SIMRouteSetting
	for _, r := range simRoutes {
		simRouteSettings = append(simRouteSettings, &SIMRouteSetting{
			SIMID:  types.StringID(r.ResourceID),
			Prefix: r.Prefix,
		})
	}

	dns, err := mgwOp.GetDNS(ctx, req.Zone, req.ID)
	if err != nil {
		return nil, err
	}

	sims, err := mgwOp.ListSIM(ctx, req.Zone, req.ID)
	if err != nil {
		return nil, err
	}
	var simSettings []*SIMSetting
	for _, s := range sims {
		simSettings = append(simSettings, &SIMSetting{
			SIMID:     types.StringID(s.ResourceID),
			IPAddress: s.IP,
		})
	}

	trafficConfig, err := mgwOp.GetTrafficConfig(ctx, req.Zone, req.ID)
	if err != nil {
		return nil, err
	}

	applyRequest := &ApplyRequest{
		Name:                            current.Name,
		Description:                     current.Description,
		Tags:                            current.Tags,
		IconID:                          current.IconID,
		PrivateInterface:                privateInterface,
		StaticRoutes:                    current.StaticRoutes,
		SIMRoutes:                       simRouteSettings,
		InternetConnectionEnabled:       current.InternetConnectionEnabled.Bool(),
		InterDeviceCommunicationEnabled: current.InterDeviceCommunicationEnabled.Bool(),
		DNS:                             dns,
		SIMs:                            simSettings,
		TrafficConfig:                   trafficConfig,
		SettingsHash:                    current.SettingsHash,
	}

	if err := service.RequestConvertTo(req, applyRequest); err != nil {
		return nil, err
	}
	return applyRequest, nil
}

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
	Description                     *string                              `request:",omitempty" validate:"omitempty,min=0,max=512"`
	Tags                            *types.Tags                          `request:",omitempty"`
	IconID                          *types.ID                            `request:",omitempty"`
	PrivateInterface                *PrivateInterfaceSettingUpdate       `request:",omitempty,recursive"`
	StaticRoutes                    *[]*sacloud.MobileGatewayStaticRoute `request:",omitempty"`
	SIMRoutes                       *[]*SIMRouteSetting                  `request:",omitempty"`
	InternetConnectionEnabled       *bool                                `request:",omitempty"`
	InterDeviceCommunicationEnabled *bool                                `request:",omitempty"`
	DNS                             *DNSSettingUpdate                    `request:",omitempty,recursive"`
	SIMs                            *[]*SIMSetting                       `request:",omitempty"`
	TrafficConfig                   *TrafficConfigUpdate                 `request:",omitempty,recursive"`

	SettingsHash string
	NoWait       bool
}

// PrivateInterfaceSetting represents API parameter/response structure
type PrivateInterfaceSettingUpdate struct {
	SwitchID       *types.ID `request:",omitempty"`
	IPAddress      *string   `request:",omitempty" validate:"omitempty,ipv4"`
	NetworkMaskLen *int      `request:",omitempty"`
}

type DNSSettingUpdate struct {
	DNS1 *string `request:",omitempty" validate:"required_with=DNS2,omitempty,ipv4"`
	DNS2 *string `request:",omitempty" validate:"required_with=DNS1,omitempty,ipv4"`
}

type TrafficConfigUpdate struct {
	TrafficQuotaInMB       *int    `request:",omitempty"`
	BandWidthLimitInKbps   *int    `request:",omitempty"`
	EmailNotifyEnabled     *bool   `request:",omitempty"`
	SlackNotifyEnabled     *bool   `request:",omitempty"`
	SlackNotifyWebhooksURL *string `request:",omitempty"`
	AutoTrafficShaping     *bool   `request:",omitempty"`
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
	for _, nic := range current.InterfaceSettings {
		if nic.Index == 1 && len(current.Interfaces) > 1 {
			privateInterface = &PrivateInterfaceSetting{
				SwitchID:       current.Interfaces[nic.Index].SwitchID,
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

	currentDNS, err := mgwOp.GetDNS(ctx, req.Zone, req.ID)
	if err != nil {
		return nil, err
	}
	var dns *DNSSetting
	if currentDNS != nil {
		dns = &DNSSetting{
			DNS1: currentDNS.DNS1,
			DNS2: currentDNS.DNS2,
		}
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

	currentTrafficConfig, err := mgwOp.GetTrafficConfig(ctx, req.Zone, req.ID)
	if err != nil {
		return nil, err
	}
	var trafficConfig *TrafficConfig
	if currentTrafficConfig != nil {
		trafficConfig = &TrafficConfig{}
		if err := service.RequestConvertTo(currentTrafficConfig, trafficConfig); err != nil {
			return nil, err
		}
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

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
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type CreateRequest struct {
	Zone string `validate:"required"`

	Name                            string `validate:"required"`
	Description                     string `validate:"min=0,max=512"`
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

	NoWait          bool
	BootAfterCreate bool
}

func (req *CreateRequest) Validate() error {
	return validate.Struct(req)
}

func (req *CreateRequest) ApplyRequest() *ApplyRequest {
	return &ApplyRequest{
		Zone:                            req.Zone,
		Name:                            req.Name,
		Description:                     req.Description,
		Tags:                            req.Tags,
		IconID:                          req.IconID,
		PrivateInterface:                req.PrivateInterface,
		StaticRoutes:                    req.StaticRoutes,
		SIMRoutes:                       req.SIMRoutes,
		InternetConnectionEnabled:       req.InternetConnectionEnabled,
		InterDeviceCommunicationEnabled: req.InterDeviceCommunicationEnabled,
		DNS:                             req.DNS,
		SIMs:                            req.SIMs,
		TrafficConfig:                   req.TrafficConfig,
		NoWait:                          req.NoWait,
		BootAfterCreate:                 req.BootAfterCreate,
	}
}

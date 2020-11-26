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

type ApplyRequest struct {
	Zone string `request:"-" validate:"required"`

	ID                              types.ID `request:"-"`
	Name                            string   `validate:"required"`
	Description                     string   `validate:"min=0,max=512"`
	Tags                            types.Tags
	IconID                          types.ID
	PrivateInterface                *PrivateInterfaceSetting `validate:"omitempty"`
	StaticRoutes                    []*sacloud.MobileGatewayStaticRoute
	SimRoutes                       []*SIMRouteSetting
	InternetConnectionEnabled       types.StringFlag
	InterDeviceCommunicationEnabled types.StringFlag
	DNS                             *sacloud.MobileGatewayDNSSetting
	SIMs                            []*SIMSetting
	TrafficConfig                   *sacloud.MobileGatewayTrafficControl
}

func (req *ApplyRequest) Validate() error {
	return validate.Struct(req)
}

// PrivateInterfaceSetting represents API parameter/response structure
type PrivateInterfaceSetting struct {
	SwitchID       types.ID
	IPAddress      []string `validate:"ipv4"`
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

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

package vpcrouter

import (
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type CreateStandardRequest struct {
	Zone string `request:"-" validate:"required"`

	Name        string `validate:"required"`
	Description string `validate:"min=0,max=512"`
	Tags        types.Tags
	IconID      types.ID

	AdditionalNICSettings []*AdditionalStandardNICSetting
	RouterSetting         *RouterSetting
	NoWait                bool
	BootAfterCreate       bool
}

func (req *CreateStandardRequest) Validate() error {
	return validate.Struct(req)
}

func (req *CreateStandardRequest) ApplyRequest() *ApplyRequest {
	var additionalNICs []AdditionalNICSettingHolder
	for _, nic := range req.AdditionalNICSettings {
		additionalNICs = append(additionalNICs, nic)
	}
	return &ApplyRequest{
		Zone:                  req.Zone,
		Name:                  req.Name,
		Description:           req.Description,
		Tags:                  req.Tags,
		IconID:                req.IconID,
		PlanID:                types.VPCRouterPlans.Standard,
		NICSetting:            &StandardNICSetting{},
		AdditionalNICSettings: additionalNICs,
		RouterSetting:         req.RouterSetting,
		NoWait:                req.NoWait,
		BootAfterCreate:       req.BootAfterCreate,
	}
}

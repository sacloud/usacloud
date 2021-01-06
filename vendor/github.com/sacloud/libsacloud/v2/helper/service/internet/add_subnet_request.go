// Copyright 2016-2021 The Libsacloud Authors
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

package internet

import (
	"github.com/sacloud/libsacloud/v2/helper/service"
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type AddSubnetRequest struct {
	Zone string   `request:"-" validate:"required"`
	ID   types.ID `request:"-" validate:"required"`

	NetworkMaskLen int    `validate:"required,min=24,max=28"`
	NextHop        string `validate:"required,ipv4"`
}

func (req *AddSubnetRequest) Validate() error {
	return validate.Struct(req)
}

func (req *AddSubnetRequest) ToRequestParameter(current *sacloud.Internet) (*sacloud.InternetAddSubnetRequest, error) {
	r := &sacloud.InternetAddSubnetRequest{}
	if err := service.RequestConvertTo(current, r); err != nil {
		return nil, err
	}
	if err := service.RequestConvertTo(req, r); err != nil {
		return nil, err
	}
	return r, nil
}

// Copyright 2017-2022 The Usacloud Authors
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

package server

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/validate"
)

func Test_validateBootParameter(t *testing.T) {
	tests := []struct {
		name      string
		parameter *bootParameter
		wantErr   bool
	}{
		{
			name:      "minimum",
			parameter: &bootParameter{ZoneParameter: cflag.ZoneParameter{Zone: "is1a"}},
			wantErr:   false,
		},
		{
			name: "with valid cloud-config",
			parameter: &bootParameter{
				ZoneParameter: cflag.ZoneParameter{Zone: "is1a"},
				UserData: `
#cloud-config
hostname: foobar`,
			},
			wantErr: false,
		},
		{
			name: "with invalid cloud-config",
			parameter: &bootParameter{
				ZoneParameter: cflag.ZoneParameter{Zone: "is1a"},
				UserData:      "f,f",
			},
			wantErr: true,
		},
	}

	validate.InitializeValidator(sacloud.SakuraCloudZones)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateBootParameter(nil, tt.parameter); (err != nil) != tt.wantErr {
				t.Errorf("validateBootParameter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

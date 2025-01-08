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

package vpcrouter

import (
	"reflect"
	"testing"

	"github.com/sacloud/iaas-api-go"
)

func Test_routerSettingUpdate_Customize(t *testing.T) {
	tests := []struct {
		name     string
		settings routerSettingUpdate
		wantErr  bool
		want     routerSettingUpdate
	}{
		{
			name:     "empty",
			settings: routerSettingUpdate{},
			wantErr:  false,
			want:     routerSettingUpdate{},
		},
		{
			name: "with wireguard json",
			settings: routerSettingUpdate{
				WireGuardData: `{"IPAddress":"192.0.2.1/24","Peers":[{"IPAddress":"192.0.2.11","Name":"example","PublicKey":"your-key"}]}`,
			},
			wantErr: false,
			want: routerSettingUpdate{
				WireGuardData: `{"IPAddress":"192.0.2.1/24","Peers":[{"IPAddress":"192.0.2.11","Name":"example","PublicKey":"your-key"}]}`,
				WireGuard: &iaas.VPCRouterWireGuard{
					IPAddress: "192.0.2.1/24",
					Peers: []*iaas.VPCRouterWireGuardPeer{
						{
							Name:      "example",
							IPAddress: "192.0.2.11",
							PublicKey: "your-key",
						},
					},
				},
			},
		},
		{
			name: "with static nat json",
			settings: routerSettingUpdate{
				StaticNATData: `[{"GlobalAddress":"192.0.2.1","PrivateAddress":"192.0.2.2"}]`,
			},
			wantErr: false,
			want: routerSettingUpdate{
				StaticNATData: `[{"GlobalAddress":"192.0.2.1","PrivateAddress":"192.0.2.2"}]`,
				StaticNAT: &[]*iaas.VPCRouterStaticNAT{
					{
						GlobalAddress:  "192.0.2.1",
						PrivateAddress: "192.0.2.2",
						Description:    "",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.settings.Customize(nil); (err != nil) != tt.wantErr {
				t.Errorf("Customize() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.settings, tt.want) {
				t.Errorf("unexpected router setting state: got = %v, actual = %v", tt.settings, tt.want)
			}
		})
	}
}

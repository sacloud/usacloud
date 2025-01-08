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

package validate

import (
	"os"
	"testing"
)

func TestPublicKeyFormat(t *testing.T) {
	dummyPubContent, err := os.ReadFile("dummy.pub")
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		t *Target
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "target is nil",
			args:    args{},
			wantErr: false,
		},
		{
			name: "value is nil",
			args: args{
				t: &Target{
					FlagName: "--foobar",
					Value:    nil,
				},
			},
			wantErr: false,
		},
		{
			name: "value is empty",
			args: args{
				t: &Target{
					FlagName: "--foobar",
					Value:    "",
				},
			},
			wantErr: false,
		},
		{
			name: "value is invalid path",
			args: args{
				t: &Target{
					FlagName: "--foobar",
					Value:    "/invalid/file/path",
				},
			},
			wantErr: true,
		},
		{
			name: "value is invalid key content",
			args: args{
				t: &Target{
					FlagName: "--foobar",
					Value:    "invalid content",
				},
			},
			wantErr: true,
		},
		{
			name: "value has valid key",
			args: args{
				t: &Target{
					FlagName: "--foobar",
					Value:    "invalid content",
				},
			},
			wantErr: true,
		},
		{
			name: "value is valid path",
			args: args{
				t: &Target{
					FlagName: "--foobar",
					Value:    "dummy.pub",
				},
			},
			wantErr: false,
		},
		{
			name: "value is valid conrent",
			args: args{
				t: &Target{
					FlagName: "--foobar",
					Value:    string(dummyPubContent),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PublicKeyFormat(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("PublicKeyFormat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

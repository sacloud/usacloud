// Copyright 2017-2022 The sacloud/usacloud Authors
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

package query

import "testing"

func TestByGoJQ(t *testing.T) {
	type args struct {
		input   interface{}
		query   string
		printer func(interface{}) error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "with map",
			args: args{
				input: map[string]interface{}{
					"foo": "bar",
				},
				query:   ".",
				printer: func(interface{}) error { return nil },
			},
			wantErr: false,
		},
		{
			name: "with slice",
			args: args{
				input:   []interface{}{"1", "2"},
				query:   ".",
				printer: func(interface{}) error { return nil },
			},
			wantErr: false,
		},
		{
			name: "with struct",
			args: args{
				input:   struct{ Foo string }{Foo: "bar"},
				query:   ".",
				printer: func(interface{}) error { return nil },
			},
			wantErr: false,
		},
		{
			name: "with primitive",
			args: args{
				input:   1,
				query:   ".",
				printer: func(interface{}) error { return nil },
			},
			wantErr: false,
		},
		{
			name: "with invalid query",
			args: args{
				input:   map[string]interface{}{"foo": "bar"},
				query:   "...",
				printer: func(interface{}) error { return nil },
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ByGoJQ(tt.args.input, tt.args.query, tt.args.printer); (err != nil) != tt.wantErr {
				t.Errorf("ByGoJQ() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

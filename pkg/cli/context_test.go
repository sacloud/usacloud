// Copyright 2017-2021 The Usacloud Authors
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

package cli

import (
	"testing"

	"github.com/sacloud/usacloud/pkg/config"
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/stretchr/testify/require"
)

type dummyOption struct {
	outputType        string
	format            string
	quiet             bool
	query             string
	queryDriver       string
	defaultOutputType string
}

func (o *dummyOption) OutputTypeFlagValue() string { return o.outputType }
func (o *dummyOption) FormatFlagValue() string     { return o.format }
func (o *dummyOption) QuietFlagValue() bool        { return o.quiet }
func (o dummyOption) QueryFlagValue() string       { return o.query }
func (o dummyOption) QueryDriverFlagValue() string { return o.queryDriver }

func Test_getOutputWriter(t *testing.T) {
	io := newIO()

	type args struct {
		io           IO
		globalOption *config.Config
		columnDefs   []output.ColumnDef
		rawOptions   interface{}
	}
	tests := []struct {
		name string
		args args
		want output.Output
	}{
		{
			name: "nil options",
			args: args{
				io:           io,
				globalOption: &config.Config{},
				columnDefs:   nil,
				rawOptions:   nil,
			},
			want: output.NewDiscardOutput(),
		},
		{
			name: "no output options",
			args: args{
				io:           io,
				globalOption: &config.Config{},
				columnDefs:   nil,
				rawOptions:   &struct{}{},
			},
			want: output.NewDiscardOutput(),
		},
		{
			name: "empty options",
			args: args{
				io:           io,
				globalOption: &config.Config{},
				columnDefs:   nil,
				rawOptions:   &dummyOption{},
			},
			want: output.NewTableOutput(io.Out(), io.Err(), nil),
		},
		{
			name: "with --quiet option",
			args: args{
				io:           io,
				globalOption: &config.Config{},
				columnDefs:   nil,
				rawOptions: &dummyOption{
					quiet: true,
					// quietがtrueの場合、以下は無視される
					outputType:        "json",
					format:            "dummy",
					query:             "dummy",
					queryDriver:       "jmespath",
					defaultOutputType: "table",
				},
			},
			want: output.NewIDOutput(io.Out(), io.Err()),
		},
		{
			name: "with format option",
			args: args{
				io:           io,
				globalOption: &config.Config{},
				columnDefs:   nil,
				rawOptions: &dummyOption{
					format: "dummy",
					// quietがfalse、かつformatが空でない場合、以下は無視される
					outputType:        "json",
					query:             "dummy",
					queryDriver:       "jmespath",
					defaultOutputType: "table",
				},
			},
			want: output.NewFreeOutput(io.Out(), io.Err(), &dummyOption{
				format:            "dummy",
				outputType:        "json",
				query:             "dummy",
				queryDriver:       "jmespath",
				defaultOutputType: "table",
			}),
		},
		{
			name: "with query option",
			args: args{
				io:           io,
				globalOption: &config.Config{},
				columnDefs:   nil,
				rawOptions: &dummyOption{
					query:       "dummy",
					queryDriver: "jmespath",
					// quietがfalse && formatが空でない && queryが空でない場合、以下は無視される
					outputType:        "yaml",
					defaultOutputType: "table",
				},
			},
			want: output.NewJSONOutput(io.Out(), io.Err(), false, "dummy", "jmespath"),
		},
		{
			name: "with json output-type",
			args: args{
				io:           io,
				globalOption: &config.Config{},
				columnDefs:   nil,
				rawOptions:   &dummyOption{outputType: "json"},
			},
			want: output.NewJSONOutput(io.Out(), io.Err(), false, "", ""),
		},
		{
			name: "with yaml output-type",
			args: args{
				io:           io,
				globalOption: &config.Config{},
				columnDefs:   nil,
				rawOptions:   &dummyOption{outputType: "yaml"},
			},
			want: output.NewYAMLOutput(io.Out(), io.Err()),
		},
		{
			name: "with table output-type",
			args: args{
				io:           io,
				globalOption: &config.Config{},
				columnDefs: []output.ColumnDef{
					{Name: "col1"},
					{Name: "col2"},
				},
				rawOptions: &dummyOption{outputType: "table"},
			},
			want: output.NewTableOutput(io.Out(), io.Err(), []output.ColumnDef{
				{Name: "col1"},
				{Name: "col2"},
			}),
		},
		{
			name: "with DefaultOutputType and empty output-type",
			args: args{
				io:           io,
				globalOption: &config.Config{DefaultOutputType: "json"},
				columnDefs:   nil,
				rawOptions:   &dummyOption{},
			},
			want: output.NewJSONOutput(io.Out(), io.Err(), false, "", ""),
		},
		{
			name: "with both of DefaultOutputType and output-type",
			args: args{
				io:           io,
				globalOption: &config.Config{DefaultOutputType: "json"},
				columnDefs:   nil,
				rawOptions:   &dummyOption{outputType: "yaml"},
			},
			want: output.NewYAMLOutput(io.Out(), io.Err()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getOutputWriter(tt.args.io, tt.args.globalOption, tt.args.columnDefs, tt.args.rawOptions)
			require.EqualValues(t, tt.want, got, tt.name)
		})
	}
}

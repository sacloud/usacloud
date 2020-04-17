// Copyright 2017-2020 The Usacloud Authors
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

// +build !windows

package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type dummyOption struct {
	outputType        string
	column            []string
	format            string
	formatFile        string
	quiet             bool
	query             string
	queryFile         string
	defaultOutputType string
}

func (o *dummyOption) GetOutputType() string { return o.outputType }
func (o *dummyOption) GetColumn() []string   { return o.column }
func (o *dummyOption) GetFormat() string     { return o.format }
func (o *dummyOption) GetFormatFile() string { return o.formatFile }
func (o *dummyOption) GetQuiet() bool        { return o.quiet }
func (o dummyOption) GetQuery() string       { return o.query }
func (o dummyOption) GetQueryFile() string   { return o.queryFile }

func TestValidateOutputOption(t *testing.T) {

	expects := []struct {
		testName string
		option   *dummyOption
		expect   bool
	}{
		{
			testName: "Should get no error with default values",
			option:   &dummyOption{},
			expect:   true,
		},
		// outputType with format/format-file
		{
			testName: "Should get error when OutputType is csv and have format",
			option: &dummyOption{
				defaultOutputType: "table",
				outputType:        "csv",
				format:            "fuga",
			},
			expect: false,
		},
		// outputType with format/format-file

		{
			testName: "Should get error when OutputType is csv and have format-file",
			option: &dummyOption{
				defaultOutputType: "table",
				outputType:        "csv",
				formatFile:        "/etc/hosts",
			},
			expect: false,
		},

		{
			testName: "Should get error when OutputType is tsv and have format",
			option: &dummyOption{
				defaultOutputType: "table",
				outputType:        "tsv",
				format:            "fuga",
			},
			expect: false,
		},
		{
			testName: "Should get no error when DefaultOutputType is csv and have format",
			option: &dummyOption{
				defaultOutputType: "csv",
				outputType:        "csv",
				format:            "fuga",
			},
			expect: true,
		},
		{
			testName: "Should get no error when DefaultOutputType is csv and have format-file",
			option: &dummyOption{
				defaultOutputType: "csv",
				outputType:        "csv",
				formatFile:        "/etc/hosts",
			},
			expect: true,
		},
		{
			testName: "Should get error when OutputType is tsv and have format-file",
			option: &dummyOption{
				defaultOutputType: "table",
				outputType:        "tsv",
				formatFile:        "/etc/hosts",
			},
			expect: false,
		},
		{
			testName: "Should get no error when DefaultOutputType is tsv and have format",
			option: &dummyOption{
				defaultOutputType: "tsv",
				outputType:        "tsv",
				format:            "fuga",
			},
			expect: true,
		},
		{
			testName: "Should get no error when DefaultOutputType is tsv and have format-file",
			option: &dummyOption{
				defaultOutputType: "tsv",
				outputType:        "tsv",
				formatFile:        "/etc/hosts",
			},
			expect: true,
		},
		{
			testName: "Should get error when OutputType is json and have format",
			option: &dummyOption{
				defaultOutputType: "table",
				outputType:        "json",
				format:            "fuga",
			},
			expect: false,
		},
		{
			testName: "Should get error when OutputType is json and have format-file",
			option: &dummyOption{
				defaultOutputType: "table",
				outputType:        "json",
				formatFile:        "/etc/hosts",
			},
			expect: false,
		},
		// format and format-file
		{
			testName: "Should get error with format and format-file",
			option: &dummyOption{
				format:     "a",
				formatFile: "b",
			},
			expect: false,
		},
		{
			testName: "Should get no error when have format only",
			option: &dummyOption{
				outputType: "table",
				format:     "a",
			},
			expect: true,
		},
		{
			testName: "Should get no error when have format-file only",
			option: &dummyOption{
				outputType: "table",
				formatFile: "/etc/hosts",
			},
			expect: true,
		},
		// column and output-type
		{
			testName: "Should get error when have column and output-type is empty",
			option: &dummyOption{
				defaultOutputType: "table",
				column:            []string{"col1", "col2"},
			},
			expect: false,
		},
		{
			testName: "Should get no error when have column and output-type is csv",
			option: &dummyOption{
				outputType: "csv",
				column:     []string{"col1", "col2"},
			},
			expect: true,
		},
		{
			testName: "Should get no error when have column and output-type is tsv",
			option: &dummyOption{
				outputType: "tsv",
				column:     []string{"col1", "col2"},
			},
			expect: true,
		},
		{
			testName: "Should get error when have column and output-type is json",
			option: &dummyOption{
				outputType: "json",
				column:     []string{"col1", "col2"},
			},
			expect: false,
		},
		// column with format/format-file
		{
			testName: "Should get error when have both column and format",
			option: &dummyOption{
				outputType: "csv",
				column:     []string{"col1", "col2"},
				format:     "a",
			},
			expect: false,
		},
		{
			testName: "Should get error when have both column and format",
			option: &dummyOption{
				outputType: "csv",
				column:     []string{"col1", "col2"},
				formatFile: "/etc/hosts",
			},
			expect: false,
		},
		// quiet with output-type
		{
			testName: "Should get no error with same output-type both of default and param",
			option: &dummyOption{
				defaultOutputType: "csv",
				outputType:        "csv",
				quiet:             true,
			},
			expect: true,
		},
		{
			testName: "Should get error with output-type",
			option: &dummyOption{
				defaultOutputType: "csv",
				outputType:        "table",
				quiet:             false,
			},
			expect: true,
		},
		// quiet with format/format-file
		{
			testName: "Should get error with format and quiet",
			option: &dummyOption{
				format: "a",
				quiet:  true,
			},
			expect: false,
		},
		{
			testName: "Should get error with format-file and quiet",
			option: &dummyOption{
				formatFile: "/etc/hosts",
				quiet:      true,
			},
			expect: false,
		},
		// format-file
		{
			testName: "Should get error when format-file is not exists",
			option: &dummyOption{
				outputType: "table",
				formatFile: "dummy-not-exists-format-file",
			},
			expect: false,
		},
		{
			testName: "Should get no error when format-file is exists",
			option: &dummyOption{
				outputType: "table",
				formatFile: "/etc/hosts",
			},
			expect: true,
		},
		{
			testName: "Should get no error when have query and output-type is json",
			option: &dummyOption{
				outputType: "json",
				query:      "[].ID",
			},
			expect: true,
		},
		{
			testName: "Should get error when have query and output-type is not json",
			option: &dummyOption{
				outputType: "csv",
				query:      "[].ID",
			},
			expect: false,
		},
	}

	// do table-driven test
	for _, expect := range expects {
		// TODO グローバルオプションの扱いが確定したら修正する
		GlobalOption = &CLIOptions{}
		t.Run(expect.testName, func(t *testing.T) {
			if expect.option.defaultOutputType == "" {
				GlobalOption.DefaultOutputType = "table"
			} else {
				GlobalOption.DefaultOutputType = expect.option.defaultOutputType
			}

			res := ValidateOutputOption(expect.option)
			assert.Equal(t, expect.expect, len(res) == 0)
		})
	}
}

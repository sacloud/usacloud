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

//go:build !windows
// +build !windows

package cli

import (
	"testing"

	"github.com/sacloud/usacloud/pkg/config"
	"github.com/stretchr/testify/assert"
)

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
			testName: "Should get no error when have format only",
			option: &dummyOption{
				outputType: "table",
				format:     "a",
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
			testName: "Should get no error when have query and output-type is json",
			option: &dummyOption{
				outputType: "json",
				query:      "[].ID",
			},
			expect: true,
		},
	}

	// do table-driven test
	for _, expect := range expects {
		options := &config.Config{}
		t.Run(expect.testName, func(t *testing.T) {
			if expect.option.defaultOutputType == "" {
				options.DefaultOutputType = "table"
			} else {
				options.DefaultOutputType = expect.option.defaultOutputType
			}

			res := ValidateOutputOption(expect.option, options.DefaultOutputType)
			assert.Equal(t, expect.expect, len(res) == 0)
		})
	}
}

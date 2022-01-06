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

package cli

import (
	"fmt"

	"github.com/jmespath/go-jmespath"
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/util"
)

func ValidateSetProhibited(fieldName string, object interface{}) []error {
	if !util.IsEmpty(object) {
		return []error{fmt.Errorf("%q: can't set on current context", fieldName)}
	}
	return []error{}
}

func ValidateOutputOption(o output.Option, defaultOutputType string) []error {
	outputType := o.OutputTypeFlagValue()
	//columns := o.GetColumn()
	format := o.FormatFlagValue()
	quiet := o.QuietFlagValue()
	query := o.QueryFlagValue()

	// format(or format-file) with output-type
	if outputType != defaultOutputType {
		cannotWith := map[string]string{
			"--format": format,
		}
		for k, v := range cannotWith {
			if v != "" {
				return []error{fmt.Errorf("%q: can't set with --output-type", k)}
			}
		}
	}

	// with quiet
	if quiet {
		cannotWith := map[string]string{
			"--format": format,
		}
		for k, v := range cannotWith {
			if v != "" {
				return []error{fmt.Errorf("%q: can't set with %s", "--quiet", k)}
			}
		}

		// quiet with output-type
		if outputType != defaultOutputType {
			return []error{fmt.Errorf("%q: can't set with %s", "--quiet", outputType)}
		}
	}

	// query only allow when outputType is json
	if outputType != "json" && len(query) > 0 {
		return []error{fmt.Errorf("%q: can't set when --output-type is not json", "--query")}
	}

	if outputType == "json" && len(query) > 0 {
		_, err := jmespath.Compile(query)
		if err != nil {
			return []error{fmt.Errorf("%q: invalid JMESPath: %s", "--query", err)}
		}
	}

	if format != "" {
		_, err := util.StringFromPathOrContent(format)
		if err != nil {
			return []error{fmt.Errorf("%q: invalid format: %s", "--format", err)}
		}
	}

	if query != "" {
		_, err := util.StringFromPathOrContent(query)
		if err != nil {
			return []error{fmt.Errorf("%q: invalid JMESPath: %s", "--query", err)}
		}
	}

	return []error{}
}

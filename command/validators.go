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

package command

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/sacloud/go-jmespath"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func ValidateInStrValues(fieldName string, object interface{}, allowValues ...string) []error {
	return schema.ValidateInStrValues(allowValues...)(fieldName, object)
}

func ValidateRequired(fieldName string, object interface{}) []error {
	if IsEmpty(object) {
		return []error{fmt.Errorf("%q: is required", fieldName)}
	}
	return []error{}
}

func ValidateSakuraID(fieldName string, object interface{}) []error {
	return schema.ValidateSakuraID()(fieldName, object)
}

func ValidateSetProhibited(fieldName string, object interface{}) []error {
	if !IsEmpty(object) {
		return []error{fmt.Errorf("%q: can't set on current context", fieldName)}
	}
	return []error{}
}

func ValidateIPv4Address(fieldName string, object interface{}) []error {
	return schema.ValidateIPv4Address()(fieldName, object)
}

func ValidateIPv6Address(fieldName string, object interface{}) []error {
	return schema.ValidateIPv6Address()(fieldName, object)
}

func ValidateExistsFileOrStdIn(fieldName string, object interface{}) []error {
	return schema.ValidateMultiOr(schema.ValidateFileExists(), schema.ValidateStdinExists())(fieldName, object)
}

func ValidateConflicts(fieldName string, object interface{}, values map[string]interface{}) []error {

	if !IsEmpty(object) {
		for _, v := range values {
			if !IsEmpty(v) {
				keys := []string{}
				for k := range values {
					keys = append(keys, fmt.Sprintf("%q", k))
				}
				return []error{fmt.Errorf("%q: is conflict with %s", fieldName, strings.Join(keys, " or "))}
			}
		}
	}
	return []error{}

}

func ValidateConflictValues(fieldName string, object interface{}, values map[string]interface{}) []error {

	if !IsEmpty(object) {
		for _, v := range values {
			if !IsEmpty(v) {
				keys := []string{}
				for k := range values {
					keys = append(keys, fmt.Sprintf("%q", k))
				}
				return []error{fmt.Errorf("%q(%#v): is conflict with %s", fieldName, object, strings.Join(keys, " or "))}
			}
		}
	}
	return []error{}

}

func ValidateBetween(fieldName string, object interface{}, min int, max int) []error {

	if object == nil {
		object = []int64{}
	}

	isSlice := func(object interface{}) bool {
		_, ok1 := object.([]int64)
		_, ok2 := object.([]string)

		return ok1 || ok2
	}

	if isSlice(object) {
		sliceLen := 0
		if s, ok := object.([]int64); ok {
			sliceLen = len(s)
		} else {
			s := object.([]string)
			sliceLen = len(s)
		}

		if max <= 0 {
			if sliceLen < min {
				return []error{fmt.Errorf("%q: slice length must be %d or more", fieldName, min)}
			}
		} else {
			if !(min <= sliceLen && sliceLen <= max) {
				return []error{fmt.Errorf("%q: slice length must be beetween %d and %d", fieldName, min, max)}
			}

		}
	}

	return []error{}
}

func ValidateOutputOption(o output.Option) []error {

	defaultOutputType := GlobalOption.DefaultOutputType
	outputType := o.GetOutputType()
	columns := o.GetColumn()
	format := o.GetFormat()
	formatFile := o.GetFormatFile()
	quiet := o.GetQuiet()
	query := o.GetQuery()
	queryFile := o.GetQueryFile()

	// format and format-file
	if format != "" && formatFile != "" {
		return []error{fmt.Errorf("%q: can't set with --format-file", "--format")}
	}
	// query and query-file
	if query != "" && queryFile != "" {
		return []error{fmt.Errorf("%q: can't set with --query-file", "--query")}
	}

	// format(or format-file) with output-type
	if outputType != defaultOutputType {
		cannotWith := map[string]string{
			"--format":      format,
			"--format-file": formatFile,
		}
		for k, v := range cannotWith {
			if v != "" {
				return []error{fmt.Errorf("%q: can't set with --output-type", k)}
			}
		}
	}

	// format-file is exists?
	if formatFile != "" {
		errs := schema.ValidateFileExists()("--format-file", formatFile)
		if len(errs) > 0 {
			return errs
		}
	}

	// with quiet
	if quiet {
		cannotWith := map[string]string{
			"--format":      format,
			"--format-file": formatFile,
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

	// columns only allow when outputType is tsv/csv
	if outputType != "tsv" && outputType != "csv" && len(columns) > 0 {
		return []error{fmt.Errorf("%q: can't set when --output-type is csv/tsv", "--column")}
	}

	// query-file is exists?
	if queryFile != "" {
		errs := schema.ValidateFileExists()("--query-file", formatFile)
		if len(errs) > 0 {
			return errs
		}
	}
	// query only allow when outputType is json
	if outputType != "json" && len(query) > 0 {
		return []error{fmt.Errorf("%q: can't set when --output-type is not json", "--query")}
	}
	if outputType != "json" && len(queryFile) > 0 {
		return []error{fmt.Errorf("%q: can't set when --output-type is not json", "--query-file")}
	}

	if outputType == "json" && len(query) > 0 {
		_, err := jmespath.Compile(query)
		if err != nil {
			return []error{fmt.Errorf("%q: invalid JMESPath: %s", "--query", err)}
		}
	}

	if outputType == "json" && len(queryFile) > 0 {
		bQuery, err := ioutil.ReadFile(queryFile)
		if err != nil {
			return []error{fmt.Errorf("%q: can't open query file: %s", "--query-file", err)}
		}
		_, err = jmespath.Compile(string(bQuery))
		if err != nil {
			return []error{fmt.Errorf("%q: invalid JMESPath: %s", "--query-file", err)}
		}
	}

	return []error{}
}

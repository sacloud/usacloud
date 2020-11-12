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

package validate

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/sacloud/usacloud/pkg/naming"
)

var validate = validator.New()

func Exec(parameter interface{}) error {
	err := validate.Struct(parameter)
	if err != nil {
		// see https://github.com/go-playground/validator/blob/f6584a41c8acc5dfc0b62f7962811f5231c11530/_examples/simple/main.go#L59-L65
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, errorMessageFromErr(err))
		}
		return fmt.Errorf("validation error:\n%s", strings.Join(errorMessages, "\n"))
	}
	return nil
}

func errorMessageFromErr(err validator.FieldError) string {
	flagName := naming.ToCLIFlag(err.StructField())
	param := err.Param()
	detail := err.ActualTag()
	if param != "" {
		detail += "=" + param
	}
	msg := fmt.Sprintf("\t%s: %s", flagName, detail)
	return msg
}

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

package validate

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/sacloud/usacloud/pkg/naming"
	"github.com/sacloud/usacloud/tools/clitag"
)

var validate = validator.New()

func Exec(parameter interface{}) error {
	err := validate.Struct(parameter)
	if err != nil {
		// see https://github.com/go-playground/validator/blob/f6584a41c8acc5dfc0b62f7962811f5231c11530/_examples/simple/main.go#L59-L65
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		var errs []error
		for _, err := range err.(validator.ValidationErrors) {
			errs = append(errs, errorFromValidationErr(parameter, err))
		}
		return NewValidationError(errs...)
	}
	return nil
}

func NewValidationError(errs ...error) error {
	if len(errs) == 0 {
		return nil
	}

	var errStrings []string
	for _, err := range errs {
		errStrings = append(errStrings, "\t"+err.Error())
	}
	return fmt.Errorf("validation error:\n%s", strings.Join(errStrings, "\n"))
}

func NewFlagError(flagName, message string) error {
	return fmt.Errorf("%s: %s", flagName, message)
}

func errorFromValidationErr(parameter interface{}, err validator.FieldError) error {
	flagName := flagName(parameter, err)
	param := err.Param()
	detail := err.ActualTag()
	if param != "" {
		detail += "=" + param
	}

	// detailがvalidatorのタグ名だけの場合の対応をここで行う。
	switch detail {
	case "file":
		detail = fmt.Sprintf("invalid file path: %v", err.Value())
	}

	return NewFlagError(flagName, detail)
}

func flagName(parameter interface{}, fieldErr validator.FieldError) string {
	fields, err := clitag.Parse(parameter)
	if err != nil {
		panic(err) // 基本的にコード生成時にclitag.Parseを呼び出し済みなため発生し得ない
	}
	namespaces := strings.Split(fieldErr.StructNamespace(), ".")
	errFieldName := strings.Join(namespaces[1:], ".") // .で区切った先頭を除いたもの

	for _, f := range fields {
		if f.FieldName == errFieldName {
			return naming.ToCLIFlag(f.FlagName)
		}
	}
	return naming.ToCLIFlag(fieldErr.StructField())
}

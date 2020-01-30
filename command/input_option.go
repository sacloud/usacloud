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

	"github.com/sacloud/usacloud/schema"
)

type InputOption interface {
	GetParamTemplate() string
	GetParamTemplateFile() string
}

func ValidateInputOption(o InputOption) []error {

	t := o.GetParamTemplate()
	tf := o.GetParamTemplateFile()

	// tmpl and tmpl-file
	if t != "" && tf != "" {
		return []error{fmt.Errorf("%q: can't set with --param-template-file", "--param-template")}
	}

	if tf != "" {
		errs := schema.ValidateFileExists()("--param-template-file", tf)
		if len(errs) > 0 {
			return errs
		}
	}

	return []error{}

}

func GetParamTemplateValue(o InputOption) (string, error) {
	t := o.GetParamTemplate()
	tf := o.GetParamTemplateFile()

	if t == "" && tf == "" {
		return "", nil
	}

	if t != "" {
		return t, nil
	}
	b, err := ioutil.ReadFile(o.GetParamTemplateFile())
	if err != nil {
		return "", fmt.Errorf("Read ParameterTemplateFile[%s] is failed: %s", o.GetParamTemplateFile(), err)
	}
	return string(b), nil
}

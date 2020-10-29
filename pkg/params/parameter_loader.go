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

package params

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/sacloud/usacloud/pkg/schema"
)

type ParameterLoader interface {
	GetParameters() string
	GetParameterFile() string
}

func validateParameterOptions(o ParameterLoader) []error {
	p := o.GetParameters()
	pf := o.GetParameterFile()

	// tmpl and tmpl-file
	if p != "" && pf != "" {
		return []error{fmt.Errorf("%q: can't set with --parameters", "--parameter-file")}
	}

	if pf != "" {
		errs := schema.ValidateFileExists()("--param-template-file", pf)
		if len(errs) > 0 {
			return errs
		}
	}

	return []error{}
}

func loadParameters(input interface{}) error {
	o, ok := input.(ParameterLoader)
	if !ok {
		return nil
	}

	p := o.GetParameters()
	pf := o.GetParameterFile()

	if p == "" && pf == "" {
		return nil
	}

	strParameter := p
	if strParameter == "" {
		// Note: この段階ではファイルサイズのバリデーションは済んでいる前提
		b, err := ioutil.ReadFile(pf)
		if err != nil {
			return fmt.Errorf("reading parameters from %s is failed: %s", pf, err)
		}
		strParameter = string(b)
	}

	return json.Unmarshal([]byte(strParameter), o)
}

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

package server

import (
	"fmt"

	"github.com/ghodss/yaml"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/util"
	"github.com/sacloud/usacloud/pkg/validate"
)

var bootCommand = &core.Command{
	Name:     "boot",
	Aliases:  []string{"power-on"},
	Category: "power",
	Order:    10,

	ColumnDefs: defaultColumnDefs,

	SelectorType: core.SelectorTypeRequireMulti,

	ParameterInitializer: func() interface{} {
		return newBootParameter()
	},
	ValidateFunc: validateBootParameter,
}

type bootParameter struct {
	cflag.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`

	UserData string `mapconv:",omitempty,filters=path_or_content"`

	cflag.NoWaitParameter `cli:",squash" mapconv:",squash"`
}

func newBootParameter() *bootParameter {
	return &bootParameter{}
}

func init() {
	Resource.AddCommand(bootCommand)
}

func validateBootParameter(_ cli.Context, parameter interface{}) error {
	if err := validate.Exec(parameter); err != nil {
		return err
	}

	p, ok := parameter.(*bootParameter)
	if !ok {
		return fmt.Errorf("invalid parameter: %v", parameter)
	}

	var errs []error
	if p.UserData != "" {
		userData, err := util.BytesFromPathOrContent(p.UserData)
		if err != nil {
			errs = append(errs, err)
		}

		if len(userData) > 0 {
			v := make(map[string]interface{})
			err := yaml.Unmarshal(userData, &v)
			if err != nil {
				errs = append(errs, err)
			}
		}
	}
	return validate.NewValidationError(errs...)
}

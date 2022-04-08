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

package config

import (
	"fmt"

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/core"
)

var createCommand = &core.Command{
	Name:       "create",
	Category:   "basic",
	Order:      25,
	NoProgress: true,

	ParameterInitializer: func() interface{} {
		return newCreateParameter()
	},

	Func: createProfile,
	ValidateFunc: func(ctx cli.Context, parameter interface{}) error {
		p, ok := parameter.(*createParameter)
		if !ok {
			return fmt.Errorf("invalid parameter: %v", parameter)
		}
		if len(ctx.Args()) > 0 && p.Name == "" {
			p.Name = ctx.Args()[0]
		}
		if p.Name == "" {
			return fmt.Errorf("--name or name argument required")
		}

		config, err := getProfileConfigValue(p.Name)
		if err != nil {
			return err
		}
		if config != nil {
			return fmt.Errorf("profile %q already exists", p.Name)
		}
		return nil
	},
}

type createParameter struct {
	EditParameter `cli:",squash"`
}

func newCreateParameter() *createParameter {
	return &createParameter{}
}

func init() {
	Resource.AddCommand(createCommand)
}

func createProfile(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*createParameter)
	if !ok {
		return nil, fmt.Errorf("invalid parameter: %v", parameter)
	}
	return editProfile(ctx, &p.EditParameter)
}

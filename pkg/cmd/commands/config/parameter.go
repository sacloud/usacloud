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

	"github.com/sacloud/libsacloud/v2/sacloud/profile"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/validate"
)

type ProfileParameter struct {
	Name string `validate:"omitempty,profile_name"`
}

func (p *ProfileParameter) GetName() string {
	return p.Name
}

func (p *ProfileParameter) SetName(n string) {
	p.Name = n
}

type profileParameterHolder interface {
	GetName() string
	SetName(string)
}

func validateProfileParameter(ctx cli.Context, parameter interface{}) error {
	p, ok := parameter.(profileParameterHolder)
	if !ok {
		return fmt.Errorf("invalid parameter: %v", parameter)
	}
	if len(ctx.Args()) > 0 && p.GetName() == "" {
		p.SetName(ctx.Args()[0])
	}
	if p.GetName() == "" {
		current, err := profile.CurrentName()
		if err != nil {
			return err
		}
		p.SetName(current)
	}
	if err := validate.Exec(p); err != nil {
		return err
	}

	profiles, err := profile.List()
	if err != nil {
		return err
	}

	for _, name := range profiles {
		if name == p.GetName() {
			return nil
		}
	}
	return validate.NewValidationError(validate.NewFlagError("--name", fmt.Sprintf("profile %q not exists", p.GetName())))
}

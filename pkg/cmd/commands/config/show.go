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

package config

import (
	"encoding/json"
	"fmt"

	"github.com/sacloud/usacloud/pkg/config"

	"github.com/sacloud/libsacloud/v2/sacloud/profile"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/core"
)

var showCommand = &core.Command{
	Name:     "show",
	Category: "basic",
	Order:    16,

	ParameterInitializer: func() interface{} {
		return &showParameter{}
	},
	NoProgress: true,

	Func:                 showFunc,
	ValidateFunc:         validateProfileParameter,
	CustomCompletionFunc: profileCompletion,
}

func init() {
	Resource.AddCommand(showCommand)
}

type showParameter struct {
	ProfileParameter `cli:",squash"`
}

func showFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*showParameter)
	if !ok {
		return nil, fmt.Errorf("invalid parameter: %v", parameter)
	}

	var profileValue config.Config
	if err := profile.Load(p.Name, &profileValue); err != nil {
		return nil, err
	}

	data, err := json.MarshalIndent(profileValue, "", "  ")
	if err != nil {
		return nil, err
	}

	out := ctx.IO().Out()
	fmt.Fprintln(out, string(data)) // nolint
	return nil, nil
}

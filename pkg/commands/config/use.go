// Copyright 2017-2025 The sacloud/usacloud Authors
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

	"github.com/sacloud/api-client-go/profile"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/core"
)

var useCommand = &core.Command{
	Name:     "use",
	Category: "basic",
	Order:    17,

	ParameterInitializer: func() interface{} {
		return &useParameter{}
	},
	NoProgress: true,

	Func:                 useFunc,
	ValidateFunc:         validateProfileParameter,
	CustomCompletionFunc: profileCompletion,
}

func init() {
	Resource.AddCommand(useCommand)
}

type useParameter struct {
	ProfileParameter `cli:",squash"`
}

func useFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*useParameter)
	if !ok {
		return nil, fmt.Errorf("invalid parameter: %v", parameter)
	}
	return nil, profile.SetCurrentName(p.Name)
}

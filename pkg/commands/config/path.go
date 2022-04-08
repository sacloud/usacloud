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

	"github.com/sacloud/api-client-go/profile"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/core"
)

var pathCommand = &core.Command{
	Name:     "path",
	Category: "basic",
	Order:    19,

	ParameterInitializer: func() interface{} {
		return &pathParameter{}
	},
	NoProgress: true,

	Func:                 pathFunc,
	ValidateFunc:         validateProfileParameter,
	CustomCompletionFunc: profileCompletion,
}

func init() {
	Resource.AddCommand(pathCommand)
}

type pathParameter struct {
	ProfileParameter `cli:",squash"`
}

func pathFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*pathParameter)
	if !ok {
		return nil, fmt.Errorf("invalid parameter: %v", parameter)
	}

	profilePath, err := profile.ConfigFilePath(p.Name)
	if err != nil {
		return nil, err
	}
	out := ctx.IO().Out()
	fmt.Fprintln(out, profilePath) // nolint
	return nil, nil
}

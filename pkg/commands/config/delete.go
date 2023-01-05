// Copyright 2017-2023 The sacloud/usacloud Authors
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
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/core"
)

var deleteCommand = &core.Command{
	Name:     "delete",
	Category: "basic",
	Order:    30,

	ParameterInitializer: func() interface{} {
		return &deleteParameter{}
	},
	NoProgress: true,

	Func:                 deleteFunc,
	ValidateFunc:         validateProfileParameter,
	CustomCompletionFunc: profileCompletion,
}

func init() {
	Resource.AddCommand(deleteCommand)
}

type deleteParameter struct {
	ProfileParameter       `cli:",squash"`
	cflag.ConfirmParameter `cli:",squash"`
}

func deleteFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*deleteParameter)
	if !ok {
		return nil, fmt.Errorf("invalid parameter: %v", parameter)
	}

	return nil, profile.Remove(p.Name)
}

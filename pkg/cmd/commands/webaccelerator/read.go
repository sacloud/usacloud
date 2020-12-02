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

package webaccelerator

import (
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
)

var readCommand = &core.Command{
	Name:       "read",
	Aliases:    []string{"show"},
	Category:   "basic",
	Order:      30,
	NoProgress: true,

	ColumnDefs: defaultColumnDefs,

	SelectorType: core.SelectorTypeRequireMulti,

	ParameterInitializer: func() interface{} {
		return newReadParameter()
	},
	ListAllFunc: listAllFunc,
	Func:        readFunc,
}

type readParameter struct {
	cflag.IDParameter     `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter `cli:",squash" mapconv:"-"`
}

func newReadParameter() *readParameter {
	return &readParameter{}
}

func init() {
	Resource.AddCommand(readCommand)
}

func readFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*readParameter)
	if !ok {
		return nil, fmt.Errorf("got invalid parameter type: %#v", parameter)
	}
	webAccelOp := sacloud.NewWebAccelOp(ctx.Client())
	result, err := webAccelOp.Read(ctx, p.ID)
	if err != nil {
		return nil, err
	}
	return []interface{}{result}, nil
}

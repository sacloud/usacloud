//go:build !wasm
// +build !wasm

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

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/helper/wait"
	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/connect"
	"github.com/sacloud/usacloud/pkg/core"
)

var vncCommand = &core.Command{
	Name:         "vnc",
	Category:     "connect",
	Order:        20,
	SelectorType: core.SelectorTypeRequireSingle,
	NoProgress:   true,

	ParameterInitializer: func() interface{} {
		return newVNCParameter()
	},

	Func: vncFunc,
}

func init() {
	Resource.AddCommand(vncCommand)
}

func vncFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*vncParameter)
	if !ok {
		return nil, fmt.Errorf("invalid parameter: %v", parameter)
	}

	instance := ctx.Resource().(*iaas.Server)
	if !instance.InstanceStatus.IsUp() && p.WaitUntilReady {
		if _, err := wait.UntilServerIsUp(ctx, iaas.NewServerOp(ctx.Client().(iaas.APICaller)), p.Zone, types.StringID(p.ID)); err != nil {
			return nil, err
		}
	}

	vncInfo, err := iaas.NewServerOp(ctx.Client().(iaas.APICaller)).GetVNCProxy(ctx, p.Zone, types.StringID(p.ID))
	if err != nil {
		return nil, err
	}
	return nil, connect.StartDefaultVNCClient(vncInfo)
}

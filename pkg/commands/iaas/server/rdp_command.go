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
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/connect"
	"github.com/sacloud/usacloud/pkg/core"
)

var rdpCommand = &core.Command{
	Name:         "rdp",
	Aliases:      []string{"remote-desktop"},
	Category:     "connect",
	Order:        30,
	SelectorType: core.SelectorTypeRequireSingle,
	NoProgress:   true,

	ParameterInitializer: func() interface{} {
		return newRDPParameter()
	},

	Func: rdpFunc,
}

func init() {
	Resource.AddCommand(rdpCommand)
}

func rdpFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*rdpParameter)
	if !ok {
		return nil, fmt.Errorf("invalid parameter: %v", parameter)
	}

	instance := ctx.Resource().(*iaas.Server)
	if len(instance.Interfaces) == 0 {
		return nil, fmt.Errorf("server[%q] has no network interfaces", p.ID)
	}

	if !instance.InstanceStatus.IsUp() && p.WaitUntilReady {
		lastState, err := wait.UntilServerIsUp(ctx, iaas.NewServerOp(ctx.Client().(iaas.APICaller)), p.Zone, p.ID)
		if err != nil {
			return nil, err
		}
		instance = lastState
	}

	ip := instance.Interfaces[0].IPAddress
	if ip == "" {
		ip = instance.Interfaces[0].UserIPAddress
	}
	if ip == "" {
		return nil, fmt.Errorf("server[%q] has no valid ip addresses", p.ID)
	}

	opener := connect.RDPOpener{
		IPAddress: ip,
		User:      p.User,
		Port:      p.Port,
	}

	return nil, opener.StartDefaultClient()
}

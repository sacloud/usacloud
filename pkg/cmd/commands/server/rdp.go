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

package server

import (
	"fmt"

	"github.com/sacloud/libsacloud/v2/helper/wait"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/connect"
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

type rdpParameter struct {
	cflag.ZoneParameter   `cli:",squash" mapconv:",squash"`
	cflag.IDParameter     `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter `cli:",squash" mapconv:"-"`

	User           string
	Port           int
	WaitUntilReady bool `cli:",aliases=wait"`
}

func newRDPParameter() *rdpParameter {
	return &rdpParameter{
		User: "Administrator",
		Port: 3389,
	}
}

func init() {
	Resource.AddCommand(rdpCommand)
}

func rdpFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*rdpParameter)
	if !ok {
		return nil, fmt.Errorf("invalid parameter: %v", parameter)
	}

	instance := ctx.Resource().(*sacloud.Server)
	if len(instance.Interfaces) == 0 {
		return nil, fmt.Errorf("server[%q] has no network interfaces", p.ID)
	}

	if !instance.InstanceStatus.IsUp() && p.WaitUntilReady {
		lastState, err := wait.UntilServerIsUp(ctx, sacloud.NewServerOp(ctx.Client()), p.Zone, p.ID)
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

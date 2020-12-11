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
	"os/exec"
	"strings"

	"github.com/sacloud/libsacloud/v2/helper/query"
	"github.com/sacloud/libsacloud/v2/helper/service/server"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
)

var sshCommand = &core.Command{
	Name:         "ssh",
	Category:     "connect",
	Order:        10,
	SelectorType: core.SelectorTypeRequireSingle,
	NoProgress:   true,

	ParameterInitializer: func() interface{} {
		return newSSHParameter()
	},

	Func: sshFunc,
}

type sshParameter struct {
	cflag.ZoneParameter  `cli:",squash" mapconv:",squash"`
	cflag.IDParameter    `cli:",squash" mapconv:",squash"`
	cflag.InputParameter `cli:",squash" mapconv:"-"`

	Key      string `cli:",short=i" validate:"omitempty,file"`
	User     string `cli:",short=l"`
	Port     int    `cli:",short=p" validate:"required"`
	Password string `cli:",aliases=pass-phrase"`
}

func newSSHParameter() *sshParameter {
	return &sshParameter{
		Port: 22,
	}
}

func init() {
	Resource.AddCommand(sshCommand)
}

func sshFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*sshParameter)
	if !ok {
		return nil, fmt.Errorf("invalid parameter: %v", parameter)
	}

	svc := server.New(ctx.Client())
	instance, err := svc.ReadWithContext(ctx, &server.ReadRequest{
		Zone: p.Zone,
		ID:   p.ID,
	})
	if err != nil {
		return nil, err
	}

	if len(instance.Interfaces) == 0 {
		return nil, fmt.Errorf("server[%q] has no network interfaces", p.ID)
	}

	ip := instance.Interfaces[0].IPAddress
	if ip == "" {
		ip = instance.Interfaces[0].UserIPAddress
	}
	if ip == "" {
		return nil, fmt.Errorf("server[%q] has no valid ip addresses", p.ID)
	}

	user := p.User
	if user == "" {
		u, err := query.ServerDefaultUserName(ctx, p.Zone, query.NewServerSourceReader(ctx.Client()), p.ID)
		if err != nil {
			return nil, err
		}
		user = u
	}

	args := []string{fmt.Sprintf("%s@%s", user, ip)}
	if p.Key != "" {
		args = append(args, "-i", p.Key)
	}
	if p.Port != 22 {
		args = append(args, "-p", fmt.Sprintf("%d", p.Port))
	}
	if len(ctx.Args()) > 0 {
		args = append(args, ctx.Args()[1:]...)
	}

	fmt.Fprintf(ctx.IO().Err(), "connecting server...\n\tcommand: ssh %s \n", strings.Join(args, " ")) // nolint
	cmd := exec.Command("ssh", args...)
	cmd.Stdout = ctx.IO().Out()
	cmd.Stderr = ctx.IO().Err()
	cmd.Stdin = ctx.IO().In()
	return nil, cmd.Run()
}

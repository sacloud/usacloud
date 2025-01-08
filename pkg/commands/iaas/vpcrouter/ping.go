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

package vpcrouter

import (
	"fmt"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/core"
)

var pingCommand = &core.Command{
	Name:     "ping",
	Category: "other",
	Order:    30,

	SelectorType: core.SelectorTypeRequireSingle,
	NoProgress:   true,

	ParameterInitializer: func() interface{} {
		return newPingParameter()
	},

	Func: pingFunc,
}

type pingParameter struct {
	cflag.ZoneParameter   `cli:",squash" mapconv:",squash"`
	cflag.IDParameter     `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter `cli:",squash" mapconv:"-"`

	Destination string `cli:",aliases=dest,short=d,desc=destination address(IPv4)" validate:"required,ipv4"`
}

func newPingParameter() *pingParameter {
	return &pingParameter{}
}

func init() {
	Resource.AddCommand(pingCommand)
}

func pingFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*pingParameter)
	if !ok {
		return nil, fmt.Errorf("invalid parameter: %v", parameter)
	}

	vpcRouterOp := iaas.NewVPCRouterOp(ctx.Client().(iaas.APICaller))
	results, err := vpcRouterOp.Ping(ctx, p.Zone, types.StringID(p.ID), p.Destination)
	if err != nil {
		return nil, err
	}

	for _, result := range results.Result {
		fmt.Fprintln(ctx.IO().Out(), result)
	}
	return nil, nil
}

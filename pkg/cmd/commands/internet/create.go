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

package internet

import (
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/cmd/examples"
)

var createCommand = &core.Command{
	Name:     "create",
	Category: "basic",
	Order:    20,

	ColumnDefs: defaultColumnDefs,

	ParameterInitializer: func() interface{} {
		return newCreateParameter()
	},
}

type createParameter struct {
	cflag.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameParameter   `cli:",squash" mapconv:",squash"`
	cflag.DescParameter   `cli:",squash" mapconv:",squash"`
	cflag.TagsParameter   `cli:",squash" mapconv:",squash"`
	cflag.IconIDParameter `cli:",squash" mapconv:",squash"`

	NetworkMaskLen int `cli:"netmask,aliases=network-mask-len,options=internet_network_mask_len" validate:"required,internet_network_mask_len"`
	BandWidthMbps  int `cli:"band-width,aliases=band-width-mbps,options=internet_bandwidth" validate:"required,internet_bandwidth"`

	EnableIPv6            bool `cli:"enable-ipv6"`
	cflag.NoWaitParameter `cli:",squash" mapconv:",squash"`
	NotFoundRetry         int
}

func newCreateParameter() *createParameter {
	return &createParameter{
		NetworkMaskLen: 28,
		BandWidthMbps:  100,
		NotFoundRetry:  10,
	}
}

func init() {
	Resource.AddCommand(createCommand)
}

func (p *createParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &createParameter{
		ZoneParameter:   examples.Zones(ctx.Option().Zones),
		NameParameter:   examples.Name,
		DescParameter:   examples.Description,
		TagsParameter:   examples.Tags,
		IconIDParameter: examples.IconID,
		NetworkMaskLen:  28,
		BandWidthMbps:   100,
		EnableIPv6:      true,
		NoWaitParameter: cflag.NoWaitParameter{
			NoWait: false,
		},
		NotFoundRetry: 10,
	}
}

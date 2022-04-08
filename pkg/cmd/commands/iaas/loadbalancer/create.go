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

package loadbalancer

import (
	"fmt"
	"net/http"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/cmd/examples"
	"github.com/sacloud/usacloud/pkg/util"
	"github.com/sacloud/usacloud/pkg/validate"
)

var createCommand = &core.Command{
	Name:     "create",
	Category: "basic",
	Order:    20,

	ColumnDefs: defaultColumnDefs,

	ParameterInitializer: func() interface{} {
		return newCreateParameter()
	},

	ValidateFunc: validateCreateParameter,
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

	PlanID string `cli:"plan,options=loadbalancer_plan,category=plan,order=10" mapconv:",filters=loadbalancer_plan_to_value" validate:"required,loadbalancer_plan"`

	VRID           int      `cli:"vrid,category=network,order=10"`
	SwitchID       types.ID `cli:",category=network,order=20" validate:"required"`
	IPAddresses    []string `cli:"ip-address,aliases=ipaddress,category=network,order=30" validate:"required,min=1,max=2,dive,ipv4"`
	NetworkMaskLen int      `cli:"netmask,aliases=network-mask-len,category=network,order=40" validate:"required,min=1,max=32"`
	DefaultRoute   string   `cli:"gateway,aliases=default-route,category=network,order=50" validate:"omitempty,ipv4"`
	Port           int      `cli:",category=network,order=60" validate:"omitempty,min=1,max=65535"`

	VirtualIPAddressesData string                              `cli:"virtual-ip-addresses,aliases=vips,category=network,order=70" mapconv:"-" json:"-"`
	VirtualIPAddresses     iaas.LoadBalancerVirtualIPAddresses `cli:"-"`

	cflag.NoWaitParameter `cli:",squash" mapconv:",squash"`
}

func validateCreateParameter(ctx cli.Context, parameter interface{}) error {
	if err := validate.Exec(parameter); err != nil {
		return err
	}

	p, ok := parameter.(*createParameter)
	if !ok {
		return fmt.Errorf("invalid parameter: %v", parameter)
	}

	switch p.PlanID {
	case "standard":
		if len(p.IPAddresses) != 1 {
			return validate.NewValidationError(validate.NewFlagError("--ip-addresses", "for the standard plan, specify only one IP address"))
		}
	default:
		if len(p.IPAddresses) != 2 {
			return validate.NewValidationError(validate.NewFlagError("--ip-addresses", "for the highspec plan, specify two IP addresses"))
		}
	}
	return nil
}

func newCreateParameter() *createParameter {
	return &createParameter{
		PlanID: "standard",
	}
}

func init() {
	Resource.AddCommand(createCommand)
}

// Customize パラメータ変換処理
func (p *createParameter) Customize(_ cli.Context) error {
	if p.VirtualIPAddressesData != "" {
		var vips iaas.LoadBalancerVirtualIPAddresses
		if err := util.MarshalJSONFromPathOrContent(p.VirtualIPAddressesData, &vips); err != nil {
			return err
		}
		p.VirtualIPAddresses = append(p.VirtualIPAddresses, vips...)
	}

	return nil
}

func (p *createParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &createParameter{
		ZoneParameter:   examples.Zones(ctx.Option().Zones),
		NameParameter:   examples.Name,
		DescParameter:   examples.Description,
		TagsParameter:   examples.Tags,
		IconIDParameter: examples.IconID,
		PlanID:          examples.OptionsString("loadbalancer_plan"),
		VRID:            1,
		SwitchID:        examples.ID,
		IPAddresses:     examples.IPAddresses,
		NetworkMaskLen:  examples.NetworkMaskLen,
		DefaultRoute:    examples.DefaultRoute,
		Port:            80,
		VirtualIPAddresses: iaas.LoadBalancerVirtualIPAddresses{
			{
				VirtualIPAddress: examples.VirtualIPAddress,
				Port:             80,
				DelayLoop:        10,
				SorryServer:      "192.0.2.1",
				Description:      "example",
				Servers: iaas.LoadBalancerServers{
					{
						IPAddress: "192.0.2.101",
						Port:      80,
						Enabled:   true,
						HealthCheck: &iaas.LoadBalancerServerHealthCheck{
							Protocol:     types.ELoadBalancerHealthCheckProtocol(examples.OptionsString("loadbalancer_server_protocol")),
							Path:         "/",
							ResponseCode: http.StatusOK,
						},
					},
				},
			},
		},
		NoWaitParameter: cflag.NoWaitParameter{
			NoWait: false,
		},
	}
}

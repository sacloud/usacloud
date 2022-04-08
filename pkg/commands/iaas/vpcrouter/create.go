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

package vpcrouter

import (
	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/iaas-service-go/vpcrouter/builder"
	cflag2 "github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/examples"
	"github.com/sacloud/usacloud/pkg/util"
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
	cflag2.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag2.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag2.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag2.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag2.NameParameter   `cli:",squash" mapconv:",squash"`
	cflag2.DescParameter   `cli:",squash" mapconv:",squash"`
	cflag2.TagsParameter   `cli:",squash" mapconv:",squash"`
	cflag2.IconIDParameter `cli:",squash" mapconv:",squash"`

	Plan    string `cli:"plan,options=vpc_router_plan_premium,category=plan" mapconv:"PlanID,filters=vpc_router_plan_premium_to_value" validate:"required,vpc_router_plan_premium"`
	Version int    `validate:"required,oneof=1 2"`

	PublicNetworkInterface builder.PremiumNICSetting `cli:",category=network,order=10" mapconv:"NICSetting,omitempty"`

	PrivateNetworkInterfacesData string                                 `cli:"private-network-interfaces,category=network,order=20" mapconv:"-" json:"-"`
	PrivateNetworkInterfaces     []*builder.AdditionalPremiumNICSetting `cli:"-" mapconv:"AdditionalNICSettings"`

	RouterSetting routerSetting `cli:",squash" mapconv:",recursive"`

	BootAfterCreate        bool
	cflag2.NoWaitParameter `cli:",squash" mapconv:",squash"`
}

func newCreateParameter() *createParameter {
	return &createParameter{
		Plan:    "premium",
		Version: 2,
	}
}

func init() {
	Resource.AddCommand(createCommand)
}

// Customize パラメータ変換処理
func (p *createParameter) Customize(ctx cli.Context) error {
	if p.PrivateNetworkInterfacesData != "" {
		var nics []*builder.AdditionalPremiumNICSetting
		if err := util.MarshalJSONFromPathOrContent(p.PrivateNetworkInterfacesData, &nics); err != nil {
			return err
		}
		p.PrivateNetworkInterfaces = append(p.PrivateNetworkInterfaces, nics...)
	}
	return p.RouterSetting.Customize(ctx)
}

func (p *createParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &createParameter{
		ZoneParameter:   examples.Zones(ctx.Option().Zones),
		NameParameter:   examples.Name,
		DescParameter:   examples.Description,
		TagsParameter:   examples.Tags,
		IconIDParameter: examples.IconID,
		Plan:            examples.OptionsString("vpc_router_plan_premium"),
		Version:         2,
		PublicNetworkInterface: builder.PremiumNICSetting{
			SwitchID:         examples.ID,
			IPAddresses:      examples.IPAddresses,
			VirtualIPAddress: examples.VirtualIPAddress,
			IPAliases:        []string{"192.0.2.102"},
		},
		PrivateNetworkInterfaces: []*builder.AdditionalPremiumNICSetting{
			{
				SwitchID:         examples.ID,
				IPAddresses:      []string{"192.168.0.11", "192.168.0.12"},
				VirtualIPAddress: "192.168.0.1",
				NetworkMaskLen:   24,
				Index:            1,
			},
		},
		RouterSetting: routerSetting{
			VRID:                      1,
			InternetConnectionEnabled: true,
			StaticNAT: []*iaas.VPCRouterStaticNAT{
				{
					GlobalAddress:  examples.VirtualIPAddress,
					PrivateAddress: "192.168.0.1",
					Description:    "example",
				},
			},
			PortForwarding: []*iaas.VPCRouterPortForwarding{
				{
					Protocol:       types.EVPCRouterPortForwardingProtocol(examples.OptionsString("vpc_router_port_forwarding_protocol")),
					GlobalPort:     22,
					PrivateAddress: "192.168.0.11",
					PrivatePort:    22,
					Description:    "example",
				},
			},
			Firewall: []*iaas.VPCRouterFirewall{
				{
					Send: []*iaas.VPCRouterFirewallRule{
						{
							Protocol:           types.Protocol(examples.OptionsString("vpc_router_firewall_protocol")),
							SourceNetwork:      "192.0.2.1 | 192.0.2.0/24",
							SourcePort:         "1024 | 1024-2048",
							DestinationNetwork: "192.0.2.1 | 192.0.2.0/24",
							DestinationPort:    "1024 | 1024-2048",
							Action:             types.Action(examples.OptionsString("packetfilter_action")),
							Logging:            true,
							Description:        "example",
						},
					},
					Receive: []*iaas.VPCRouterFirewallRule{
						{
							Protocol:           types.Protocol(examples.OptionsString("vpc_router_firewall_protocol")),
							SourceNetwork:      "192.0.2.1 | 192.0.2.0/24",
							SourcePort:         "1024 | 1024-2048",
							DestinationNetwork: "192.0.2.1 | 192.0.2.0/24",
							DestinationPort:    "1024 | 1024-2048",
							Action:             types.Action(examples.OptionsString("packetfilter_action")),
							Logging:            true,
							Description:        "example",
						},
					},
					Index: 0,
				},
			},
			DHCPServer: []*iaas.VPCRouterDHCPServer{
				{
					Interface:  "eth1",
					RangeStart: "192.168.0.240",
					RangeStop:  "192.168.0.244",
					DNSServers: []string{"133.242.0.3", "133.242.0.4"},
				},
			},
			DHCPStaticMapping: []*iaas.VPCRouterDHCPStaticMapping{
				{
					MACAddress: "9C:A3:BA:xx:xx:xx",
					IPAddress:  "192.168.0.245",
				},
			},

			DNSForwarding: &iaas.VPCRouterDNSForwarding{
				Interface:  "eth1",
				DNSServers: []string{"133.242.0.3", "133.242.0.4"},
			},

			PPTPServer: &iaas.VPCRouterPPTPServer{
				RangeStart: "192.168.0.246",
				RangeStop:  "192.168.0.249",
			},
			L2TPIPsecServer: &iaas.VPCRouterL2TPIPsecServer{
				RangeStart:      "192.168.0.250",
				RangeStop:       "192.168.0.254",
				PreSharedSecret: "presharedsecret",
			},
			WireGuard: &iaas.VPCRouterWireGuard{
				IPAddress: "192.168.0.240/28",
				Peers: []*iaas.VPCRouterWireGuardPeer{
					{
						Name:      "client1",
						IPAddress: "192.168.0.242",
						PublicKey: "your-key",
					},
				},
			},
			RemoteAccessUsers: []*iaas.VPCRouterRemoteAccessUser{
				{
					UserName: "username",
					Password: "password",
				},
			},
			SiteToSiteIPsecVPN: []*iaas.VPCRouterSiteToSiteIPsecVPN{
				{
					Peer:            "192.0.2.1",
					PreSharedSecret: "presharedsecret",
					RemoteID:        "192.0.2.1",
					Routes:          []string{"10.0.0.0/8"},
					LocalPrefix:     []string{"192.168.0.0/24"},
				},
			},
			StaticRoute: []*iaas.VPCRouterStaticRoute{
				{
					Prefix:  "172.16.0.0/16",
					NextHop: "192.168.0.21",
				},
			},
			SyslogHost: "192.168.0.1",
		},
		BootAfterCreate: true,
		NoWaitParameter: cflag2.NoWaitParameter{
			NoWait: false,
		},
	}
}

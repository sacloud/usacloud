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

package vpcrouter

import (
	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/iaas-service-go/vpcrouter/builder"
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/examples"
	"github.com/sacloud/usacloud/pkg/util"
)

var createStandardCommand = &core.Command{
	Name:     "create-standard",
	Category: "basic",
	Order:    25,

	ColumnDefs: defaultColumnDefs,

	ParameterInitializer: func() interface{} {
		return newCreateStandardParameter()
	},
}

type createStandardParameter struct {
	cflag.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameParameter   `cli:",squash" mapconv:",squash"`
	cflag.DescParameter   `cli:",squash" mapconv:",squash"`
	cflag.TagsParameter   `cli:",squash" mapconv:",squash"`
	cflag.IconIDParameter `cli:",squash" mapconv:",squash"`

	Version int `validate:"required,oneof=1 2"`

	PrivateNetworkInterfacesData string                                  `cli:"private-network-interfaces" mapconv:"-"`
	PrivateNetworkInterfaces     []*builder.AdditionalStandardNICSetting `cli:"-" mapconv:"AdditionalNICSettings"`

	RouterSetting routerSetting `cli:",squash" mapconv:",omitempty"`

	BootAfterCreate       bool
	cflag.NoWaitParameter `cli:",squash" mapconv:",squash"`
}

func newCreateStandardParameter() *createStandardParameter {
	return &createStandardParameter{
		Version: 2,
	}
}

func init() {
	Resource.AddCommand(createStandardCommand)
}

// Customize パラメータ変換処理
func (p *createStandardParameter) Customize(ctx cli.Context) error {
	if p.PrivateNetworkInterfacesData != "" {
		var nics []*builder.AdditionalStandardNICSetting
		if err := util.MarshalJSONFromPathOrContent(p.PrivateNetworkInterfacesData, &nics); err != nil {
			return err
		}
		p.PrivateNetworkInterfaces = append(p.PrivateNetworkInterfaces, nics...)
	}
	return nil
}

func (p *createStandardParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &createStandardParameter{
		ZoneParameter:   examples.Zones(ctx.Option().Zones),
		NameParameter:   examples.Name,
		DescParameter:   examples.Description,
		TagsParameter:   examples.Tags,
		IconIDParameter: examples.IconID,
		Version:         2,
		PrivateNetworkInterfaces: []*builder.AdditionalStandardNICSetting{
			{
				SwitchID:       examples.ID,
				IPAddress:      "192.168.0.11",
				NetworkMaskLen: 24,
				Index:          1,
			},
		},
		RouterSetting: routerSetting{
			InternetConnectionEnabled: true,
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
			SiteToSiteIPsecVPN: &iaas.VPCRouterSiteToSiteIPsecVPN{
				Config: []*iaas.VPCRouterSiteToSiteIPsecVPNConfig{
					{
						Peer:            "192.0.2.1",
						PreSharedSecret: "presharedsecret",
						RemoteID:        "192.0.2.1",
						Routes:          []string{"10.0.0.0/8"},
						LocalPrefix:     []string{"192.168.0.0/24"},
					},
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
		NoWaitParameter: cflag.NoWaitParameter{
			NoWait: false,
		},
	}
}

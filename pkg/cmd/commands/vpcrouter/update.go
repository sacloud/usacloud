// Copyright 2017-2021 The Usacloud Authors
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
	"github.com/sacloud/libsacloud/v2/helper/service/vpcrouter"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/pointer"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/cmd/examples"
	"github.com/sacloud/usacloud/pkg/util"
)

var updateCommand = &core.Command{
	Name:         "update",
	Category:     "basic",
	Order:        40,
	SelectorType: core.SelectorTypeRequireMulti,

	ColumnDefs: defaultColumnDefs,

	ParameterInitializer: func() interface{} {
		return newUpdateParameter()
	},
}

type updateParameter struct {
	cflag.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.DescUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.TagsUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.IconIDUpdateParameter `cli:",squash" mapconv:",omitempty,squash"`

	PublicNetworkInterface vpcrouter.PremiumNICSettingUpdate `cli:",category=network,order=10" mapconv:",omitempty"`

	PrivateNetworkInterfacesData string                                          `cli:"private-network-interfaces,category=network,order=20" mapconv:"-" json:"-"`
	PrivateNetworkInterfaces     *[]*vpcrouter.AdditionalPremiumNICSettingUpdate `cli:"-" mapconv:"AdditionalNICSettings"`

	RouterSetting routerSettingUpdate `cli:",squash" mapconv:",omitempty,recursive"`

	cflag.NoWaitParameter `cli:",squash" mapconv:",squash"`
}

func newUpdateParameter() *updateParameter {
	return &updateParameter{}
}

func init() {
	Resource.AddCommand(updateCommand)
}

// Customize パラメータ変換処理
func (p *updateParameter) Customize(ctx cli.Context) error {
	if p.PrivateNetworkInterfacesData != "" {
		var nics []*vpcrouter.AdditionalPremiumNICSettingUpdate
		if err := util.MarshalJSONFromPathOrContent(p.PrivateNetworkInterfacesData, &nics); err != nil {
			return err
		}
		if p.PrivateNetworkInterfaces == nil {
			p.PrivateNetworkInterfaces = &[]*vpcrouter.AdditionalPremiumNICSettingUpdate{}
		}
		*p.PrivateNetworkInterfaces = append(*p.PrivateNetworkInterfaces, nics...)
	}
	return p.RouterSetting.Customize(ctx)
}

func (p *updateParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &updateParameter{
		ZoneParameter:         examples.Zones(ctx.Option().Zones),
		NameUpdateParameter:   examples.NameUpdate,
		DescUpdateParameter:   examples.DescriptionUpdate,
		TagsUpdateParameter:   examples.TagsUpdate,
		IconIDUpdateParameter: examples.IconIDUpdate,
		PublicNetworkInterface: vpcrouter.PremiumNICSettingUpdate{
			IPAddresses:      &examples.IPAddresses,
			VirtualIPAddress: &examples.VirtualIPAddress,
			IPAliases:        pointer.NewStringSlice([]string{"192.0.2.102"}),
		},
		PrivateNetworkInterfaces: &[]*vpcrouter.AdditionalPremiumNICSettingUpdate{
			{
				SwitchID:         &examples.ID,
				IPAddresses:      pointer.NewStringSlice([]string{"192.168.0.11", "192.168.0.12"}),
				VirtualIPAddress: pointer.NewString("192.168.0.1"),
				NetworkMaskLen:   pointer.NewInt(24),
				Index:            1,
			},
		},
		RouterSetting: routerSettingUpdate{
			InternetConnectionEnabled: pointer.NewBool(true),
			StaticNAT: &[]*sacloud.VPCRouterStaticNAT{
				{
					GlobalAddress:  examples.VirtualIPAddress,
					PrivateAddress: "192.168.0.1",
					Description:    "example",
				},
			},
			PortForwarding: &[]*sacloud.VPCRouterPortForwarding{
				{
					Protocol:       types.EVPCRouterPortForwardingProtocol(examples.OptionsString("vpc_router_port_forwarding_protocol")),
					GlobalPort:     22,
					PrivateAddress: "192.168.0.11",
					PrivatePort:    22,
					Description:    "example",
				},
			},
			Firewall: &[]*sacloud.VPCRouterFirewall{
				{
					Send: []*sacloud.VPCRouterFirewallRule{
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
					Receive: []*sacloud.VPCRouterFirewallRule{
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
			DHCPServer: &[]*sacloud.VPCRouterDHCPServer{
				{
					Interface:  "eth1",
					RangeStart: "192.168.0.240",
					RangeStop:  "192.168.0.244",
					DNSServers: []string{"133.242.0.3", "133.242.0.4"},
				},
			},
			DHCPStaticMapping: &[]*sacloud.VPCRouterDHCPStaticMapping{
				{
					MACAddress: "9C:A3:BA:xx:xx:xx",
					IPAddress:  "192.168.0.245",
				},
			},
			DNSForwarding: &sacloud.VPCRouterDNSForwarding{
				Interface:  "eth1",
				DNSServers: []string{"133.242.0.3", "133.242.0.4"},
			},
			PPTPServer: &sacloud.VPCRouterPPTPServer{
				RangeStart: "192.168.0.246",
				RangeStop:  "192.168.0.249",
			},
			L2TPIPsecServer: &sacloud.VPCRouterL2TPIPsecServer{
				RangeStart:      "192.168.0.250",
				RangeStop:       "192.168.0.254",
				PreSharedSecret: "presharedsecret",
			},
			WireGuard: &sacloud.VPCRouterWireGuard{
				IPAddress: "192.168.0.240/28",
				Peers: []*sacloud.VPCRouterWireGuardPeer{
					{
						Name:      "client1",
						IPAddress: "192.168.0.242",
						PublicKey: "your-key",
					},
				},
			},
			RemoteAccessUsers: &[]*sacloud.VPCRouterRemoteAccessUser{
				{
					UserName: "username",
					Password: "password",
				},
			},
			SiteToSiteIPsecVPN: &[]*sacloud.VPCRouterSiteToSiteIPsecVPN{
				{
					Peer:            "192.0.2.1",
					PreSharedSecret: "presharedsecret",
					RemoteID:        "192.0.2.1",
					Routes:          []string{"10.0.0.0/8"},
					LocalPrefix:     []string{"192.168.0.0/24"},
				},
			},
			StaticRoute: &[]*sacloud.VPCRouterStaticRoute{
				{
					Prefix:  "172.16.0.0/16",
					NextHop: "192.168.0.21",
				},
			},
			SyslogHost: pointer.NewString("192.168.0.1"),
		},
		NoWaitParameter: cflag.NoWaitParameter{
			NoWait: false,
		},
	}
}

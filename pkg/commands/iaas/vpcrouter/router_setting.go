// Copyright 2017-2022 The sacloud/usacloud Authors
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
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/util"
)

type routerSetting struct {
	VRID                      int `json:",omitempty"`
	InternetConnectionEnabled bool

	StaticNATData string                     `cli:"static-nat" mapconv:"-" json:"-"`
	StaticNAT     []*iaas.VPCRouterStaticNAT `cli:"-" json:",omitempty"`

	PortForwardingData string                          `cli:"port-forwarding" mapconv:"-" json:"-"`
	PortForwarding     []*iaas.VPCRouterPortForwarding `cli:"-" json:",omitempty"`

	FirewallData string                    `cli:"firewall" mapconv:"-" json:"-"`
	Firewall     []*iaas.VPCRouterFirewall `cli:"-" json:",omitempty"`

	DHCPServerData string                      `cli:"dhcp-server" mapconv:"-" json:"-"`
	DHCPServer     []*iaas.VPCRouterDHCPServer `cli:"-" json:",omitempty"`

	DHCPStaticMappingData string                             `cli:"dhcp-static-mapping" mapconv:"-" json:"-"`
	DHCPStaticMapping     []*iaas.VPCRouterDHCPStaticMapping `cli:"-" json:",omitempty"`

	DNSForwardingData string                       `cli:"dns-forwarding" mapconv:"-" json:"-"`
	DNSForwarding     *iaas.VPCRouterDNSForwarding `cli:"-" json:",omitempty"`

	PPTPServerData string                    `cli:"pptp" mapconv:"-" json:"-"`
	PPTPServer     *iaas.VPCRouterPPTPServer `cli:"-" json:",omitempty"`

	L2TPIPsecServerData string                         `cli:"l2tp" mapconv:"-" json:"-"`
	L2TPIPsecServer     *iaas.VPCRouterL2TPIPsecServer `cli:"-" json:",omitempty"`

	RemoteAccessUsersData string                            `cli:"users" mapconv:"-" json:"-"`
	RemoteAccessUsers     []*iaas.VPCRouterRemoteAccessUser `cli:"-" json:",omitempty"`

	WireGuardData string                   `cli:"wireguard" mapconv:"-" json:"-"`
	WireGuard     *iaas.VPCRouterWireGuard `cli:"-" json:",omitempty"`

	SiteToSiteIPsecVPNData string                            `cli:"site-to-site-vpn" mapconv:"-" json:"-"`
	SiteToSiteIPsecVPN     *iaas.VPCRouterSiteToSiteIPsecVPN `cli:"-" json:",omitempty"`

	StaticRouteData string                       `cli:"static-route" mapconv:"-" json:"-"`
	StaticRoute     []*iaas.VPCRouterStaticRoute `cli:"-" json:",omitempty"`

	SyslogHost string `json:",omitempty"`
}

func (r *routerSetting) Customize(_ cli.Context) error {
	if r.StaticNATData != "" {
		var staticNat []*iaas.VPCRouterStaticNAT
		if err := util.MarshalJSONFromPathOrContent(r.StaticNATData, &staticNat); err != nil {
			return err
		}
		r.StaticNAT = append(r.StaticNAT, staticNat...)
	}

	if r.PortForwardingData != "" {
		var portForwarding []*iaas.VPCRouterPortForwarding
		if err := util.MarshalJSONFromPathOrContent(r.PortForwardingData, &portForwarding); err != nil {
			return err
		}
		r.PortForwarding = append(r.PortForwarding, portForwarding...)
	}

	if r.FirewallData != "" {
		var firewall []*iaas.VPCRouterFirewall
		if err := util.MarshalJSONFromPathOrContent(r.FirewallData, &firewall); err != nil {
			return err
		}
		r.Firewall = append(r.Firewall, firewall...)
	}

	if r.DHCPServerData != "" {
		var dhcpServer []*iaas.VPCRouterDHCPServer
		if err := util.MarshalJSONFromPathOrContent(r.DHCPServerData, &dhcpServer); err != nil {
			return err
		}
		r.DHCPServer = append(r.DHCPServer, dhcpServer...)
	}

	if r.DHCPStaticMappingData != "" {
		var dhcpStaticMapping []*iaas.VPCRouterDHCPStaticMapping
		if err := util.MarshalJSONFromPathOrContent(r.DHCPStaticMappingData, &dhcpStaticMapping); err != nil {
			return err
		}
		r.DHCPStaticMapping = append(r.DHCPStaticMapping, dhcpStaticMapping...)
	}
	if r.DNSForwardingData != "" {
		var df iaas.VPCRouterDNSForwarding
		if err := util.MarshalJSONFromPathOrContent(r.DNSForwardingData, &df); err != nil {
			return err
		}
		r.DNSForwarding = &df
	}

	if r.PPTPServerData != "" {
		var pptp iaas.VPCRouterPPTPServer
		if err := util.MarshalJSONFromPathOrContent(r.PPTPServerData, &pptp); err != nil {
			return err
		}
		r.PPTPServer = &pptp
	}
	if r.L2TPIPsecServerData != "" {
		var l2tp iaas.VPCRouterL2TPIPsecServer
		if err := util.MarshalJSONFromPathOrContent(r.L2TPIPsecServerData, &l2tp); err != nil {
			return err
		}
		r.L2TPIPsecServer = &l2tp
	}
	if r.WireGuardData != "" {
		var wireGuard iaas.VPCRouterWireGuard
		if err := util.MarshalJSONFromPathOrContent(r.WireGuardData, &wireGuard); err != nil {
			return err
		}
		r.WireGuard = &wireGuard
	}

	if r.RemoteAccessUsersData != "" {
		var users []*iaas.VPCRouterRemoteAccessUser
		if err := util.MarshalJSONFromPathOrContent(r.RemoteAccessUsersData, &users); err != nil {
			return err
		}
		r.RemoteAccessUsers = append(r.RemoteAccessUsers, users...)
	}

	if r.SiteToSiteIPsecVPNData != "" {
		var s2s *iaas.VPCRouterSiteToSiteIPsecVPN
		if err := util.MarshalJSONFromPathOrContent(r.SiteToSiteIPsecVPNData, &s2s); err != nil {
			return err
		}
		r.SiteToSiteIPsecVPN = s2s
	}

	if r.StaticRouteData != "" {
		var staticRoutes []*iaas.VPCRouterStaticRoute
		if err := util.MarshalJSONFromPathOrContent(r.StaticRouteData, &staticRoutes); err != nil {
			return err
		}
		r.StaticRoute = append(r.StaticRoute, staticRoutes...)
	}

	return nil
}

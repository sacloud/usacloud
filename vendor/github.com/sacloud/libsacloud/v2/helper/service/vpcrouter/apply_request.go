// Copyright 2016-2020 The Libsacloud Authors
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
	"github.com/sacloud/libsacloud/v2/helper/builder"
	vpcRouterBuilder "github.com/sacloud/libsacloud/v2/helper/builder/vpcrouter"
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// ApplyRequest Applyサービスへのパラメータ
//
// NOTE: helper/builderを統合するまでの経過措置として、実際の処理をここで実装せずにhelper/builder/vpcrouter#Builderへ移譲する
// 利用者がbuilderに依存することをへの対応としてNICSettingHolderなどの型はここで実装する
type ApplyRequest struct {
	Zone string   `request:"-" validate:"required"`
	ID   types.ID `request:"-"`

	Name        string `validate:"required"`
	Description string `validate:"min=0,max=512"`
	Tags        types.Tags
	IconID      types.ID

	PlanID                types.ID                     `validate:"required"`
	NICSetting            NICSettingHolder             // StandardNICSetting または PremiumNICSetting を指定する
	AdditionalNICSettings []AdditionalNICSettingHolder // AdditionalStandardNICSetting または AdditionalPremiumNICSetting を指定する
	RouterSetting         *RouterSetting
	NoWait                bool
	BootAfterCreate       bool
}

func (req *ApplyRequest) Validate() error {
	return validate.Struct(req)
}

// RouterSetting VPCルータの設定
type RouterSetting struct {
	VRID                      int
	InternetConnectionEnabled types.StringFlag
	StaticNAT                 []*sacloud.VPCRouterStaticNAT
	PortForwarding            []*sacloud.VPCRouterPortForwarding
	Firewall                  []*sacloud.VPCRouterFirewall
	DHCPServer                []*sacloud.VPCRouterDHCPServer
	DHCPStaticMapping         []*sacloud.VPCRouterDHCPStaticMapping
	PPTPServer                *sacloud.VPCRouterPPTPServer
	L2TPIPsecServer           *sacloud.VPCRouterL2TPIPsecServer
	RemoteAccessUsers         []*sacloud.VPCRouterRemoteAccessUser
	SiteToSiteIPsecVPN        []*sacloud.VPCRouterSiteToSiteIPsecVPN
	StaticRoute               []*sacloud.VPCRouterStaticRoute
	SyslogHost                string
}

func (req *ApplyRequest) Builder(caller sacloud.APICaller) *vpcRouterBuilder.Builder {
	return &vpcRouterBuilder.Builder{
		Name:                  req.Name,
		Description:           req.Description,
		Tags:                  req.Tags,
		IconID:                req.IconID,
		PlanID:                req.PlanID,
		NICSetting:            req.nicSetting(),
		AdditionalNICSettings: req.additionalNICSetting(),
		RouterSetting:         req.routerSetting(),
		NoWait:                req.NoWait,
		Client:                sacloud.NewVPCRouterOp(caller),
		SetupOptions: &builder.RetryableSetupParameter{
			BootAfterBuild: req.BootAfterCreate,
		},
	}
}

func (req *ApplyRequest) routerSetting() *vpcRouterBuilder.RouterSetting {
	if req.RouterSetting == nil {
		return nil
	}

	return &vpcRouterBuilder.RouterSetting{
		VRID:                      req.RouterSetting.VRID,
		InternetConnectionEnabled: req.RouterSetting.InternetConnectionEnabled,
		StaticNAT:                 req.RouterSetting.StaticNAT,
		PortForwarding:            req.RouterSetting.PortForwarding,
		Firewall:                  req.RouterSetting.Firewall,
		DHCPServer:                req.RouterSetting.DHCPServer,
		DHCPStaticMapping:         req.RouterSetting.DHCPStaticMapping,
		PPTPServer:                req.RouterSetting.PPTPServer,
		L2TPIPsecServer:           req.RouterSetting.L2TPIPsecServer,
		RemoteAccessUsers:         req.RouterSetting.RemoteAccessUsers,
		SiteToSiteIPsecVPN:        req.RouterSetting.SiteToSiteIPsecVPN,
		StaticRoute:               req.RouterSetting.StaticRoute,
		SyslogHost:                req.RouterSetting.SyslogHost,
	}
}

func (req *ApplyRequest) nicSetting() vpcRouterBuilder.NICSettingHolder {
	switch v := req.NICSetting.(type) {
	case *StandardNICSetting:
		return &vpcRouterBuilder.StandardNICSetting{}
	case *PremiumNICSetting:
		return &vpcRouterBuilder.PremiumNICSetting{
			SwitchID:         v.SwitchID,
			IPAddresses:      v.IPAddresses,
			VirtualIPAddress: v.VirtualIPAddress,
			IPAliases:        v.IPAliases,
		}
	default:
		return nil
	}
}

func (req *ApplyRequest) additionalNICSetting() []vpcRouterBuilder.AdditionalNICSettingHolder {
	var settings []vpcRouterBuilder.AdditionalNICSettingHolder
	for _, s := range req.AdditionalNICSettings {
		switch v := s.(type) {
		case *AdditionalStandardNICSetting:
			settings = append(settings, &vpcRouterBuilder.AdditionalStandardNICSetting{
				SwitchID:       v.SwitchID,
				IPAddress:      v.IPAddress,
				NetworkMaskLen: v.NetworkMaskLen,
				Index:          v.Index,
			})
		case *AdditionalPremiumNICSetting:
			settings = append(settings, &vpcRouterBuilder.AdditionalPremiumNICSetting{
				SwitchID:         v.SwitchID,
				IPAddresses:      v.IPAddresses,
				VirtualIPAddress: v.VirtualIPAddress,
				NetworkMaskLen:   v.NetworkMaskLen,
				Index:            v.Index,
			})
		default:
			continue
		}
	}
	return settings
}

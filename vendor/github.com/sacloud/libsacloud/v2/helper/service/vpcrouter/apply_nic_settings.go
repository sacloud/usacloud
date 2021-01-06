// Copyright 2016-2021 The Libsacloud Authors
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
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// NICSettingHolder VPCルータのeth0の設定 SharedNICSettingまたはRouterNICSettingを指定する
type NICSettingHolder interface {
	connectedSwitch() *sacloud.ApplianceConnectedSwitch
	ipAddresses() []string
	interfaceSetting() *sacloud.VPCRouterInterfaceSetting
}

// StandardNICSetting VPCルータのeth0を共有セグメントに接続するためのSetting(スタンダードプラン)
type StandardNICSetting struct{}

func (s *StandardNICSetting) connectedSwitch() *sacloud.ApplianceConnectedSwitch {
	return &sacloud.ApplianceConnectedSwitch{Scope: types.Scopes.Shared}
}

func (s *StandardNICSetting) ipAddresses() []string {
	return nil
}

func (s *StandardNICSetting) interfaceSetting() *sacloud.VPCRouterInterfaceSetting {
	return nil
}

// PremiumNICSetting VPCルータのeth0をスイッチ+ルータに接続するためのSetting(プレミアム/ハイスペックプラン)
type PremiumNICSetting struct {
	SwitchID         types.ID
	IPAddresses      []string
	VirtualIPAddress string
	IPAliases        []string
}

func (s *PremiumNICSetting) connectedSwitch() *sacloud.ApplianceConnectedSwitch {
	return &sacloud.ApplianceConnectedSwitch{ID: s.SwitchID}
}

func (s *PremiumNICSetting) ipAddresses() []string {
	return s.IPAddresses
}

func (s *PremiumNICSetting) interfaceSetting() *sacloud.VPCRouterInterfaceSetting {
	return &sacloud.VPCRouterInterfaceSetting{
		IPAddress:        s.IPAddresses,
		VirtualIPAddress: s.VirtualIPAddress,
		IPAliases:        s.IPAliases,
		Index:            0,
	}
}

// AdditionalNICSettingHolder VPCルータのeth1-eth7の設定
type AdditionalNICSettingHolder interface {
	switchInfo() (switchID types.ID, index int)
	interfaceSetting() *sacloud.VPCRouterInterfaceSetting
}

// AdditionalStandardNICSetting VPCルータのeth1-eth7の設定(スタンダードプラン向け)
type AdditionalStandardNICSetting struct {
	SwitchID       types.ID
	IPAddress      string
	NetworkMaskLen int
	Index          int
}

func (s *AdditionalStandardNICSetting) switchInfo() (switchID types.ID, index int) {
	return s.SwitchID, s.Index
}

func (s *AdditionalStandardNICSetting) interfaceSetting() *sacloud.VPCRouterInterfaceSetting {
	return &sacloud.VPCRouterInterfaceSetting{
		IPAddress:      []string{s.IPAddress},
		NetworkMaskLen: s.NetworkMaskLen,
		Index:          s.Index,
	}
}

// AdditionalPremiumNICSetting VPCルータのeth1-eth7の設定(プレミアム/ハイスペックプラン向け)
type AdditionalPremiumNICSetting struct {
	SwitchID         types.ID
	IPAddresses      []string
	VirtualIPAddress string
	NetworkMaskLen   int
	Index            int
}

func (s *AdditionalPremiumNICSetting) switchInfo() (switchID types.ID, index int) {
	return s.SwitchID, s.Index
}

func (s *AdditionalPremiumNICSetting) interfaceSetting() *sacloud.VPCRouterInterfaceSetting {
	return &sacloud.VPCRouterInterfaceSetting{
		IPAddress:        s.IPAddresses,
		VirtualIPAddress: s.VirtualIPAddress,
		NetworkMaskLen:   s.NetworkMaskLen,
		Index:            s.Index,
	}
}

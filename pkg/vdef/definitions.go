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

package vdef

import (
	"fmt"
	"reflect"

	"github.com/sacloud/iaas-api-go/mapconv"
	"github.com/sacloud/iaas-api-go/ostype"
	"github.com/sacloud/iaas-api-go/types"
)

type definition struct {
	key   interface{}
	value interface{}
}

func (d *definition) stringKey() string {
	switch s := d.key.(type) {
	case string:
		return s
	case fmt.Stringer:
		return s.String()
	default:
		return fmt.Sprintf("%v", s)
	}
}

// definitions usacloudで使う名称(key)/値(value)のペア
var definitions = map[string][]*definition{
	"backup_start_minute": {
		{key: 0, value: 0},
		{key: 15, value: 15},
		{key: 30, value: 30},
		{key: 45, value: 45},
	},
	"cdrom_sizes": {
		{key: 5, value: 5},
		{key: 10, value: 10},
		{key: 20, value: 20},
	},
	"certificate_authority_issuance_method": {
		{key: types.CertificateAuthorityIssuanceMethods.URL.String(), value: types.CertificateAuthorityIssuanceMethods.URL},
		{key: types.CertificateAuthorityIssuanceMethods.EMail.String(), value: types.CertificateAuthorityIssuanceMethods.EMail},
		{key: types.CertificateAuthorityIssuanceMethods.CSR.String(), value: types.CertificateAuthorityIssuanceMethods.CSR},
		{key: types.CertificateAuthorityIssuanceMethods.PublicKey.String(), value: types.CertificateAuthorityIssuanceMethods.PublicKey},
	},
	"container_registry_access_level": {
		{key: types.ContainerRegistryAccessLevels.ReadWrite.String(), value: types.ContainerRegistryAccessLevels.ReadWrite},
		{key: types.ContainerRegistryAccessLevels.ReadOnly.String(), value: types.ContainerRegistryAccessLevels.ReadOnly},
		{key: types.ContainerRegistryAccessLevels.None.String(), value: types.ContainerRegistryAccessLevels.None},
	},
	"container_registry_permission": {
		{key: types.ContainerRegistryPermissions.All.String(), value: types.ContainerRegistryPermissions.All},
		{key: types.ContainerRegistryPermissions.ReadWrite.String(), value: types.ContainerRegistryPermissions.ReadWrite},
		{key: types.ContainerRegistryPermissions.ReadOnly.String(), value: types.ContainerRegistryPermissions.ReadOnly},
	},
	"database_plan": {
		{key: "10g", value: types.DatabasePlans.DB10GB},
		{key: "30g", value: types.DatabasePlans.DB30GB},
		{key: "90g", value: types.DatabasePlans.DB90GB},
		{key: "240g", value: types.DatabasePlans.DB240GB},
		{key: "500g", value: types.DatabasePlans.DB500GB},
		{key: "1t", value: types.DatabasePlans.DB1TB},
	},
	"database_type": {
		{key: "postgresql", value: "postgres"},
		{key: "postgres", value: "postgres"},
		{key: "mariadb", value: "mariadb"},
	},
	"disk_plan": {
		{key: "ssd", value: types.DiskPlans.SSD},
		{key: "hdd", value: types.DiskPlans.HDD},
	},
	"disk_connection": {
		{key: types.DiskConnections.VirtIO.String(), value: types.DiskConnections.VirtIO.String()},
		{key: types.DiskConnections.IDE.String(), value: types.DiskConnections.IDE.String()},
	},
	"disk_encryption_algorithm": {
		{key: types.DiskEncryptionAlgorithms.None.String(), value: types.DiskEncryptionAlgorithms.None.String()},
		{key: types.DiskEncryptionAlgorithms.AES256XTS.String(), value: types.DiskEncryptionAlgorithms.AES256XTS.String()},
	},
	"dns_record_type": {
		{key: types.DNSRecordTypes.A.String(), value: types.DNSRecordTypes.A},
		{key: types.DNSRecordTypes.AAAA.String(), value: types.DNSRecordTypes.AAAA},
		{key: types.DNSRecordTypes.ALIAS.String(), value: types.DNSRecordTypes.ALIAS},
		{key: types.DNSRecordTypes.CNAME.String(), value: types.DNSRecordTypes.CNAME},
		{key: types.DNSRecordTypes.NS.String(), value: types.DNSRecordTypes.NS},
		{key: types.DNSRecordTypes.MX.String(), value: types.DNSRecordTypes.MX},
		{key: types.DNSRecordTypes.TXT.String(), value: types.DNSRecordTypes.TXT},
		{key: types.DNSRecordTypes.SRV.String(), value: types.DNSRecordTypes.SRV},
		{key: types.DNSRecordTypes.CAA.String(), value: types.DNSRecordTypes.CAA},
		{key: types.DNSRecordTypes.PTR.String(), value: types.DNSRecordTypes.PTR},
	},
	"enhanced_db_type": {
		{key: types.EnhancedDBTypesTiDB.String(), value: types.EnhancedDBTypesTiDB},
		{key: types.EnhancedDBTypesMariaDB.String(), value: types.EnhancedDBTypesMariaDB},
	},
	"enhanced_db_region": {
		{key: types.EnhancedDBRegionsIs1.String(), value: types.EnhancedDBRegionsIs1},
		{key: types.EnhancedDBRegionsTk1.String(), value: types.EnhancedDBRegionsTk1},
	},
	"gslb_protocol": {
		{key: types.GSLBHealthCheckProtocols.HTTP.String(), value: types.GSLBHealthCheckProtocols.HTTP},
		{key: types.GSLBHealthCheckProtocols.HTTPS.String(), value: types.GSLBHealthCheckProtocols.HTTPS},
		{key: types.GSLBHealthCheckProtocols.Ping.String(), value: types.GSLBHealthCheckProtocols.Ping},
		{key: types.GSLBHealthCheckProtocols.TCP.String(), value: types.GSLBHealthCheckProtocols.TCP},
	},
	"interface_driver": {
		{key: "virtio", value: types.InterfaceDrivers.VirtIO},
		{key: "e1000", value: types.InterfaceDrivers.E1000},
	},
	"internet_network_mask_len": {
		{key: 28, value: 28},
		{key: 27, value: 27},
		{key: 26, value: 26},
	},
	"internet_bandwidth": {
		{key: 100, value: 100},
		{key: 250, value: 250},
		{key: 500, value: 500},
		{key: 1000, value: 1000},
		{key: 1500, value: 1500},
		{key: 2000, value: 2000},
		{key: 2500, value: 2500},
		{key: 3000, value: 3000},
		{key: 3500, value: 3500},
		{key: 4000, value: 4000},
		{key: 4500, value: 4500},
		{key: 5000, value: 5000},
		{key: 5500, value: 5500},
		{key: 6000, value: 6000},
		{key: 6500, value: 6500},
		{key: 7000, value: 7000},
		{key: 7500, value: 7500},
		{key: 8000, value: 8000},
		{key: 8500, value: 8500},
		{key: 9000, value: 9000},
		{key: 9500, value: 9500},
		{key: 10000, value: 10000},
	},
	"loadbalancer_plan": {
		{key: "standard", value: types.LoadBalancerPlans.Standard},
		{key: "highspec", value: types.LoadBalancerPlans.HighSpec},
	},
	"loadbalancer_server_protocol": {
		{key: types.LoadBalancerHealthCheckProtocols.HTTP.String(), value: types.LoadBalancerHealthCheckProtocols.HTTP},
		{key: types.LoadBalancerHealthCheckProtocols.HTTPS.String(), value: types.LoadBalancerHealthCheckProtocols.HTTPS},
		{key: types.LoadBalancerHealthCheckProtocols.TCP.String(), value: types.LoadBalancerHealthCheckProtocols.TCP},
		{key: types.LoadBalancerHealthCheckProtocols.Ping.String(), value: types.LoadBalancerHealthCheckProtocols.Ping},
	},
	"nfs_plan": {
		{key: "ssd", value: types.NFSPlans.SSD},
		{key: "hdd", value: types.NFSPlans.HDD},
	},
	"note_class": {
		{key: "shell", value: "shell"},
		{key: "yaml_cloud_config", value: "sheyaml_cloud_configll"},
	},
	"os_type": ostypeDefinition(),
	"packetfilter_protocol": {
		{key: types.Protocols.HTTP.String(), value: types.Protocols.HTTP},
		{key: types.Protocols.HTTPS.String(), value: types.Protocols.HTTPS},
		{key: types.Protocols.TCP.String(), value: types.Protocols.TCP},
		{key: types.Protocols.UDP.String(), value: types.Protocols.UDP},
		{key: types.Protocols.ICMP.String(), value: types.Protocols.ICMP},
		{key: types.Protocols.Fragment.String(), value: types.Protocols.Fragment},
		{key: types.Protocols.IP.String(), value: types.Protocols.IP},
	},
	"packetfilter_action": {
		{key: types.Actions.Allow.String(), value: types.Actions.Allow},
		{key: types.Actions.Deny.String(), value: types.Actions.Deny},
	},
	"private_host_class": {
		{key: "dynamic", value: "dynamic"},
		{key: "ms_windows", value: "ms_windows"},
	},
	"proxylb_proxy_mode": {
		{key: types.ProxyLBProxyModes.HTTP.String(), value: types.ProxyLBProxyModes.HTTP},
		{key: types.ProxyLBProxyModes.HTTPS.String(), value: types.ProxyLBProxyModes.HTTPS},
		{key: types.ProxyLBProxyModes.TCP.String(), value: types.ProxyLBProxyModes.TCP},
	},
	"proxylb_plan": {
		{key: types.ProxyLBPlans.CPS100.String(), value: types.ProxyLBPlans.CPS100},
		{key: types.ProxyLBPlans.CPS500.String(), value: types.ProxyLBPlans.CPS500},
		{key: types.ProxyLBPlans.CPS1000.String(), value: types.ProxyLBPlans.CPS1000},
		{key: types.ProxyLBPlans.CPS5000.String(), value: types.ProxyLBPlans.CPS5000},
		{key: types.ProxyLBPlans.CPS10000.String(), value: types.ProxyLBPlans.CPS10000},
		{key: types.ProxyLBPlans.CPS50000.String(), value: types.ProxyLBPlans.CPS50000},
		{key: types.ProxyLBPlans.CPS100000.String(), value: types.ProxyLBPlans.CPS100000},
		{key: types.ProxyLBPlans.CPS400000.String(), value: types.ProxyLBPlans.CPS400000},
	},
	"proxylb_protocol": {
		{key: types.ProxyLBProtocols.HTTP.String(), value: types.ProxyLBProtocols.HTTP},
		{key: types.ProxyLBProtocols.TCP.String(), value: types.ProxyLBProtocols.TCP},
	},
	"proxylb_region": {
		{key: types.ProxyLBRegions.TK1.String(), value: types.ProxyLBRegions.TK1},
		{key: types.ProxyLBRegions.IS1.String(), value: types.ProxyLBRegions.IS1},
		{key: types.ProxyLBRegions.Anycast.String(), value: types.ProxyLBRegions.Anycast},
	},
	"proxylb_ssl_policy": {
		{key: "TLS-1-2-2019-04", value: "TLS-1-2-2019-04"},
		{key: "TLS-1-2-2021-06", value: "TLS-1-2-2021-06"},
		{key: "TLS-1-3-2021-06", value: "TLS-1-3-2021-06"},
	},
	"rest_method": {
		{key: "get", value: "get"},
		{key: "post", value: "post"},
		{key: "patch", value: "patch"},
		{key: "put", value: "put"},
		{key: "delete", value: "delete"},
		{key: "GET", value: "get"},
		{key: "POST", value: "post"},
		{key: "PATCH", value: "patch"},
		{key: "PUT", value: "put"},
		{key: "DELETE", value: "delete"},
	},
	"scope": {
		{key: types.Scopes.User.String(), value: types.Scopes.User},
		{key: types.Scopes.Shared.String(), value: types.Scopes.Shared},
	},
	"server_plan_commitment": {
		{key: types.Commitments.Standard.String(), value: types.Commitments.Standard},
		{key: types.Commitments.DedicatedCPU.String(), value: types.Commitments.DedicatedCPU},
	},
	"server_plan_generation": {
		{key: "default", value: types.PlanGenerations.Default},
		{key: "g100", value: types.PlanGenerations.G100},
		{key: "g200", value: types.PlanGenerations.G200},
	},
	"simple_monitor_protocol": {
		{key: types.SimpleMonitorProtocols.HTTP.String(), value: types.SimpleMonitorProtocols.HTTP},
		{key: types.SimpleMonitorProtocols.HTTPS.String(), value: types.SimpleMonitorProtocols.HTTPS},
		{key: types.SimpleMonitorProtocols.Ping.String(), value: types.SimpleMonitorProtocols.Ping},
		{key: types.SimpleMonitorProtocols.TCP.String(), value: types.SimpleMonitorProtocols.TCP},
		{key: types.SimpleMonitorProtocols.DNS.String(), value: types.SimpleMonitorProtocols.DNS},
		{key: types.SimpleMonitorProtocols.SSH.String(), value: types.SimpleMonitorProtocols.SSH},
		{key: types.SimpleMonitorProtocols.SMTP.String(), value: types.SimpleMonitorProtocols.SMTP},
		{key: types.SimpleMonitorProtocols.POP3.String(), value: types.SimpleMonitorProtocols.POP3},
		{key: types.SimpleMonitorProtocols.SNMP.String(), value: types.SimpleMonitorProtocols.SNMP},
		{key: types.SimpleMonitorProtocols.SSLCertificate.String(), value: types.SimpleMonitorProtocols.SSLCertificate},
		{key: types.SimpleMonitorProtocols.FTP.String(), value: types.SimpleMonitorProtocols.FTP},
	},
	"simple_monitor_ftps": {
		{key: types.SimpleMonitorFTPSValues.Explicit.String(), value: types.SimpleMonitorFTPSValues.Explicit},
		{key: types.SimpleMonitorFTPSValues.Implicit.String(), value: types.SimpleMonitorFTPSValues.Implicit},
	},
	"vpc_router_plan": {
		{key: "standard", value: types.VPCRouterPlans.Standard},
		{key: "premium", value: types.VPCRouterPlans.Premium},
		{key: "highspec", value: types.VPCRouterPlans.HighSpec},
		{key: "highspec4000", value: types.VPCRouterPlans.HighSpec4000},
	},
	"vpc_router_plan_premium": {
		{key: "premium", value: types.VPCRouterPlans.Premium},
		{key: "highspec", value: types.VPCRouterPlans.HighSpec},
		{key: "highspec4000", value: types.VPCRouterPlans.HighSpec4000},
	},
	"vpc_router_port_forwarding_protocol": {
		{key: string(types.VPCRouterPortForwardingProtocols.TCP), value: types.VPCRouterPortForwardingProtocols.TCP},
		{key: string(types.VPCRouterPortForwardingProtocols.UDP), value: types.VPCRouterPortForwardingProtocols.UDP},
	},
	"vpc_router_firewall_protocol": {
		{key: string(types.VPCRouterFirewallProtocols.TCP), value: types.VPCRouterFirewallProtocols.TCP},
		{key: string(types.VPCRouterFirewallProtocols.UDP), value: types.VPCRouterFirewallProtocols.UDP},
		{key: string(types.VPCRouterFirewallProtocols.ICMP), value: types.VPCRouterFirewallProtocols.ICMP},
		{key: string(types.VPCRouterFirewallProtocols.IP), value: types.VPCRouterFirewallProtocols.IP},
	},
	"weekdays": {
		{key: "all", value: "all"},
		{key: types.DaysOfTheWeek.Sunday.String(), value: types.DaysOfTheWeek.Sunday},
		{key: types.DaysOfTheWeek.Monday.String(), value: types.DaysOfTheWeek.Monday},
		{key: types.DaysOfTheWeek.Tuesday.String(), value: types.DaysOfTheWeek.Tuesday},
		{key: types.DaysOfTheWeek.Wednesday.String(), value: types.DaysOfTheWeek.Wednesday},
		{key: types.DaysOfTheWeek.Thursday.String(), value: types.DaysOfTheWeek.Thursday},
		{key: types.DaysOfTheWeek.Friday.String(), value: types.DaysOfTheWeek.Friday},
		{key: types.DaysOfTheWeek.Saturday.String(), value: types.DaysOfTheWeek.Saturday},
	},
}

func Keys(key string) ([]string, bool) {
	values, ok := definitions[key]
	if !ok {
		return nil, false
	}
	var results []string
	for _, v := range values {
		results = append(results, v.stringKey())
	}
	return results, true
}

func ostypeDefinition() []*definition {
	var def []*definition
	for _, name := range ostype.OSTypeShortNames {
		def = append(def, &definition{key: name, value: ostype.StrToOSType(name)})
	}
	return def
}

func init() {
	registerFunctions()
}

func registerFunctions() {
	// definitionsから各種定義を登録(Note: 同名のものがあった場合は上書き)
	registerConverterFilters()
	registerTemplateFuncMap()
	registerValidators()
	registerCLITagOptions()
}

func registerConverterFilters() {
	for name, defs := range definitions {
		if _, ok := ConverterFilters[name+"_to_value"]; !ok {
			ConverterFilters[name+"_to_value"] = convertFuncToValue(name, defs)
		}
		if _, ok := ConverterFilters[name+"_to_key"]; !ok {
			ConverterFilters[name+"_to_key"] = convertFuncToKey(name, defs)
		}
	}
}

func registerTemplateFuncMap() {
	for name, defs := range definitions {
		if _, ok := TemplateFuncMap[name+"_to_value"]; !ok {
			TemplateFuncMap[name+"_to_value"] = templateFuncToValue(defs)
		}
		if _, ok := TemplateFuncMap[name+"_to_key"]; !ok {
			TemplateFuncMap[name+"_to_key"] = templateFuncToKey(defs)
		}
	}
}

func registerValidators() {
	// definitionsの各値からキーを取り出し、"oneof=keyのスペース区切り"というルールを登録する
	for name, defs := range definitions {
		if _, ok := validatorAliases[name]; ok {
			continue
		}
		var allows []string
		for _, def := range defs {
			allows = append(allows, def.stringKey())
		}
		validatorAliases[name] = fmt.Sprintf("oneof=%s", joinWithSpace(allows))
	}
}

func registerCLITagOptions() {
	// definitionsの各値からキーを取り出し、FlagOptionsMapに登録する
	for name, defs := range definitions {
		if _, ok := FlagOptionsMap[name]; ok {
			continue
		}
		var allows []string
		for _, def := range defs {
			allows = append(allows, def.stringKey())
		}
		FlagOptionsMap[name] = allows
	}
}

func convertFuncToValue(defName string, defs []*definition) mapconv.FilterFunc {
	var fn func(interface{}) (interface{}, error)
	fn = func(v interface{}) (interface{}, error) {
		// スライスの場合は再帰処理
		vt := reflect.ValueOf(v)
		if vt.Kind() == reflect.Slice || vt.Kind() == reflect.Array {
			var results []interface{}
			for i := 0; i < vt.Len(); i++ {
				res, err := fn(vt.Index(i).Interface())
				if err != nil {
					return nil, err
				}
				results = append(results, res)
			}
			return results, nil
		}

		var result interface{}
		for _, def := range defs {
			if reflect.DeepEqual(v, def.key) {
				result = def.value
				break
			}
		}
		if result == nil {
			return nil, fmt.Errorf("key %v not found in %s", v, defName)
		}
		return result, nil
	}
	return fn
}

func convertFuncToKey(defName string, defs []*definition) mapconv.FilterFunc {
	var fn func(interface{}) (interface{}, error)
	fn = func(v interface{}) (interface{}, error) {
		// スライスの場合は再帰処理
		vt := reflect.ValueOf(v)
		if vt.Kind() == reflect.Slice || vt.Kind() == reflect.Array {
			var results []interface{}
			for i := 0; i < vt.Len(); i++ {
				res, err := fn(vt.Index(i).Interface())
				if err != nil {
					return nil, err
				}
				results = append(results, res)
			}
			return results, nil
		}

		var result interface{}
		for _, def := range defs {
			if reflect.DeepEqual(v, def.value) {
				result = def.key
				break
			}
		}
		if result == nil {
			return nil, fmt.Errorf("value %v not found in %s", v, defName)
		}
		return result, nil
	}
	return fn
}

func templateFuncToValue(defs []*definition) func(interface{}) interface{} {
	return func(raw interface{}) interface{} {
		in := raw
		if v, ok := raw.(fmt.Stringer); ok {
			in = v.String()
		}
		var result interface{}
		for _, def := range defs {
			switch ky := def.key.(type) {
			case fmt.Stringer:
				if reflect.DeepEqual(in, ky.String()) {
					result = def.value
					break
				}
			default:
				if reflect.DeepEqual(in, def.key) {
					result = def.value
					break
				}
			}
		}
		return result
	}
}

func templateFuncToKey(defs []*definition) func(interface{}) interface{} {
	return func(raw interface{}) interface{} {
		in := raw
		if v, ok := raw.(fmt.Stringer); ok {
			in = v.String()
		}
		var result interface{}
		for _, def := range defs {
			switch val := def.value.(type) {
			case fmt.Stringer:
				if reflect.DeepEqual(in, val.String()) {
					result = def.key
					break
				}
			default:
				if reflect.DeepEqual(in, def.value) {
					result = def.key
					break
				}
			}
		}
		return result
	}
}

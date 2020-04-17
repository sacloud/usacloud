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

package define

import (
	"fmt"
	"strings"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/schema"
)

func ProxyLBResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find", "select"},
			Params:             proxyLBListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: proxyLBListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"create": {
			Type:             schema.CommandCreate,
			Params:           proxyLBCreateParam(),
			IncludeFields:    proxyLBDetailIncludes(),
			ExcludeFields:    proxyLBDetailExcludes(),
			UseCustomCommand: true,
			Category:         "basics",
			Order:            20,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        proxyLBReadParam(),
			IncludeFields: proxyLBDetailIncludes(),
			ExcludeFields: proxyLBDetailExcludes(),
			Category:      "basics",
			Order:         30,
		},
		"update": {
			Type:             schema.CommandUpdate,
			Params:           proxyLBUpdateParam(),
			IncludeFields:    proxyLBDetailIncludes(),
			ExcludeFields:    proxyLBDetailExcludes(),
			UseCustomCommand: true,
			Category:         "basics",
			Order:            40,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"rm"},
			Params:        proxyLBDeleteParam(),
			IncludeFields: proxyLBDetailIncludes(),
			ExcludeFields: proxyLBDetailExcludes(),
			Category:      "basics",
			Order:         50,
		},
		"plan-change": {
			Type:             schema.CommandManipulateMulti,
			Params:           proxyLBPlanChangeParam(),
			Usage:            "Change ProxyLB plan",
			IncludeFields:    proxyLBDetailIncludes(),
			ExcludeFields:    proxyLBDetailExcludes(),
			UseCustomCommand: true,
			Category:         "basics",
			Order:            60,
		},
		"bind-port-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             proxyLBBindPortListParam(),
			Aliases:            []string{"bind-port-list"},
			TableType:          output.TableSimple,
			TableColumnDefines: proxyLBBindPortListColumns(),
			UseCustomCommand:   true,
			NeedlessConfirm:    true,
			Category:           "bind-port",
			Order:              10,
		},
		"bind-port-add": {
			Type:               schema.CommandManipulateSingle,
			Params:             proxyLBBindPortAddParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: proxyLBBindPortListColumns(),
			UseCustomCommand:   true,
			Category:           "bind-port",
			Order:              20,
		},
		"bind-port-update": {
			Type:               schema.CommandManipulateSingle,
			Params:             proxyLBBindPortUpdateParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: proxyLBBindPortListColumns(),
			UseCustomCommand:   true,
			Category:           "bind-port",
			Order:              30,
		},
		"bind-port-delete": {
			Type:               schema.CommandManipulateSingle,
			Params:             proxyLBBindPortDeleteParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: proxyLBBindPortListColumns(),
			UseCustomCommand:   true,
			ConfirmMessage:     "delete bind-port",
			Category:           "bind-port",
			Order:              40,
		},
		"response-header-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             proxyLBResponseHeaderListParam(),
			Aliases:            []string{"response-header-list"},
			TableType:          output.TableSimple,
			TableColumnDefines: proxyLBResponseHeaderListColumns(),
			UseCustomCommand:   true,
			NeedlessConfirm:    true,
			Category:           "response-header",
			Order:              10,
		},
		"response-header-add": {
			Type:               schema.CommandManipulateSingle,
			Params:             proxyLBResponseHeaderAddParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: proxyLBResponseHeaderListColumns(),
			UseCustomCommand:   true,
			Category:           "response-header",
			Order:              20,
		},
		"response-header-update": {
			Type:               schema.CommandManipulateSingle,
			Params:             proxyLBResponseHeaderUpdateParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: proxyLBResponseHeaderListColumns(),
			UseCustomCommand:   true,
			Category:           "response-header",
			Order:              30,
		},
		"response-header-delete": {
			Type:               schema.CommandManipulateSingle,
			Params:             proxyLBResponseHeaderDeleteParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: proxyLBResponseHeaderListColumns(),
			UseCustomCommand:   true,
			ConfirmMessage:     "delete response-header",
			Category:           "response-header",
			Order:              40,
		},
		"acme-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             proxyLBACMEInfoParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: proxyLBACMEInfoColumns(),
			UseCustomCommand:   true,
			NeedlessConfirm:    true,
			Category:           "acme",
			Order:              10,
		},
		"acme-setting": {
			Type:               schema.CommandManipulateSingle,
			Params:             proxyLBACMESettingParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: proxyLBACMEInfoColumns(),
			UseCustomCommand:   true,
			Category:           "acme",
			Order:              20,
		},
		"acme-renew": {
			Type:             schema.CommandManipulateSingle,
			Params:           proxyLBACMERenewParam(),
			NoOutput:         true,
			UseCustomCommand: true,
			Category:         "acme",
			Order:            30,
		},
		"server-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             proxyLBServerListParam(),
			Aliases:            []string{"server-list"},
			TableType:          output.TableSimple,
			TableColumnDefines: proxyLBServerListColumns(),
			UseCustomCommand:   true,
			NeedlessConfirm:    true,
			Category:           "servers",
			Order:              10,
		},
		"server-add": {
			Type:               schema.CommandManipulateSingle,
			Params:             proxyLBServerAddParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: proxyLBServerListColumns(),
			UseCustomCommand:   true,
			Category:           "servers",
			Order:              20,
		},
		"server-update": {
			Type:               schema.CommandManipulateSingle,
			Params:             proxyLBServerUpdateParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: proxyLBServerListColumns(),
			UseCustomCommand:   true,
			Category:           "servers",
			Order:              30,
		},
		"server-delete": {
			Type:               schema.CommandManipulateSingle,
			Params:             proxyLBServerDeleteParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: proxyLBServerListColumns(),
			UseCustomCommand:   true,
			ConfirmMessage:     "delete server",
			Category:           "servers",
			Order:              40,
		},
		"certificate-info": {
			Type:               schema.CommandManipulateSingle,
			Params:             proxyLBCertListParam(),
			Aliases:            []string{"certificate-list", "cert-list", "cert-info"},
			TableType:          output.TableSimple,
			TableColumnDefines: proxyLBCertListColumns(),
			UseCustomCommand:   true,
			NeedlessConfirm:    true,
			Category:           "certificate",
			Order:              10,
		},
		"certificate-add": {
			Type:               schema.CommandManipulateSingle,
			Params:             proxyLBCertAddParam(),
			Aliases:            []string{"cert-add"},
			TableType:          output.TableSimple,
			TableColumnDefines: proxyLBCertListColumns(),
			UseCustomCommand:   true,
			Category:           "certificate",
			Order:              20,
		},
		"certificate-update": {
			Type:               schema.CommandManipulateSingle,
			Params:             proxyLBCertUpdateParam(),
			Aliases:            []string{"cert-update"},
			TableType:          output.TableSimple,
			TableColumnDefines: proxyLBCertListColumns(),
			UseCustomCommand:   true,
			Category:           "certificate",
			Order:              30,
		},
		"certificate-delete": {
			Type:               schema.CommandManipulateSingle,
			Params:             proxyLBCertDeleteParam(),
			Aliases:            []string{"cert-delete"},
			TableType:          output.TableSimple,
			TableColumnDefines: proxyLBCertListColumns(),
			UseCustomCommand:   true,
			ConfirmMessage:     "delete certificate",
			Category:           "certificate",
			Order:              40,
		},
		"monitor": {
			Type:               schema.CommandRead,
			Params:             proxyLBMonitorParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: proxyLBMonitorColumns(),
			UseCustomCommand:   true,
			Category:           "monitor",
		},
	}

	return &schema.Resource{
		Commands:            commands,
		CommandCategories:   proxyLBCommandCategories,
		Aliases:             []string{"enhanced-load-balancer", "proxylb"},
		ResourceCategory:    CategoryCommonServiceItem,
		ListResultFieldName: "CommonServiceProxyLBItems",
	}
}

var proxyLBCommandCategories = []schema.Category{
	{
		Key:         "basics",
		DisplayName: "Basics",
		Order:       10,
	},
	{
		Key:         "bind-port",
		DisplayName: "Bind Port(s) Management",
		Order:       20,
	},
	{
		Key:         "response-header",
		DisplayName: "Additional Response Header(s) Management",
		Order:       22,
	},
	{
		Key:         "acme",
		DisplayName: "ACME settings",
		Order:       25,
	},
	{
		Key:         "servers",
		DisplayName: "Real Server(s) Management",
		Order:       30,
	},
	{
		Key:         "certificate",
		DisplayName: "Certificate(s) Management",
		Order:       40,
	},
	{
		Key:         "monitor",
		DisplayName: "Monitoring",
		Order:       50,
	},
}

func proxyLBListParam() map[string]*schema.Schema {
	return mergeParameterMap(CommonListParam, paramTagsCond)
}

func proxyLBListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{
			Name:    "HealthCheck",
			Sources: []string{"Settings.ProxyLB.HealthCheck.Protocol"},
		},
		{
			Name:    "VIP",
			Sources: []string{"Status.VirtualIPAddress"},
		},
		{
			Name:    "FQDN",
			Sources: []string{"Status.FQDN"},
		},
		{
			Name:    "StickySession",
			Sources: []string{"Settings.ProxyLB.StickySession.Enabled"},
		},
	}
}

func proxyLBBindPortListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "__ORDER__"}, // magic column name(generated on demand)
		{Name: "ProxyMode"},
		{Name: "Port"},
		{
			Name:    "RedirectToHTTPS",
			Sources: []string{"RedirectToHttps"},
		},
		{
			Name:    "SupportHTTP2",
			Sources: []string{"SupportHttp2"},
		},
	}
}

func proxyLBResponseHeaderListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "__ORDER__"}, // magic column name(generated on demand)
		{Name: "Header"},
		{Name: "Value"},
	}
}

func proxyLBACMEInfoColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Enabled"},
		{Name: "CommonName"},
	}
}

func proxyLBServerListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "__ORDER__"}, // magic column name(generated on demand)
		{Name: "IPAddress"},
		{Name: "Port"},
		{Name: "Enabled"},
	}
}

func proxyLBCertListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "__ORDER__"}, // magic column name(generated on demand)
		{Name: "CommonName"},
		{Name: "IPAddress"},
		{Name: "Issuer"},
		{Name: "PrimaryCert.CertificateEndDate"},
	}
}

func proxyLBDetailIncludes() []string {
	return []string{}
}

func proxyLBDetailExcludes() []string {
	return []string{}
}

func proxyLBCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"plan": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set plan",
			DefaultValue: 1000,
			ValidateFunc: validateInIntValues(sacloud.AllowProxyLBPlans...),
			Category:     "ProxyLB",
			Order:        5,
		},
		"protocol": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set healthcheck protocol[http/tcp]",
			DefaultValue: "tcp",
			Required:     true,
			ValidateFunc: validateInStrValues(sacloud.AllowProxyLBHealthCheckProtocols...),
			Category:     "ProxyLB",
			Order:        10,
		},
		"host-header": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set host header of http/https healthcheck request",
			Category:    "ProxyLB",
			Order:       20,
		},
		"path": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set path of http/https healthcheck request",
			DefaultValue: "/",
			Category:     "ProxyLB",
			Order:        30,
		},
		"delay-loop": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set delay-loop of healthcheck",
			ValidateFunc: validateIntRange(10, 60),
			Required:     true,
			DefaultValue: 10,
			Category:     "ProxyLB",
			Order:        60,
		},
		"sticky-session": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "enable sticky-session",
			Category:    "ProxyLB",
			Order:       70,
		},
		"sorry-server-ipaddress": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set sorry-server ip address",
			Category:    "ProxyLB",
			Order:       80,
		},
		"sorry-server-port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set sorry-server ports",
			ValidateFunc: validateIntRange(1, 65535),
			Category:     "ProxyLB",
			Order:        85,
		},
		"timeout": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set timeout",
			DefaultValue: 10,
			ValidateFunc: validateIntRange(10, 600),
			Category:     "ProxyLB",
			Order:        90,
		},
		"name":        paramRequiredName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func proxyLBReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func proxyLBUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"protocol": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set healthcheck protocol[http/tcp]",
			ValidateFunc: validateInStrValues(sacloud.AllowProxyLBHealthCheckProtocols...),
			Category:     "ProxyLB",
			Order:        10,
		},
		"host-header": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set host header of http/https healthcheck request",
			Category:    "ProxyLB",
			Order:       20,
		},
		"path": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set path of http/https healthcheck request",
			Category:    "ProxyLB",
			Order:       30,
		},
		"delay-loop": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set delay-loop of healthcheck",
			ValidateFunc: validateIntRange(10, 60),
			Category:     "ProxyLB",
			Order:        60,
		},
		"sticky-session": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "enable sticky-session",
			Category:    "ProxyLB",
			Order:       70,
		},
		"sorry-server-ipaddress": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set sorry-server ip address",
			Category:    "ProxyLB",
			Order:       80,
		},
		"sorry-server-port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set sorry-server ports",
			ValidateFunc: validateIntRange(1, 65535),
			Category:     "ProxyLB",
			Order:        80,
		},
		"timeout": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set timeout",
			DefaultValue: 10,
			ValidateFunc: validateIntRange(10, 600),
			Category:     "ProxyLB",
			Order:        90,
		},
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func proxyLBDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func proxyLBPlanChangeParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"plan": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set plan",
			Required:     true,
			ValidateFunc: validateInIntValues(sacloud.AllowProxyLBPlans...),
			Category:     "ProxyLB",
			Order:        10,
		},
	}
}

func proxyLBBindPortListParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func proxyLBBindPortAddParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"mode": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  fmt.Sprintf("set bind mode[%s]", strings.Join(sacloud.AllowProxyLBBindModes, "/")),
			Required:     true,
			ValidateFunc: validateInStrValues(sacloud.AllowProxyLBBindModes...),
			Category:     "bind-port",
			Order:        10,
		},
		"port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Required:     true,
			Description:  "set port number",
			ValidateFunc: validateIntRange(1, 65535),
			Category:     "bind-port",
			Order:        20,
		},
		"redirect-to-https": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "enable to redirect to https",
			Category:    "bind-port",
			Order:       30,
		},
		"support-http2": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "enable http/2",
			Category:    "bind-port",
			Order:       40,
		},
	}
}
func proxyLBBindPortUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target server",
			Required:    true,
			Category:    "server",
			Order:       1,
		},
		"mode": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  fmt.Sprintf("set bind mode[%s]", strings.Join(sacloud.AllowProxyLBBindModes, "/")),
			ValidateFunc: validateInStrValues(sacloud.AllowProxyLBBindModes...),
			Category:     "bind-port",
			Order:        10,
		},
		"port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set port number",
			ValidateFunc: validateIntRange(1, 65535),
			Category:     "bind-port",
			Order:        20,
		},
		"redirect-to-https": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "enable to redirect to https",
			Category:    "bind-port",
			Order:       30,
		},
		"support-http2": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "enable http/2",
			Category:    "bind-port",
			Order:       40,
		},
	}
}

func proxyLBBindPortDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target bind-port",
			Required:    true,
			Category:    "bind-port",
			Order:       1,
		},
	}
}

func proxyLBResponseHeaderListParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"port-index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target bind-port",
			Required:    true,
			Category:    "response-header",
			Order:       2,
		},
	}
}

func proxyLBResponseHeaderAddParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"port-index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target bind-port",
			Required:    true,
			Category:    "response-header",
			Order:       2,
		},
		"header": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set Header",
			Required:    true,
			Category:    "response-header",
			Order:       10,
		},
		"value": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set Value",
			Required:    true,
			Category:    "response-header",
			Order:       20,
		},
	}
}
func proxyLBResponseHeaderUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target server",
			Required:    true,
			Category:    "server",
			Order:       1,
		},
		"port-index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target bind-port",
			Required:    true,
			Category:    "response-header",
			Order:       2,
		},
		"header": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set Header",
			Category:    "response-header",
			Order:       10,
		},
		"value": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set Value",
			Category:    "response-header",
			Order:       20,
		},
	}
}

func proxyLBResponseHeaderDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target bind-port",
			Required:    true,
			Category:    "bind-port",
			Order:       1,
		},
		"port-index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target bind-port",
			Required:    true,
			Category:    "response-header",
			Order:       2,
		},
	}
}

func proxyLBACMEInfoParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func proxyLBACMESettingParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"accept-tos": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "the flag of accept Let's Encrypt's terms of services: https://letsencrypt.org/repository/",
			Category:    "acme",
			Order:       10,
		},
		"common-name": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set common name",
			Category:    "acme",
			Order:       20,
		},
		"disable": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "the flag of disable Let's Encrypt",
			Category:    "acme",
			Order:       30,
		},
	}
}

func proxyLBACMERenewParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func proxyLBServerListParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func proxyLBServerAddParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ipaddress": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Required:     true,
			Description:  "set target ipaddress",
			ValidateFunc: validateIPv4Address(),
			Category:     "server",
			Order:        10,
		},
		"disabled": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set disabled",
			Category:    "server",
			Order:       20,
		},
		"port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Required:     true,
			Description:  "set server ports",
			ValidateFunc: validateIntRange(1, 65535),
			Category:     "server",
			Order:        80,
		},
	}
}
func proxyLBServerUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target server",
			Required:    true,
			Category:    "server",
			Order:       1,
		},
		"ipaddress": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set target ipaddress",
			ValidateFunc: validateIPv4Address(),
			Category:     "server",
			Order:        10,
		},
		"disabled": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set disabled",
			Category:    "server",
			Order:       20,
		},
		"port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set server ports",
			ValidateFunc: validateIntRange(1, 65535),
			Category:     "server",
			Order:        80,
		},
	}
}
func proxyLBServerDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"index": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "index of target server",
			Required:    true,
			Category:    "server",
			Order:       1,
		},
	}
}

func proxyLBCertListParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func proxyLBCertAddParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"server-certificate": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"server-cert"},
			Required:    true,
			Category:    "server",
			Order:       10,
		},
		"intermediate-certificate": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"issuer-cert"},
			Category:    "server",
			Order:       20,
		},
		"private-key": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Required:    true,
			Category:    "server",
			Order:       30,
		},
	}
}

func proxyLBCertUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"server-certificate": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"server-cert"},
			Category:    "server",
			Order:       10,
		},
		"intermediate-certificate": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"issuer-cert"},
			Category:    "server",
			Order:       20,
		},
		"private-key": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Category:    "server",
			Order:       30,
		},
	}
}

func proxyLBCertDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func proxyLBMonitorParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"start": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set start-time",
			ValidateFunc: validateDateTimeString(),
			Category:     "monitor",
			Order:        10,
		},
		"end": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set end-time",
			ValidateFunc: validateDateTimeString(),
			Category:     "monitor",
			Order:        20,
		},
		"key-format": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set monitoring value key-format",
			DefaultValue: "sakuracloud.proxylb.{{.ID}}",
			Required:     true,
			Category:     "monitor",
			Order:        30,
		},
	}
}

func proxyLBMonitorColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Key"},
		{Name: "TimeStamp"},
		{Name: "UnixTime"},
		{Name: "ActiveConnections"},
		{Name: "ConnectionsPerSec"},
	}
}

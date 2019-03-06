package define

import (
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func ProxyLBResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find", "selector"},
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
		Aliases:             []string{"enhanced-load-balancer", "proxylb"},
		ResourceCategory:    CategoryCommonServiceItem,
		ListResultFieldName: "CommonServiceProxyLBItems",
	}
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
	}
}

func proxyLBBindPortListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "__ORDER__"}, // magic column name(generated on demand)
		{Name: "ProxyMode"},
		{Name: "Port"},
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
		{Name: "CertificateEndDate"},
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
			CompleteFunc: completeInIntValues(sacloud.AllowProxyLBPlans...),
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
			CompleteFunc: completeInStrValues(sacloud.AllowProxyLBHealthCheckProtocols...),
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
			CompleteFunc: completeInStrValues(sacloud.AllowProxyLBHealthCheckProtocols...),
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
		"name":        paramName,
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func proxyLBDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func proxyLBBindPortListParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func proxyLBBindPortAddParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"mode": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set bind mode[http/https]",
			Required:     true,
			ValidateFunc: validateInStrValues(sacloud.AllowProxyLBBindModes...),
			CompleteFunc: completeInStrValues(sacloud.AllowProxyLBBindModes...),
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
			Description:  "set bind mode[http/https]",
			ValidateFunc: validateInStrValues(sacloud.AllowProxyLBBindModes...),
			CompleteFunc: completeInStrValues(sacloud.AllowProxyLBBindModes...),
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

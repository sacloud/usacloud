package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func SimpleMonitorResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find", "selector"},
			Params:             simpleMonitorListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: simpleMonitorListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"create": {
			Type:             schema.CommandCreate,
			Params:           simpleMonitorCreateParam(),
			ParamCategories:  simpleMonitorCreateParamCategories,
			IncludeFields:    simpleMonitorDetailIncludes(),
			ExcludeFields:    simpleMonitorDetailExcludes(),
			UseCustomCommand: true,
			Category:         "basics",
			Order:            20,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        simpleMonitorReadParam(),
			IncludeFields: simpleMonitorDetailIncludes(),
			ExcludeFields: simpleMonitorDetailExcludes(),
			Category:      "basics",
			Order:         30,
		},
		"update": {
			Type:             schema.CommandUpdate,
			Params:           simpleMonitorUpdateParam(),
			ParamCategories:  simpleMonitorUpdateParamCategories,
			IncludeFields:    simpleMonitorDetailIncludes(),
			ExcludeFields:    simpleMonitorDetailExcludes(),
			UseCustomCommand: true,
			Category:         "basics",
			Order:            40,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"rm"},
			Params:        simpleMonitorDeleteParam(),
			IncludeFields: simpleMonitorDetailIncludes(),
			ExcludeFields: simpleMonitorDetailExcludes(),
			Category:      "basics",
			Order:         50,
		},
	}

	return &schema.Resource{
		Commands:         commands,
		ResourceCategory: CategoryCommonServiceItem,
	}
}

var simpleMonitorCreateParamCategories = []schema.Category{
	{
		Key:         "monitor",
		DisplayName: "Simple-Monitor options",
		Order:       10,
	},
	{
		Key:         "health-check",
		DisplayName: "Health-Check(Common) options",
		Order:       20,
	},
	{
		Key:         "http-check",
		DisplayName: "Health-Check(HTTP/HTTPS) options",
		Order:       22,
	},
	{
		Key:         "dns-check",
		DisplayName: "Health-Check(DNS) options",
		Order:       24,
	},
	{
		Key:         "ssl-check",
		DisplayName: "Health-Check(SSL Certificate) options",
		Order:       26,
	},
	{
		Key:         "notify",
		DisplayName: "Notify options",
		Order:       30,
	},
	{
		Key:         "common",
		DisplayName: "Common options",
		Order:       100,
	},
}

var simpleMonitorUpdateParamCategories = []schema.Category{
	{
		Key:         "health-check",
		DisplayName: "Health-Check(Common) options",
		Order:       20,
	},
	{
		Key:         "http-check",
		DisplayName: "Health-Check(HTTP/HTTPS) options",
		Order:       22,
	},
	{
		Key:         "dns-check",
		DisplayName: "Health-Check(DNS) options",
		Order:       24,
	},
	{
		Key:         "ssl-check",
		DisplayName: "Health-Check(SSL Certificate) options",
		Order:       26,
	},
	{
		Key:         "notify",
		DisplayName: "Notify options",
		Order:       30,
	},
	{
		Key:         "common",
		DisplayName: "Common options",
		Order:       100,
	},
}

func simpleMonitorListParam() map[string]*schema.Schema {
	return mergeParameterMap(CommonListParam, paramTagsCond)
}

func simpleMonitorListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{
			Name:    "Target",
			Sources: []string{"Status.Target"},
		},
		{
			Name:    "Protocol",
			Sources: []string{"Settings.SimpleMonitor.HealthCheck.Protocol"},
		},
		{
			Name:    "Email",
			Sources: []string{"Settings.SimpleMonitor.NotifyEmail.Enabled"},
		},
		{
			Name:    "Slack",
			Sources: []string{"Settings.SimpleMonitor.NotifySlack.Enabled"},
		},
	}
}

func simpleMonitorDetailIncludes() []string {
	return []string{}
}

func simpleMonitorDetailExcludes() []string {
	return []string{}
}

var allowSimpleMonitorProtocol = []string{"http", "https", "ping", "tcp", "dns", "ssh", "smtp", "pop3", "ssl-certificate"}

func simpleMonitorCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"target": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set monitoring target IP or Hostname",
			Required:    true,
			Category:    "monitor",
			Order:       10,
		},
		"protocol": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set monitoring protocol[http/https/ping/tcp/dns/ssh/smtp/pop3/ssl-certificate]",
			// TODO SNMP is not supported on current version.
			ValidateFunc: validateInStrValues(allowSimpleMonitorProtocol...),
			CompleteFunc: completeInStrValues(allowSimpleMonitorProtocol...),
			Required:     true,
			DefaultValue: "ping",
			Category:     "health-check",
			Order:        10,
		},
		"port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set port of tcp monitoring",
			ValidateFunc: validateIntRange(1, 65535),
			Category:     "health-check",
			Order:        20,
		},
		"delay-loop": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set delay-loop of monitoring(minute)",
			ValidateFunc: validateIntRange(1, 60),
			DefaultValue: 1,
			Required:     true,
			Category:     "health-check",
			Order:        30,
		},
		"disabled": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set monitoring disable",
			Category:    "health-check",
			Order:       40,
		},
		"host-header": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set host header of http/https monitoring request",
			Category:    "http-check",
			Order:       10,
		},
		"path": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set path of http/https monitoring request",
			Category:    "http-check",
			Order:       20,
		},
		"response-code": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "set response-code of http/https monitoring request",
			Category:    "http-check",
			Order:       30,
		},

		"dns-qname": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set DNS query target name",
			Category:    "dns-check",
			Order:       10,
		},
		"dns-excepted": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set DNS query excepted value",
			Category:    "dns-check",
			Order:       20,
		},
		"remaining-days": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set SSL-Certificate remaining days",
			ValidateFunc: validateIntRange(1, 9999),
			DefaultValue: 30,
			Category:     "ssl-check",
			Order:        10,
		},
		"notify-email": {
			Type:         schema.TypeBool,
			HandlerType:  schema.HandlerNoop,
			Description:  "enable e-mail notification",
			DefaultValue: true,
			Category:     "notify",
			Order:        10,
		},
		"email-type": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set e-mail type",
			ValidateFunc: validateInStrValues("text", "html"),
			CompleteFunc: completeInStrValues("text", "html"),
			DefaultValue: "text",
			Category:     "notify",
			Order:        20,
		},
		"slack-webhook": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set slack-webhook URL",
			Category:    "notify",
			Order:       30,
		},
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func simpleMonitorReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func simpleMonitorUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"protocol": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set monitoring protocol[http/https/ping/tcp/dns/ssh/smtp/pop3/ssl-certificate]",
			// TODO SNMP is not supported on current version.
			ValidateFunc: validateInStrValues(allowSimpleMonitorProtocol...),
			CompleteFunc: completeInStrValues(allowSimpleMonitorProtocol...),
			Category:     "health-check",
			Order:        10,
		},
		"port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set port of tcp monitoring",
			ValidateFunc: validateIntRange(1, 65535),
			Category:     "health-check",
			Order:        20,
		},
		"delay-loop": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set delay-loop of monitoring(minute)",
			ValidateFunc: validateIntRange(1, 60),
			Category:     "health-check",
			Order:        30,
		},
		"disabled": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set monitoring enable/disable",
			Category:    "health-check",
			Order:       40,
		},
		"host-header": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set host header of http/https monitoring request",
			Category:    "http-check",
			Order:       10,
		},
		"path": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set path of http/https monitoring request",
			Category:    "http-check",
			Order:       20,
		},
		"response-code": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "set response-code of http/https monitoring request",
			Category:    "http-check",
			Order:       30,
		},
		"dns_qname": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set DNS query target name",
			Category:    "dns-check",
			Order:       10,
		},
		"dns_excepted": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set DNS query excepted value",
			Category:    "dns-check",
			Order:       20,
		},
		"remaining-days": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set SSL-Certificate remaining days",
			ValidateFunc: validateIntRange(1, 9999),
			Category:     "ssl-check",
			Order:        10,
		},
		"notify-email": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "enable e-mail notification",
			Category:    "notify",
			Order:       10,
		},
		"email-type": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set e-mail type",
			ValidateFunc: validateInStrValues("text", "html"),
			CompleteFunc: completeInStrValues("text", "html"),
			Category:     "notify",
			Order:        20,
		},
		"slack-webhook": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set slack-webhook URL",
			Category:    "notify",
			Order:       30,
		},
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     paramIconResourceID,
	}
}

func simpleMonitorDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

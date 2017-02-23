package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func SimpleMonitorResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:                schema.CommandList,
			ListResultFieldName: "SimpleMonitors",
			Aliases:             []string{"l", "ls", "find"},
			Params:              simpleMonitorListParam(),
			TableType:           output.TableSimple,
			TableColumnDefines:  simpleMonitorListColumns(),
		},
		"create": {
			Type:             schema.CommandCreate,
			Aliases:          []string{"c"},
			Params:           simpleMonitorCreateParam(),
			IncludeFields:    simpleMonitorDetailIncludes(),
			ExcludeFields:    simpleMonitorDetailExcludes(),
			UseCustomCommand: true,
		},
		"read": {
			Type:          schema.CommandRead,
			Aliases:       []string{"r"},
			Params:        simpleMonitorReadParam(),
			IncludeFields: simpleMonitorDetailIncludes(),
			ExcludeFields: simpleMonitorDetailExcludes(),
		},
		"update": {
			Type:             schema.CommandUpdate,
			Aliases:          []string{"u"},
			Params:           simpleMonitorUpdateParam(),
			IncludeFields:    simpleMonitorDetailIncludes(),
			ExcludeFields:    simpleMonitorDetailExcludes(),
			UseCustomCommand: true,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"d", "rm"},
			Params:        simpleMonitorDeleteParam(),
			IncludeFields: simpleMonitorDetailIncludes(),
			ExcludeFields: simpleMonitorDetailExcludes(),
		},
	}

	return &schema.Resource{
		Commands: commands,
	}
}

func simpleMonitorListParam() map[string]*schema.Schema {
	return CommonListParam
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

func simpleMonitorCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"target": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set monitoring target IP or Hostname",
			Required:    true,
		},
		"protocol": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set monitoring protocol[http/https/ping/tcp/dns/ssh/smtp/pop3]",
			// TODO SNMP is not supported on current version.
			ValidateFunc: validateInStrValues("http", "https", "ping", "tcp", "dns", "ssh", "smtp", "pop3"),
			Required:     true,
			DefaultValue: "ping",
		},
		"host-header": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set host header of http/https monitoring request",
		},
		"path": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set path of http/https monitoring request",
		},
		"response-code": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "set response-code of http/https monitoring request",
		},
		"port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set port of tcp monitoring",
			ValidateFunc: validateIntRange(1, 65535),
		},
		"delay-loop": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set delay-loop of monitoring(minute)",
			ValidateFunc: validateIntRange(1, 60),
			DefaultValue: 1,
			Required:     true,
		},
		"enabled": {
			Type:         schema.TypeBool,
			HandlerType:  schema.HandlerNoop,
			Description:  "set monitoring enable/disable",
			DefaultValue: true,
			Required:     true,
		},
		"dns-qname": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set DNS query target name",
		},
		"dns-excepted": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set DNS query excepted value",
		},
		"notify-email": {
			Type:         schema.TypeBool,
			HandlerType:  schema.HandlerNoop,
			Description:  "enable e-mail notification",
			DefaultValue: true,
		},
		"email-type": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set e-mail type",
			ValidateFunc: validateInStrValues("text", "html"),
		},
		"slack-webhook": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set slack-webhook URL",
		},
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     getParamSubResourceID("Icon"),
	}
}

func simpleMonitorReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func simpleMonitorUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
		"protocol": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set monitoring protocol[http/https/ping/tcp/dns/ssh/smtp/pop3]",
			// TODO SNMP is not supported on current version.
			ValidateFunc: validateInStrValues("http", "https", "ping", "tcp", "dns", "ssh", "smtp", "pop3"),
		},
		"host-header": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set host header of http/https monitoring request",
		},
		"path": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set path of http/https monitoring request",
		},
		"response-code": {
			Type:        schema.TypeInt,
			HandlerType: schema.HandlerNoop,
			Description: "set response-code of http/https monitoring request",
		},
		"port": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set port of tcp monitoring",
			ValidateFunc: validateIntRange(1, 65535),
		},
		"delay-loop": {
			Type:         schema.TypeInt,
			HandlerType:  schema.HandlerNoop,
			Description:  "set delay-loop of monitoring(minute)",
			ValidateFunc: validateIntRange(1, 60),
		},
		"enabled": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "set monitoring enable/disable",
		},
		"dns_qname": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set DNS query target name",
		},
		"dns_excepted": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set DNS query excepted value",
		},
		"notify-email": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "enable e-mail notification",
		},
		"email-type": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set e-mail type",
			ValidateFunc: validateInStrValues("text", "html"),
		},
		"slack-webhook": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set slack-webhook URL",
		},
		"description": paramDescription,
		"tags":        paramTags,
		"icon-id":     getParamSubResourceID("Icon"),
	}
}

func simpleMonitorDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

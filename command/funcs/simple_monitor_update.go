package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func SimpleMonitorUpdate(ctx command.Context, params *params.UpdateSimpleMonitorParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSimpleMonitorAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("SimpleMonitorUpdate is failed: %s", e)
	}

	protocol := p.Settings.SimpleMonitor.HealthCheck.Protocol
	if ctx.IsSet("protocol") {
		protocol = params.Protocol
	}

	// set params
	switch protocol {
	case "http", "https":

		port := p.Settings.SimpleMonitor.HealthCheck.Port
		if ctx.IsSet("port") && params.Port > 0 {
			port = fmt.Sprintf("%d", params.Port)
		}

		path := p.Settings.SimpleMonitor.HealthCheck.Path
		if ctx.IsSet("path") {
			path = params.Path
		}

		responseCode := p.Settings.SimpleMonitor.HealthCheck.Status
		if ctx.IsSet("response-code") && params.ResponseCode > 0 {
			responseCode = fmt.Sprintf("%d", params.ResponseCode)
		}

		hostHeader := p.Settings.SimpleMonitor.HealthCheck.Host
		if ctx.IsSet("host-header") {
			hostHeader = params.HostHeader
		}

		// set health check
		setHealthCheck := p.SetHealthCheckHTTP
		if params.Protocol == "https" {
			setHealthCheck = p.SetHealthCheckHTTPS
		}

		if path == "" {
			return fmt.Errorf("path is required when protocol is http/https")
		}
		setHealthCheck(
			port,
			path,
			responseCode,
			hostHeader,
		)

	case "ping":
		p.SetHealthCheckPing()
	case "tcp", "ssh", "smtp", "pop3":
		port := p.Settings.SimpleMonitor.HealthCheck.Port
		if ctx.IsSet("port") {
			port = fmt.Sprintf("%d", params.Port)
		}
		if port == "0" {
			return fmt.Errorf("port is required when protocol is tcp/ssh/smtp/pop3")
		}

		var setHealchCheck func(string)
		switch params.Protocol {
		case "tcp":
			setHealchCheck = p.SetHealthCheckTCP
		case "ssh":
			setHealchCheck = p.SetHealthCheckSSH
		case "smtp":
			setHealchCheck = p.SetHealthCheckSMTP
		case "pop3":
			setHealchCheck = p.SetHealthCheckPOP3
		}

		setHealchCheck(port)

	case "dns":

		qname := p.Settings.SimpleMonitor.HealthCheck.QName
		if ctx.IsSet("dns-qname") {
			qname = params.DnsQname
		}

		excepted := p.Settings.SimpleMonitor.HealthCheck.ExpectedData
		if ctx.IsSet("dns-excepted") {
			excepted = params.DnsExcepted
		}

		if qname == "" {
			return fmt.Errorf("dns-qname is required when protocol is dns")
		}
		p.SetHealthCheckDNS(qname, excepted)
	}

	if ctx.IsSet("delay-loop") {
		p.SetDelayLoop(params.DelayLoop * 60)
	}

	if ctx.IsSet("disabled") {

		enabled := "True"
		if params.Disabled {
			enabled = "False"
		}
		p.Settings.SimpleMonitor.Enabled = enabled
	}

	if ctx.IsSet("notify-email") {
		if params.NotifyEmail {
			if params.EmailType == "" {
				return fmt.Errorf("email-type is required when notify-email is true")
			}
			html := false
			if params.EmailType == "html" {
				html = true
			}
			p.EnableNotifyEmail(html)
		} else {
			p.DisableNotifyEmail()
		}
	}

	if ctx.IsSet("slack-webhook") {
		if params.SlackWebhook != "" {
			p.EnableNofitySlack(params.SlackWebhook)
		} else {
			p.DisableNotifySlack()
		}
	}

	if ctx.IsSet("description") {
		p.SetDescription(params.Description)
	}
	if ctx.IsSet("tags") {
		p.SetTags(params.Tags)
	}
	if ctx.IsSet("icon-id") {
		p.SetIconByID(params.IconId)
	}

	// call Update(id)
	res, err := api.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("SimpleMonitorUpdate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}

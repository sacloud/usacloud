package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func SimpleMonitorCreate(ctx command.Context, params *params.CreateSimpleMonitorParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSimpleMonitorAPI()
	p := api.New(params.Target)

	p.SetTarget(params.Target)
	switch params.Protocol {
	case "http", "https":
		if params.Path == "" {
			return fmt.Errorf("path is required when protocol is http/https")
		}
		port := ""
		if params.Port > 0 {
			port = fmt.Sprintf("%d", params.Port)
		}
		responseCode := ""
		if params.ResponseCode > 0 {
			responseCode = fmt.Sprintf("%d", params.ResponseCode)
		}

		// set health check
		if params.Protocol == "http" {
			p.SetHealthCheckHTTP(
				port,
				params.Path,
				responseCode,
				params.HostHeader,
				params.Username,
				params.Password,
			)
		} else {
			p.SetHealthCheckHTTPS(
				port,
				params.Path,
				responseCode,
				params.HostHeader,
				params.Sni,
				params.Username,
				params.Password,
			)
		}

	case "ping":
		p.SetHealthCheckPing()
	case "tcp", "ssh", "smtp", "pop3":
		if params.Port == 0 {
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

		setHealchCheck(fmt.Sprintf("%d", params.Port))

	case "dns":
		if params.DnsQname == "" {
			return fmt.Errorf("dns-qname is required when protocol is dns")
		}
		p.SetHealthCheckDNS(params.DnsQname, params.DnsExcepted)
	case "ssl-certificate":
		p.SetHealthCheckSSLCertificate(params.RemainingDays)
	}

	p.SetDelayLoop(params.DelayLoop * 60)

	enabled := "True"
	if params.Disabled {
		enabled = "False"
	}
	p.Settings.SimpleMonitor.Enabled = enabled

	if params.NotifyEmail {
		if params.EmailType == "" {
			return fmt.Errorf("email-type is required when notify-email is true")
		}
		html := false
		if params.EmailType == "html" {
			html = true
		}
		p.EnableNotifyEmail(html)
	}

	if params.SlackWebhook != "" {
		p.EnableNofitySlack(params.SlackWebhook)
	}

	p.SetIconByID(params.IconId)
	p.SetDescription(params.Description)
	p.SetTags(params.Tags)

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("SimpleMonitorCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}

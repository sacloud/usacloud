package funcs

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/helper/printer"
	"time"
)

func VPCRouterLogs(ctx command.Context, params *params.LogsVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()

	// validate params
	if params.Follow && params.LogName == "all" {
		fmt.Fprintf(command.GlobalOption.Err, "[WARN] -f/--follow option can not use with --log-name=%q, ignored", "all")
		params.Follow = false
	}
	if ctx.IsSet("refresh-interval") && params.LogName == "all" {
		fmt.Fprintf(command.GlobalOption.Err, "[WARN] --refresh-interval option can not use with --log-name=%q, ignored", "all")
	}

	logBuf := internal.NewHashQueue(500)
	out := command.GlobalOption.Out

	if params.ListLogNames {
		for _, name := range define.AllowVPCRouterLogNames {
			fmt.Fprintln(out, name)
		}
		return nil
	}

	for {
		// call Read(id)
		res, err := api.Status(params.Id)
		if err != nil {
			return fmt.Errorf("VPCRouterLog is failed: %s", err)
		}

		logs := map[string][]string{}
		switch params.LogName {
		case "all":
			logs["vpn"] = res.VPNLogs
			logs["firewall-send"] = res.FirewallSendLogs
			logs["firewall-receive"] = res.FirewallReceiveLogs
		case "vpn":
			logs["vpn"] = res.VPNLogs
		case "firewall-send":
			logs["firewall-send"] = res.FirewallSendLogs
		case "firewall-receive":
			logs["firewall-receive"] = res.FirewallReceiveLogs
		}

		for key, lines := range logs {
			if params.LogName == "all" {
				printer.Fprintf(out, color.New(color.FgHiGreen), "\n==> [%s]:start\n", key)
			}
			for _, line := range lines {
				if logBuf.PutIfAbsent(line) {
					fmt.Fprintf(out, "%s\n", line)
				}
			}
			if params.LogName == "all" {
				printer.Fprintf(out, color.New(color.FgHiGreen), "\n<== [%s]:end\n", key)
			}
		}

		if !params.Follow {
			break
		}
		time.Sleep(time.Duration(params.RefreshInterval) * time.Second)
	}
	return nil
}

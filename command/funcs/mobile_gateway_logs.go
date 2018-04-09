package funcs

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
)

func MobileGatewayLogs(ctx command.Context, params *params.LogsMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()

	// call Read(id)
	_, err := api.Read(params.Id)
	if err != nil {
		return fmt.Errorf("MobileGatewayLogs is failed: %s", err)
	}

	logBuf := internal.NewHashQueue(500)
	out := command.GlobalOption.Out

	if params.Follow {
		for {
			// call Read(id)
			logs, err := api.Logs(params.Id, nil)
			if err != nil {
				return fmt.Errorf("MobileGatewayLogs is failed: %s", err)
			}

			for _, log := range logs {
				data, err := json.MarshalIndent(log, "", "\t")
				if err != nil {
					return fmt.Errorf("MobileGatewayLogs is failed: %s", err)
				}
				line := string(data)
				if logBuf.PutIfAbsent(line) {
					fmt.Fprintf(out, "%s\n", line)
				}
			}
			time.Sleep(time.Duration(params.RefreshInterval) * time.Second)
		}
	} else {
		// call Read(id)
		logs, err := api.Logs(params.Id, nil)
		if err != nil {
			return fmt.Errorf("MobileGatewayLogs is failed: %s", err)
		}

		if len(logs) == 0 {
			fmt.Fprintf(command.GlobalOption.Err, "Result is empty\n")
			return nil
		}

		data, err := json.MarshalIndent(logs, "", "\t")
		if err != nil {
			return fmt.Errorf("MobileGatewayLogs is failed: %s", err)
		}
		fmt.Fprintf(out, "%s\n", string(data))
	}

	return nil

}

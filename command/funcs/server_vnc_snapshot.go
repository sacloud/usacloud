package funcs

import (
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
)

func ServerVncSnapshot(ctx command.Context, params *params.VncSnapshotServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerVncSnapshot is failed: %s", e)
	}
	if !p.IsUp() {
		return fmt.Errorf("ServerVncSnapshot is failed: %s", "server is not running")
	}

	if !p.IsUp() && params.WaitForBoot {

		err := internal.ExecWithProgress(
			fmt.Sprintf("Still booting[ID:%d]...", params.Id),
			fmt.Sprintf("Connect to server[ID:%d]", params.Id),
			command.GlobalOption.Progress,
			func(compChan chan bool, errChan chan error) {
				// call manipurate functions
				err := api.SleepUntilUp(params.Id, client.DefaultTimeoutDuration)
				if err != nil {
					errChan <- err
					return
				}
				compChan <- true
			},
		)
		if err != nil {
			return fmt.Errorf("ServerVncSnapshot is failed: %s", e)
		}
	}

	snapshotReq := api.NewVNCSnapshotRequest()
	vncSnapshotResponse, err := api.GetVNCSnapshot(params.Id, snapshotReq)
	if err != nil {
		return fmt.Errorf("ServerVncSnapshot is failed: %s", err)
	}
	vncImage, err := base64.StdEncoding.DecodeString(vncSnapshotResponse.Image)
	if err != nil {
		return fmt.Errorf("ServerVncSnapshot is failed: %s", err)
	}
	filename := fmt.Sprintf("%d_%s.gif", params.Id, time.Now().Format("20060102-150405"))
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("ServerVncSnapshot is failed: %s", err)
	}
	defer file.Close()

	_, err = file.Write(vncImage)
	if err != nil {
		return fmt.Errorf("ServerVncSnapshot is failed: %s", err)
	}

	out := command.GlobalOption.Out
	fmt.Fprintln(out, "Snapshot created:", filename)

	return nil
}

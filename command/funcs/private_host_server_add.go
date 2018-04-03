package funcs

import (
	"fmt"

	sacloudAPI "github.com/sacloud/libsacloud/api"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func PrivateHostServerAdd(ctx command.Context, params *params.ServerAddPrivateHostParam) error {

	client := ctx.GetAPIClient()
	api := client.GetPrivateHostAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("PrivateHostServerAdd is failed: %s", e)
	}

	// check server status
	server, err := client.GetServerAPI().Read(params.ServerId)
	if err != nil {
		if notFoundErr, ok := err.(sacloudAPI.Error); ok {
			if notFoundErr.ResponseCode() == 404 {
				return fmt.Errorf("PrivateHostServerAdd is failed: Server[%d] is not found", params.ServerId)
			}
		}
		return fmt.Errorf("PrivateHostServerAdd is failed: %s", err)
	}
	if server.IsUp() {
		return fmt.Errorf("PrivateHostServerAdd is failed: Server[%d] is running", params.ServerId)
	}

	// update server
	server.SetPrivateHostByID(p.ID)
	_, err = client.GetServerAPI().Update(server.ID, server)
	if err != nil {
		return fmt.Errorf("PrivateHostServerAdd is failed: %s", err)
	}

	return ctx.GetOutput().Print(server)
}

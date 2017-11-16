package funcs

import (
	"fmt"
	sacloudAPI "github.com/sacloud/libsacloud/api"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func PrivateHostServerDelete(ctx command.Context, params *params.ServerDeletePrivateHostParam) error {

	client := ctx.GetAPIClient()
	api := client.GetPrivateHostAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("PrivateHostServerDelete is failed: %s", e)
	}

	// check server status
	server, err := client.GetServerAPI().Read(params.ServerId)
	if err != nil {
		if notFoundErr, ok := err.(sacloudAPI.Error); ok {
			if notFoundErr.ResponseCode() == 404 {
				return fmt.Errorf("PrivateHostServerDelete is failed: Server[%d] is not found", params.ServerId)
			}
		}
		return fmt.Errorf("PrivateHostServerDelete is failed: %s", err)
	}
	if server.IsUp() {
		return fmt.Errorf("PrivateHostServerDelete is failed: Server[%d] is running", params.ServerId)
	}
	if server.PrivateHost == nil {
		return fmt.Errorf("PrivateHostServerDelete is failed: Server[%d] is not on PrivateHost", params.ServerId)
	}
	if server.PrivateHost.ID != params.Id {
		return fmt.Errorf("PrivateHostServerDelete is failed: Server[%d] is on a different PrivateHost[%d]", params.ServerId, server.PrivateHost.ID)
	}

	// update server
	server.ClearPrivateHost()
	server, err = client.GetServerAPI().Update(server.ID, server)
	if err != nil {
		return fmt.Errorf("PrivateHostServerDelete is failed: %s", err)
	}

	return ctx.GetOutput().Print(server)

}

package command

import (
	"fmt"
	"github.com/sacloud/usacloud/command/internal"
)

func DiskEdit(ctx Context, params *EditDiskParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDiskAPI()
	p := api.NewCondig()

	// set params
	if ctx.IsSet("hostname") {
		p.SetHostName(params.Hostname)
	}
	if ctx.IsSet("password") {
		p.SetPassword(params.Password)
	}
	if ctx.IsSet("ssh-key") {
		p.SetSSHKeys(StringIDs(params.SshKeyIds))
	}
	if ctx.IsSet("disable-password-auth") {
		p.SetDisablePWAuth(params.DisablePasswordAuth)
	}
	if ctx.IsSet("startup-script") {
		p.SetNotes(StringIDs(params.StartupScriptIds))
	}
	if ctx.IsSet("ipaddress") {
		p.SetUserIPAddress(params.Ipaddress)
	}
	if ctx.IsSet("default-route") {
		p.SetDefaultRoute(params.DefaultRoute)
	}
	if ctx.IsSet("nw-masklen") {
		p.SetNetworkMaskLen(fmt.Sprintf("%d", params.NwMasklen))
	}

	// wait for copy with progress
	err := internal.ExecWithProgress(
		fmt.Sprintf("Still editing[ID:%d]...", params.Id),
		fmt.Sprintf("Edit disk[ID:%d]", params.Id),
		GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {
			// call manipurate functions
			_, err := api.Config(params.Id, p)
			if err != nil {
				errChan <- err
				return
			}
			compChan <- true
		},
	)

	if err != nil {
		return fmt.Errorf("DiskEdit is failed: %s", err)
	}

	// read
	res, err := api.Read(params.Id)
	if err != nil {
		return fmt.Errorf("DiskEdit is failed: %s", err)
	}
	return ctx.GetOutput().Print(res)
}

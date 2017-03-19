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
	spinner := internal.NewSpinner(
		"Editing...",
		"Edit disk is complete.\n",
		internal.CharSetProgress,
		GlobalOption.Progress)
	spinner.Start()
	compChan := make(chan bool)
	errChan := make(chan error)
	go func() {
		// call manipurate functions
		_, err := api.Config(params.Id, p)
		if err != nil {
			errChan <- err
			return
		}
		compChan <- true
	}()
edit:
	for {
		select {
		case <-compChan:
			spinner.Stop()
			break edit
		case err := <-errChan:
			return fmt.Errorf("DiskEdit is failed: %s", err)
		}
	}

	// read
	res, err := api.Read(params.Id)
	if err != nil {
		return fmt.Errorf("DiskEdit is failed: %s", err)
	}
	return ctx.GetOutput().Print(res)

}

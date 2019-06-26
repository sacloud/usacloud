package funcs

import (
	"errors"
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ProxyLBAcmeRenew(ctx command.Context, params *params.AcmeRenewProxyLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ProxyLBAcmeRenew is failed: %s", e)
	}

	if !p.Settings.ProxyLB.LetsEncrypt.Enabled {
		return errors.New("Let's Encrypt is disabbled. Please update ACME settings")
	}

	// call manipurate functions
	_, err := api.RenewLetsEncryptCert(params.Id)
	if err != nil {
		return fmt.Errorf("ProxyLBAcmeRenew is failed: %s", err)
	}
	return nil
}

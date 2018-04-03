package funcs

import (
	"fmt"
	"io/ioutil"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func WebAccelCertificateUpdate(ctx command.Context, params *params.CertificateUpdateWebAccelParam) error {

	client := ctx.GetAPIClient()
	api := client.GetWebAccelAPI()

	// param validate
	if params.Cert == "" && params.CertContent == "" {
		return fmt.Errorf("%q or %q is required", "cert", "cert-content")
	}

	// site state validate
	site, err := api.Read(fmt.Sprintf("%d", params.Id))
	if err != nil {
		return fmt.Errorf("Site[%q] not found", params.Id)
	}
	if !site.HasCertificate {
		msg := "Site[%q] don't have certificate. certificate-update command is supporting only update"
		return fmt.Errorf(msg, params.Id)
	}

	p := &sacloud.WebAccelCertRequest{}

	// set params
	if params.Cert != "" {
		b, err := ioutil.ReadFile(params.Cert)
		if err != nil {
			return fmt.Errorf("WebAccelCertificateUpdate is failed: %s", err)
		}
		p.CertificateChain = string(b)
	}
	if params.CertContent != "" {
		p.CertificateChain = params.CertContent
	}
	if params.Key != "" {
		b, err := ioutil.ReadFile(params.Key)
		if err != nil {
			return fmt.Errorf("WebAccelCertificateUpdate is failed: %s", err)
		}
		p.Key = string(b)
	}
	if params.KeyContent != "" {
		p.Key = params.CertContent
	}

	// Update
	res, err := api.UpdateCertificate(fmt.Sprintf("%d", params.Id), p)
	if err != nil {
		return fmt.Errorf("WebAccelCertificateUpdate is failed: %s", err)
	}
	return ctx.GetOutput().Print(res)

}

package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ProxyLBCertificateUpdate(ctx command.Context, params *params.CertificateUpdateProxyLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ProxyLBCertificateUpdate is failed: %s", e)
	}

	cert, err := api.GetCertificates(p.ID)
	if err != nil {
		return fmt.Errorf("ProxyLBCertificateUpdate is failed: %s", err)
	}
	if ctx.IsSet("server-certificate") {
		cert.ServerCertificate = params.ServerCertificate
	}
	if ctx.IsSet("intermediate-certificate") {
		cert.IntermediateCertificate = params.IntermediateCertificate
	}
	if ctx.IsSet("private-key") {
		cert.PrivateKey = params.PrivateKey
	}

	if _, err := api.SetCertificates(p.ID, cert); err != nil {
		return fmt.Errorf("ProxyLBCertificateUpdate is failed: %s", err)
	}

	// refresh
	cert, err = api.GetCertificates(p.ID)
	if err != nil {
		return fmt.Errorf("ProxyLBCertificateUpdate is failed: %s", err)
	}
	return ctx.GetOutput().Print(buildProxyLBCertInfo(cert))

}

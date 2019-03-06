package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ProxyLBCertificateAdd(ctx command.Context, params *params.CertificateAddProxyLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ProxyLBCertificateAdd is failed: %s", e)
	}

	cert := &sacloud.ProxyLBCertificates{
		ServerCertificate:       params.ServerCertificate,
		IntermediateCertificate: params.IntermediateCertificate,
		PrivateKey:              params.PrivateKey,
	}

	if _, err := api.SetCertificates(p.ID, cert); err != nil {
		return fmt.Errorf("ProxyLBCertificateAdd is failed: %s", err)
	}

	// refresh cert
	cert, err := api.GetCertificates(p.ID)
	if err != nil {
		return fmt.Errorf("ProxyLBCertificateAdd is failed: %s", err)
	}

	return ctx.GetOutput().Print(buildProxyLBCertInfo(cert))
}

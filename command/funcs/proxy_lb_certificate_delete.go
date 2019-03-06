package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ProxyLBCertificateDelete(ctx command.Context, params *params.CertificateDeleteProxyLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ProxyLBCertificateDelete is failed: %s", e)
	}

	cert, err := api.GetCertificates(p.ID)
	if err != nil {
		return fmt.Errorf("ProxyLBCertificateDelete is failed: %s", err)
	}

	if _, err := api.DeleteCertificates(p.ID); err != nil {
		return fmt.Errorf("ProxyLBCertificateDelete is failed: %s", err)
	}

	// refresh
	cert, err = api.GetCertificates(p.ID)
	if err != nil {
		return fmt.Errorf("ProxyLBCertificateDelete is failed: %s", err)
	}

	var list []interface{}
	if cert.ServerCertificate != "" && cert.PrivateKey != "" {
		list = append(list, buildProxyLBCertInfo(cert))
	}

	return ctx.GetOutput().Print(list...)
}

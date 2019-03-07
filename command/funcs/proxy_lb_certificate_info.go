package funcs

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strings"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ProxyLBCertificateInfo(ctx command.Context, params *params.CertificateInfoProxyLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()
	if _, e := api.Read(params.Id); e != nil {
		return fmt.Errorf("ProxyLBCertificateInfo is failed: %s", e)
	}

	cert, err := api.GetCertificates(params.Id)
	if err != nil {
		return fmt.Errorf("ProxyLBCertificateInfo is failed: %s", err)
	}

	var list []interface{}
	if cert.ServerCertificate != "" && cert.PrivateKey != "" {
		list = append(list, buildProxyLBCertInfo(cert))
	}

	return ctx.GetOutput().Print(list...)
}

type proxyLBCertInfo struct {
	*sacloud.ProxyLBCertificates
	CommonName string
	AltNames   string
	Issuer     string
}

func buildProxyLBCertInfo(cert *sacloud.ProxyLBCertificates) *proxyLBCertInfo {
	ci := &proxyLBCertInfo{ProxyLBCertificates: cert}
	block, _ := pem.Decode([]byte(cert.ServerCertificate)) // ignore err
	if block != nil {
		c, err := x509.ParseCertificate(block.Bytes) // ignore err
		if err == nil {

			ci.CommonName = c.Subject.CommonName
			ci.Issuer = c.Issuer.CommonName
			var altNames []string
			for _, ip := range c.IPAddresses {
				altNames = append(altNames, ip.String())
			}

			for _, en := range c.Subject.ExtraNames {
				altNames = append(altNames, en.Value.(string))
			}

			ci.AltNames = strings.Join(altNames, ",")
		}
	}
	return ci
}

// Copyright 2017-2020 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package proxylb

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strings"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func CertificateInfo(ctx cli.Context, params *params.CertificateInfoProxyLBParam) error {

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
	if cert.PrimaryCert != nil && cert.PrimaryCert.ServerCertificate != "" && cert.PrimaryCert.PrivateKey != "" {
		list = append(list, buildProxyLBCertInfo(cert))
	}

	return ctx.Output().Print(list...)
}

type proxyLBCertInfo struct {
	*sacloud.ProxyLBCertificates
	CommonName string
	AltNames   string
	Issuer     string
}

func buildProxyLBCertInfo(cert *sacloud.ProxyLBCertificates) *proxyLBCertInfo {
	ci := &proxyLBCertInfo{ProxyLBCertificates: cert}
	block, _ := pem.Decode([]byte(cert.PrimaryCert.ServerCertificate)) // ignore err
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

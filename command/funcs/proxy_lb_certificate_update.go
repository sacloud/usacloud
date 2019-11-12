// Copyright 2017-2019 The Usacloud Authors
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

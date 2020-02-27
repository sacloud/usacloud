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

package funcs

import (
	"fmt"
	"io/ioutil"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func WebAccelCertificateNew(ctx command.Context, params *params.CertificateNewWebAccelParam) error {

	client := ctx.GetAPIClient()
	api := client.GetWebAccelAPI()

	// param validate
	if params.Cert == "" && params.CertContent == "" {
		return fmt.Errorf("%q or %q is required", "cert", "cert-content")
	}

	// site state validate
	_, err := api.Read(fmt.Sprintf("%d", params.Id))
	if err != nil {
		return fmt.Errorf("Site[%q] not found", params.Id)
	}

	p := &sacloud.WebAccelCertRequest{}

	// set params
	if params.Cert != "" {
		b, err := ioutil.ReadFile(params.Cert)
		if err != nil {
			return fmt.Errorf("WebAccelCertificateNew is failed: %s", err)
		}
		p.CertificateChain = string(b)
	}
	if params.CertContent != "" {
		p.CertificateChain = params.CertContent
	}
	if params.Key != "" {
		b, err := ioutil.ReadFile(params.Key)
		if err != nil {
			return fmt.Errorf("WebAccelCertificateNew is failed: %s", err)
		}
		p.Key = string(b)
	}
	if params.KeyContent != "" {
		p.Key = params.CertContent
	}

	// Update
	res, err := api.CreateCertificate(params.Id, p)
	if err != nil {
		return fmt.Errorf("WebAccelCertificateNew is failed: %s", err)
	}
	return ctx.GetOutput().Print(res)

}

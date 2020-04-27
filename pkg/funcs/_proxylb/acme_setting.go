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
	"errors"
	"fmt"

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func ACMESetting(ctx cli.Context, params *params.ACMESettingProxyLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ProxyLBACMESetting is failed: %s", e)
	}

	// validate params
	if !params.Disable && !params.AcceptTos {
		return errors.New("--accept-tos=true is required to enable Let's Encrypt setting")
	}

	// set params
	if params.Disable {
		p.Settings.ProxyLB.LetsEncrypt.Enabled = false
		p.Settings.ProxyLB.LetsEncrypt.CommonName = ""
	} else if params.AcceptTos {
		p.Settings.ProxyLB.LetsEncrypt.Enabled = true
		p.Settings.ProxyLB.LetsEncrypt.CommonName = params.CommonName
	}

	// call manipurate functions
	res, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("ProxyLBACMESetting is failed: %s", err)
	}
	return ctx.Output().Print(&res.Settings.ProxyLB.LetsEncrypt)
}

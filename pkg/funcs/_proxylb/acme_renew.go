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

func ACMERenew(ctx cli.Context, params *params.ACMERenewProxyLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ProxyLBACMERenew is failed: %s", e)
	}

	if !p.Settings.ProxyLB.LetsEncrypt.Enabled {
		return errors.New("let's encrypt setting is disabled. Please update the settings")
	}

	// call manipurate functions
	_, err := api.RenewLetsEncryptCert(params.Id)
	if err != nil {
		return fmt.Errorf("ProxyLBACMERenew is failed: %s", err)
	}
	return nil
}

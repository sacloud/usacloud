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

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func SSHKeyCreate(ctx cli.Context, params *params.CreateSSHKeyParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSSHKeyAPI()
	p := api.New()

	// validate
	if params.PublicKey == "" && params.PublicKeyContent == "" {
		return fmt.Errorf("%q or %q is required", "public-key", "public-key-content")
	}

	// set params
	if params.PublicKey != "" {
		b, err := ioutil.ReadFile(params.PublicKey)
		if err != nil {
			return fmt.Errorf("SSHKeyCreate is failed: %s", err)
		}
		p.PublicKey = string(b)
	}

	if params.PublicKeyContent != "" {
		p.PublicKey = params.PublicKeyContent
	}

	p.SetName(params.Name)
	p.SetDescription(params.Description)

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("SSHKeyCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}

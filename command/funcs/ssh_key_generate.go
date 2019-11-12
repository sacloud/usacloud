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
	"io"
	"os"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func SSHKeyGenerate(ctx command.Context, params *params.GenerateSSHKeyParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSSHKeyAPI()

	// call manipurate functions
	key, err := api.Generate(params.Name, params.PassPhrase, params.Description)
	if err != nil {
		return fmt.Errorf("SSHKeyGenerate is failed: %s", err)
	}

	var w io.Writer
	if params.PrivateKeyOutput != "" {
		// file
		f, err := os.OpenFile(params.PrivateKeyOutput, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			return fmt.Errorf("SSHKeyGenerate is failed: %s", err)
		}
		w = f
		defer f.Close()

		_, err = w.Write([]byte(key.PrivateKey))
		if err != nil {
			return fmt.Errorf("SSHKeyGenerate is failed: %s", err)
		}
	}

	err = ctx.GetOutput().Print(key)
	if err != nil {
		return fmt.Errorf("SSHKeyGenerate is failed: %s", err)
	}

	if params.PrivateKeyOutput == "" {
		// output privatekey to os.StdOut
		fmt.Fprintf(command.GlobalOption.Out, sshPrivateKeyStdOutFormat, key.PrivateKey)
	}

	return err
}

var sshPrivateKeyStdOutFormat = `
=======================================
Please save the following private-key.
=======================================

%s
`

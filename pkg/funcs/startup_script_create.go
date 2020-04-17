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

func StartupScriptCreate(ctx cli.Context, params *params.CreateStartupScriptParam) error {

	client := ctx.GetAPIClient()
	api := client.GetNoteAPI()
	p := api.New()

	// set params
	p.SetName(params.Name)
	p.SetClassByStr(params.Class)
	p.SetTags(params.Tags)
	p.SetIconByID(params.IconId)

	if params.Script != "" {
		b, err := ioutil.ReadFile(params.Script)
		if err != nil {
			return fmt.Errorf("StartupScriptCreate is failed: %s", err)
		}
		p.Content = string(b)
	}

	if params.ScriptContent != "" {
		p.Content = params.ScriptContent
	}

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("StartupScriptCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)
}

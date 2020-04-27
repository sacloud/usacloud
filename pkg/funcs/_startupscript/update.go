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

package startupscript

import (
	"fmt"
	"io/ioutil"

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func Update(ctx cli.Context, params *params.UpdateStartupScriptParam) error {

	client := ctx.GetAPIClient()
	api := client.GetNoteAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("StartupScriptUpdate is failed: %s", e)
	}

	// set params

	if ctx.IsSet("tags") {
		p.SetTags(params.Tags)
	}
	if ctx.IsSet("icon-id") {
		p.SetIconByID(params.IconId)
	}
	if ctx.IsSet("name") {
		p.SetName(params.Name)
	}
	if ctx.IsSet("class") {
		p.SetClassByStr(params.Class)
	}

	if ctx.IsSet("script") {
		b, err := ioutil.ReadFile(params.Script)
		if err != nil {
			return fmt.Errorf("StartupScriptUpdate is failed: %s", err)
		}
		p.Content = string(b)
	}

	if ctx.IsSet("script-content") {
		p.Content = params.ScriptContent
	}

	// call Update(id)
	res, err := api.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("StartupScriptUpdate is failed: %s", err)
	}

	return ctx.Output().Print(res)

}

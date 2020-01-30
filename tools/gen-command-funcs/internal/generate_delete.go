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

package internal

import (
	"bytes"
	"text/template"

	"github.com/sacloud/usacloud/schema"
	"github.com/sacloud/usacloud/tools"
)

func GenerateDeleteCommand(ctx *tools.GenerateContext, command *schema.Command) (string, error) {
	b := bytes.NewBufferString("")
	t := template.New("c")
	template.Must(t.Parse(deleteCommandTemplate))

	setParamActions, err := generateSetParamActions(ctx, command)
	if err != nil {
		return "", err
	}

	err = t.Execute(b, map[string]interface{}{
		"ResourceName":    ctx.CommandResourceName(),
		"SetParamActions": setParamActions,
		"FuncName":        ctx.CommandFuncName(),
	})
	return b.String(), err
}

var deleteCommandTemplate = `
	client := ctx.GetAPIClient()
	api := client.Get{{.ResourceName}}API()

	// set params
	{{.SetParamActions}}

	// call Delete(id)
	res, err := api.Delete(params.Id)
	if err != nil {
		return fmt.Errorf("{{.FuncName}} is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)
`

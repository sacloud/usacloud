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

package main

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/sacloud/usacloud/pkg/schema"
	"github.com/sacloud/usacloud/tools"
)

func generateSetParamActions(ctx *tools.GenerateContext, command *tools.Command) (string, error) {
	b := bytes.NewBufferString("")

	for _, p := range command.Params {
		if p.Name == "id" {
			continue
		}

		t := template.New("c")
		template.Must(t.Parse(setParamTemplates[p.HandlerType]))

		customFunc := fmt.Sprintf(`params.GetCommandDef().Params["%s"].CustomHandler`, p.Name)
		needIsSetCheck := command.Type == schema.CommandUpdate

		err := t.Execute(b, map[string]interface{}{
			"FlagName":       p.FlagName(),
			"ParamName":      p.FieldName(),
			"SetterFuncName": p.SetterFuncName(),
			"CustomFunc":     customFunc,
			"NeedIsSetCheck": needIsSetCheck,
		})
		if err != nil {
			return "", err
		}
	}
	return b.String(), nil
}

var setParamTemplates = map[schema.HandlerType]string{
	schema.HandlerPathThrough: `
	{{if .NeedIsSetCheck}}if params.Changed("{{.FlagName}}") { {{ end }}
		p.{{.SetterFuncName}}(params.{{.ParamName}})
	{{if .NeedIsSetCheck}} } {{ end }}`,
	schema.HandlerPathThroughEach: `
	{{if .NeedIsSetCheck}}if ctx.IsSet("{{.FlagName}}") { {{ end }}
		for _ , v := range params.{{.ParamName}} {
			p.{{.SetterFuncName}}(v)
		}
	{{if .NeedIsSetCheck}} } {{ end }}`,
	schema.HandlerNoop: "",
}

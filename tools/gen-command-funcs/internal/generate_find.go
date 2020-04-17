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
	"fmt"
	"text/template"

	"github.com/sacloud/usacloud/pkg/schema"
	"github.com/sacloud/usacloud/tools"
)

func GenerateFindCommand(ctx *tools.GenerateContext, command *tools.Command) (string, error) {
	b := bytes.NewBufferString("")
	t := template.New("c")
	template.Must(t.Parse(findCommandTemplate))

	setParamActions, err := generateFindSetParamActions(ctx, command)
	if err != nil {
		return "", err
	}

	filterActions, err := generateFilterActions(ctx, command)
	if err != nil {
		return "", err
	}

	err = t.Execute(b, map[string]interface{}{
		"FinderFieldName": command.ResourceName(),
		"SetParamActions": setParamActions,
		"FuncName":        command.FuncName(),
		"ListResultField": command.FindResultFieldName(),
		"FilterActions":   filterActions,
	})
	return b.String(), err
}

var findCommandTemplate = `
	client := ctx.GetAPIClient()
	finder := client.Get{{.FinderFieldName}}API()

	finder.SetEmpty()

	{{.SetParamActions}}

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return fmt.Errorf("{{.FuncName}} is failed: %s", err)
	}

	list := []interface{}{}
	for i, _ := range res.{{.ListResultField}}{
		{{.FilterActions}}
		list = append(list, &res.{{.ListResultField}}[i])
	}
	return ctx.GetOutput().Print(list...)
`

func generateFindSetParamActions(ctx *tools.GenerateContext, command *tools.Command) (string, error) {

	b := bytes.NewBufferString("")

	for _, p := range command.Params {
		if p.HandlerType == schema.HandlerFilterFunc {
			continue
		}

		t := template.New("c")
		template.Must(t.Parse(findSetParamTemplates[p.HandlerType]))

		customHandlerName := fmt.Sprintf(`params.GetCommandDef().Params["%s"].CustomHandler`, p.Name)

		err := t.Execute(b, map[string]interface{}{
			"ParamName":         p.FieldName(),
			"SetterFuncName":    p.SetterFuncName(),
			"Destination":       p.DestinationName(),
			"CustomHandlerName": customHandlerName,
		})
		if err != nil {
			return "", err
		}
	}
	return b.String(), nil
}

var findSetParamTemplates = map[schema.HandlerType]string{
	schema.HandlerPathThrough: `
	if !utils.IsEmpty(params.{{.ParamName}}) {
		finder.{{.SetterFuncName}}(params.{{.ParamName}})
	}`,
	schema.HandlerPathThroughEach: `
	if !utils.IsEmpty(params.{{.ParamName}}) {
		for _ , v := range params.{{.ParamName}} {
			finder.{{.SetterFuncName}}(v)
		}
	}`,
	schema.HandlerSort: `
	if !utils.IsEmpty(params.{{.ParamName}}) {
		for _ , v := range params.{{.ParamName}} {
			setSortBy(finder , v)
		}
	}`,
	schema.HandlerFilterBy: `
	if !utils.IsEmpty(params.{{.ParamName}}) {
		finder.SetFilterBy("{{.ParamName}}", params.{{.ParamName}})
	}`,
	schema.HandlerAndParams: `
	if !utils.IsEmpty(params.{{.ParamName}}) {
		for _ , v := range params.{{.ParamName}} {
			finder.SetFilterMultiBy("{{.Destination}}", v)
		}
	}`,
	schema.HandlerOrParams: `
	if !utils.IsEmpty(params.{{.ParamName}}) {
		for _ , v := range params.{{.ParamName}} {
			finder.SetFilterBy("{{.Destination}}", v)
		}
	}`,
	schema.HandlerCustomFunc: `
	if !utils.IsEmpty(params.{{.ParamName}}) {
		{{.CustomHandlerName}}("{{.ParamName}}" , params , finder)
	}`,
}

func generateFilterActions(ctx *tools.GenerateContext, command *tools.Command) (string, error) {

	b := bytes.NewBufferString("")

	for _, p := range command.Params {
		if p.HandlerType != schema.HandlerFilterFunc {
			continue
		}

		t := template.New("c")
		template.Must(t.Parse(findFilterFuncTemplate))

		customHandlerName := fmt.Sprintf(`params.GetCommandDef().Params["%s"].FilterFunc`, p.Name)

		err := t.Execute(b, map[string]interface{}{
			"ParamName":       p.FieldName(),
			"ListResultField": command.FindResultFieldName(),
			"FilterFuncName":  customHandlerName,
		})
		if err != nil {
			return "", err
		}
	}
	return b.String(), nil
}

var findFilterFuncTemplate = `
	if !{{.FilterFuncName}}(list, &res.{{.ListResultField}}[i], params.{{.ParamName}}) {
		continue
	}
`

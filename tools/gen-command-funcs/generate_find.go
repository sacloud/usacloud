package main

import (
	"bytes"
	"fmt"
	"github.com/sacloud/usacloud/schema"
	"text/template"
)

func generateFindCommand(command *schema.Command) (string, error) {
	b := bytes.NewBufferString("")
	t := template.New("c")
	template.Must(t.Parse(findCommandTemplate))

	setParamActions, err := generateFindSetParamActions(command)
	if err != nil {
		return "", err
	}

	err = t.Execute(b, map[string]interface{}{
		"FinderFieldName": ctx.CommandResourceName(),
		"SetParamActions": setParamActions,
		"FuncName":        ctx.CommandFuncName(),
		"ListResultField": ctx.FindResultFieldName(),
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
		list = append(list, &res.{{.ListResultField}}[i])
	}

	return ctx.GetOutput().Print(list...)
`

func generateFindSetParamActions(command *schema.Command) (string, error) {

	b := bytes.NewBufferString("")

	for _, param := range command.BuildedParams() {
		k := param.ParamKey
		p := param.Param

		ctx.P = k

		t := template.New("c")
		template.Must(t.Parse(findSetParamTemplates[p.HandlerType]))

		customHandlerName := fmt.Sprintf(`params.getCommandDef().Params["%s"].CustomHandler`, ctx.P)

		err := t.Execute(b, map[string]interface{}{
			"ParamName":         ctx.InputParamFieldName(),
			"SetterFuncName":    ctx.InputParamSetterFuncName(),
			"Destination":       ctx.InputParamDestinationName(),
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
	if !isEmpty(params.{{.ParamName}}) {
		finder.{{.SetterFuncName}}(params.{{.ParamName}})
	}`,
	schema.HandlerPathThroughEach: `
	if !isEmpty(params.{{.ParamName}}) {
		for _ , v := range params.{{.ParamName}} {
			finder.{{.SetterFuncName}}(v)
		}
	}`,
	schema.HandlerSort: `
	if !isEmpty(params.{{.ParamName}}) {
		for _ , v := range params.{{.ParamName}} {
			setSortBy(finder , v)
		}
	}`,
	schema.HandlerFilterBy: `
	if !isEmpty(params.{{.ParamName}}) {
		finder.SetFilterBy("{{.ParamName}}", params.{{.ParamName}})
	}`,
	schema.HandlerAndParams: `
	if !isEmpty(params.{{.ParamName}}) {
		for _ , v := range params.{{.ParamName}} {
			finder.SetFilterMultiBy("{{.Destination}}", v)
		}
	}`,
	schema.HandlerOrParams: `
	if !isEmpty(params.{{.ParamName}}) {
		for _ , v := range params.{{.ParamName}} {
			finder.SetFilterBy("{{.Destination}}", v)
		}
	}`,
	schema.HandlerCustomFunc: `
	if !isEmpty(params.{{.ParamName}}) {
		{{.CustomHandlerName}}("{{.ParamName}}" , params , finder)
	}`,
}

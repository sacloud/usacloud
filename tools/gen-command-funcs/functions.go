package main

import (
	"bytes"
	"fmt"
	"github.com/sacloud/usacloud/schema"
	"text/template"
)

func generateSetParamActions(command *schema.Command) (string, error) {

	b := bytes.NewBufferString("")

	for k, p := range command.Params {
		ctx.P = k
		if k == "id" {
			continue
		}

		t := template.New("c")
		template.Must(t.Parse(setParamTemplates[p.HandlerType]))

		customFunc := fmt.Sprintf(`params.getCommandDef().Params["%s"].CustomHandler`, ctx.P)

		err := t.Execute(b, map[string]interface{}{
			"FlagName":       ctx.InputParamFlagName(),
			"ParamName":      ctx.InputParamFieldName(),
			"SetterFuncName": ctx.InputParamSetterFuncName(),
			"CustomFunc":     customFunc,
		})
		if err != nil {
			return "", err
		}
	}
	return b.String(), nil
}

var setParamTemplates = map[schema.HandlerType]string{
	schema.HandlerPathThrough: `
	if ctx.IsSet("{{.FlagName}}") {
		p.{{.SetterFuncName}}(params.{{.ParamName}})
	}`,
	schema.HandlerPathThroughEach: `
	if ctx.IsSet("{{.FlagName}}") {
		for _ , v := range params.{{.ParamName}} {
			p.{{.SetterFuncName}}(v)
		}
	}`,
	schema.HandlerCustomFunc: `
	if ctx.IsSet("{{.FlagName}}") {
		{{.CustomFunc}}("{{.ParamName}}" , params , p)
	}`,
}

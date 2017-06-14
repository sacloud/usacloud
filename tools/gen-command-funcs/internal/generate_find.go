package internal

import (
	"bytes"
	"fmt"
	"github.com/sacloud/usacloud/schema"
	"github.com/sacloud/usacloud/tools"
	"text/template"
)

func GenerateFindCommand(ctx *tools.GenerateContext, command *schema.Command) (string, error) {
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
		"FinderFieldName": ctx.CommandResourceName(),
		"SetParamActions": setParamActions,
		"FuncName":        ctx.CommandFuncName(),
		"ListResultField": ctx.FindResultFieldName(),
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

func generateFindSetParamActions(ctx *tools.GenerateContext, command *schema.Command) (string, error) {

	b := bytes.NewBufferString("")

	for _, param := range command.BuildedParams() {
		k := param.ParamKey
		p := param.Param

		ctx.P = k

		if p.HandlerType == schema.HandlerFilterFunc {
			continue
		}

		t := template.New("c")
		template.Must(t.Parse(findSetParamTemplates[p.HandlerType]))

		customHandlerName := fmt.Sprintf(`params.GetCommandDef().Params["%s"].CustomHandler`, ctx.P)

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
	if !command.IsEmpty(params.{{.ParamName}}) {
		finder.{{.SetterFuncName}}(params.{{.ParamName}})
	}`,
	schema.HandlerPathThroughEach: `
	if !command.IsEmpty(params.{{.ParamName}}) {
		for _ , v := range params.{{.ParamName}} {
			finder.{{.SetterFuncName}}(v)
		}
	}`,
	schema.HandlerSort: `
	if !command.IsEmpty(params.{{.ParamName}}) {
		for _ , v := range params.{{.ParamName}} {
			setSortBy(finder , v)
		}
	}`,
	schema.HandlerFilterBy: `
	if !command.IsEmpty(params.{{.ParamName}}) {
		finder.SetFilterBy("{{.ParamName}}", params.{{.ParamName}})
	}`,
	schema.HandlerAndParams: `
	if !command.IsEmpty(params.{{.ParamName}}) {
		for _ , v := range params.{{.ParamName}} {
			finder.SetFilterMultiBy("{{.Destination}}", v)
		}
	}`,
	schema.HandlerOrParams: `
	if !command.IsEmpty(params.{{.ParamName}}) {
		for _ , v := range params.{{.ParamName}} {
			finder.SetFilterBy("{{.Destination}}", v)
		}
	}`,
	schema.HandlerCustomFunc: `
	if !command.IsEmpty(params.{{.ParamName}}) {
		{{.CustomHandlerName}}("{{.ParamName}}" , params , finder)
	}`,
}

func generateFilterActions(ctx *tools.GenerateContext, command *schema.Command) (string, error) {

	b := bytes.NewBufferString("")

	for _, param := range command.BuildedParams() {
		k := param.ParamKey
		p := param.Param

		ctx.P = k

		if p.HandlerType != schema.HandlerFilterFunc {
			continue
		}

		t := template.New("c")
		template.Must(t.Parse(findFilterFuncTemplate))

		customHandlerName := fmt.Sprintf(`params.GetCommandDef().Params["%s"].FilterFunc`, ctx.P)

		err := t.Execute(b, map[string]interface{}{
			"ParamName":       ctx.InputParamFieldName(),
			"ListResultField": ctx.FindResultFieldName(),
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

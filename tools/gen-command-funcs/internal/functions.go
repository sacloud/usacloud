package internal

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/sacloud/usacloud/schema"
	"github.com/sacloud/usacloud/tools"
)

func generateSetParamActions(ctx *tools.GenerateContext, command *schema.Command) (string, error) {

	b := bytes.NewBufferString("")

	for _, param := range command.BuildedParams() {
		p := param.Param
		k := param.ParamKey

		ctx.P = k
		if k == "id" {
			continue
		}

		t := template.New("c")
		template.Must(t.Parse(setParamTemplates[p.HandlerType]))

		customFunc := fmt.Sprintf(`params.GetCommandDef().Params["%s"].CustomHandler`, ctx.P)
		needIsSetCheck := command.Type == schema.CommandUpdate

		err := t.Execute(b, map[string]interface{}{
			"FlagName":       ctx.InputParamFlagName(),
			"ParamName":      ctx.InputParamFieldName(),
			"SetterFuncName": ctx.InputParamSetterFuncName(),
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
	{{if .NeedIsSetCheck}}if ctx.IsSet("{{.FlagName}}") { {{ end }}
		p.{{.SetterFuncName}}(params.{{.ParamName}})
	{{if .NeedIsSetCheck}} } {{ end }}`,
	schema.HandlerPathThroughEach: `
	{{if .NeedIsSetCheck}}if ctx.IsSet("{{.FlagName}}") { {{ end }}
		for _ , v := range params.{{.ParamName}} {
			p.{{.SetterFuncName}}(v)
		}
	{{if .NeedIsSetCheck}} } {{ end }}`,
	schema.HandlerCustomFunc: `
	{{if .NeedIsSetCheck}}if ctx.IsSet("{{.FlagName}}") { {{ end }}
		{{.CustomFunc}}("{{.ParamName}}" , params , p)
	{{if .NeedIsSetCheck}} } {{ end }}`,
	schema.HandlerNoop: "",
}

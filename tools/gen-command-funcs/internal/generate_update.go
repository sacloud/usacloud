package internal

import (
	"bytes"
	"github.com/sacloud/usacloud/schema"
	"github.com/sacloud/usacloud/tools"
	"text/template"
)

func GenerateUpdateCommand(ctx *tools.GenerateContext, command *schema.Command) (string, error) {
	b := bytes.NewBufferString("")
	t := template.New("c")
	template.Must(t.Parse(updateCommandTemplate))

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

var updateCommandTemplate = `
	client := ctx.GetAPIClient()
	api := client.Get{{.ResourceName}}API()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("{{.FuncName}} is failed: %s", e)
	}

	// set params
	{{.SetParamActions}}

	// call Update(id)
	res, err := api.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("{{.FuncName}} is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)
`

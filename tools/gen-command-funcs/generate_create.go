package main

import (
	"bytes"
	"github.com/sacloud/usacloud/schema"
	"text/template"
)

func generateCreateCommand(command *schema.Command) (string, error) {
	b := bytes.NewBufferString("")
	t := template.New("c")
	template.Must(t.Parse(createCommandTemplate))

	setParamActions, err := generateSetParamActions(command)
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

var createCommandTemplate = `
	client := ctx.GetAPIClient()
	api := client.Get{{.ResourceName}}API()
	p := api.New()

	// set params
	{{.SetParamActions}}

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("{{.FuncName}} is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)
`

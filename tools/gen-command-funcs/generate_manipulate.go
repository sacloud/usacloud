package main

import (
	"bytes"
	"github.com/sacloud/usacloud/schema"
	"text/template"
)

func generateManipulateCommand(command *schema.Command) (string, error) {
	b := bytes.NewBufferString("")
	t := template.New("c")
	template.Must(t.Parse(manipulateCommandTemplate))

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

var manipulateCommandTemplate = `
	client := ctx.GetAPIClient()
	api := client.Get{{.ResourceName}}API()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("{{.FuncName}} is failed: %s", e)
	}

	// set params
	{{.SetParamActions}}

	// TODO Remove and implements here!!
	return fmt.Errorf("Not Implements {{.FuncName}}")

	// call manipurate functions
	// res, err := api.XXXX(params.Id, p)
	// if err != nil {
	// 	return fmt.Errorf("{{.FuncName}} is failed: %s", err)
	// }
	// return ctx.GetOutput().Print(res)
`

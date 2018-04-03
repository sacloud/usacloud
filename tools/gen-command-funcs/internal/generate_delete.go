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

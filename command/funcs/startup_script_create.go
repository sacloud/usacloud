package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"io/ioutil"
)

func StartupScriptCreate(ctx command.Context, params *params.CreateStartupScriptParam) error {

	client := ctx.GetAPIClient()
	api := client.GetNoteAPI()
	p := api.New()

	// set params
	p.SetName(params.Name)
	p.SetTags(params.Tags)
	p.SetIconByID(params.IconId)

	if params.Script != "" {
		b, err := ioutil.ReadFile(params.Script)
		if err != nil {
			return fmt.Errorf("StartupScriptCreate is failed: %s", err)
		}
		p.Content = string(b)
	}

	if params.ScriptContent != "" {
		p.Content = params.ScriptContent
	}

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("StartupScriptCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)
}

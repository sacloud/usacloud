package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"io/ioutil"
)

func StartupScriptUpdate(ctx command.Context, params *params.UpdateStartupScriptParam) error {

	client := ctx.GetAPIClient()
	api := client.GetNoteAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("StartupScriptUpdate is failed: %s", e)
	}

	// set params

	if ctx.IsSet("tags") {
		p.SetTags(params.Tags)
	}
	if ctx.IsSet("icon-id") {
		p.SetIconByID(params.IconId)
	}
	if ctx.IsSet("name") {
		p.SetName(params.Name)
	}
	if ctx.IsSet("class") {
		p.SetClassByStr(params.Class)
	}

	if ctx.IsSet("script") {
		b, err := ioutil.ReadFile(params.Script)
		if err != nil {
			return fmt.Errorf("StartupScriptUpdate is failed: %s", err)
		}
		p.Content = string(b)
	}

	if ctx.IsSet("script-content") {
		p.Content = params.ScriptContent
	}

	// call Update(id)
	res, err := api.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("StartupScriptUpdate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}

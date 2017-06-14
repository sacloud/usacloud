package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func SimpleMonitorListCompleteArgs(ctx command.Context, params *params.ListSimpleMonitorParam, cur, prev, commandName string) {

}

func SimpleMonitorCreateCompleteArgs(ctx command.Context, params *params.CreateSimpleMonitorParam, cur, prev, commandName string) {

}

func SimpleMonitorReadCompleteArgs(ctx command.Context, params *params.ReadSimpleMonitorParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSimpleMonitorAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.SimpleMonitors {
		fmt.Println(res.SimpleMonitors[i].ID)
		var target interface{} = &res.SimpleMonitors[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func SimpleMonitorUpdateCompleteArgs(ctx command.Context, params *params.UpdateSimpleMonitorParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSimpleMonitorAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.SimpleMonitors {
		fmt.Println(res.SimpleMonitors[i].ID)
		var target interface{} = &res.SimpleMonitors[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func SimpleMonitorDeleteCompleteArgs(ctx command.Context, params *params.DeleteSimpleMonitorParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSimpleMonitorAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.SimpleMonitors {
		fmt.Println(res.SimpleMonitors[i].ID)
		var target interface{} = &res.SimpleMonitors[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

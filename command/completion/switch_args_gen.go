package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func SwitchListCompleteArgs(ctx command.Context, params *params.ListSwitchParam, cur, prev, commandName string) {

}

func SwitchCreateCompleteArgs(ctx command.Context, params *params.CreateSwitchParam, cur, prev, commandName string) {

}

func SwitchReadCompleteArgs(ctx command.Context, params *params.ReadSwitchParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSwitchAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Switches {
		fmt.Println(res.Switches[i].ID)
		var target interface{} = &res.Switches[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func SwitchUpdateCompleteArgs(ctx command.Context, params *params.UpdateSwitchParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSwitchAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Switches {
		fmt.Println(res.Switches[i].ID)
		var target interface{} = &res.Switches[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func SwitchDeleteCompleteArgs(ctx command.Context, params *params.DeleteSwitchParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSwitchAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Switches {
		fmt.Println(res.Switches[i].ID)
		var target interface{} = &res.Switches[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func SwitchBridgeConnectCompleteArgs(ctx command.Context, params *params.BridgeConnectSwitchParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSwitchAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Switches {
		fmt.Println(res.Switches[i].ID)
		var target interface{} = &res.Switches[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func SwitchBridgeDisconnectCompleteArgs(ctx command.Context, params *params.BridgeDisconnectSwitchParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetSwitchAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Switches {
		fmt.Println(res.Switches[i].ID)
		var target interface{} = &res.Switches[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

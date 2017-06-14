package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func BridgeListCompleteArgs(ctx command.Context, params *params.ListBridgeParam, cur, prev, commandName string) {

}

func BridgeCreateCompleteArgs(ctx command.Context, params *params.CreateBridgeParam, cur, prev, commandName string) {

}

func BridgeReadCompleteArgs(ctx command.Context, params *params.ReadBridgeParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetBridgeAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Bridges {
		fmt.Println(res.Bridges[i].ID)
		var target interface{} = &res.Bridges[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func BridgeUpdateCompleteArgs(ctx command.Context, params *params.UpdateBridgeParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetBridgeAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Bridges {
		fmt.Println(res.Bridges[i].ID)
		var target interface{} = &res.Bridges[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func BridgeDeleteCompleteArgs(ctx command.Context, params *params.DeleteBridgeParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetBridgeAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Bridges {
		fmt.Println(res.Bridges[i].ID)
		var target interface{} = &res.Bridges[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

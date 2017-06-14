package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func IconListCompleteArgs(ctx command.Context, params *params.ListIconParam, cur, prev, commandName string) {

}

func IconCreateCompleteArgs(ctx command.Context, params *params.CreateIconParam, cur, prev, commandName string) {

}

func IconReadCompleteArgs(ctx command.Context, params *params.ReadIconParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetIconAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Icons {
		fmt.Println(res.Icons[i].ID)
		var target interface{} = &res.Icons[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func IconUpdateCompleteArgs(ctx command.Context, params *params.UpdateIconParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetIconAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Icons {
		fmt.Println(res.Icons[i].ID)
		var target interface{} = &res.Icons[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func IconDeleteCompleteArgs(ctx command.Context, params *params.DeleteIconParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetIconAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Icons {
		fmt.Println(res.Icons[i].ID)
		var target interface{} = &res.Icons[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

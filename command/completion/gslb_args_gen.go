package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func GSLBListCompleteArgs(ctx command.Context, params *params.ListGSLBParam, cur, prev, commandName string) {

}

func GSLBServerInfoCompleteArgs(ctx command.Context, params *params.ServerInfoGSLBParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetGSLBAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceGSLBItems {
		fmt.Println(res.CommonServiceGSLBItems[i].ID)
		var target interface{} = &res.CommonServiceGSLBItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func GSLBCreateCompleteArgs(ctx command.Context, params *params.CreateGSLBParam, cur, prev, commandName string) {

}

func GSLBServerAddCompleteArgs(ctx command.Context, params *params.ServerAddGSLBParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetGSLBAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceGSLBItems {
		fmt.Println(res.CommonServiceGSLBItems[i].ID)
		var target interface{} = &res.CommonServiceGSLBItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func GSLBReadCompleteArgs(ctx command.Context, params *params.ReadGSLBParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetGSLBAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceGSLBItems {
		fmt.Println(res.CommonServiceGSLBItems[i].ID)
		var target interface{} = &res.CommonServiceGSLBItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func GSLBServerUpdateCompleteArgs(ctx command.Context, params *params.ServerUpdateGSLBParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetGSLBAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceGSLBItems {
		fmt.Println(res.CommonServiceGSLBItems[i].ID)
		var target interface{} = &res.CommonServiceGSLBItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func GSLBServerDeleteCompleteArgs(ctx command.Context, params *params.ServerDeleteGSLBParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetGSLBAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceGSLBItems {
		fmt.Println(res.CommonServiceGSLBItems[i].ID)
		var target interface{} = &res.CommonServiceGSLBItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func GSLBUpdateCompleteArgs(ctx command.Context, params *params.UpdateGSLBParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetGSLBAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceGSLBItems {
		fmt.Println(res.CommonServiceGSLBItems[i].ID)
		var target interface{} = &res.CommonServiceGSLBItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func GSLBDeleteCompleteArgs(ctx command.Context, params *params.DeleteGSLBParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetGSLBAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceGSLBItems {
		fmt.Println(res.CommonServiceGSLBItems[i].ID)
		var target interface{} = &res.CommonServiceGSLBItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func AutoBackupListCompleteArgs(ctx command.Context, params *params.ListAutoBackupParam, cur, prev, commandName string) {

}

func AutoBackupCreateCompleteArgs(ctx command.Context, params *params.CreateAutoBackupParam, cur, prev, commandName string) {

}

func AutoBackupReadCompleteArgs(ctx command.Context, params *params.ReadAutoBackupParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetAutoBackupAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceAutoBackupItems {
		fmt.Println(res.CommonServiceAutoBackupItems[i].ID)
		var target interface{} = &res.CommonServiceAutoBackupItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func AutoBackupUpdateCompleteArgs(ctx command.Context, params *params.UpdateAutoBackupParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetAutoBackupAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceAutoBackupItems {
		fmt.Println(res.CommonServiceAutoBackupItems[i].ID)
		var target interface{} = &res.CommonServiceAutoBackupItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func AutoBackupDeleteCompleteArgs(ctx command.Context, params *params.DeleteAutoBackupParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetAutoBackupAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.CommonServiceAutoBackupItems {
		fmt.Println(res.CommonServiceAutoBackupItems[i].ID)
		var target interface{} = &res.CommonServiceAutoBackupItems[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

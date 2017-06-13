package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func LicenseListCompleteArgs(ctx command.Context, params *params.ListLicenseParam, cur, prev, commandName string) {

}

func LicenseCreateCompleteArgs(ctx command.Context, params *params.CreateLicenseParam, cur, prev, commandName string) {

}

func LicenseReadCompleteArgs(ctx command.Context, params *params.ReadLicenseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetLicenseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Licenses {
		fmt.Println(res.Licenses[i].ID)
		var target interface{} = &res.Licenses[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func LicenseUpdateCompleteArgs(ctx command.Context, params *params.UpdateLicenseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetLicenseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Licenses {
		fmt.Println(res.Licenses[i].ID)
		var target interface{} = &res.Licenses[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

func LicenseDeleteCompleteArgs(ctx command.Context, params *params.DeleteLicenseParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetLicenseAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Licenses {
		fmt.Println(res.Licenses[i].ID)
		var target interface{} = &res.Licenses[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}

	}

}

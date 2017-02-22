// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-funcs'; DO NOT EDIT

package command

import (
	"fmt"
)

func ProductLicenseList(ctx Context, params *ListProductLicenseParam) error {

	client := ctx.GetAPIClient()
	finder := client.GetProductLicenseAPI()

	finder.SetEmpty()

	if !isEmpty(params.From) {
		finder.SetOffset(params.From)
	}
	if !isEmpty(params.Max) {
		finder.SetLimit(params.Max)
	}
	if !isEmpty(params.Sort) {
		for _, v := range params.Sort {
			setSortBy(finder, v)
		}
	}
	if !isEmpty(params.Name) {
		for _, v := range params.Name {
			finder.SetFilterBy("Name", v)
		}
	}
	if !isEmpty(params.Id) {
		for _, v := range params.Id {
			finder.SetFilterMultiBy("ID", v)
		}
	}

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return fmt.Errorf("ProductLicenseList is failed: %s", err)
	}

	list := []interface{}{}
	for i := range res.LicenseInfo {
		list = append(list, &res.LicenseInfo[i])
	}

	return ctx.GetOutput().Print(list...)

}

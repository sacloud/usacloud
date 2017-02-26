// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-funcs'; DO NOT EDIT

package command

import (
	"fmt"
)

func ProductInternetList(ctx Context, params *ListProductInternetParam) error {

	client := ctx.GetAPIClient()
	finder := client.GetProductInternetAPI()

	finder.SetEmpty()

	if !isEmpty(params.From) {
		finder.SetOffset(params.From)
	}
	if !isEmpty(params.Id) {
		for _, v := range params.Id {
			finder.SetFilterMultiBy("ID", v)
		}
	}
	if !isEmpty(params.Max) {
		finder.SetLimit(params.Max)
	}
	if !isEmpty(params.Name) {
		for _, v := range params.Name {
			finder.SetFilterBy("Name", v)
		}
	}
	if !isEmpty(params.Sort) {
		for _, v := range params.Sort {
			setSortBy(finder, v)
		}
	}

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return fmt.Errorf("ProductInternetList is failed: %s", err)
	}

	list := []interface{}{}
	for i := range res.InternetPlans {
		list = append(list, &res.InternetPlans[i])
	}

	return ctx.GetOutput().Print(list...)

}

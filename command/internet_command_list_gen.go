package command

import (
	"fmt"
)

func InternetList(ctx Context, params *ListInternetParam) error {

	client := ctx.GetAPIClient()
	finder := client.GetInternetAPI()

	finder.SetEmpty()

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
	if !isEmpty(params.From) {
		finder.SetOffset(params.From)
	}

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return fmt.Errorf("InternetList is failed: %s", err)
	}

	list := []interface{}{}
	for i := range res.Internet {
		list = append(list, &res.Internet[i])
	}

	return ctx.GetOutput().Print(list...)

}

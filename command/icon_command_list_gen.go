package command

import (
	"fmt"
)

func IconList(ctx Context, params *ListIconParam) error {

	client := ctx.GetAPIClient()
	finder := client.GetIconAPI()

	finder.SetEmpty()

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
	if !isEmpty(params.Max) {
		finder.SetLimit(params.Max)
	}
	if !isEmpty(params.Sort) {
		for _, v := range params.Sort {
			setSortBy(finder, v)
		}
	}

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return fmt.Errorf("IconList is failed: %s", err)
	}

	list := []interface{}{}
	for i := range res.Icons {
		list = append(list, &res.Icons[i])
	}

	return ctx.GetOutput().Print(list...)

}

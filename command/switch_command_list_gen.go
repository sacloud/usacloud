package command

import (
	"fmt"
)

func SwitchList(ctx Context, params *ListSwitchParam) error {

	client := ctx.GetAPIClient()
	finder := client.GetSwitchAPI()

	finder.SetEmpty()

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
	if !isEmpty(params.Max) {
		finder.SetLimit(params.Max)
	}

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return fmt.Errorf("SwitchList is failed: %s", err)
	}

	list := []interface{}{}
	for i := range res.Switches {
		list = append(list, &res.Switches[i])
	}

	return ctx.GetOutput().Print(list...)

}

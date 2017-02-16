package command

import (
	"fmt"
)

func ISOImageList(ctx Context, params *ListISOImageParam) error {

	client := ctx.GetAPIClient()
	finder := client.GetCDROMAPI()

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
		return fmt.Errorf("ISOImageList is failed: %s", err)
	}

	list := []interface{}{}
	for i := range res.CDROMs {
		list = append(list, &res.CDROMs[i])
	}

	return ctx.GetOutput().Print(list...)

}

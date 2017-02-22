// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-funcs'; DO NOT EDIT

package command

import (
	"fmt"
)

func ArchiveList(ctx Context, params *ListArchiveParam) error {

	client := ctx.GetAPIClient()
	finder := client.GetArchiveAPI()

	finder.SetEmpty()

	if !isEmpty(params.Sort) {
		for _, v := range params.Sort {
			setSortBy(finder, v)
		}
	}
	if !isEmpty(params.Scope) {
		finder.SetFilterBy("Scope", params.Scope)
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
		return fmt.Errorf("ArchiveList is failed: %s", err)
	}

	list := []interface{}{}
	for i := range res.Archives {
		list = append(list, &res.Archives[i])
	}

	return ctx.GetOutput().Print(list...)

}

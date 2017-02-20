package command

import (
	"fmt"
)

func InterfaceList(ctx Context, params *ListInterfaceParam) error {

	client := ctx.GetAPIClient()
	finder := client.GetInterfaceAPI()

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
		return fmt.Errorf("InterfaceList is failed: %s", err)
	}

	list := []interface{}{}
	ignoreTags := []string{"@appliance-database", "@appliance-loadbalancer", "@appliance-vpcrouter"}

Outer:
	for i, nic := range res.Interfaces {
		// customize: ignore appliance interface
		for _, t := range ignoreTags {
			if nic.Server.HasTag(t) {
				continue Outer
			}
		}

		list = append(list, &res.Interfaces[i])
	}

	return ctx.GetOutput().Print(list...)

}

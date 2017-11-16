package funcs

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func PrivateHostList(ctx command.Context, params *params.ListPrivateHostParam) error {

	client := ctx.GetAPIClient()
	finder := client.GetPrivateHostAPI()

	finder.SetEmpty()

	if !command.IsEmpty(params.Name) {
		for _, v := range params.Name {
			finder.SetFilterBy("Name", v)
		}
	}
	if !command.IsEmpty(params.Id) {
		for _, v := range params.Id {
			finder.SetFilterMultiBy("ID", v)
		}
	}
	if !command.IsEmpty(params.From) {
		finder.SetOffset(params.From)
	}
	if !command.IsEmpty(params.Max) {
		finder.SetLimit(params.Max)
	}
	if !command.IsEmpty(params.Sort) {
		for _, v := range params.Sort {
			setSortBy(finder, v)
		}
	}

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return fmt.Errorf("PrivateHostList is failed: %s", err)
	}

	list := []interface{}{}

	for i, h := range res.PrivateHosts {

		if !params.GetCommandDef().Params["tags"].FilterFunc(list, &res.PrivateHosts[i], params.Tags) {
			continue
		}

		list = append(list, &struct {
			*sacloud.PrivateHost
			AssignedCore   int
			TotalCore      int
			AssignedMemory int
			TotalMemory    int
		}{
			PrivateHost:    &res.PrivateHosts[i],
			AssignedCore:   h.AssignedCPU,
			TotalCore:      h.Plan.CPU,
			AssignedMemory: h.GetAssignedMemoryGB(),
			TotalMemory:    h.Plan.GetMemoryGB(),
		})
	}
	return ctx.GetOutput().Print(list...)

}

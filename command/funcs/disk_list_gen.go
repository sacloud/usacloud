// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-funcs'; DO NOT EDIT

package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func DiskList(ctx command.Context, params *params.ListDiskParam) error {

	client := ctx.GetAPIClient()
	finder := client.GetDiskAPI()

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
	if !command.IsEmpty(params.Scope) {
		finder.SetFilterBy("Scope", params.Scope)
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
		return fmt.Errorf("DiskList is failed: %s", err)
	}

	list := []interface{}{}
	for i := range res.Disks {

		if !params.GetCommandDef().Params["tags"].FilterFunc(list, &res.Disks[i], params.Tags) {
			continue
		}

		if !params.GetCommandDef().Params["source-archive-id"].FilterFunc(list, &res.Disks[i], params.SourceArchiveId) {
			continue
		}

		if !params.GetCommandDef().Params["source-disk-id"].FilterFunc(list, &res.Disks[i], params.SourceDiskId) {
			continue
		}

		if !params.GetCommandDef().Params["storage"].FilterFunc(list, &res.Disks[i], params.Storage) {
			continue
		}

		list = append(list, &res.Disks[i])
	}
	return ctx.GetOutput().Print(list...)

}

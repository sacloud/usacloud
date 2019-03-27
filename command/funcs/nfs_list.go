package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func NFSList(ctx command.Context, params *params.ListNFSParam) error {

	client := ctx.GetAPIClient()
	finder := client.GetNFSAPI()

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
		return fmt.Errorf("NFSList is failed: %s", err)
	}

	// get plans
	plans, err := client.NFS.GetNFSPlans()
	if err != nil {
		return fmt.Errorf("NFSList is failed: %s", err)
	}

	list := []interface{}{}
	type nfsValues struct {
		*sacloud.NFS
		PlanName string
		Size     int
	}
	for i := range res.NFS {

		if !params.GetCommandDef().Params["tags"].FilterFunc(list, &res.NFS[i], params.Tags) {
			continue
		}

		var planName string
		var size int

		plan, planDetail := plans.FindByPlanID(res.NFS[i].GetPlanID())
		if planDetail != nil {
			planName = plan.String()
			size = planDetail.Size
		}

		list = append(list, &nfsValues{
			NFS:      &res.NFS[i],
			PlanName: planName,
			Size:     size,
		})
	}
	return ctx.GetOutput().Print(list...)

}

package command

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
)

func DiskCreate(ctx Context, params *CreateDiskParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDiskAPI()
	p := api.New()

	// set params

	p.SetDescription(params.Description)
	p.SetIconByID(params.IconId)
	p.SetDiskPlan(params.Plan)
	p.SetSizeGB(params.Size)
	p.SetDistantFrom(params.DistantFrom)
	p.SetName(params.Name)
	p.SetTags(params.Tags)
	p.SetDiskConnection(sacloud.EDiskConnection(params.Connection))
	p.SetSourceArchive(params.SourceArchiveId)
	p.SetSourceDisk(params.SourceDiskId)

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("DiskCreate is failed: %s", err)
	}

	if !params.Async {
		err = api.SleepWhileCopying(res.ID, client.DefaultTimeoutDuration)
		if err != nil {
			return fmt.Errorf("DiskCreate is failed: %s", err)
		}
		res, err = api.Read(res.ID)
		if err != nil {
			return fmt.Errorf("DiskCreate is failed: %s", err)
		}
	}

	return ctx.GetOutput().Print(res)

}

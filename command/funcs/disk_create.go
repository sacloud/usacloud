package funcs

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
)

func DiskCreate(ctx command.Context, params *params.CreateDiskParam) error {

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

	// wait for copy with progress
	var res *sacloud.Disk
	var err error
	err = internal.ExecWithProgress(
		"Still creating...",
		"Create disk",
		command.GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {
			// call Create(id)
			res, err = api.Create(p)
			if err != nil {
				errChan <- err
				return
			}
			err = api.SleepWhileCopying(res.ID, client.DefaultTimeoutDuration)
			if err != nil {
				errChan <- err
				return
			}
			res, err = api.Read(res.ID)
			if err != nil {
				errChan <- err
				return
			}
			compChan <- true
		},
	)

	if err != nil {
		return fmt.Errorf("DiskCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}

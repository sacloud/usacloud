package command

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command/internal"
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

	// wait for copy with progress
	var res *sacloud.Disk
	var err error
	compChan := make(chan bool)
	errChan := make(chan error)
	spinner := internal.NewProgress(
		"Still creating...",
		"Create disk",
		GlobalOption.Progress)
	spinner.Start()

	// call Create(id)
	go func() {
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
	}()

copy:
	for {
		select {
		case <-compChan:
			spinner.Stop()
			break copy
		case err := <-errChan:
			return fmt.Errorf("DiskCreate is failed: %s", err)
		}
	}

	return ctx.GetOutput().Print(res)

}

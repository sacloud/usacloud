package command

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
)

func AutoBackupCreate(ctx Context, params *CreateAutoBackupParam) error {

	client := ctx.GetAPIClient()
	api := client.GetAutoBackupAPI()
	p := api.New(params.Name, params.DiskId)

	// set params
	p.SetName(params.Name)
	p.SetTags(params.Tags)
	p.SetDescription(params.Description)
	p.SetIconByID(params.IconId)
	p.SetBackupMaximumNumberOfArchives(params.Generation)
	p.SetBackupHour(params.StartHour)

	exists := false
	for _, v := range params.Weekdays {
		if v == "all" {
			exists = true
			break
		}
	}
	if exists {
		p.SetBackupSpanWeekdays(sacloud.AllowAutoBackupWeekdays())
	} else {
		p.SetBackupSpanWeekdays(params.Weekdays)
	}

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("AutoBackupCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}

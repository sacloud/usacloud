package command

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
)

func AutoBackupUpdate(ctx Context, params *UpdateAutoBackupParam) error {

	client := ctx.GetAPIClient()
	api := client.GetAutoBackupAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("AutoBackupUpdate is failed: %s", e)
	}

	// set params

	if ctx.IsSet("name") {
		p.SetName(params.Name)
	}
	if ctx.IsSet("tags") {
		p.SetTags(params.Tags)
	}
	if ctx.IsSet("start-hour") {
		p.SetBackupHour(params.StartHour)
	}
	if ctx.IsSet("description") {
		p.SetDescription(params.Description)
	}
	if ctx.IsSet("icon-id") {
		p.SetIconByID(params.IconId)
	}
	if ctx.IsSet("generation") {
		p.SetBackupMaximumNumberOfArchives(params.Generation)
	}

	if ctx.IsSet("weekdays") {
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
	}

	// call Update(id)
	res, err := api.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("AutoBackupUpdate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}

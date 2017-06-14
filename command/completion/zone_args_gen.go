package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ZoneListCompleteArgs(ctx command.Context, params *params.ListZoneParam, cur, prev, commandName string) {

}

func ZoneReadCompleteArgs(ctx command.Context, params *params.ReadZoneParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	if cur != "" && !isSakuraID(cur) {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetZoneAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.Zones {
		fmt.Println(res.Zones[i].ID)

	}

}

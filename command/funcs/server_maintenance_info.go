package funcs

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ServerMaintenanceInfo(ctx command.Context, params *params.MaintenanceInfoServerParam) error {

	client := ctx.GetAPIClient()
	finder := client.GetServerAPI()

	type mainteInfo struct {
		*sacloud.Server
		MaintenanceInfo *sacloud.NewsFeed
		StartDate       string
		EndDate         string
		InfoURL         string
	}

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return fmt.Errorf("ServerMaintenanceInfo is failed: %s", err)
	}

	list := []interface{}{}
	timeLayout := "01/02 15:04"
	for i, s := range res.Servers {
		if s.MaintenanceScheduled() {

			info, err := client.NewsFeed.GetFeedByURL(s.GetMaintenanceInfoURL())
			if err != nil {
				return fmt.Errorf("GetFeedByURL(%s) is failed: %s", s.GetMaintenanceInfoURL(), err)
			}
			if info == nil {
				continue
			}

			v := &mainteInfo{
				Server:          &res.Servers[i],
				MaintenanceInfo: info,
				StartDate:       info.EventStart().Format(timeLayout),
				EndDate:         info.EventEnd().Format(timeLayout),
				InfoURL:         s.GetMaintenanceInfoURL(),
			}

			list = append(list, v)
		}
	}
	return ctx.GetOutput().Print(list...)

}

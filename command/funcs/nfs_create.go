package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
)

func NFSCreate(ctx command.Context, params *params.CreateNFSParam) error {

	client := ctx.GetAPIClient()
	api := client.GetNFSAPI()

	p := &sacloud.CreateNFSValue{
		SwitchID:     fmt.Sprintf("%d", params.SwitchId),
		IPAddress:    params.Ipaddress,
		MaskLen:      params.NwMaskLen,
		DefaultRoute: params.DefaultRoute,
		Name:         params.Name,
		Description:  params.Description,
		Tags:         params.Tags,
		Icon:         sacloud.NewResource(params.IconId),
	}

	var plan sacloud.NFSPlan
	switch params.Plan {
	case "ssd":
		plan = sacloud.NFSPlanSSD
	case "hdd":
		plan = sacloud.NFSPlanHDD
	default:
		return fmt.Errorf("NFSCreate is failed: invalid plan %s", params.Plan)
	}

	// call Create(id)
	res, err := api.CreateWithPlan(p, plan, sacloud.NFSSize(params.Size))
	if err != nil {
		return fmt.Errorf("NFSCreate is failed: %s", err)
	}

	// wait for boot
	err = internal.ExecWithProgress(
		fmt.Sprintf("Still creating[ID:%d]...", res.ID),
		fmt.Sprintf("Create nfs[ID:%d]", res.ID),
		command.GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {
			// call manipurate functions
			err := api.SleepWhileCopying(res.ID, client.DefaultTimeoutDuration, 20)
			if err != nil {
				errChan <- err
				return
			}
			err = api.SleepUntilUp(res.ID, client.DefaultTimeoutDuration)
			if err != nil {
				errChan <- err
				return
			}
			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("NFSCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}

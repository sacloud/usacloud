package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ProductInternetListCompleteArgs(ctx command.Context, params *params.ListProductInternetParam, cur, prev, commandName string) {

}

func ProductInternetReadCompleteArgs(ctx command.Context, params *params.ReadProductInternetParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	if cur != "" && !isSakuraID(cur) {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetProductInternetAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.InternetPlans {
		fmt.Println(res.InternetPlans[i].ID)

	}

}

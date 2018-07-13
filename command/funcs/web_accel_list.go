package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func WebAccelList(ctx command.Context, params *params.ListWebAccelParam) error {

	client := ctx.GetAPIClient()
	finder := client.GetWebAccelAPI()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return fmt.Errorf("WebAccelList is failed: %s", err)
	}

	list := []interface{}{}
	for i := range res.WebAccelSites {

		list = append(list, &res.WebAccelSites[i])
	}
	return ctx.GetOutput().Print(list...)

}

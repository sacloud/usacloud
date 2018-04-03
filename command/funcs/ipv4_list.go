package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func Ipv4List(ctx command.Context, params *params.ListIpv4Param) error {

	client := ctx.GetAPIClient()
	finder := client.GetIPAddressAPI()

	finder.SetEmpty()

	if !command.IsEmpty(params.From) {
		finder.SetOffset(params.From)
	}
	if !command.IsEmpty(params.Max) {
		finder.SetLimit(params.Max)
	}
	if !command.IsEmpty(params.Sort) {
		for _, v := range params.Sort {
			setSortBy(finder, v)
		}
	}

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return fmt.Errorf("IPv4List is failed: %s", err)
	}

	list := []interface{}{}
	for i := range res.IPAddress {

		list = append(list, &res.IPAddress[i])
	}
	return ctx.GetOutput().Print(list...)

}

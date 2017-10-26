package funcs

import (
	"fmt"
	"github.com/sacloud/libsacloud/api"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func Ipv6List(ctx command.Context, params *params.ListIpv6Param) error {

	client := ctx.GetAPIClient()
	finder := client.GetIPv6AddrAPI()

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
		return fmt.Errorf("Ipv6List is failed: %s", err)
	}

	list := []interface{}{}
	ipv6NetCache := map[int64]*sacloud.IPv6Net{}
	for i := range res.IPv6Addrs {

		n, err := findIPv6NetIfAbsent(client, res.IPv6Addrs[i].IPv6Net.ID, ipv6NetCache)
		if err != nil {
			return fmt.Errorf("Ipv6List is failed: %s", err)
		}
		res.IPv6Addrs[i].IPv6Net = n

		// filter by internet(switch+router) id
		if !params.GetCommandDef().Params["internet-id"].FilterFunc(list, &res.IPv6Addrs[i], params.InternetId) {
			continue
		}
		// filter by ipv6net id
		if !params.GetCommandDef().Params["ipv6net-id"].FilterFunc(list, &res.IPv6Addrs[i], params.Ipv6netId) {
			continue
		}

		list = append(list, &res.IPv6Addrs[i])
	}
	return ctx.GetOutput().Print(list...)

}

func findIPv6NetIfAbsent(client *api.Client, id int64, cache map[int64]*sacloud.IPv6Net) (*sacloud.IPv6Net, error) {
	if n, ok := cache[id]; ok {
		return n, nil
	}
	ipv6net, err := client.GetIPv6NetAPI().Read(id)
	if err != nil {
		return nil, err
	}
	return ipv6net, nil
}

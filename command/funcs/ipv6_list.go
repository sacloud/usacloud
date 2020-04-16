// Copyright 2017-2020 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/api"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/pkg/utils"
)

func IPv6List(ctx command.Context, params *params.ListIPv6Param) error {

	client := ctx.GetAPIClient()
	finder := client.GetIPv6AddrAPI()

	finder.SetEmpty()

	if !utils.IsEmpty(params.From) {
		finder.SetOffset(params.From)
	}
	if !utils.IsEmpty(params.Max) {
		finder.SetLimit(params.Max)
	}
	if !utils.IsEmpty(params.Sort) {
		for _, v := range params.Sort {
			setSortBy(finder, v)
		}
	}

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return fmt.Errorf("IPv6List is failed: %s", err)
	}

	list := []interface{}{}
	ipv6NetCache := map[sacloud.ID]*sacloud.IPv6Net{}
	for i := range res.IPv6Addrs {

		n, err := findIPv6NetIfAbsent(client, res.IPv6Addrs[i].IPv6Net.ID, ipv6NetCache)
		if err != nil {
			return fmt.Errorf("IPv6List is failed: %s", err)
		}
		res.IPv6Addrs[i].IPv6Net = n

		// filter by internet(switch+router) id
		if !params.GetCommandDef().Params["internet-id"].FilterFunc(list, &res.IPv6Addrs[i], params.InternetId) {
			continue
		}
		// filter by ipv6net id
		if !params.GetCommandDef().Params["ipv6net-id"].FilterFunc(list, &res.IPv6Addrs[i], params.IPv6netId) {
			continue
		}

		list = append(list, &res.IPv6Addrs[i])
	}
	return ctx.GetOutput().Print(list...)

}

func findIPv6NetIfAbsent(client *api.Client, id sacloud.ID, cache map[sacloud.ID]*sacloud.IPv6Net) (*sacloud.IPv6Net, error) {
	if n, ok := cache[id]; ok {
		return n, nil
	}
	ipv6net, err := client.GetIPv6NetAPI().Read(id)
	if err != nil {
		return nil, err
	}
	return ipv6net, nil
}

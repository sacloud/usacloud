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

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func InternetIpv6Enable(ctx command.Context, params *params.Ipv6EnableInternetParam) error {

	client := ctx.GetAPIClient()
	api := client.GetInternetAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("InternetIpv6Enable is failed: %s", e)
	}

	// has switch?
	sw := p.GetSwitch()
	if sw == nil {
		return fmt.Errorf("InternetIpv6Enable is failed: %s", "Invalid state: missing Switch resource")
	}

	if len(sw.IPv6Nets) > 0 {
		fmt.Fprintln(command.GlobalOption.Err, "IPv6 is already enabled on this resource")
		return nil
	}

	ipv6net, err := api.EnableIPv6(params.Id)
	if err != nil {
		return fmt.Errorf("InternetIpv6Enable is failed: %s", err)
	}

	return ctx.GetOutput().Print(&ipv6net)

}

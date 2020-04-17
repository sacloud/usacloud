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

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func ProxyLBResponseHeaderInfo(ctx cli.Context, params *params.ResponseHeaderInfoProxyLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ProxyLBResponseHeaderInfo is failed: %s", e)
	}
	if len(p.Settings.ProxyLB.BindPorts) == 0 {
		return fmt.Errorf("ProxyLB don't have any bind-ports")
	}

	// validate index
	if params.PortIndex <= 0 || params.PortIndex-1 >= len(p.Settings.ProxyLB.BindPorts) {
		return fmt.Errorf("port-index(%d) is out of range", params.PortIndex)
	}

	bindPort := p.Settings.ProxyLB.BindPorts[params.PortIndex-1]

	if len(bindPort.AddResponseHeader) == 0 {
		return fmt.Errorf("Port %s:%d don't have any additional response headers", bindPort.ProxyMode, bindPort.Port)
	}

	var list []interface{}
	for i := range bindPort.AddResponseHeader {
		list = append(list, &bindPort.AddResponseHeader[i])
	}
	return ctx.GetOutput().Print(list...)

}

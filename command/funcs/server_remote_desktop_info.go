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

	"github.com/sacloud/libsacloud/utils/server"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ServerRemoteDesktopInfo(ctx command.Context, params *params.RemoteDesktopInfoServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerRdpInfo is failed: %s", e)
	}

	if !p.IsUp() {
		return fmt.Errorf("ServerRdpInfo is failed: %s", "server is not running")
	}

	// collect server IPAddress
	ip := p.Interfaces[0].IPAddress
	if ip == "" {
		ip = p.Interfaces[0].UserIPAddress
	}
	if ip == "" {
		return fmt.Errorf("ServerRdpInfo is failed: collecting IPAddress from server is failed: %#v", p)
	}

	rdpClient := &server.RDPOpener{
		User:      params.User,
		Port:      params.Port,
		IPAddress: ip,
	}

	fmt.Fprint(command.GlobalOption.Out, rdpClient.RDPFileContent())
	return nil

}

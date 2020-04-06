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

func IPv4PtrAdd(ctx command.Context, params *params.PtrAddIPv4Param) error {

	client := ctx.GetAPIClient()
	api := client.GetIPAddressAPI()

	targetIP, err := getIPv4AddrFromArgs(ctx.Args())
	if err != nil {
		return err
	}

	ip, err := api.Read(targetIP)
	if err != nil {
		return fmt.Errorf("IPv4PtrAdd is failed: %s", err)
	}

	if ip.HostName != "" {
		return fmt.Errorf("PTR record has already been set for IPAddress %q", targetIP)
	}

	ip, err = api.Update(targetIP, params.Hostname)
	if err != nil {
		return fmt.Errorf("IPv4PtrAdd is failed: %s", err)
	}

	return ctx.GetOutput().Print(ip)
}

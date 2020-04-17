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

func InternetIPv6Disable(ctx cli.Context, params *params.IPv6DisableInternetParam) error {

	client := ctx.GetAPIClient()
	api := client.GetInternetAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("InternetIPv6Disable is failed: %s", e)
	}

	// has switch?
	sw := p.GetSwitch()
	if sw == nil {
		return fmt.Errorf("InternetIPv6Disable is failed: %s", "Invalid state: missing Switch resource")
	}

	if len(sw.IPv6Nets) == 0 {
		fmt.Fprintln(ctx.IO().Err(), "IPv6 is already disabled on this resource")
		return nil
	}

	_, err := api.DisableIPv6(params.Id, sw.IPv6Nets[0].ID)
	if err != nil {
		return fmt.Errorf("InternetIPv6Disable is failed: %s", err)
	}

	return nil

}

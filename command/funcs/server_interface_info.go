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

func ServerInterfaceInfo(ctx command.Context, params *params.InterfaceInfoServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerInterfaceInfo is failed: %s", e)
	}

	interfaces := p.GetInterfaces()
	if len(interfaces) == 0 {
		fmt.Fprintf(ctx.IO().Err(), "Server don't have any interfaces\n")
		return nil
	}

	// build parameters to display table
	list := []interface{}{}
	for i := range interfaces {
		list = append(list, &interfaces[i])
	}

	return ctx.GetOutput().Print(list...)

}

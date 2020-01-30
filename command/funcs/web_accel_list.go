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

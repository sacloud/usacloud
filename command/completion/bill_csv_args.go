// Copyright 2017-2019 The Usacloud Authors
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

package completion

import (
	"fmt"
	"strings"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func BillCsvCompleteArgs(ctx command.Context, params *params.CsvBillParam, cur, prev, commandName string) {

	if !command.GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetBillAPI()

	// auth-status
	auth, err := client.GetAuthStatusAPI().Read()
	if err != nil {
		return
	}
	accountID := auth.Account.ID
	if accountID == 0 {
		return
	}

	if !strings.Contains(auth.ExternalPermission, "bill") {
		return
	}

	// call Find()
	res, err := finder.ByContract(accountID)
	if err != nil {
		return
	}

	for i := range res.Bills {
		fmt.Println(res.Bills[i].BillID)
	}

}

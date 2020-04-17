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
	"os"
	"strings"

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func BillCsv(ctx cli.Context, params *params.CsvBillParam) error {

	client := ctx.GetAPIClient()
	finder := client.GetBillAPI()

	// auth-status
	auth, err := client.GetAuthStatusAPI().Read()
	if err != nil {
		return fmt.Errorf("BillCsv is failed: %s", err)
	}
	accountID := auth.Account.ID
	if accountID == 0 {
		return fmt.Errorf("BillCsv is failed: %s", "invalid account id")
	}
	memberCD := auth.Member.Code

	if !strings.Contains(auth.ExternalPermission, "bill") {
		return fmt.Errorf("Don't have permission to view bills")
	}

	// validate param
	billID := params.BillId
	if !ctx.IsSet("bill-id") {
		bills, err := finder.ByContract(accountID)
		if err != nil {
			return fmt.Errorf("BillCsv is failed: %s", err)
		}
		if len(bills.Bills) == 0 {
			return fmt.Errorf("BillCsv is failed: Empty result")
		}
		billID = bills.Bills[0].BillID
	}

	// call Find()
	res, err := finder.GetDetailCSV(memberCD, billID)
	if err != nil {
		return fmt.Errorf("BillCsv is failed: %s", err)
	}

	var out = ctx.IO().Out()
	if params.BillOutput != "" {
		file, err := os.Create(params.BillOutput)
		if err != nil {
			return fmt.Errorf("BillCsv is failed: %s", err)
		}
		out = file
		defer file.Close()
	}

	// write to out

	if !params.NoHeader {
		_, err := fmt.Fprintf(out, "\"%s\"\n", strings.Join(res.HeaderRow, "\",\""))
		if err != nil {
			return fmt.Errorf("BillCsv is failed: %s", err)
		}
	}

	for _, line := range res.BodyRows {
		_, err := fmt.Fprintf(out, "\"%s\"\n", strings.Join(line, "\",\""))
		if err != nil {
			return fmt.Errorf("BillCsv is failed: %s", err)
		}
	}

	return nil
}

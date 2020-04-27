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

package bill

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/sacloud/libsacloud/v2/sacloud"

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func Csv(ctx cli.Context, params *params.CsvBillParam) error {
	client := sacloud.NewBillOp(ctx.Client())
	authClient := sacloud.NewAuthStatusOp(ctx.Client())

	// auth-status
	auth, err := authClient.Read(ctx)
	if err != nil {
		return err
	}
	accountID := auth.AccountID
	if accountID.IsEmpty() {
		return errors.New("invalid account id")
	}
	memberCD := auth.MemberCode

	if !auth.ExternalPermission.PermittedBill() {
		return errors.New("you don't have permission to view bills")
	}

	// validate param
	billID := params.BillId
	if !params.Changed("bill-id") {
		bills, err := client.ByContract(ctx, accountID)
		if err != nil {
			return err
		}
		if len(bills.Bills) == 0 {
			return errors.New("no results")
		}
		billID = bills.Bills[0].ID
	}

	// call Find()
	res, err := client.DetailsCSV(ctx, memberCD, billID)
	if err != nil {
		return err
	}

	var out = ctx.IO().Out()
	if params.BillOutput != "" {
		file, err := os.Create(params.BillOutput)
		if err != nil {
			return err
		}
		out = file
		defer file.Close()
	}

	// write to out

	if !params.NoHeader {
		_, err := fmt.Fprintf(out, "\"%s\"\n", strings.Join(res.HeaderRow, "\",\""))
		if err != nil {
			return err
		}
	}

	for _, line := range res.BodyRows {
		_, err := fmt.Fprintf(out, "\"%s\"\n", strings.Join(line, "\",\""))
		if err != nil {
			return err
		}
	}

	return nil
}

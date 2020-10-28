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

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func List(ctx cli.Context, params *params.ListBillParam) error {
	client := sacloud.NewBillOp(ctx.Client())
	authClient := sacloud.NewAuthStatusOp(ctx.Client())

	// validate
	if params.Month > 0 && params.Year == 0 {
		return errors.New("year is required when month is set")
	}

	// auth-status
	auth, err := authClient.Read(ctx)
	if err != nil {
		return err
	}
	if auth.AccountID.IsEmpty() {
		return errors.New("invalid account id")
	}

	if !auth.ExternalPermission.PermittedBill() {
		return errors.New("you don't have permission to view bills")
	}

	var bills []*sacloud.Bill
	switch {
	case params.Month > 0:
		res, err := client.ByContractYearMonth(ctx, auth.AccountID, params.Year, params.Month)
		if err != nil {
			return err
		}
		bills = res.Bills
	case params.Year > 0:
		res, err := client.ByContractYear(ctx, auth.AccountID, params.Year)
		if err != nil {
			return err
		}
		bills = res.Bills
	default:
		res, err := client.ByContract(ctx, auth.AccountID)
		if err != nil {
			return err
		}
		bills = res.Bills
	}

	return ctx.Output().Print(bills)
}

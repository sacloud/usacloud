// Copyright 2016-2021 The Libsacloud Authors
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
	"context"
	"errors"

	"github.com/sacloud/libsacloud/v2/sacloud"
)

func (s *Service) Csv(req *CsvRequest) (*sacloud.BillDetailCSV, error) {
	return s.CsvWithContext(context.Background(), req)
}

func (s *Service) CsvWithContext(ctx context.Context, req *CsvRequest) (*sacloud.BillDetailCSV, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	billOp := sacloud.NewBillOp(s.caller)
	authOp := sacloud.NewAuthStatusOp(s.caller)

	// check auth status
	auth, err := authOp.Read(ctx)
	if err != nil {
		return nil, err
	}
	if auth.AccountID.IsEmpty() {
		return nil, errors.New("invalid account id")
	}
	if !auth.ExternalPermission.PermittedBill() {
		return nil, errors.New("you don't have a permission")
	}

	// get latest bill ID if empty
	billID := req.ID
	if billID.IsEmpty() {
		bills, err := billOp.ByContract(ctx, auth.AccountID)
		if err != nil {
			return nil, err
		}
		if len(bills.Bills) == 0 {
			return nil, sacloud.NewNoResultsError()
		}
		billID = bills.Bills[0].ID
	}

	return billOp.DetailsCSV(ctx, auth.MemberCode, billID)
}

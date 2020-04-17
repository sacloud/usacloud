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
	"time"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func CouponList(ctx cli.Context, params *params.ListCouponParam) error {
	client := ctx.GetAPIClient()
	finder := client.GetCouponAPI()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return fmt.Errorf("CouponList is failed: %s", err)
	}

	var list []interface{}
	for _, coupon := range res {
		if params.Usable && !isUsableCoupon(coupon) {
			continue
		}
		list = append(list, coupon)
	}
	return ctx.GetOutput().Print(list...)
}

func isUsableCoupon(coupon *sacloud.Coupon) bool {
	return coupon.Discount > 0 && coupon.AppliedAt.Before(time.Now()) && coupon.UntilAt.After(time.Now())
}

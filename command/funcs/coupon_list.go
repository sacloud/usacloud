package funcs

import (
	"fmt"
	"time"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func CouponList(ctx command.Context, params *params.ListCouponParam) error {
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

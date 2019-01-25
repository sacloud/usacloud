package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func CouponResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandCustom,
			Aliases:            []string{"ls", "find"},
			Params:             couponListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: couponListColumns(),
			UseCustomCommand:   true,
			NeedlessConfirm:    true,
		},
	}

	return &schema.Resource{
		Commands:         commands,
		DefaultCommand:   "list",
		ResourceCategory: CategoryCoupon,
	}
}

func couponListParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"usable": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Description: "show usable coupons only",
			Category:    "output",
			Order:       10,
		},
	}
}

func couponListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "CouponID"},
		{Name: "Discount"},
		{Name: "AppliedAt"},
		{Name: "UntilAt"},
	}
}

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

package define

import (
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/schema"
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
		IsGlobal:         true,
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

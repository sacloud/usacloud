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

package _define

import (
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/schema"
)

func PriceResource() *schema.Resource {
	commands := map[string]*schema.Command{
		"list": {
			Type:    schema.CommandList,
			Aliases: []string{"ls", "find"},
			Params:  priceListParam(),
			// TableType:          output.TableSimple,
			TableColumnDefines: priceListColumns(),
			Category:           "basics",
			Order:              10,
		},
	}

	return &schema.Resource{
		Commands:            commands,
		DefaultCommand:      "list",
		Aliases:             []string{"public-price"},
		AltResource:         "PublicPrice",
		ListResultFieldName: "ServiceClasses",
		ResourceCategory:    CategoryInformation,
		IsGlobal:            true,
	}
}

func priceListParam() map[string]*schema.Parameter {
	return CommonListParam
}

func priceListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "DisplayName"},
		{
			Name:    "Zone",
			Sources: []string{"Price.Zone"},
		},
		{
			Name:    "Price:Hourly",
			Sources: []string{"Price.Hourly"},
		},
		{
			Name:    "Price:Daily",
			Sources: []string{"Price.Daily"},
		},
		{
			Name:    "Price:Monthly",
			Sources: []string{"Price.Monthly"},
		},
	}
}

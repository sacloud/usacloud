package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func PriceResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"l", "ls", "find"},
			Params:             priceListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: priceListColumns(),
		},
	}

	return &schema.Resource{
		Commands:            commands,
		DefaultCommand:      "list",
		Aliases:             []string{"public-price"},
		AltResource:         "PublicPrice",
		ListResultFieldName: "ServiceClasses",
		ResourceCategory:    CategoryInformation,
	}
}

func priceListParam() map[string]*schema.Schema {
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
			Name:    "Price-Hourly",
			Sources: []string{"Price.Hourly"},
		},
		{
			Name:    "Price-Daily",
			Sources: []string{"Price.Daily"},
		},
		{
			Name:    "Price-Monthly",
			Sources: []string{"Price.Monthly"},
		},
	}
}

package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func ProductLicenseResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:                schema.CommandList,
			ListResultFieldName: "LicenseInfo",
			Aliases:             []string{"l", "ls", "find"},
			Params:              productLicenseListParam(),
			TableType:           output.TableSimple,
			TableColumnDefines:  productLicenseListColumns(),
		},
		"read": {
			Type:          schema.CommandRead,
			Aliases:       []string{"r"},
			Params:        productLicenseReadParam(),
			IncludeFields: productLicenseDetailIncludes(),
			ExcludeFields: productLicenseDetailExcludes(),
		},
	}

	return &schema.Resource{
		Commands:         commands,
		Aliases:          []string{"license-info"},
		ResourceCategory: CategoryInformation,
	}
}

func productLicenseListParam() map[string]*schema.Schema {
	return CommonListParam
}

func productLicenseListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{Name: "TermsOfUse"},
	}
}

func productLicenseDetailIncludes() []string {
	return []string{}
}

func productLicenseDetailExcludes() []string {
	return []string{}
}

func productLicenseReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": getParamResourceShortID("resource ID", 5),
	}
}

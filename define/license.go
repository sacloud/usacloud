package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func LicenseResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find"},
			Params:             licenseListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: licenseListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"create": {
			Type:          schema.CommandCreate,
			Params:        licenseCreateParam(),
			IncludeFields: licenseDetailIncludes(),
			ExcludeFields: licenseDetailExcludes(),
			Category:      "basics",
			Order:         20,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        licenseReadParam(),
			IncludeFields: licenseDetailIncludes(),
			ExcludeFields: licenseDetailExcludes(),
			NoSelector:    true,
			Category:      "basics",
			Order:         30,
		},
		"update": {
			Type:          schema.CommandUpdate,
			Params:        licenseUpdateParam(),
			IncludeFields: licenseDetailIncludes(),
			ExcludeFields: licenseDetailExcludes(),
			NoSelector:    true,
			Category:      "basics",
			Order:         40,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"rm"},
			Params:        licenseDeleteParam(),
			IncludeFields: licenseDetailIncludes(),
			ExcludeFields: licenseDetailExcludes(),
			NoSelector:    true,
			Category:      "basics",
			Order:         50,
		},
	}

	return &schema.Resource{
		Commands:            commands,
		ResourceCategory:    CategoryCommonItem,
		ListResultFieldName: "Licenses",
	}
}

func licenseListParam() map[string]*schema.Schema {
	return CommonListParam
}

func licenseListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{
			Name:    "LicenseInfo:ID",
			Sources: []string{"LicenseInfo.ID"},
		},
		{
			Name:    "LicenseInfo:Name",
			Sources: []string{"LicenseInfo.Name"},
		},
	}
}

func licenseDetailIncludes() []string {
	return []string{}
}

func licenseDetailExcludes() []string {
	return []string{}
}

func licenseCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": paramRequiredName,
		"license-info-id": {
			Type:            schema.TypeInt64,
			HandlerType:     schema.HandlerPathThrough,
			DestinationProp: "SetLicenseInfoByID",
			Description:     "set LicenseInfo ID",
			CompleteFunc:    completeLicenseInfoID(),
			Category:        "license",
			Order:           10,
		},
	}
}

func licenseReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func licenseUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": paramName,
	}
}

func licenseDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

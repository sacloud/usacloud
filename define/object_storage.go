package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func ObjectStorageResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:                   schema.CommandCustom,
			Aliases:                []string{"l", "ls"},
			Params:                 objectStorageListParam(),
			TableType:              output.TableSimple,
			TableColumnDefines:     objectStorageListColumns(),
			UseCustomCommand:       true,
			UseCustomArgCompletion: true,
			ArgsUsage:              "<remote path>",
			SkipAuth:               true,
		},
		"put": {
			Type:                   schema.CommandCustom,
			Params:                 objectStoragePutParam(),
			UseCustomCommand:       true,
			UseCustomArgCompletion: true,
			ArgsUsage:              "<local file/directory> <remote path>",
			SkipAuth:               true,
		},
		"get": {
			Type:                   schema.CommandCustom,
			Params:                 objectStorageGetParam(),
			UseCustomCommand:       true,
			UseCustomArgCompletion: true,
			ArgsUsage:              "<remote path> <local file/directory>",
			SkipAuth:               true,
		},
		"delete": {
			Type:                   schema.CommandCustom,
			Aliases:                []string{"rm", "del"},
			Params:                 objectStorageDelParam(),
			UseCustomCommand:       true,
			UseCustomArgCompletion: true,
			ArgsUsage:              "<remote path>",
			SkipAuth:               true,
			NeedConfirm:            true,
		},
	}

	return &schema.Resource{
		Commands:         commands,
		Aliases:          []string{"ojs"},
		ResourceCategory: CategoryOther,
	}
}

func objectStorageListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "Key"},
		{Name: "Size"},
		{Name: "ETag"},
	}
}

func objectStorageDetailIncludes() []string {
	return []string{}
}

func objectStorageDetailExcludes() []string {
	return []string{}
}

func objectStorageListParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access-key": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set access-key",
			EnvVars:     []string{"AWS_ACCESS_KEY_ID", "SACLOUD_OJS_ACCESS_KEY_ID"},
			Required:    true,
		},
		"secret-key": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set access-key",
			EnvVars:     []string{"AWS_SECRET_ACCESS_KEY", "SACLOUD_OJS_SECRET_ACCESS_KEY"},
			Required:    true,
		},
		"bucket": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set bucket",
			EnvVars:     []string{"SACLOUD_OJS_BUCKET_NAME"},
		},
	}
}

func objectStoragePutParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access-key": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set access-key",
			EnvVars:     []string{"AWS_ACCESS_KEY_ID", "SACLOUD_OJS_ACCESS_KEY_ID"},
			Required:    true,
		},
		"secret-key": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set access-key",
			EnvVars:     []string{"AWS_SECRET_ACCESS_KEY", "SACLOUD_OJS_SECRET_ACCESS_KEY"},
			Required:    true,
		},
		"bucket": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set bucket",
			EnvVars:     []string{"SACLOUD_OJS_BUCKET_NAME"},
		},
		"content-type": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set content-type",
			DefaultValue: "application/octet-stream",
		},
		"recursive": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"r"},
			Description: "put objects recursive",
		},
	}
}

func objectStorageGetParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access-key": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set access-key",
			EnvVars:     []string{"AWS_ACCESS_KEY_ID", "SACLOUD_OJS_ACCESS_KEY_ID"},
			Required:    true,
		},
		"secret-key": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set access-key",
			EnvVars:     []string{"AWS_SECRET_ACCESS_KEY", "SACLOUD_OJS_SECRET_ACCESS_KEY"},
			Required:    true,
		},
		"bucket": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set bucket",
			EnvVars:     []string{"SACLOUD_OJS_BUCKET_NAME"},
		},
		"recursive": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"r"},
			Description: "get objects recursive",
		},
	}
}

func objectStorageDelParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access-key": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set access-key",
			EnvVars:     []string{"AWS_ACCESS_KEY_ID", "SACLOUD_OJS_ACCESS_KEY_ID"},
			Required:    true,
		},
		"secret-key": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set access-key",
			EnvVars:     []string{"AWS_SECRET_ACCESS_KEY", "SACLOUD_OJS_SECRET_ACCESS_KEY"},
			Required:    true,
		},
		"bucket": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "set bucket",
			EnvVars:     []string{"SACLOUD_OJS_BUCKET_NAME"},
		},
		"recursive": {
			Type:        schema.TypeBool,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"r"},
			Description: "delete objects recursive",
		},
	}
}

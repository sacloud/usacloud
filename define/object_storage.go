package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func ObjectStorageResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandCustom,
			Aliases:            []string{"l", "ls"},
			Params:             objectStorageListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: objectStorageListColumns(),
			UseCustomCommand:   true,
			ArgsUsage:          "[PATH]",
		},
		"put": {
			Type:             schema.CommandCustom,
			Params:           objectStoragePutParam(),
			UseCustomCommand: true,
			ArgsUsage:        "[FILE] [PATH]",
		},
		"get": {
			Type:             schema.CommandCustom,
			Params:           objectStorageGetParam(),
			UseCustomCommand: true,
			ArgsUsage:        "[PATH] [FILE]",
		},
		"delete": {
			Type:             schema.CommandCustom,
			Aliases:          []string{"rm", "del"},
			Params:           objectStorageDelParam(),
			UseCustomCommand: true,
			ArgsUsage:        "[PATH]",
		},
	}

	return &schema.Resource{
		Commands: commands,
		Aliases:  []string{"ojs"},
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
		//"path": {
		//	Type:        schema.TypeString,
		//	HandlerType: schema.HandlerNoop,
		//	Description: "set path",
		//},
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
		//TODO add recursive flag
		//"path": {
		//	Type:        schema.TypeString,
		//	HandlerType: schema.HandlerNoop,
		//	Description: "set path",
		//},
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
		//TODO add recursive flag
		//"path": {
		//	Type:        schema.TypeString,
		//	HandlerType: schema.HandlerNoop,
		//	Description: "set path",
		//},
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
		//TODO add recursive flag
		//"path": {
		//	Type:        schema.TypeString,
		//	HandlerType: schema.HandlerNoop,
		//	Description: "set path",
		//},
	}
}

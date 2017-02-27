package define

import (
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

func SSHKeyResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:                schema.CommandList,
			ListResultFieldName: "SSHKeys",
			Aliases:             []string{"l", "ls", "find"},
			Params:              sshKeyListParam(),
			TableType:           output.TableSimple,
			TableColumnDefines:  sshKeyListColumns(),
		},
		"create": {
			Type:             schema.CommandCreate,
			Aliases:          []string{"c"},
			Params:           sshKeyCreateParam(),
			IncludeFields:    sshKeyDetailIncludes(),
			ExcludeFields:    sshKeyDetailExcludes(),
			UseCustomCommand: true,
		},
		"read": {
			Type:          schema.CommandRead,
			Aliases:       []string{"r"},
			Params:        sshKeyReadParam(),
			IncludeFields: sshKeyDetailIncludes(),
			ExcludeFields: sshKeyDetailExcludes(),
		},
		"update": {
			Type:          schema.CommandUpdate,
			Aliases:       []string{"u"},
			Params:        sshKeyUpdateParam(),
			IncludeFields: sshKeyDetailIncludes(),
			ExcludeFields: sshKeyDetailExcludes(),
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"d", "rm"},
			Params:        sshKeyDeleteParam(),
			IncludeFields: sshKeyDetailIncludes(),
			ExcludeFields: sshKeyDetailExcludes(),
		},
		"generate": {
			Type:             schema.CommandManipulate,
			Aliases:          []string{"g", "gen"},
			Params:           sshKeyGenerateParam(),
			IncludeFields:    sshKeyDetailIncludes(),
			ExcludeFields:    sshKeyDetailExcludes(),
			UseCustomCommand: true,
		},
	}

	return &schema.Resource{
		Commands:         commands,
		ResourceCategory: CategoryCommonItem,
	}
}

func sshKeyListParam() map[string]*schema.Schema {
	return CommonListParam
}

func sshKeyListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{
			Name:    "Finger-Print",
			Sources: []string{"Fingerprint"},
		},
	}
}

func sshKeyDetailIncludes() []string {
	return []string{}
}

func sshKeyDetailExcludes() []string {
	return []string{
		"PublicKey",
	}
}

func sshKeyCreateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramRequiredName,
		"description": paramDescription,
		"public-key-content": {
			Type:          schema.TypeString,
			HandlerType:   schema.HandlerNoop,
			Description:   "set public-key",
			ConflictsWith: []string{"public-key"},
		},
		"public-key": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set public-key from file",
			ValidateFunc: validateFileExists(),
		},
	}
}

func sshKeyGenerateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramRequiredName,
		"description": paramDescription,
		"pass-phrase": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set ssh-key pass phrase",
			ValidateFunc: validateStrLen(8, 64),
		},
		"private-key-output": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"file"},
			Description: "set ssh-key privatekey output path",
		},
	}
}

func sshKeyReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

func sshKeyUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id":          paramID,
		"name":        paramName,
		"description": paramDescription,
	}
}

func sshKeyDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": paramID,
	}
}

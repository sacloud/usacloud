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

func SSHKeyResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:               schema.CommandList,
			Aliases:            []string{"ls", "find"},
			Params:             sshKeyListParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: sshKeyListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"create": {
			Type:             schema.CommandCreate,
			Params:           sshKeyCreateParam(),
			IncludeFields:    sshKeyDetailIncludes(),
			ExcludeFields:    sshKeyDetailExcludes(),
			UseCustomCommand: true,
			Category:         "basics",
			Order:            20,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        sshKeyReadParam(),
			IncludeFields: sshKeyDetailIncludes(),
			ExcludeFields: sshKeyDetailExcludes(),
			NoSelector:    true,
			Category:      "basics",
			Order:         30,
		},
		"update": {
			Type:          schema.CommandUpdate,
			Params:        sshKeyUpdateParam(),
			IncludeFields: sshKeyDetailIncludes(),
			ExcludeFields: sshKeyDetailExcludes(),
			NoSelector:    true,
			Category:      "basics",
			Order:         40,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"rm"},
			Params:        sshKeyDeleteParam(),
			IncludeFields: sshKeyDetailIncludes(),
			ExcludeFields: sshKeyDetailExcludes(),
			NoSelector:    true,
			Category:      "basics",
			Order:         50,
		},
		"generate": {
			Type:             schema.CommandCustom,
			Aliases:          []string{"gen"},
			Params:           sshKeyGenerateParam(),
			IncludeFields:    sshKeyDetailIncludes(),
			ExcludeFields:    sshKeyDetailExcludes(),
			UseCustomCommand: true,
			Category:         "basics",
			Order:            60,
		},
	}

	return &schema.Resource{
		Commands:            commands,
		ResourceCategory:    CategoryCommonItem,
		ListResultFieldName: "SSHKeys",
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
		"public-key-content": {
			Type:          schema.TypeString,
			HandlerType:   schema.HandlerNoop,
			Description:   "set public-key",
			ConflictsWith: []string{"public-key"},
			Category:      "input",
			Order:         10,
		},
		"public-key": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set public-key from file",
			ValidateFunc: validateFileExists(),
			Category:     "upload",
			Order:        10,
		},
		"name":        paramRequiredName,
		"description": paramDescription,
	}
}

func sshKeyGenerateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"pass-phrase": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "set ssh-key pass phrase",
			ValidateFunc: validateStrLen(8, 64),
			Category:     "generate",
			Order:        10,
		},
		"private-key-output": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Aliases:     []string{"file"},
			Description: "set ssh-key privatekey output path",
			Category:    "output",
			Order:       10,
		},
		"name":        paramRequiredName,
		"description": paramDescription,
	}
}

func sshKeyReadParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func sshKeyUpdateParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name":        paramName,
		"description": paramDescription,
	}
}

func sshKeyDeleteParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

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

func BridgeResource() *schema.Resource {
	commands := map[string]*schema.Command{
		"list": {
			Type:    schema.CommandList,
			Aliases: []string{"ls", "find"},
			Params:  bridgeListParam(),
			// TableType:          output.TableSimple,
			TableColumnDefines: bridgeListColumns(),
			Category:           "basics",
			Order:              10,
		},
		"create": {
			Type:          schema.CommandCreate,
			Params:        bridgeCreateParam(),
			IncludeFields: bridgeDetailIncludes(),
			ExcludeFields: bridgeDetailExcludes(),
			Category:      "basics",
			Order:         20,
		},
		"read": {
			Type:          schema.CommandRead,
			Params:        bridgeReadParam(),
			IncludeFields: bridgeDetailIncludes(),
			ExcludeFields: bridgeDetailExcludes(),
			NoSelector:    true,
			Category:      "basics",
			Order:         30,
		},
		"update": {
			Type:          schema.CommandUpdate,
			Params:        bridgeUpdateParam(),
			IncludeFields: bridgeDetailIncludes(),
			ExcludeFields: bridgeDetailExcludes(),
			NoSelector:    true,
			Category:      "basics",
			Order:         40,
		},
		"delete": {
			Type:          schema.CommandDelete,
			Aliases:       []string{"rm"},
			Params:        bridgeDeleteParam(),
			IncludeFields: bridgeDetailIncludes(),
			ExcludeFields: bridgeDetailExcludes(),
			NoSelector:    true,
			Category:      "basics",
			Order:         50,
		},
	}

	return &schema.Resource{
		Commands:         commands,
		ResourceCategory: CategoryNetworking,
	}
}

func bridgeListParam() map[string]*schema.Parameter {
	return CommonListParam
}

func bridgeListColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name"},
		{
			Name:    "Switch(this zone only)",
			Sources: []string{"SwitchInZone.ID", "SwitchInZone.Name"},
			Format:  "%s(%s)",
		},
	}
}

func bridgeDetailIncludes() []string {
	return []string{}
}

func bridgeDetailExcludes() []string {
	return []string{
		"Region.Description",
		"Region.NameServers",
	}
}

func bridgeCreateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"name":        paramRequiredName,
		"description": paramDescription,
	}
}

func bridgeReadParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

func bridgeUpdateParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{
		"name":        paramName,
		"description": paramDescription,
	}
}

func bridgeDeleteParam() map[string]*schema.Parameter {
	return map[string]*schema.Parameter{}
}

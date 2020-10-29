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

func AuthStatusResource() *schema.Resource {
	commands := map[string]*schema.Command{
		"show": {
			Type:               schema.CommandCustom,
			Params:             authShowParam(),
			TableType:          output.TableSimple,
			TableColumnDefines: authShowColumns(),
			UseCustomCommand:   true,
			NeedlessConfirm:    true,
		},
	}

	return &schema.Resource{
		Commands:         commands,
		DefaultCommand:   "show",
		ResourceCategory: CategoryAuth,
		IsGlobal:         true,
	}
}

func authShowParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

func authShowColumns() []output.ColumnDef {
	return []output.ColumnDef{
		{Name: "AccountID"},
		{Name: "AccountCode"},
		{Name: "AccountName"},
		{Name: "MemberCode"},
		{Name: "Permission"},
		{Name: "ExternalPermission"},
	}
}

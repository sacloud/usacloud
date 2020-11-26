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

package sim

import (
	"github.com/sacloud/usacloud/pkg/cmd/ccol"
	"github.com/sacloud/usacloud/pkg/output"
)

var defaultColumnDefs = []output.ColumnDef{
	ccol.ID,
	ccol.Name,
	ccol.Tags,
	ccol.Description,
	{Name: "Status", Template: `{{ if .Info }}{{ .SessionStatus }}{{ end }}`},
	{Name: "Activated", Template: `{{ if .Info }}{{ .Activated }}{{ end }}`},
	{Name: "IMEILock", Template: `{{ if .Info }}{{ .IMEILock }}{{ end }}`},
	{Name: "IPAddress", Template: `{{ if .Info }}{{ .IP }}{{ end }}`},
}

// Copyright 2017-2022 The sacloud/usacloud Authors
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

package certificateauthority

import (
	"github.com/sacloud/usacloud/pkg/ccol"
	"github.com/sacloud/usacloud/pkg/output"
)

var (
	defaultColumnDefs = []output.ColumnDef{
		ccol.ID,
		ccol.Name,
		ccol.Tags,
		ccol.Description,
	}

	// Note: CAは他のリソースと異なりReadやCreate/Updateの戻り値が詳細情報を含んでおり、
	//       参照可能な項目は`.CertificateAuthority`というプレフィックス配下にある
	//       このためdefaultColumnDefsが利用できないためここで改めて定義している
	defaultServiceColumnDefs = []output.ColumnDef{
		{Name: "ID"},
		{Name: "Name", Template: "{{ .CertificateAuthority.Name | ellipsis 30 }}"},
		{Name: "Tags", Template: "{{ if .CertificateAuthority.Tags }}{{ .CertificateAuthority.Tags | ellipsis 20 }}{{ end }}"},
		{Name: "Description", Template: "{{ .CertificateAuthority.Description | ellipsis 20 | to_single_line}}"},
	}
)

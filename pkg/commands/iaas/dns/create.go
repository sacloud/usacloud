// Copyright 2017-2022 The Usacloud Authors
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

package dns

import (
	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
	cflag2 "github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/examples"
	"github.com/sacloud/usacloud/pkg/util"
)

var createCommand = &core.Command{
	Name:     "create",
	Category: "basic",
	Order:    20,

	ColumnDefs: defaultColumnDefs,

	ParameterInitializer: func() interface{} {
		return newCreateParameter()
	},
}

type createParameter struct {
	cflag2.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag2.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag2.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag2.NameParameter   `cli:",squash" mapconv:",squash"`
	cflag2.DescParameter   `cli:",squash" mapconv:",squash"`
	cflag2.TagsParameter   `cli:",squash" mapconv:",squash"`
	cflag2.IconIDParameter `cli:",squash" mapconv:",squash"`

	RecordsData string          `cli:"records" mapconv:"-" json:"-"`
	Records     iaas.DNSRecords `cli:"-"`
}

func newCreateParameter() *createParameter {
	return &createParameter{}
}

func init() {
	Resource.AddCommand(createCommand)
}

// Customize パラメータ変換処理
func (p *createParameter) Customize(_ cli.Context) error {
	if p.RecordsData != "" {
		var records iaas.DNSRecords
		if err := util.MarshalJSONFromPathOrContent(p.RecordsData, &records); err != nil {
			return err
		}
		p.Records = append(p.Records, records...)
	}

	return nil
}

func (p *createParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &createParameter{
		NameParameter:   examples.Name,
		DescParameter:   examples.Description,
		TagsParameter:   examples.Tags,
		IconIDParameter: examples.IconID,
		Records: iaas.DNSRecords{
			{
				Name:  "www",
				Type:  types.EDNSRecordType(examples.OptionsString("dns_record_type")),
				RData: examples.IPAddress,
				TTL:   300,
			},
			{
				Name:  "@",
				Type:  types.EDNSRecordType(examples.OptionsString("dns_record_type")),
				RData: examples.IPAddress,
				TTL:   300,
			},
		},
	}
}

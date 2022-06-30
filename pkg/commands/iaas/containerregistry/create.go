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

package containerregistry

import (
	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/iaas-service-go/containerregistry/builder"
	"github.com/sacloud/usacloud/pkg/cflag"
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
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameParameter   `cli:",squash" mapconv:",squash"`
	cflag.DescParameter   `cli:",squash" mapconv:",squash"`
	cflag.TagsParameter   `cli:",squash" mapconv:",squash"`
	cflag.IconIDParameter `cli:",squash" mapconv:",squash"`

	AccessLevel    string `cli:",options=container_registry_access_level" mapconv:",filters=container_registry_access_level_to_value" validate:"required,container_registry_access_level"`
	SubDomainLabel string `cli:"subdomain-label" validate:"required"`
	VirtualDomain  string `validate:"omitempty,fqdn"`

	UsersData string          `cli:"users" mapconv:"-" json:"-"`
	Users     []*builder.User `cli:"-"` // --parametersでファイルからパラメータ指定する場合向け
}

func newCreateParameter() *createParameter {
	return &createParameter{}
}

func init() {
	Resource.AddCommand(createCommand)
}

// Customize パラメータ変換処理
func (p *createParameter) Customize(_ cli.Context) error {
	var users []*builder.User
	if p.UsersData != "" {
		if err := util.MarshalJSONFromPathOrContent(p.UsersData, &users); err != nil {
			return err
		}
	}

	p.Users = append(p.Users, users...)
	return nil
}

func (p *createParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &createParameter{
		NameParameter:   examples.Name,
		DescParameter:   examples.Description,
		TagsParameter:   examples.Tags,
		IconIDParameter: examples.IconID,
		AccessLevel:     examples.OptionsString("container_registry_access_level"),
		SubDomainLabel:  "your-sub-domain",
		VirtualDomain:   "your-domain.example.com",
		Users: []*builder.User{
			{
				UserName:   "example-user-name",
				Password:   "example-password",
				Permission: types.EContainerRegistryPermission(examples.OptionsString("container_registry_permission")),
			},
		},
	}
}

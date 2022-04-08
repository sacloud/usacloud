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

package containerregistry

import (
	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/iaas-service-go/containerregistry/builder"
	"github.com/sacloud/packages-go/pointer"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/cmd/examples"
	"github.com/sacloud/usacloud/pkg/util"
)

var updateCommand = &core.Command{
	Name:         "update",
	Category:     "basic",
	Order:        40,
	SelectorType: core.SelectorTypeRequireMulti,

	ColumnDefs: defaultColumnDefs,

	ParameterInitializer: func() interface{} {
		return newUpdateParameter()
	},
}

type updateParameter struct {
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.DescUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.TagsUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.IconIDUpdateParameter `cli:",squash" mapconv:",omitempty,squash"`

	AccessLevel    *string `cli:",options=container_registry_access_level" mapconv:",omitempty,filters=dereference,container_registry_access_level_to_value" validate:"omitempty,container_registry_access_level"`
	SubDomainLabel *string `cli:"subdomain-label" validate:"omitempty"`
	VirtualDomain  *string `validate:"omitempty,fqdn"`

	UsersData *string          `cli:"users" mapconv:"-" json:"-"`
	Users     *[]*builder.User `cli:"-"` // --parametersでファイルからパラメータ指定する場合向け
}

func newUpdateParameter() *updateParameter {
	return &updateParameter{}
}

func init() {
	Resource.AddCommand(updateCommand)
}

// Customize パラメータ変換処理
func (p *updateParameter) Customize(_ cli.Context) error {
	var users []*builder.User
	if p.UsersData != nil && *p.UsersData != "" {
		if err := util.MarshalJSONFromPathOrContent(*p.UsersData, &users); err != nil {
			return err
		}
		p.Users = &users
	}

	return nil
}

func (p *updateParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &updateParameter{
		NameUpdateParameter:   examples.NameUpdate,
		DescUpdateParameter:   examples.DescriptionUpdate,
		TagsUpdateParameter:   examples.TagsUpdate,
		IconIDUpdateParameter: examples.IconIDUpdate,
		AccessLevel:           pointer.NewString(examples.OptionsString("container_registry_access_level")),
		SubDomainLabel:        pointer.NewString("your-sub-domain"),
		VirtualDomain:         pointer.NewString("your-domain.example.com"),
		Users: &[]*builder.User{
			{
				UserName:   "example-user-name",
				Password:   "example-password",
				Permission: types.EContainerRegistryPermission(examples.OptionsString("container_registry_permission")),
			},
		},
	}
}

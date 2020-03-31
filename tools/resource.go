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

package tools

import (
	"fmt"

	"github.com/sacloud/usacloud/schema"
)

// Resource コード生成時に利用するリソース定義
type Resource struct {
	*schema.Resource

	Name     string
	Commands []*Command
}

// NewResource コード生成時に利用するリソース定義オブジェクトの作成
func NewResource(name string, resource *schema.Resource) *Resource {
	r := &Resource{
		Resource: resource,
		Name:     name,
	}
	var commands []*Command
	for _, c := range r.Resource.SortedCommands() {
		commands = append(commands, NewCommand(c.CommandKey, c.Command, c.Category, r))
	}
	r.Commands = commands

	return r
}

func (r *Resource) CLIName() string {
	return ToDashedName(r.Name)
}

func (r *Resource) AliasesLiteral() string {
	return FlattenStringList(r.Resource.Aliases)
}

func (r *Resource) Usage() string {
	usage := r.Resource.Usage
	if usage == "" {
		usage = fmt.Sprintf("A manage commands of %s", r.Name)
	}
	return usage
}

func (r *Resource) CLIVariableName() string {
	return fmt.Sprintf("%sCmd", ToCamelWithFirstLower(r.Name))
}

func (r *Resource) CLISourceFileName() string {
	return fmt.Sprintf("zz_%s_gen.go", ToSnakeCaseName(r.Name))
}

func (r *Resource) ParameterSourceFileName() string {
	return fmt.Sprintf("zz_%s_gen.go", ToSnakeCaseName(r.Name))
}

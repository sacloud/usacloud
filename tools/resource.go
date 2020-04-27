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
	"sort"

	"github.com/sacloud/usacloud/pkg/schema"
)

// Resource コード生成時に利用するリソース定義
type Resource struct {
	*schema.Resource

	Name                string
	Commands            []*Command
	CategorizedCommands []*CategorizedCommands
}

type CategorizedCommands struct {
	*schema.Category
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
	r.buildCategorizedCommands()
	return r
}

func (r *Resource) buildCategorizedCommands() {
	m := map[string]*CategorizedCommands{}
	for _, c := range r.Commands {
		cat := c.Category
		cc, ok := m[cat.Key]
		if !ok {
			cc = &CategorizedCommands{
				Category: cat,
			}
		}
		cc.Commands = append(cc.Commands, c)
		m[cat.Key] = cc
	}
	r.CategorizedCommands = []*CategorizedCommands{}
	for _, cat := range m {
		r.CategorizedCommands = append(r.CategorizedCommands, cat)
	}
	sort.Slice(r.CategorizedCommands, func(i, j int) bool {
		if r.CategorizedCommands[i].Order == r.CategorizedCommands[j].Order {
			return r.CategorizedCommands[i].Key < r.CategorizedCommands[j].Key
		}
		return r.CategorizedCommands[i].Order < r.CategorizedCommands[j].Order
	})
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

func (r *Resource) CLIVariableFuncName() string {
	return fmt.Sprintf("%sCmd", ToCamelWithFirstLower(r.Name))
}

func (r *Resource) CLISourceFileName() string {
	return fmt.Sprintf("zz_%s_gen.go", ToSnakeCaseName(r.Name))
}

func (r *Resource) CLIResourceFinderSourceFileName() string {
	return fmt.Sprintf("zz_%s_finder_gen.go", ToSnakeCaseName(r.Name))
}

func (r *Resource) CLIUsageFileName() string {
	return fmt.Sprintf("zz_%s_usage_gen.go", ToSnakeCaseName(r.Name))
}

func (r *Resource) CLINormalizeFlagsFileName() string {
	return fmt.Sprintf("zz_%s_normalize_flag_names_gen.go", ToSnakeCaseName(r.Name))
}

func (r *Resource) ParameterSourceFileName() string {
	return fmt.Sprintf("zz_%s_gen.go", ToSnakeCaseName(r.Name))
}

func (r *Resource) CommandOrderFunc() string {
	return fmt.Sprintf("%sCommandOrder", ToCamelWithFirstLower(r.Name))
}

func (r *Resource) PackageDirName() string {
	n := ToLowerName(r.Name)
	switch n {
	case "switch":
		return "swytch"
	case "interface":
		return "iface"
	default:
		return n
	}
}

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

	"github.com/sacloud/usacloud/pkg/utils"

	"github.com/sacloud/usacloud/schema"
)

// Command コード生成時に利用するコマンド定義
type Command struct {
	*schema.Command

	Name              string
	Resource          *Resource
	Category          *schema.Category
	Params            []*Parameter
	CategorizedParams []*CategorizedParameters
}

type CategorizedParameters struct {
	*schema.Category
	Params []*Parameter
}

func NewCommand(name string, command *schema.Command, category *schema.Category, resource *Resource) *Command {
	c := &Command{
		Command:  command,
		Name:     name,
		Resource: resource,
		Category: category,
	}
	var params []*Parameter
	for _, p := range c.Command.BuiltParams() {
		params = append(params, NewParameter(p.ParamKey, p.Param, p.Category, c))
	}

	c.Params = params
	c.buildCategorizedParams()
	return c
}

func (c *Command) buildCategorizedParams() {
	m := map[string]*CategorizedParameters{}
	for _, p := range c.Params {
		c := p.Category
		cp, ok := m[c.Key]
		if !ok {
			cp = &CategorizedParameters{
				Category: c,
			}
		}
		cp.Params = append(cp.Params, p)
		m[c.Key] = cp
	}
	c.CategorizedParams = []*CategorizedParameters{}
	for _, cat := range m {
		c.CategorizedParams = append(c.CategorizedParams, cat)
	}
	sort.Slice(c.CategorizedParams, func(i, j int) bool {
		return c.CategorizedParams[i].Order < c.CategorizedParams[j].Order
	})
}

func (c *Command) ExperimentWarning() string {
	if c.Command.ExperimentWarning != "" {
		return c.Command.ExperimentWarning
	}
	return c.Resource.ExperimentWarning
}

func (c *Command) Usage() string {
	usage := c.Command.Usage
	if usage == "" {
		usage = fmt.Sprintf("%s %s", ToCamelCaseName(c.Name), ToCamelCaseName(c.Resource.Name))
	}
	if c.Resource.DefaultCommand == c.Name {
		usage = fmt.Sprintf("%s (default)", usage)
	}
	return usage
}

func (c *Command) ArgsUsage() string {
	argsUsage := c.Command.ArgsUsage
	if argsUsage == "" && c.Type.IsRequiredIDType() {
		t := c.Type
		switch {
		case t.IsNeedIDOnlyType():
			argsUsage = "<ID>"
		case t.IsNeedSingleIDType():
			argsUsage = "<ID or Name(only single target)>"
		default:
			argsUsage = "<ID or Name(allow multiple target)>"

		}
	}
	return argsUsage
}

func (c *Command) AliasesLiteral() string {
	return FlattenStringList(c.Command.Aliases)
}

func (c *Command) HasOutputOption() bool {
	return !c.NoOutput
}

func (c *Command) CLIVariableFuncName() string {
	return fmt.Sprintf("%s%sCmd", ToCamelWithFirstLower(c.Resource.Name), ToCamelCaseName(c.Name))
}

func (c *Command) CLIv2CommandsFileName() string {
	return fmt.Sprintf("%s_gen.go", ToSnakeCaseName(c.Resource.Name))
}

func (c *Command) InputParameterVariable() string {
	return fmt.Sprintf("%s%sParam", ToCamelWithFirstLower(c.Resource.Name), ToCamelCaseName(c.Name))
}

func (c *Command) InputParameterTypeName() string {
	return fmt.Sprintf("%s%sParam", ToCamelCaseName(c.Name), ToCamelCaseName(c.Resource.Name))
}

func (c *Command) FunctionName() string {
	return fmt.Sprintf("%s%s", ToCamelCaseName(c.Resource.Name), ToCamelCaseName(c.Name))
}

func (c *Command) NeedConfirm() bool {
	return c.Type.IsNeedConfirmType() && !c.NeedlessConfirm
}

func (c *Command) ConfirmMessage() string {
	if c.Command.ConfirmMessage == "" {
		return ToDashedName(c.Name)
	}
	return c.Command.ConfirmMessage
}

func (c *Command) RequireID() bool {
	return c.Command.Type.IsRequiredIDType()
}

func (c *Command) SingleArgToIdParam() bool {
	return c.Command.Type.IsNeedIDOnlyType()
}

func (c *Command) MultipleArgToIdParams() bool {
	return c.RequireID() && !c.SingleArgToIdParam()
}

func (c *Command) ArgToIdFunc() string {
	return fmt.Sprintf("find%s%sTargets", ToCamelCaseName(c.Resource.Name), ToCamelCaseName(c.Name))
}

func (c *Command) FlagOrderFunc() string {
	return fmt.Sprintf("%s%sFlagOrder", ToCamelWithFirstLower(c.Resource.Name), ToCamelCaseName(c.Name))
}

func (c *Command) TargetAPIName() string {
	return utils.FirstNonEmptyString(c.AltResource, c.Resource.AltResource, ToCamelCaseName(c.Resource.Name))
}

func (c *Command) FindResultFieldName() string {
	return utils.FirstNonEmptyString(c.ListResultFieldName, c.Resource.ListResultFieldName, ToCamelCaseName(c.Resource.Name)+"s")
}

func (c *Command) RequireSingleID() bool {
	return c.Type.IsNeedSingleIDType()
}

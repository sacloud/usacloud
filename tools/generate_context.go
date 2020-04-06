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
	"go/build"
	"path/filepath"
	"sort"

	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

type GenerateContext struct {
	// for v0
	R           string
	C           string
	P           string
	ResourceDef map[string]*schema.Resource
	// for v1
	Resources            []*Resource
	CategorizedResources []*CategorizedResources
}

type CategorizedResources struct {
	*schema.Category
	Resources []*Resource
}

func NewGenerateContext() *GenerateContext {
	ctx := &GenerateContext{
		ResourceDef: define.Resources,
	}
	for rn, r := range define.Resources {
		ctx.Resources = append(ctx.Resources, NewResource(rn, r))
	}
	sort.Slice(ctx.Resources, func(i, j int) bool {
		if ctx.Resources[i].ResourceCategory.Order == ctx.Resources[j].ResourceCategory.Order {
			if ctx.Resources[i].ResourceCategory.Key == ctx.Resources[j].ResourceCategory.Key {
				return ctx.Resources[i].Name < ctx.Resources[j].Name
			}
			return ctx.Resources[i].ResourceCategory.Key < ctx.Resources[j].ResourceCategory.Key
		}
		return ctx.Resources[i].ResourceCategory.Order < ctx.Resources[j].ResourceCategory.Order
	})

	ctx.buildCategorizedResources()
	return ctx
}

func (c *GenerateContext) buildCategorizedResources() {
	m := map[string]*CategorizedResources{}
	for _, r := range c.Resources {
		c := &r.ResourceCategory
		cr, ok := m[c.Key]
		if !ok {
			cr = &CategorizedResources{
				Category: c,
			}
		}
		cr.Resources = append(cr.Resources, r)
		m[c.Key] = cr
	}
	c.CategorizedResources = []*CategorizedResources{}
	for _, cat := range m {
		c.CategorizedResources = append(c.CategorizedResources, cat)
	}
	sort.Slice(c.CategorizedResources, func(i, j int) bool {
		if c.CategorizedResources[i].Order == c.CategorizedResources[j].Order {
			return c.CategorizedResources[i].Key < c.CategorizedResources[j].Key
		}
		return c.CategorizedResources[i].Order < c.CategorizedResources[j].Order
	})
}

func (c *GenerateContext) SetCurrentR(k string) {
	c.R = k
}

/*
   Get name functions
*/

func (c *GenerateContext) CamelR() string {
	return ToCamelCaseName(c.R)
}

func (c *GenerateContext) Camelr() string {
	return ToCamelWithFirstLower(c.R)
}

func (c *GenerateContext) DashR() string {
	return ToDashedName(c.R)
}

func (c *GenerateContext) SnakeR() string {
	return ToSnakeCaseName(c.R)
}

func (c *GenerateContext) CamelC() string {
	return ToCamelCaseName(c.C)
}

func (c *GenerateContext) Camelc() string {
	return ToCamelWithFirstLower(c.C)
}

func (c *GenerateContext) DashC() string {
	return ToDashedName(c.C)
}

func (c *GenerateContext) SnakeC() string {
	return ToSnakeCaseName(c.C)
}

func (c *GenerateContext) CamelP() string {
	return ToCamelCaseName(c.P)
}

func (c *GenerateContext) Camelp() string {
	return ToCamelWithFirstLower(c.P)
}

func (c *GenerateContext) DashP() string {
	return ToDashedName(c.P)
}

func (c *GenerateContext) SnakeP() string {
	return ToSnakeCaseName(c.P)
}

func (c *GenerateContext) Gopath() string {
	gopath := build.Default.GOPATH
	gopath = filepath.SplitList(gopath)[0]
	return gopath
}

/*
   Get current context schema
*/

func (c *GenerateContext) CurrentResource() *schema.Resource {
	return c.ResourceDef[c.R]
}

func (c *GenerateContext) CurrentCommand() *schema.Command {
	return c.CurrentResource().Commands[c.C]
}

func (c *GenerateContext) CurrentParam() *schema.Schema {
	return c.CurrentCommand().Params[c.P]
}

/*
   Get contextual name functions
*/

func (c *GenerateContext) InputModelFileName() string {
	return fmt.Sprintf("params_%s_gen.go", c.SnakeR())
}

func (c *GenerateContext) InputModelTypeName() string {
	return fmt.Sprintf("%s%sParam", c.CamelC(), c.CamelR())
}

func (c *GenerateContext) InputParamFieldName() string {
	return c.CamelP()
}

func (c *GenerateContext) InputParamFlagName() string {
	return c.DashP()
}

func (c *GenerateContext) InputParamSetterFuncName() string {
	n := c.CurrentParam().DestinationProp
	if n == "" {
		n = fmt.Sprintf("Set%s", c.CamelP())
	}
	return n
}

func (c *GenerateContext) InputParamDestinationName() string {
	n := c.CurrentParam().DestinationProp
	if n == "" {
		n = c.CamelP()
	}
	return n
}

func (c *GenerateContext) InputParamCLIFlagName() string {
	return ToCLIFlagName(c.P)
}

func (c *GenerateContext) InputParamVariableName() string {
	return fmt.Sprintf("%s%sParam", ToCamelWithFirstLower(c.R), ToCamelCaseName(c.C))
}

func (c *GenerateContext) CommandFuncName() string {
	return fmt.Sprintf("%s%s", c.CamelR(), c.CamelC())
}

func (c *GenerateContext) CommandFileName(useCustomCommand bool) string {
	if useCustomCommand {
		return fmt.Sprintf("%s_%s.go", c.SnakeR(), c.SnakeC())
	}
	return fmt.Sprintf("%s_%s_gen.go", c.SnakeR(), c.SnakeC())
}

func (c *GenerateContext) CLICommandsFileName() string {
	return fmt.Sprintf("cli_%s_gen.go", c.SnakeR())
}

func (c *GenerateContext) CLIv2CommandsFileName() string {
	return fmt.Sprintf("%s_gen.go", c.SnakeR())
}

func (c *GenerateContext) CommandResourceName() string {
	n := c.CurrentCommand().AltResource
	if n == "" {
		n = c.CurrentResource().AltResource
	}
	if n == "" {
		n = c.CamelR()
	}
	return n
}

func (c *GenerateContext) FindResultFieldName() string {
	n := c.CurrentCommand().ListResultFieldName
	if n == "" {
		n = c.CurrentResource().ListResultFieldName
	}

	if n == "" {
		n = c.CamelR() + "s"
	}
	return n
}

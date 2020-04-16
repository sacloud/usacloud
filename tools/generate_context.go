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

	"github.com/sacloud/usacloud/version"

	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

type GenerateContext struct {
	Resources            []*Resource
	CategorizedResources []*CategorizedResources
}

type CategorizedResources struct {
	*schema.Category
	Resources []*Resource
}

func NewGenerateContext() *GenerateContext {
	ctx := &GenerateContext{}
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

func (c *GenerateContext) Copyright() string {
	return fmt.Sprintf("Copyright %s The Usacloud Authors", version.CopyrightYear)
}

func (c *GenerateContext) Gopath() string {
	gopath := build.Default.GOPATH
	gopath = filepath.SplitList(gopath)[0]
	return gopath
}

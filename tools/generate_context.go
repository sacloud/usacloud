// Copyright 2017-2025 The sacloud/usacloud Authors
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
	"log"
	"path/filepath"

	"github.com/sacloud/usacloud/pkg"
	"github.com/sacloud/usacloud/pkg/version"
)

type GenerateContext struct {
	Resources []*Resource
}

func NewGenerateContext() *GenerateContext {
	// command schema validation
	for _, r := range pkg.Resources() {
		for _, c := range r.Commands() {
			if err := c.ValidateSchema(); err != nil {
				log.Fatal(err)
			}
		}
	}
	return &GenerateContext{
		Resources: NewResources(pkg.Resources()),
	}
}

func (c *GenerateContext) Copyright() string {
	return fmt.Sprintf("Copyright %s The Usacloud Authors", version.CopyrightYear)
}

func (c *GenerateContext) Gopath() string {
	gopath := build.Default.GOPATH
	gopath = filepath.SplitList(gopath)[0]
	return gopath
}

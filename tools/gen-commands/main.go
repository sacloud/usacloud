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

package main

import (
	_ "embed"
	"log"
	"path/filepath"

	"github.com/sacloud/usacloud/tools"
	"github.com/sacloud/usacloud/tools/utils"
)

var (
	destination = "pkg"
	ctx         = tools.NewGenerateContext()

	//go:embed flags.tmpl
	flagsTemplate string

	//go:embed init_service.tmpl
	initServiceTemplate string

	//go:embed service_command.tmpl
	serviceCommandTemplate string
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("gen-commands: ")

	for _, resource := range ctx.Resources {
		for _, command := range resource.Commands {
			// flag関連ソースの生成
			filePath := filepath.Join(destination, "commands", resource.PlatformName, resource.PackageDirName(), command.CLICommandGeneratedSourceFile())
			utils.WriteFileWithTemplate(&utils.TemplateConfig{
				OutputPath: filepath.Join(utils.ProjectRootPath(), filePath),
				Template:   flagsTemplate,
				Parameter:  command,
			})
		}
		for _, child := range resource.Children() {
			childResource := tools.NewResource(child)
			for _, command := range childResource.Commands {
				// flag関連ソースの生成
				filePath := filepath.Join(destination, "commands", childResource.PlatformName, childResource.PackageDirName(), command.CLICommandGeneratedSourceFile())
				utils.WriteFileWithTemplate(&utils.TemplateConfig{
					OutputPath: filepath.Join(utils.ProjectRootPath(), filePath),
					Template:   flagsTemplate,
					Parameter:  command,
				})
			}
		}

		if resource.PlatformName != "" {
			// サービス初期化コードの生成
			{
				filePath := filepath.Join(destination, "commands", resource.PlatformName, "services_gen.go")
				utils.WriteFileWithTemplate(&utils.TemplateConfig{
					OutputPath: filepath.Join(utils.ProjectRootPath(), filePath),
					Template:   initServiceTemplate,
					Parameter:  resource,
				})
			}

			// service呼び出し関連ソースの生成
			{
				// service呼び出し関連ソースの生成
				filePath := filepath.Join(destination, "services", resource.PlatformName, resource.ServiceSourceFileName())
				utils.WriteFileWithTemplate(&utils.TemplateConfig{
					OutputPath: filepath.Join(utils.ProjectRootPath(), filePath),
					Template:   serviceCommandTemplate,
					Parameter:  resource,
				})

				for _, child := range resource.Children() {
					childResource := tools.NewResource(child)
					// リソース単位のファイルを生成
					filePath := filepath.Join(destination, "services", resource.PlatformName, resource.ChildResourceServiceSourceFileName(childResource))
					utils.WriteFileWithTemplate(&utils.TemplateConfig{
						OutputPath: filepath.Join(utils.ProjectRootPath(), filePath),
						Template:   serviceCommandTemplate,
						Parameter:  childResource,
					})
				}
			}
		}
	}
}

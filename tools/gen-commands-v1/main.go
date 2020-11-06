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

package main

import (
	"log"
	"path/filepath"

	"github.com/sacloud/usacloud/tools"
	"github.com/sacloud/usacloud/tools/utils"
)

var (
	destination = "pkg/cmd"
	ctx         = tools.NewGenerateContext()
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("gen-commands-v1: ")

	for _, resource := range ctx.Resources {
		hasV1Command := false
		for _, command := range resource.Commands {
			// コマンド単位のファイルを生成

			// TODO 実験的実装、パラメータが設定されている場合だけ処理する
			if command.Parameters == nil {
				continue
			}
			hasV1Command = true

			// flag関連ソースの生成
			filePath := filepath.Join(destination, resource.PackageDirName(), command.CLICommandGeneratedSourceFile())
			utils.WriteFileWithTemplate(&utils.TemplateConfig{
				OutputPath: filepath.Join(utils.ProjectRootPath(), filePath),
				Template:   flagsTemplate,
				Parameter:  command,
			})
		}

		// リソース単位のファイルを生成
		if hasV1Command {
			// libsacloud service呼び出し関連ソースの生成
			filePath := filepath.Join(destination, resource.PackageDirName(), resource.ServiceSourceFileName())
			utils.WriteFileWithTemplate(&utils.TemplateConfig{
				OutputPath: filepath.Join(utils.ProjectRootPath(), filePath),
				Template:   serviceCommandTemplate,
				Parameter:  resource,
			})

			// cli commands関連ソースの生成
			filePath = filepath.Join(destination, resource.PackageDirName(), resource.CLICommandsSourceFileName())
			utils.WriteFileWithTemplate(&utils.TemplateConfig{
				OutputPath: filepath.Join(utils.ProjectRootPath(), filePath),
				Template:   cliCommandsTemplate,
				Parameter:  resource,
			})
		}
	}
}

// Copyright 2017-2021 The Usacloud Authors
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

package utils

import (
	"bytes"
	"go/build"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

// TemplateConfig ソース生成を行うためのテンプレート設定
type TemplateConfig struct {
	OutputPath         string
	Template           string
	Parameter          interface{}
	PreventOverwriting bool
}

// WriteFileWithTemplate 指定の設定に従いファイル出力
func WriteFileWithTemplate(config *TemplateConfig) bool {
	buf := bytes.NewBufferString("")
	t := template.New("t")
	template.Must(t.Parse(config.Template))
	if err := t.Execute(buf, config.Parameter); err != nil {
		log.Fatalf("writing output: %s", err)
	}

	// create dir
	if _, err := os.Stat(filepath.Dir(config.OutputPath)); err != nil && os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(config.OutputPath), 0755); err != nil {
			log.Fatal(err)
		}
	}

	if config.PreventOverwriting {
		if _, err := os.Stat(config.OutputPath); err == nil {
			return false
		}
	}

	// write to file
	if err := os.WriteFile(config.OutputPath, Sformat(buf.Bytes()), 0644); err != nil {
		log.Fatalf("writing output: %s", err)
	}
	log.Printf("generated: %s\n", config.OutputPath)
	return true
}

// Gopath returns GOPATH
func Gopath() string {
	gopath := build.Default.GOPATH
	gopath = filepath.SplitList(gopath)[0]
	return gopath
}

// ProjectRootPath プロジェクトルートパス
func ProjectRootPath() string {
	return filepath.Join(Gopath(), "src/github.com/sacloud/usacloud")
}

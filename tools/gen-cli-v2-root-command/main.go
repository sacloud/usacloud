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
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/sacloud/usacloud/tools"
)

var (
	destination = "src/github.com/sacloud/usacloud/cmdv2/commands"
	ctx         = tools.NewGenerateContext()
)

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprint(os.Stderr, "\tgen-cli-v2-root-command\n")
	os.Exit(2)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("gen-cli-v2-root-command: ")

	filePath := filepath.Join(destination, "root_init_gen.go")
	fileFullPath := filepath.Join(ctx.Gopath(), filePath)

	src, err := generateSource()
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}

	err = ioutil.WriteFile(fileFullPath, tools.Sformat([]byte(src)), 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}
	fmt.Printf("generated: %s\n", filePath)
}

func generateSource() (string, error) {
	buf := bytes.NewBufferString("")
	t := template.New("t")
	template.Must(t.Parse(srcTemplate))
	err := t.Execute(buf, ctx)
	return buf.String(), err
}

var srcTemplate = `// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-v2-root-command'; DO NOT EDIT

package commands

import (
	"github.com/spf13/cobra"
)

func rootCommandOrder(cmd *cobra.Command) []*commandSet {
	var commands []*commandSet
	{{ range .CategorizedResources -}}
	{
		set := &commandSet {
			title: "{{.DisplayName}}",
		}
		{{ range .Resources -}}
		set.commands = append(set.commands, lookupCmd(cmd, "{{.CLIName}}"))
		{{ end -}}
		commands = append(commands, set)
	}
	{{ end }}
	return commands	
}

func init() {
	{{ range .Resources -}}
	{
		cmd := {{ .CLIVariableFuncName }}()
		{{ range .Commands -}}
		cmd.AddCommand({{ .CLIVariableFuncName }}())
		{{ end -}}
		buildCommandsUsage(cmd, {{ .CommandOrderFunc }}(cmd))
		rootCmd.AddCommand(cmd)
	}
	{{ end }}
	buildCommandsUsage(rootCmd, rootCommandOrder(rootCmd))
}
`

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

// Usage is a replacement usage function for the flags package.
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprint(os.Stderr, "\tgen-cli-v2-commands\n")
	os.Exit(2)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("gen-cli-v2-commands: ")

	for _, resource := range ctx.Resources {
		filePath := filepath.Join(destination, resource.CLISourceFileName())
		fileFullPath := filepath.Join(ctx.Gopath(), filePath)

		src, err := generateSource(resource)
		if err != nil {
			log.Fatalf("writing output: %s", err)
		}

		err = ioutil.WriteFile(fileFullPath, tools.Sformat([]byte(src)), 0644)
		if err != nil {
			log.Fatalf("writing output: %s", err)
		}
		fmt.Printf("generated: %s\n", filePath)
	}
}

func generateSource(resource *tools.Resource) (string, error) {
	buf := bytes.NewBufferString("")
	t := template.New("t")
	template.Must(t.Parse(srcTemplate))
	err := t.Execute(buf, resource)
	return buf.String(), err
}

var srcTemplate = `// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-v2-commands'; DO NOT EDIT

package commands

import (
	"errors"
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/sacloud/usacloud/pkg/utils"
	"github.com/spf13/cobra"
)

// {{ .CLIVariableFuncName }} represents the command to manage SAKURA Cloud {{ .Name }}
func {{ .CLIVariableFuncName }}() *cobra.Command {
	return &cobra.Command {
		Use:   "{{ .CLIName }}",
		Short: "{{ .Usage }}",
		Long: ` + "`{{.Usage}}`" + `,
		Run: func(cmd *cobra.Command, args []string) {
			{{ if .DefaultCommand }}// TODO not implements: call {{.DefaultCommand}} func as default{{ else }}cmd.HelpFunc()(cmd,args){{ end }}
		},
	}
}

{{ range .Commands }}
func {{ .CLIVariableFuncName }}() *cobra.Command {
	{{ .InputParameterVariable }} := params.New{{ .InputParameterTypeName }}()
	cmd := &cobra.Command{
		Use:   "{{ .Name }}",
		{{ if .Aliases }}Aliases: []string{ {{ .AliasesLiteral }} },{{ end }}
		Short: "{{ .Usage }}",
		Long: ` + "`{{ .Usage }}`" + `,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return {{ .InputParameterVariable }}.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, {{ .InputParameterVariable }})
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("{{.ExperimentWarning}}")

			if {{ .InputParameterVariable }}.GenerateSkeleton {
				return generateSkeleton(ctx, {{ .InputParameterVariable }})
			}
			
			{{ if .MultipleArgToIdParams -}}
			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := {{ .ArgToIdFunc }}(ctx, {{ .InputParameterVariable }})
			if err != nil {
				return err
			}
			{{ end -}}

			{{ if .NeedConfirm }}
			// confirm
			if !{{.InputParameterVariable}}.Assumeyes {
				if !utils.IsTerminal(){
				    return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("{{.ConfirmMessage}}", ctx.IO().In(), ctx.IO().Out(){{ if .MultipleArgToIdParams }}, ids...{{ end }})
				if err != nil || !result {
					return err
				}
			}
			{{ end }}

			{{ if .MultipleArgToIdParams -}}
			var wg sync.WaitGroup
			var errs []error
			for _ , id := range ids {
				wg.Add(1)
				{{ .InputParameterVariable }}.SetId(id)
				go func(p *params.{{ .InputParameterTypeName }}) {
					err := funcs.{{ .FunctionName }}(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}({{ .InputParameterVariable }})
			}
			wg.Wait()
			return command.FlattenErrors(errs)
			{{ else }}
			return funcs.{{ .FunctionName }}(ctx, {{ .InputParameterVariable }}.ToV0())
			{{ end }}
		},
	}

	{{ if .Params -}}
	fs := cmd.Flags()
{{ range .Params -}}
	fs.{{ .FlagDefinitionStatement }}
{{ end -}}
	setFlagsUsage(cmd, buildFlagsUsage({{.FlagOrderFunc}}(cmd)))
{{ end -}}

	return cmd
}
{{ end }}

func init() {
	parent := {{ .CLIVariableFuncName }}()
{{ range .Commands -}}
	parent.AddCommand({{ .CLIVariableFuncName }}())
{{ end -}}
	rootCmd.AddCommand(parent)
}
`

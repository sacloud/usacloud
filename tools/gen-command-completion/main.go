package main

import (
	"bytes"
	"fmt"
	"github.com/sacloud/usacloud/schema"
	"github.com/sacloud/usacloud/tools"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	destination = "src/github.com/sacloud/usacloud/command/completion"
	ctx         = tools.NewGenerateContext()
)

// Usage is a replacement usage function for the flags package.
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\tgen-command-completion\n")
	os.Exit(2)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("gen-command-funcs: ")

	for k, r := range ctx.ResourceDef {
		ctx.SetCurrentR(k)

		err := generateResource(r)
		if err != nil {
			log.Fatalf("writing output: %s", err)
		}
	}
}

func generateResource(resource *schema.Resource) error {

	normalArgCompletions := []schema.SortableCommand{}
	normalFlagCompletions := []schema.SortableCommand{}

	// build custom completions per command
	for _, comm := range resource.SortedCommands() {
		c := comm.Command
		k := comm.CommandKey
		ctx.C = k

		// Write to file.
		// like 'comp_switch_list_args.go'
		baseName := ctx.CommandArgsCompletionFileName(c.UseCustomArgCompletion)
		outputName := filepath.Join(ctx.Gopath(), destination, baseName)

		if c.UseCustomArgCompletion {
			// target file is exist?
			_, err := os.Stat(outputName)
			if err != nil {
				// generate args completion
				src, err := generateArgsComplete(comm)
				if err != nil {
					return err
				}

				err = ioutil.WriteFile(outputName, tools.Sformat([]byte(src)), 0644)
				if err != nil {
					return err
				}
				fmt.Printf("generated: %s\n", filepath.Join(destination, baseName))
			}
		} else {
			normalArgCompletions = append(normalArgCompletions, comm)
		}

		// Write to file.
		// like 'comp_switch_list_flags.go'
		baseName = ctx.CommandFlagsCompletionFileName(c.UseCustomFlagsCompletion)
		outputName = filepath.Join(ctx.Gopath(), destination, baseName)

		if c.UseCustomFlagsCompletion {
			// target file is exist?
			_, err := os.Stat(outputName)
			if err != nil {

				// generate flags completion
				src, err := generateFlagsComplete(comm)
				if err != nil {
					return err
				}

				err = ioutil.WriteFile(outputName, tools.Sformat([]byte(src)), 0644)
				if err != nil {
					return err
				}
				fmt.Printf("generated: %s\n", filepath.Join(destination, baseName))
			}
		} else {
			normalFlagCompletions = append(normalFlagCompletions, comm)
		}
	}

	// build completions per resource

	// args
	if len(normalArgCompletions) > 0 {

		baseName := ctx.ResourceArgsCompletionFileName()
		outputName := filepath.Join(ctx.Gopath(), destination, baseName)
		// generate args completion
		src, err := generateArgsComplete(normalArgCompletions...)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(outputName, tools.Sformat([]byte(src)), 0644)
		if err != nil {
			return err
		}
		fmt.Printf("generated: %s\n", filepath.Join(destination, baseName))
	}

	if len(normalFlagCompletions) > 0 {

		// flags
		baseName := ctx.ResourceFlagsCompletionFileName()
		outputName := filepath.Join(ctx.Gopath(), destination, baseName)
		// generate args completion
		src, err := generateFlagsComplete(normalFlagCompletions...)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(outputName, tools.Sformat([]byte(src)), 0644)
		if err != nil {
			return err
		}
		fmt.Printf("generated: %s\n", filepath.Join(destination, baseName))
	}

	return nil
}

func generateArgsComplete(commands ...schema.SortableCommand) (string, error) {
	b := bytes.NewBufferString("")
	t := template.New("c")
	template.Must(t.Parse(completeArgsTemplate))

	paramCommands := []map[string]interface{}{}
	useImport := false

	for _, command := range commands {
		c := command.Command
		k := command.CommandKey
		ctx.C = k

		action, err := generateArgsCompleteAction(c)
		if err != nil {
			return "", err
		}
		paramCommands = append(paramCommands, map[string]interface{}{
			"FuncName":  ctx.CompleteArgsFuncName(),
			"ParamName": ctx.InputModelTypeName(),
			"Action":    action,
		})
		if action != "" {
			useImport = true
		}
	}

	err := t.Execute(b, map[string]interface{}{
		"NeedDonotEditComment": !commands[0].Command.UseCustomArgCompletion,
		"UseImport":            useImport,
		"Commands":             paramCommands,
	})

	return b.String(), err
}

func generateArgsCompleteAction(command *schema.Command) (string, error) {

	var res string
	var err error

	if command.Type.IsRequiredIDType() {
		res, err = generateIDCompletion(command)
	} else {
		res, err = "", nil
	}

	return res, err
}

func generateIDCompletion(command *schema.Command) (string, error) {
	b := bytes.NewBufferString("")
	t := template.New("c")
	template.Must(t.Parse(completeIDTemplate))

	err := t.Execute(b, map[string]interface{}{
		"FinderFieldName": ctx.CommandResourceName(),
		"ListResultField": ctx.FindResultFieldName(),
		"IsNeedNameComp":  !command.Type.IsNeedIDOnlyType(),
	})
	return b.String(), err
}

func generateFlagsComplete(commands ...schema.SortableCommand) (string, error) {
	b := bytes.NewBufferString("")
	t := template.New("c")
	t.Funcs(template.FuncMap{
		"join": strings.Join,
	})
	template.Must(t.Parse(completeFlagsTemplate))

	paramCommands := []map[string]interface{}{}
	useImport := false

	for _, command := range commands {

		c := command.Command
		k := command.CommandKey
		ctx.C = k

		flags := []map[string]interface{}{}
		for _, param := range c.BuildedParams() {
			p := param.Param
			ctx.P = param.ParamKey

			names := []string{param.ParamKey}
			names = append(names, p.Aliases...)
			for i := range names {
				names[i] = fmt.Sprintf("%q", names[i])
			}

			flags = append(flags, map[string]interface{}{
				"ResourceKey": ctx.R,
				"CommandKey":  ctx.C,
				"ParamKey":    ctx.P,
				"Names":       names,
				"OutputFlag":  p.Category == "output",
				"InputFlag":   p.Category == "input",
			})
		}

		// output/inputフラグ以外を利用する場合はdefineパッケージをimportする
		for _, f := range flags {
			if !f["OutputFlag"].(bool) && !f["InputFlag"].(bool) {
				useImport = true
				break
			}
		}

		// outputフラグがある場合、固定でoutput-typeへの補完処理を追加する
		hasOutputFlag := false
		for _, f := range flags {
			if f["OutputFlag"].(bool) {
				hasOutputFlag = true
				break
			}
		}

		paramCommands = append(paramCommands, map[string]interface{}{
			"FuncName":      ctx.CompleteFlagsFuncName(),
			"ParamName":     ctx.InputModelTypeName(),
			"Flags":         flags,
			"HasOutputFlag": hasOutputFlag,
		})
	}

	err := t.Execute(b, map[string]interface{}{
		"NeedDonotEditComment": !commands[0].Command.UseCustomFlagsCompletion,
		"UseImport":            useImport,
		"Commands":             paramCommands,
	})

	return b.String(), err
}

var completeArgsTemplate = `{{ if .NeedDonotEditComment }}// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT{{ end }}

package completion

import ({{if .UseImport}}
    "fmt"{{end}}
    "github.com/sacloud/usacloud/command"
    "github.com/sacloud/usacloud/command/params"
)
{{range .Commands}}
func {{.FuncName}}(ctx command.Context, params *params.{{.ParamName}}, cur, prev, commandName string) {
    {{.Action}}
}
{{end}}
`

var completeIDTemplate = `
	if !command.GlobalOption.Valid {
		return
	}

	{{if not .IsNeedNameComp}}if cur != "" && !isSakuraID(cur){
		return
	}{{end}}

	client := ctx.GetAPIClient()
	finder := client.Get{{.FinderFieldName}}API()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}

	type nameHolder interface {
		GetName() string
	}

	for i := range res.{{.ListResultField}} {
		fmt.Println(res.{{.ListResultField}}[i].ID)
		{{if .IsNeedNameComp}}var target interface{} = &res.{{.ListResultField}}[i]
		if v, ok := target.(nameHolder); ok {
			fmt.Println(v.GetName())
		}
		{{end}}
	}
`

var completeFlagsTemplate = `{{ if .NeedDonotEditComment }}// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT{{ end }}

package completion

import (
	"github.com/sacloud/usacloud/schema"{{if .UseImport}}
	"github.com/sacloud/usacloud/define"{{end}}
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"fmt"
)
{{range .Commands}}
func {{.FuncName}}(ctx command.Context, params *params.{{.ParamName}} , flagName string , currentValue string) {
    	var comp schema.CompletionFunc

	switch flagName { {{range .Flags}}{{ if not .OutputFlag }}{{ if not .InputFlag }}
	case {{join .Names ", "}}:
		param := define.Resources["{{.ResourceKey}}"].Commands["{{.CommandKey}}"].BuildedParams().Get("{{.ParamKey}}")
		if param != nil {
	 		comp = param.Param.CompleteFunc
	 	}{{end}}{{end}}{{end}}
	{{ if .HasOutputFlag }}case "output-type", "out":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
{{ end -}}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}
{{end}}
`

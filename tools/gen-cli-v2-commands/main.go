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
	"strings"
	"text/template"

	"github.com/sacloud/usacloud/schema"
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

	for k, resource := range ctx.ResourceDef {
		ctx.R = k

		// Write to file.
		baseName := ctx.CLICommandsFileName()
		outputName := filepath.Join(ctx.Gopath(), destination, strings.ToLower(baseName))

		src, err := generateSource(resource)
		if err != nil {
			log.Fatalf("writing output: %s", err)
		}

		err = ioutil.WriteFile(outputName, tools.Sformat([]byte(src)), 0644)
		if err != nil {
			log.Fatalf("writing output: %s", err)
		}
		fmt.Printf("generated: %s\n", filepath.Join(destination, strings.ToLower(baseName)))
	}
}

func generateSource(resource *schema.Resource) (string, error) {
	// build commands
	var commands []map[string]interface{}
	var parameters string
	var defaultCommandFlags interface{}
	for _, comm := range resource.SortedCommands() {
		c := comm.Command
		k := comm.CommandKey

		ctx.C = k

		params, err := buildCommandsParams(resource, c)
		if err != nil {
			return "", err
		}
		commands = append(commands, params)

		// add parameters
		paramName := ctx.InputParamVariableName()
		paramTypeName := ctx.InputModelTypeName()
		parameters += fmt.Sprintf("%s := params.New%s()\n", paramName, paramTypeName)

		if resource.DefaultCommand != "" && k == resource.DefaultCommand {
			defaultCommandFlags = params["Flags"]
		}
	}

	// parameters
	usage := resource.Usage
	if usage == "" {
		usage = fmt.Sprintf("A manage commands of %s", ctx.R)
	}

	// hasIdRequiredType Command
	needsyncImport := false
	for _, comm := range resource.SortedCommands() {
		if comm.Command.Type.IsRequiredIDType() && !comm.Command.Type.IsNeedIDOnlyType() {
			needsyncImport = true
		}
		if needsyncImport {
			break
		}
	}

	buf := bytes.NewBufferString("")
	t := template.New("t")
	template.Must(t.Parse(srcTemplate))
	err := t.Execute(buf, map[string]interface{}{
		"ShortName":        ctx.Camelr(),
		"Name":             ctx.DashR(),
		"Aliases":          tools.FlattenStringList(resource.Aliases),
		"Usage":            usage,
		"DefaultCommand":   resource.DefaultCommand,
		"Flags":            defaultCommandFlags,
		"Commands":         commands,
		"Parameters":       parameters,
		"IsNeedSyncImport": needsyncImport,
	})
	return buf.String(), err
}

func buildCommandsParams(resource *schema.Resource, command *schema.Command) (map[string]interface{}, error) {
	var res map[string]interface{}

	flags, err := buildFlagsParams(command.BuildedParams())
	if err != nil {
		return res, err
	}

	usage := command.Usage
	if usage == "" {
		usage = fmt.Sprintf("%s %s", ctx.CamelC(), ctx.CamelR())
	}
	if ctx.CurrentResource().DefaultCommand == ctx.C {
		usage = fmt.Sprintf("%s (default)", usage)
	}

	argsUsage := command.ArgsUsage
	if command.ArgsUsage == "" && command.Type.IsRequiredIDType() {
		t := command.Type
		switch {
		case t.IsNeedIDOnlyType():
			argsUsage = "<ID>"
		case t.IsNeedSingleIDType():
			argsUsage = "<ID or Name(only single target)>"
		default:
			argsUsage = "<ID or Name(allow multiple target)>"

		}

	}

	experimentWarning := resource.ExperimentWarning
	if command.ExperimentWarning != "" {
		experimentWarning = command.ExperimentWarning
	}

	res = map[string]interface{}{
		"CommandObjectName":      ctx.Camelr() + ctx.CamelC() + "Cmd",
		"Name":                   ctx.C,
		"Aliases":                tools.FlattenStringList(command.Aliases),
		"Usage":                  usage,
		"ArgsUsage":              argsUsage,
		"Flags":                  flags,
		"ApplyConfigFile":        !ctx.CurrentResource().SkipApplyConfigFile,
		"ExperimentWarning":      experimentWarning,
		"InputParameterVariable": ctx.InputParamVariableName(),
		"InputParameterType":     ctx.InputModelTypeName(),
	}

	action, err := buildActionParams(command)
	if err != nil {
		return res, err
	}
	for k, v := range action {
		res[k] = v
	}

	return res, err
}

func buildFlagsParams(params schema.SortableParams) ([]map[string]interface{}, error) {

	var res []map[string]interface{}

	if len(params) == 0 {
		return res, nil
	}

	for _, param := range params {

		s := param.Param
		k := param.ParamKey

		ctx.P = k

		d := ""
		if s.DefaultValue != nil {
			switch s.DefaultValue.(type) {
			case bool:
				d = fmt.Sprintf("%t", s.DefaultValue)
			case int, int64:
				d = fmt.Sprintf("%d", s.DefaultValue)
			case string:
				d = fmt.Sprintf("\"%s\"", s.DefaultValue)
			case []string:
				d = fmt.Sprintf("cli.NewStringSlice(\"%s\")", strings.Join(s.DefaultValue.([]string), "\",\""))
			case []int64:
				tmp := []string{}
				for _, v := range s.DefaultValue.([]int64) {
					tmp = append(tmp, fmt.Sprintf("%d", v))
				}
				d = fmt.Sprintf("cli.NewInt64Slice(%s)", strings.Join(tmp, ","))
			default:
				return res, fmt.Errorf("Set DefaultValue is not implement: %v", s.DefaultValue)
			}

		}

		ts, err := getFlagTypeString(s.Type)
		if err != nil {
			return res, err
		}

		usage := s.Description
		if s.Required {
			usage = fmt.Sprintf("[Required] %s", usage)
		}

		param := map[string]interface{}{
			"FlagType":     ts,
			"Name":         ctx.InputParamFlagName(),
			"Aliases":      tools.FlattenStringList(s.Aliases),
			"Usage":        usage,
			"EnvVars":      tools.FlattenStringList(s.EnvVars),
			"DefaultValue": d,
			"DefaultText":  s.DefaultText,
			"PropName":     ctx.InputParamFieldName(),
			"Hidden":       s.Hidden,
		}
		res = append(res, param)
	}

	return res, nil
}

var setDefaultTemplate = `if c.IsSet("%s") {
	%s.%s = c.%s("%s")
}
`

var setDefaultIdTemplate = `if c.IsSet("%s") {
	%s.%s = sacloud.ID(c.%s("%s"))
}
`

var setDefaultIdListTemplate = `if c.IsSet("%s") {
	%s.%s = toSakuraIDs(c.%s("%s"))
}
`

var setDefaultWithEnvTemplate = `if c.IsSet("%s") || command.IsEmpty(%s.%s) {
	%s.%s = c.%s("%s")
}
`

var setDefaultIdWithEnvTemplate = `if c.IsSet("%s") || command.IsEmpty(%s.%s) {
	%s.%s = sacloud.ID(c.%s("%s"))
}
`
var setDefaultIdListWithEnvTemplate = `if c.IsSet("%s") || command.IsEmpty(%s.%s) {
	%s.%s = toSakuraIDs(c.%s("%s"))
}
`

func buildActionParams(command *schema.Command) (map[string]interface{}, error) {

	var res map[string]interface{}

	// build params
	paramName := ctx.InputParamVariableName()
	setDefault := ""
	for _, param := range command.BuildedParams() {
		k := param.ParamKey
		p := param.Param
		ctx.P = k

		propName := ctx.InputParamFieldName()
		flagName := ctx.InputParamFlagName()
		valueFuncName, err := getFlagValueFuncString(p.Type)
		if err != nil {
			return res, err
		}

		if valueFuncName != "" {
			if len(p.EnvVars) == 0 {
				switch p.Type {
				case schema.TypeId:
					setDefault += fmt.Sprintf(setDefaultIdTemplate,
						flagName, paramName, propName, valueFuncName, flagName)
				case schema.TypeIdList:
					setDefault += fmt.Sprintf(setDefaultIdListTemplate,
						flagName, paramName, propName, valueFuncName, flagName)
				default:
					setDefault += fmt.Sprintf(setDefaultTemplate,
						flagName, paramName, propName, valueFuncName, flagName)
				}
			} else {
				switch p.Type {
				case schema.TypeId:
					setDefault += fmt.Sprintf(setDefaultIdWithEnvTemplate,
						flagName, paramName, propName, paramName, propName, valueFuncName, flagName)
				case schema.TypeIdList:
					setDefault += fmt.Sprintf(setDefaultIdListWithEnvTemplate,
						flagName, paramName, propName, paramName, propName, valueFuncName, flagName)
				default:
					setDefault += fmt.Sprintf(setDefaultWithEnvTemplate,
						flagName, paramName, propName, paramName, propName, valueFuncName, flagName)
				}
			}
		}
	}
	action := fmt.Sprintf("%s(ctx , %s)", ctx.CommandFuncName(), paramName)

	needConfirm := false
	confirmMsg := command.ConfirmMessage
	if command.Type.IsNeedConfirmType() && !command.NeedlessConfirm {
		needConfirm = true
	}
	if confirmMsg == "" {
		confirmMsg = fmt.Sprintf("%s", ctx.DashC())
	}

	createParamFuncName := fmt.Sprintf("params.New%s", ctx.InputModelTypeName())

	res = map[string]interface{}{
		"ParamName":           paramName,
		"NoSelector":          command.NoSelector,
		"CreateParamFunc":     createParamFuncName,
		"SkipAuth":            ctx.CurrentCommand().SkipAuth,
		"SetDefault":          setDefault,
		"Action":              action,
		"NeedConfirm":         needConfirm,
		"ConfirmMsg":          confirmMsg,
		"IdParamRequired":     command.Type.IsRequiredIDType(),
		"CommandResourceName": ctx.CommandResourceName(),
		"FindResultName":      ctx.FindResultFieldName(),
		"IsNeedSingleID":      command.Type.IsNeedSingleIDType(),
		"IsNeedIDOnlyType":    command.Type.IsNeedIDOnlyType(),
		"SkipAfterSecondArgs": command.SkipAfterSecondArgs,
	}

	return res, nil
}

func getFlagTypeString(t schema.ValueType) (string, error) {
	switch t {
	case schema.TypeBool:
		return "BoolFlag", nil
	case schema.TypeInt:
		return "IntFlag", nil
	case schema.TypeInt64:
		return "Int64Flag", nil
	case schema.TypeFloat:
		return "FloatFlag", nil
	case schema.TypeString:
		return "StringFlag", nil
	case schema.TypeIntList:
		return "Int64SliceFlag", nil
	case schema.TypeStringList:
		return "StringSliceFlag", nil
	case schema.TypeId:
		return "Int64Flag", nil
	case schema.TypeIdList:
		return "Int64SliceFlag", nil
	}

	return "", fmt.Errorf("Inalid type: %v", t)
}

func getFlagValueFuncString(t schema.ValueType) (string, error) {
	switch t {
	case schema.TypeBool:
		return "Bool", nil
	case schema.TypeInt:
		return "Int", nil
	case schema.TypeInt64:
		return "Int64", nil
	case schema.TypeFloat:
		return "Float64", nil
	case schema.TypeString:
		return "String", nil
	case schema.TypeIntList:
		return "Int64Slice", nil
	case schema.TypeStringList:
		return "StringSlice", nil
	case schema.TypeId:
		return "Int64", nil
	case schema.TypeIdList:
		return "Int64Slice", nil
	}

	return "", fmt.Errorf("Inalid type: %v", t)
}

var srcTemplate = `// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-v2-commands'; DO NOT EDIT

package commands

import (
	"fmt"

	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/spf13/cobra"
)

// {{ .ShortName }}Cmd represents the {{ .ShortName }} command
var {{ .ShortName }}Cmd = &cobra.Command{
	Use:   "{{ .Name }}",
	Short: "{{ .Usage }}",
	Long: ` + "`{{.Usage}}`" + `,
	Run: func(cmd *cobra.Command, args []string) {
		{{ if .DefaultCommand }}// TODO not implements: call {{.DefaultCommand}} func as default{{ else }}cmd.HelpFunc()(cmd,args){{ end }}
	},
}

{{ range .Commands }}
var {{ .CommandObjectName }} = &cobra.Command{
	Use:   "{{ .Name }}",
	{{ if .Aliases }}Aliases: []string{ {{ .Aliases }} },{{ end }}
	Short: "{{ .Usage }}",
	Long: ` + "`{{ .Usage }}`" + `,
	RunE: func(cmd *cobra.Command, args []string) error {
		{{ .InputParameterVariable }}, err := params.New{{ .InputParameterType }}(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("{{.Name}} parameter: \n%s\n", debugMarshalIndent({{.InputParameterVariable}}))
		return err
	},
}
{{ end }}

func init() {
	parent := {{ .ShortName }}Cmd
{{ range .Commands -}}
	parent.AddCommand({{ .CommandObjectName }})
{{ end -}}

	rootCmd.AddCommand({{ .ShortName }}Cmd)
}

`

/*
	"InputParameterVariable": ctx.InputParamVariableName(),
	"InputParameterType": ctx.InputModelTypeName(),
*/

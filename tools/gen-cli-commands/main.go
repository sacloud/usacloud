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
	destination = "src/github.com/sacloud/usacloud/command"
	ctx         = tools.NewGenerateContext()
)

// Usage is a replacement usage function for the flags package.
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\tgen-cli-commands\n")
	os.Exit(2)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("gen-cli-commands: ")

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
	for k, c := range resource.Commands {
		ctx.C = k

		params, err := buildCommandsParams(c)
		if err != nil {
			return "", err
		}
		commands = append(commands, params)

		// add parameters
		paramName := ctx.InputParamVariableName()
		paramTypeName := ctx.InputModelTypeName()
		parameters += fmt.Sprintf("%s := New%s()\n", paramName, paramTypeName)
	}

	// parameters
	usage := resource.Usage
	if usage == "" {
		usage = fmt.Sprintf("A manage commands of %s", ctx.R)
	}

	buf := bytes.NewBufferString("")
	t := template.New("t")
	template.Must(t.Parse(srcTemplate))
	err := t.Execute(buf, map[string]interface{}{
		"Name":       ctx.DashR(),
		"Aliases":    tools.FlattenStringList(resource.Aliases),
		"Usage":      usage,
		"Commands":   commands,
		"Parameters": parameters,
	})
	return buf.String(), err
}

func buildCommandsParams(command *schema.Command) (map[string]interface{}, error) {

	var res map[string]interface{}

	flags, err := buildFlagsParams(command.Params)
	if err != nil {
		return res, err
	}

	usage := command.Usage
	if usage == "" {
		usage = fmt.Sprintf("%s %s", ctx.CamelC(), ctx.CamelR())
	}
	argsUsage := ""
	if command.Type.IsRequiredIDType() {
		argsUsage = "[ResourceID]"
	}

	res = map[string]interface{}{
		"Name":      ctx.C,
		"Aliases":   tools.FlattenStringList(command.Aliases),
		"Usage":     usage,
		"ArgsUsage": argsUsage,
		"Flags":     flags,
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

func buildFlagsParams(params map[string]*schema.Schema) ([]map[string]interface{}, error) {

	var res []map[string]interface{}

	if len(params) == 0 {
		return res, nil
	}

	for k, s := range params {
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
				d = fmt.Sprintf("[]string{\"%s\"}", strings.Join(s.DefaultValue.([]string), "\",\""))
			case []int, []int64:
				d = fmt.Sprintf("[]string{\"%s\"}", strings.Join(s.DefaultValue.([]string), "\",\""))
			default:
				return res, fmt.Errorf("Set DefaultValue is not implement: %v", s.DefaultValue)
			}

		}

		ts, err := getFlagTypeString(s.Type)
		if err != nil {
			return res, err
		}

		dest := ""
		if !s.Type.IsSliceType() {
			dest = fmt.Sprintf("&%s", ctx.InputParamVariableName())
		}

		usage := s.Description
		if s.Required {
			usage = fmt.Sprintf("[Required] %s", usage)
		}

		param := map[string]interface{}{
			"FlagType":        ts,
			"Name":            ctx.InputParamFlagName(),
			"Aliases":         tools.FlattenStringList(s.Aliases),
			"Usage":           usage,
			"EnvVars":         tools.FlattenStringList(s.EnvVars),
			"DefaultValue":    d,
			"DefaultText":     s.DefaultText,
			"DestinationName": dest,
			"PropName":        ctx.InputParamFieldName(),
		}
		res = append(res, param)
	}

	return res, nil
}

func buildActionParams(command *schema.Command) (map[string]interface{}, error) {

	var res map[string]interface{}

	// build params
	paramName := ctx.InputParamVariableName()
	setDefault := ""
	for k, p := range command.Params {
		ctx.P = k

		propName := ctx.InputParamFieldName()
		flagName := ctx.InputParamFlagName()
		valueFuncName, err := getSliceFlagTypeFuncString(p.Type)
		if err != nil {
			return res, err
		}

		if valueFuncName != "" {
			setDefault += fmt.Sprintf("%s.%s = c.%s(\"%s\")\n", paramName, propName, valueFuncName, flagName)
		}
	}
	action := fmt.Sprintf("return %s(ctx , %s)", ctx.CommandFuncName(), paramName)

	res = map[string]interface{}{
		"ParamName":         paramName,
		"SkipAuth":          ctx.CurrentCommand().SkipAuth,
		"NeedAsignFromArgs": ctx.CurrentCommand().Type.IsRequiredIDType(),
		"SetDefault":        setDefault,
		"Action":            action,
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
	}

	return "", fmt.Errorf("Inalid type: %v", t)
}

func getSliceFlagTypeFuncString(t schema.ValueType) (string, error) {
	switch t {
	case schema.TypeBool, schema.TypeInt, schema.TypeInt64, schema.TypeFloat, schema.TypeString:
		return "", nil
	case schema.TypeIntList:
		return "Int64Slice", nil
	case schema.TypeStringList:
		return "StringSlice", nil
	}

	return "", fmt.Errorf("Inalid type: %v", t)
}

var srcTemplate = `// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-commands'; DO NOT EDIT

package command

import (
    "gopkg.in/urfave/cli.v2"
)

func init() {
        {{.Parameters}}
	cliCommand := &cli.Command{
		Name: "{{.Name}}",
		{{- if .Aliases }}
			Aliases: []string{ {{.Aliases}} },{{ end }}
		{{- if .Usage}}
			Usage: "{{.Usage}}",{{ end }}
		Subcommands:[]*cli.Command{
			{{ range .Commands -}}
			{
				Name: "{{.Name}}",
				{{- if .Aliases }}
					Aliases: []string{ {{.Aliases}} },{{ end }}
				{{- if .Usage }}
					Usage: "{{.Usage}}",{{ end }}
				{{- if .ArgsUsage }}
					ArgsUsage: "{{.ArgsUsage}}",{{ end }}
				{{- if .Flags }}
				Flags: []cli.Flag{
					{{ range .Flags -}}
					&cli.{{.FlagType}}{
						Name:        "{{.Name}}",
						{{- if .Aliases}}
							Aliases:     []string{ {{.Aliases}} },{{ end }}
						{{- if .Usage}}
							Usage:       "{{.Usage}}",{{ end }}
						{{- if .EnvVars}}
							EnvVars:     []string{ {{.EnvVars}} },{{ end }}
						{{- if .DefaultValue}}
							Value:       {{.DefaultValue}},{{ end }}
						{{- if .DefaultText}}
							DefaultText: "{{.DefaultText}}",{{ end }}
					        {{- if .DestinationName}}
					        	Destination: {{.DestinationName}}.{{.PropName}},{{ end }}
					},
					{{ end }}
				},{{ end }}
				Action: func(c *cli.Context) error {

					{{ if .SetDefault }}
					// Set option values for slice
					{{.SetDefault}}{{ end }}

					// Validate global params
					if errors := GlobalOption.Validate({{.SkipAuth}}); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors,"GlobalOptions")
					}

					{{ if .NeedAsignFromArgs }}
					// id is can set from option or args(first)
					if c.NArg() == 1 {
						c.Set("id", c.Args().First())
					}{{ end }}

					// Validate specific for each command params
					if errors := {{.ParamName}}.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors,"Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), {{.ParamName}})

					// Run command with params
					{{.Action}}
				},
			},
			{{ end }}
		},
	}

	Commands = append(Commands, cliCommand)
}`

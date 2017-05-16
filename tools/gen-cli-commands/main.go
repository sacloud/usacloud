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
	fmt.Fprint(os.Stderr, "\tgen-cli-commands\n")
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
	for _, comm := range resource.SortedCommands() {
		c := comm.Command
		k := comm.CommandKey

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

	// build category resource map
	type categoryResourceMapValue struct {
		*schema.Category
		ResourceKey string
	}
	categoryResourceMap := categoryResourceMapValue{
		ResourceKey: ctx.DashR(),
		Category:    &resource.ResourceCategory,
	}

	// build category command map
	type categoryCommandMapValue struct {
		*schema.Category
		ResourceKey string
		CommandKey  string
	}
	categoryCommandMap := map[string]interface{}{}
	for _, comm := range resource.SortedCommands() {
		k := comm.CommandKey
		category := comm.Category
		if category == nil {
			category = schema.DefaultCommandCategory
		}
		categoryCommandMap[k] = categoryCommandMapValue{
			ResourceKey: ctx.DashR(),
			CommandKey:  k,
			Category:    category,
		}
	}

	// build category param map
	type categoryMapValue struct {
		*schema.Category
		CommandKey  string
		ResourceKey string
	}
	categoryParamMap := map[string]map[string]interface{}{}
	for _, comm := range resource.SortedCommands() {
		c := comm.Command
		k := comm.CommandKey
		categoryParamMap[k] = map[string]interface{}{}
		for _, param := range c.BuildedParams() {
			categoryParamMap[k][param.ParamKey] = &categoryMapValue{
				Category:    param.Category,
				CommandKey:  k,
				ResourceKey: ctx.DashR(),
			}
		}
	}

	// hasIdRequiredType Command
	needadditionalImport := false
	needsyncImport := false
	for _, comm := range resource.SortedCommands() {
		if comm.Command.Type.IsRequiredIDType() {
			needadditionalImport = true
		}
		if comm.Command.Type.IsRequiredIDType() && !comm.Command.Type.IsNeedIDOnlyType() {
			needsyncImport = true
		}

		if needadditionalImport && needsyncImport {
			break
		}

	}

	buf := bytes.NewBufferString("")
	t := template.New("t")
	template.Must(t.Parse(srcTemplate))
	err := t.Execute(buf, map[string]interface{}{
		"Name":                   ctx.DashR(),
		"Aliases":                tools.FlattenStringList(resource.Aliases),
		"Usage":                  usage,
		"DefaultCommand":         resource.DefaultCommand,
		"Commands":               commands,
		"Parameters":             parameters,
		"CategoryResourceMap":    categoryResourceMap,
		"CategoryCommandMap":     categoryCommandMap,
		"CategoryParamMap":       categoryParamMap,
		"IsNeedAdditionalImport": needadditionalImport,
		"IsNeedSyncImport":       needsyncImport,
	})
	return buf.String(), err
}

func buildCommandsParams(command *schema.Command) (map[string]interface{}, error) {

	var res map[string]interface{}

	flags, err := buildFlagsParams(command.BuildedParams())
	if err != nil {
		return res, err
	}

	usage := command.Usage
	if usage == "" {
		usage = fmt.Sprintf("%s %s", ctx.CamelC(), ctx.CamelR())
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
			"Hidden":          s.Hidden,
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
	for _, param := range command.BuildedParams() {
		k := param.ParamKey
		p := param.Param
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
	action := fmt.Sprintf("%s(ctx , %s)", ctx.CommandFuncName(), paramName)

	needConfirm := false
	confirmMsg := command.ConfirmMessage
	if command.Type.IsNeedConfirmType() && !command.NeedlessConfirm {
		needConfirm = true
	}
	if confirmMsg == "" {
		confirmMsg = fmt.Sprintf("%s", ctx.DashC())
	}

	res = map[string]interface{}{
		"ParamName":             paramName,
		"SkipAuth":              ctx.CurrentCommand().SkipAuth,
		"SetDefault":            setDefault,
		"Action":                action,
		"CompleteArgsFuncName":  ctx.CompleteArgsFuncName(),
		"CompleteFlagsFuncName": ctx.CompleteFlagsFuncName(),
		"NeedConfirm":           needConfirm,
		"ConfirmMsg":            confirmMsg,
		"IdParamRequired":       command.Type.IsRequiredIDType(),
		"CommandResourceName":   ctx.CommandResourceName(),
		"FindResultName":        ctx.FindResultFieldName(),
		"IsNeedSingleID":        command.Type.IsNeedSingleIDType(),
		"IsNeedIDOnlyType":      command.Type.IsNeedIDOnlyType(),
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
    "github.com/sacloud/usacloud/schema"
    "strings"
    {{ if .IsNeedAdditionalImport }}"fmt"
    {{ if .IsNeedSyncImport }}"sync"{{end}}{{end}}
)

func init() {
        {{.Parameters}}
	cliCommand := &cli.Command{
		Name: "{{.Name}}",
		{{- if .Aliases }}
			Aliases: []string{ {{.Aliases}} },{{ end }}
		{{- if .Usage}}
			Usage: "{{.Usage}}",{{ end }}
		{{- if .DefaultCommand }}
			Action: func(c *cli.Context) error {
				comm := c.App.Command("{{.DefaultCommand}}")
				if comm != nil {
					return comm.Run(c)
				}
				return cli.ShowSubcommandHelp(c)
			},{{ end }}
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
					        {{- if .Hidden}}
					        	Hidden: {{.Hidden}},{{ end }}
					},
					{{ end }}
				},{{ end }}
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					// c.Args() == arg1 arg2 arg3 -- [cur] [prev] [commandName]
					args := c.Args().Slice()
					commandName := args[c.NArg()-1]
					prev := args[c.NArg()-2]
					cur := args[c.NArg()-3]

					// set real args
					realArgs := args[0 : c.NArg()-3]

					// Validate global params
					GlobalOption.Validate(false)

					// build command context
					ctx := NewContext(c, realArgs, {{.ParamName}})
					{{ if .SetDefault }}
					// Set option values for slice
					{{.SetDefault}}{{ end }}


					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completionFlagNames(c, commandName)
								return
							} else {
								{{.CompleteArgsFuncName}}(ctx , {{.ParamName}}, cur, prev, commandName)
								return
							}
						}

						// cleanup flag name
						name := prev
						for {
							if !strings.HasPrefix(name, "-") {
								break
							}
							name = strings.Replace(name, "-", "", 1)
						}

						// flag is exists? , is BoolFlag?
						exists := false
						for _, flag := range c.App.Command(commandName).Flags {

							for _, n := range flag.Names() {
								if n == name {
									exists = true
									break
								}
							}

							if exists {
								if _, ok := flag.(*cli.BoolFlag); ok {
									if strings.HasPrefix(cur, "-") {
										completionFlagNames(c, commandName)
										return
									} else {
										{{.CompleteArgsFuncName}}(ctx , {{.ParamName}}, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									{{.CompleteFlagsFuncName}}(ctx, {{.ParamName}}, name, cur)
									return
								}
							}
						}
						// here, prev is wrong, so noop.
					} else {
						if strings.HasPrefix(cur, "-") {
							completionFlagNames(c, commandName)
							return
						} else {
							{{.CompleteArgsFuncName}}(ctx , {{.ParamName}}, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					{{ if .SetDefault }}
					// Set option values for slice
					{{.SetDefault}}{{ end }}

					// Validate global params
					if errors := GlobalOption.Validate({{.SkipAuth}}); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors,"GlobalOptions")
					}

					{{ if .IsNeedIDOnlyType }}
					if c.NArg() == 0 {
						return fmt.Errorf("ID argument is required")
					}
					c.Set("id", c.Args().First())
					{{ end }}

					// Validate specific for each command params
					if errors := {{.ParamName}}.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors,"Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), {{.ParamName}})

					{{if and .IdParamRequired (not .IsNeedIDOnlyType) }}
					if c.NArg() == 0 {
						return fmt.Errorf("ID or Name argument is required")
					}
					idOrName := c.Args().First()
					apiClient := ctx.GetAPIClient().{{.CommandResourceName}}
					ids := []int64{}

					if id, ok := toSakuraID(idOrName); ok {
						ids = append(ids, id)
					} else {
						apiClient.SetFilterBy("Name", idOrName)
						res, err := apiClient.Find()
						if err != nil {
							return fmt.Errorf("Find ID is failed: %s", err)
						}
						for _, v := range res.{{.FindResultName}} {
							ids = append(ids, v.GetID())
						}
					}

					if len(ids) == 0 {
						return fmt.Errorf("ID or Name argument is required")
					}

					{{ if .IsNeedSingleID }}
					if len(ids) != 1 {
						return fmt.Errorf("Can't run with multiple targets: %v", ids)
					}
					{{ end }}

					{{ if .NeedConfirm }}
					// confirm
					if !{{.ParamName}}.Assumeyes && !confirmContinue("{{.ConfirmMsg}}", ids...) {
						return nil
					}
					{{ end }}

					wg := sync.WaitGroup{}
					errs := []error{}

					for _, id := range ids {
						wg.Add(1)
						{{.ParamName}}.SetId(id)
						p := *{{.ParamName}} // copy struct value
						{{.ParamName}} := &p
						go func() {
							err := {{.Action}}
							if err != nil {
								errs = append(errs, err)
							}
							wg.Done()
						}()
					}
					wg.Wait()
					return flattenErrors(errs)

					{{ else }}
					{{ if .NeedConfirm }}
					// confirm
					if !{{.ParamName}}.Assumeyes && !confirmContinue("{{.ConfirmMsg}}") {
						return nil
					}
					{{ end }}
					// Run command with params
					return {{.Action}}
					{{ end }}
				},
			},
			{{ end }}
		},
	}

	// build Category-Resource mapping
	appendResourceCategoryMap("{{.CategoryResourceMap.ResourceKey}}" , &schema.Category{
		Key:		"{{.CategoryResourceMap.Key}}",
		DisplayName:	"{{.CategoryResourceMap.DisplayName}}",
		Order:		{{.CategoryResourceMap.Order}},
	})

	// build Category-Command mapping
	{{ range .CategoryCommandMap}}
	appendCommandCategoryMap("{{.ResourceKey}}", "{{.CommandKey}}", &schema.Category{
		Key:		"{{.Key}}",
		DisplayName:	"{{.DisplayName}}",
		Order:		{{.Order}},
	}){{end}}

	// build Category-Param mapping
	{{ range .CategoryParamMap}}{{ range $paramKey , $category := . }}
	appendFlagCategoryMap("{{.ResourceKey}}", "{{.CommandKey}}", "{{$paramKey}}", &schema.Category{
		Key:		"{{$category.Key}}",
		DisplayName:	"{{$category.DisplayName}}",
		Order:		{{$category.Order}},
	}){{end}}{{end}}

	// append command to GlobalContext
	Commands = append(Commands, cliCommand)
}`

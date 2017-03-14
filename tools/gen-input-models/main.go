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
	fmt.Fprintf(os.Stderr, "\tgen-input-models\n")
	os.Exit(2)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("gen-input-models: ")

	for key, resource := range ctx.ResourceDef {

		ctx.SetCurrentR(key)

		// schema validation
		errors := []error{}
		for _, comm := range resource.SortedCommands() {
			c := comm.Command
			errs := c.Validate()
			errors = append(errors, errs...)
		}
		if len(errors) > 0 {
			log.Println("*** Schema validation error ***")
			log.SetPrefix("\t")
			for _, e := range errors {
				log.Println(e.Error())
			}
			os.Exit(2)
		}

		src, err := generateResource(resource)
		if err != nil {
			log.Fatalf("writing output: %s", err)
		}

		// Write to file.
		baseName := ctx.InputModelFileName()
		outputName := filepath.Join(ctx.Gopath(), destination, strings.ToLower(baseName))

		err = ioutil.WriteFile(outputName, tools.Sformat([]byte(src)), 0644)
		if err != nil {
			log.Fatalf("writing output: %s", err)
		}
		fmt.Printf("generated: %s\n", filepath.Join(destination, strings.ToLower(baseName)))
	}
}

func generateResource(resource *schema.Resource) (string, error) {

	var commands []map[string]interface{}
	buf := bytes.NewBufferString("")
	t := template.New("t")
	template.Must(t.Parse(srcTemplate))

	// build commands
	for _, comm := range resource.SortedCommands() {
		c := comm.Command
		k := comm.CommandKey

		ctx.C = k

		params, err := buildCommandParams(c)
		if err != nil {
			return "", err
		}

		commands = append(commands, params)
	}

	err := t.Execute(buf, map[string]interface{}{
		"Commands": commands,
	})
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func buildCommandParams(command *schema.Command) (map[string]interface{}, error) {

	// add force flag if need
	if command.NeedConfirm {
		// has "force" param?
		if _, ok := command.Params["force"]; !ok {
			command.Params["force"] = &schema.Schema{
				Type:        schema.TypeBool,
				HandlerType: schema.HandlerNoop,
				Aliases:     []string{"f"},
			}
		}
	}

	var res map[string]interface{}

	fields, initializers, validators, err := buildFieldsParams(command.SortedParams())
	if err != nil {
		return res, err
	}

	res = map[string]interface{}{
		"Name":         ctx.InputModelTypeName(),
		"Fields":       fields,
		"Initializers": initializers,
		"Validators":   validators,
		"R":            ctx.R,
		"C":            ctx.C,
	}

	return res, err
}

func buildFieldsParams(params schema.SortableParams) ([]map[string]interface{}, []map[string]interface{}, string, error) {

	var fieldsRes []map[string]interface{}
	var initializerRes []map[string]interface{}
	var validatorRes string

	if len(params) == 0 {
		return fieldsRes, initializerRes, validatorRes, nil
	}

	t3 := template.New("t3")
	template.Must(t3.Parse(validatorTemplate))
	t4 := template.New("t4")
	template.Must(t4.Parse(conflictsWithTemplate))
	t5 := template.New("t5")
	template.Must(t5.Parse(betweenTemplate))

	validatorBuf := bytes.NewBufferString("")

	for _, param := range params {

		s := param.Param
		k := param.ParamKey

		ctx.P = k

		// to field
		ts, err := getFieldTypeString(s.Type)
		if err != nil {
			return fieldsRes, initializerRes, validatorRes, err
		}

		fieldsRes = append(fieldsRes, map[string]interface{}{
			"Name":      ctx.InputModelTypeName(),
			"ParamName": ctx.InputParamFieldName(),
			"TypeName":  ts,
			"Tag":       "", // TODO not yet implemented
		})

		// to initializer
		d := s.DefaultValue
		if d != nil {
			switch v := d.(type) {
			case string:
				d = fmt.Sprintf("\"%s\"", v)
			case []string:
				d = fmt.Sprintf("[]string {%s}", tools.FlattenStringList(v))
			case []int:
				d = fmt.Sprintf("[]int {%s}", tools.FlattenIntList(v))
			case []int64:
				d = fmt.Sprintf("[]int64 {%s}", tools.FlattenInt64List(v))
			case []uint:
				d = fmt.Sprintf("[]uint {%s}", tools.FlattenUintList(v))
			case []uint64:
				d = fmt.Sprintf("[]int64 {%s}", tools.FlattenUint64List(v))
			case []float32:
				d = fmt.Sprintf("[]float32 {%s}", tools.FlattenFloatList(v))
			case []float64:
				d = fmt.Sprintf("[]float64 {%s}", tools.FlattenFloat64List(v))

			}

		}
		initializerRes = append(initializerRes, map[string]interface{}{
			"Name":    ctx.InputParamFieldName(),
			"Default": d,
		})

		// to validator

		// required validator
		if s.Required {
			validatorName := "validateRequired"
			err = t3.Execute(validatorBuf, map[string]interface{}{
				"FlagName":     ctx.InputParamCLIFlagName(),
				"Name":         ctx.InputParamFieldName(),
				"ValidateFunc": validatorName,
			})
			if err != nil {
				return fieldsRes, initializerRes, validatorRes, err
			}
		}

		// custom validator
		if s.ValidateFunc != nil {
			validatorName := fmt.Sprintf(`define.Resources["%s"].Commands["%s"].Params["%s"].ValidateFunc`,
				ctx.R, ctx.C, ctx.P)
			err = t3.Execute(validatorBuf, map[string]interface{}{
				"FlagName":     ctx.InputParamCLIFlagName(),
				"Name":         ctx.InputParamFieldName(),
				"ValidateFunc": validatorName,
			})
			if err != nil {
				return fieldsRes, initializerRes, validatorRes, err
			}
		}

		// conflicts with
		if s.ConflictsWith != nil && len(s.ConflictsWith) > 0 {

			conflictsWith := map[string]interface{}{}

			for _, key := range s.ConflictsWith {
				k := tools.ToCLIFlagName(key)
				v := tools.ToCamelCaseName(key)
				conflictsWith[k] = v
			}

			err = t4.Execute(validatorBuf, map[string]interface{}{
				"FlagName":      ctx.InputParamCLIFlagName(),
				"Name":          ctx.InputParamFieldName(),
				"ConflictsWith": conflictsWith,
			})
			if err != nil {
				return fieldsRes, initializerRes, validatorRes, err
			}
		}

		// implements maxitem/minitem
		if s.MaxItems > 0 || s.MinItems > 0 {
			err = t5.Execute(validatorBuf, map[string]interface{}{
				"FlagName": ctx.InputParamCLIFlagName(),
				"Name":     ctx.InputParamFieldName(),
				"Min":      s.MinItems,
				"Max":      s.MaxItems,
			})
			if err != nil {
				return fieldsRes, initializerRes, validatorRes, err
			}
		}

	}

	validatorRes = validatorBuf.String()
	return fieldsRes, initializerRes, validatorRes, nil
}

var srcTemplate = `// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package command

import (
    "github.com/sacloud/usacloud/define"
    "github.com/sacloud/usacloud/schema"
    "github.com/sacloud/usacloud/output"
)

{{ range .Commands -}}

// {{.Name}} is input parameters for the sacloud API
type {{.Name}} struct {
    {{ range .Fields -}}
    {{.ParamName}} {{.TypeName}} {{.Tag}}
    {{ end }}
}

// New{{.Name}} return new {{.Name}}
func New{{.Name}}() *{{.Name}} {
	return &{{.Name}} {
	{{ range .Initializers -}}
	{{ if .Default }}
		{{.Name}}: {{.Default}},{{ end }}
	{{ end }}
	}
}

// Validate checks current values in model
func (p *{{.Name}}) Validate() []error{
    errors := []error{}
    {{.Validators}}
    return errors
}

func (p *{{.Name}}) getResourceDef() *schema.Resource {
    return define.Resources["{{.R}}"]
}

func (p *{{.Name}}) getCommandDef() *schema.Command {
    return p.getResourceDef().Commands["{{.C}}"]
}

func (p *{{.Name}}) GetIncludeFields() []string {
    return p.getCommandDef().IncludeFields
}

func (p *{{.Name}}) GetExcludeFields() []string {
    return p.getCommandDef().ExcludeFields
}

func (p *{{.Name}}) GetTableType() output.OutputTableType {
    return p.getCommandDef().TableType
}

func (p *{{.Name}}) GetColumnDefs() []output.ColumnDef {
    return p.getCommandDef().TableColumnDefines
}

{{ range .Fields -}}
func (p *{{.Name}}) Set{{.ParamName}}(v {{.TypeName}}) {
	p.{{.ParamName}} = v
}

func (p *{{.Name}}) Get{{.ParamName}}() {{.TypeName}} {
	return p.{{.ParamName}}
}
{{ end }}
{{- end }}
`

func getFieldTypeString(t schema.ValueType) (string, error) {
	switch t {
	case schema.TypeBool:
		return "bool", nil
	case schema.TypeInt:
		return "int", nil
	case schema.TypeInt64:
		return "int64", nil
	case schema.TypeFloat:
		return "float64", nil
	case schema.TypeString:
		return "string", nil
	case schema.TypeIntList:
		return "[]int64", nil
	case schema.TypeStringList:
		return "[]string", nil
	}

	return "", fmt.Errorf("Inalid type: %v", t)
}

var validatorTemplate = `{
    validator := {{.ValidateFunc}}
    errs := validator("{{.FlagName}}" , p.{{.Name}} )
    if errs != nil {
        errors = append(errors , errs...)
    }
}
`

var conflictsWithTemplate = `{
    errs := validateConflicts("{{.FlagName}}" , p.{{.Name}} , map[string]interface{}{
        {{range $k,$v := .ConflictsWith}}
        "{{$k}}": p.{{$v}},{{end}}
    })
    if errs != nil {
        errors = append(errors , errs...)
    }
}
`
var betweenTemplate = `{
    errs := validateBetween("{{.FlagName}}" , p.{{.Name}} , {{.Min}} , {{.Max}})
    if errs != nil {
        errors = append(errors , errs...)
    }
}
`

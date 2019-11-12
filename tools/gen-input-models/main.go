// Copyright 2017-2019 The Usacloud Authors
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
	destination = "src/github.com/sacloud/usacloud/command/params"
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

	var res map[string]interface{}

	fields, initializers, skeletonInitializers, validators, err := buildFieldsParams(command.BuildedParams())
	if err != nil {
		return res, err
	}

	res = map[string]interface{}{
		"Name":                 ctx.InputModelTypeName(),
		"Fields":               fields,
		"Initializers":         initializers,
		"SkeletonInitializers": skeletonInitializers,
		"Validators":           validators,
		"OutputExists":         !command.NoOutput,
		"R":                    ctx.R,
		"C":                    ctx.C,
	}

	return res, err
}

func buildFieldsParams(params schema.SortableParams) ([]map[string]interface{}, []map[string]interface{}, []map[string]interface{}, string, error) {

	var fieldsRes []map[string]interface{}
	var initializerRes []map[string]interface{}
	var skeletonInitializerRes []map[string]interface{}
	var validatorRes string

	if len(params) == 0 {
		return fieldsRes, initializerRes, skeletonInitializerRes, validatorRes, nil
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
			return fieldsRes, initializerRes, skeletonInitializerRes, validatorRes, err
		}

		fieldsRes = append(fieldsRes, map[string]interface{}{
			"Name":      ctx.InputModelTypeName(),
			"ParamName": ctx.InputParamFieldName(),
			"TypeName":  ts,
			"Tag":       fmt.Sprintf("`json:%q`", ctx.InputParamFlagName()),
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

		// to skeleton initializer
		skeleton, err := getSetEmptyStatement(s.Type)
		if err != nil {
			return fieldsRes, initializerRes, skeletonInitializerRes, validatorRes, err
		}
		skeletonInitializerRes = append(skeletonInitializerRes, map[string]interface{}{
			"Name":      ctx.InputParamFieldName(),
			"Statement": skeleton,
		})

		if s.Category == "output" {
			continue
		}

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
				return fieldsRes, initializerRes, skeletonInitializerRes, validatorRes, err
			}
		}

		// sakuraID(number-only,12 digit) validator
		if s.SakuraID {
			validatorName := "validateSakuraID"
			err = t3.Execute(validatorBuf, map[string]interface{}{
				"FlagName":     ctx.InputParamCLIFlagName(),
				"Name":         ctx.InputParamFieldName(),
				"ValidateFunc": validatorName,
			})
			if err != nil {
				return fieldsRes, initializerRes, skeletonInitializerRes, validatorRes, err
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
				return fieldsRes, initializerRes, skeletonInitializerRes, validatorRes, err
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
				return fieldsRes, initializerRes, skeletonInitializerRes, validatorRes, err
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
				return fieldsRes, initializerRes, skeletonInitializerRes, validatorRes, err
			}
		}

	}

	validatorRes = validatorBuf.String()
	return fieldsRes, initializerRes, skeletonInitializerRes, validatorRes, nil
}

var srcTemplate = `// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package params

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
	{{ range .Initializers -}}{{ if .Default }}
		{{.Name}}: {{.Default}},{{ end }}{{ end }}
	}
}

// FillValueToSkeleton fill values to empty fields
func (p *{{.Name}}) FillValueToSkeleton() {
	{{ range .SkeletonInitializers -}}
	if isEmpty(p.{{.Name}}){
		p.{{.Name}} = {{.Statement}}
	}
	{{ end }}
}

// Validate checks current values in model
func (p *{{.Name}}) Validate() []error{
	errors := []error{}
	{{.Validators}}
	{{ if .OutputExists }}
	{
		validator := schema.ValidateInStrValues(define.AllowOutputTypes...)
		errs := validator("--output-type" , p.OutputType )
		if errs != nil {
			errors = append(errors , errs...)
		}
	}
	{
		errs := validateInputOption(p)
		if errs != nil {
			errors = append(errors , errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors , errs...)
		}
	}
	{{ end }}
	return errors
}

func (p *{{.Name}}) GetResourceDef() *schema.Resource {
	return define.Resources["{{.R}}"]
}

func (p *{{.Name}}) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["{{.C}}"]
}

func (p *{{.Name}}) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *{{.Name}}) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *{{.Name}}) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *{{.Name}}) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
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

func getSetEmptyStatement(t schema.ValueType) (string, error) {
	switch t {
	case schema.TypeBool:
		return "false", nil
	case schema.TypeInt:
		return "0", nil
	case schema.TypeInt64:
		return "0", nil
	case schema.TypeFloat:
		return "0", nil
	case schema.TypeString:
		return `""`, nil
	case schema.TypeIntList:
		return "[]int64{0}", nil
	case schema.TypeStringList:
		return `[]string{""}`, nil
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

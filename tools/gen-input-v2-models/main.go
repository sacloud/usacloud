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

	"github.com/sacloud/libsacloud/sacloud"

	"github.com/sacloud/usacloud/schema"
	"github.com/sacloud/usacloud/tools"
)

var (
	destination = "src/github.com/sacloud/usacloud/cmdv2/params"
	ctx         = tools.NewGenerateContext()
)

// Usage is a replacement usage function for the flags package.
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\tgen-input-v2-models\n")
	os.Exit(2)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("gen-input-v2-models: ")

	for _, resource := range ctx.Resources {
		// schema validation
		var errors []error
		for _, c := range resource.Commands {
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
		baseName := resource.ParameterSourceFileName()
		filePath := filepath.Join(destination, strings.ToLower(baseName))
		fileFullPath := filepath.Join(ctx.Gopath(), filePath)

		err = ioutil.WriteFile(fileFullPath, tools.Sformat([]byte(src)), 0644)
		if err != nil {
			log.Fatalf("writing output: %s", err)
		}
		fmt.Printf("generated: %s\n", filePath)
	}
}

func generateResource(resource *tools.Resource) (string, error) {
	buf := bytes.NewBufferString("")
	t := template.New("t")
	template.Must(t.Parse(srcTemplate))
	err := t.Execute(buf, resource)
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
			case []sacloud.ID:
				d = fmt.Sprintf("[]sacloud.ID {%s}", tools.FlattenIDList(v))

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
	"io"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/pkg/utils"
	"github.com/sacloud/usacloud/pkg/validation"
)

{{ range .Commands -}}

// {{.InputParameterTypeName}} is input parameters for the sacloud API
type {{.InputParameterTypeName}} struct {
	{{ range .Params -}}
	{{.FieldName}} {{.FieldTypeName}} {{.FieldTag}}
	{{ end }}
	input Input
}

// New{{.InputParameterTypeName}} return new {{.InputParameterTypeName}}
func New{{.InputParameterTypeName}}() *{{.InputParameterTypeName}}{
	return &{{.InputParameterTypeName}} {
		{{ range .Params }}{{ if .DefaultValue }}{{.FieldName}}: {{.DefaultValueOnSource}},{{ end }}{{ end -}}
	}
}

// Initialize init {{.InputParameterTypeName}}
func (p *{{.InputParameterTypeName}}) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *{{.InputParameterTypeName}}) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *{{.InputParameterTypeName}}) fillValueToSkeleton() {
	{{ range .Params -}}
	if utils.IsEmpty(p.{{.FieldName}}){
		p.{{.FieldName}} = {{.SetEmptyStatement}}
	}
	{{ end }}
}

func (p *{{.InputParameterTypeName}}) validate() error {
	var errors []error
	{{ range .Params }}
	{{.ValidatorStatements}}
	{{ end }}
	return utils.FlattenErrors(errors)
}

func (p *{{.InputParameterTypeName}}) ResourceDef() *schema.Resource {
	return define.Resources["{{.Resource.Name}}"]
}

func (p *{{.InputParameterTypeName}}) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["{{.Name}}"]
}

func (p *{{.InputParameterTypeName}}) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *{{.InputParameterTypeName}}) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *{{.InputParameterTypeName}}) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *{{.InputParameterTypeName}}) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

{{ range .Params -}}
func (p *{{.Command.InputParameterTypeName}}) Set{{.FieldName}}(v {{.FieldTypeName}}) {
	p.{{.FieldName}} = v
}

func (p *{{.Command.InputParameterTypeName}}) Get{{.FieldName}}() {{.FieldTypeName}} {
	return p.{{.FieldName}}
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
	case schema.TypeId:
		return "sacloud.ID", nil
	case schema.TypeIdList:
		return "[]sacloud.ID", nil
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
	case schema.TypeId:
		return "sacloud.ID(0)", nil
	case schema.TypeIdList:
		return "[]sacloud.ID{}", nil
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
	errs := validation.ConflictsWith("{{.FlagName}}" , p.{{.Name}} , map[string]interface{}{
		{{range $k,$v := .ConflictsWith}}
		"{{$k}}": p.{{$v}},{{end}}
	})
	if errs != nil {
		errors = append(errors , errs...)
	}
}
`
var betweenTemplate = `{
	errs := validation.SliceLenBetween("{{.FlagName}}" , p.{{.Name}} , {{.Min}} , {{.Max}})
	if errs != nil {
		errors = append(errors , errs...)
	}
}
`

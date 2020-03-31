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

package tools

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/sacloud/libsacloud/sacloud"

	"github.com/sacloud/usacloud/schema"
)

// Parameter コード生成時に利用するパラメータ定義
type Parameter struct {
	*schema.Schema

	Name     string
	Category *schema.Category
	Command  *Command
}

func NewParameter(name string, param *schema.Schema, category *schema.Category, parent *Command) *Parameter {
	return &Parameter{
		Schema:   param,
		Name:     name,
		Category: category,
		Command:  parent,
	}
}

func (p *Parameter) FlagName() string {
	return ToDashedName(p.Name)
}

func (p *Parameter) FlagNameWithDash() string {
	return ToCLIFlagName(p.Name)
}

func (p *Parameter) FlagNameShorthands() string {
	var res []string
	for _, a := range p.Aliases {
		if len(a) == 1 {
			res = append(res, a)
		}
	}
	return strings.Join(res, ",")
}

func (p *Parameter) FlagDefinitionStatement() string {
	cliVariable := fmt.Sprintf("&%s.%s", p.Command.InputParameterVariable(), p.FieldName())
	name := p.FlagName()
	shorthands := p.FlagNameShorthands()
	value := p.DefaultValueOnSource()
	usage := p.Description

	statement := ""
	switch p.Type {
	case schema.TypeInvalid:
		panic("invalid parameter type")
	case schema.TypeBool:
		statement = `BoolVarP(%s, "%s", "%s", %s, "%s")`
	case schema.TypeInt:
		statement = `IntVarP(%s, "%s", "%s", %s, "%s")`
	case schema.TypeInt64:
		statement = `Int64VarP(%s, "%s", "%s", %s, "%s")`
	case schema.TypeFloat:
		statement = `Float64VarP(%s, "%s", "%s", %s, "%s")`
	case schema.TypeString:
		statement = `StringVarP(%s, "%s", "%s", %s, "%s")`
	case schema.TypeIntList:
		statement = `IntSliceVarP(%s, "%s", "%s", %s, "%s")`
	case schema.TypeStringList:
		statement = `StringSliceVarP(%s, "%s", "%s", %s, "%s")`
	case schema.TypeId:
		statement = `VarP(newIDValue(%s,%s), "%s", "%s", "%s")`
		return fmt.Sprintf(statement, value, cliVariable, name, shorthands, usage)
	case schema.TypeIdList:
		statement = `VarP(newIDSliceValue(%s,%s), "%s", "%s", "%s")`
		return fmt.Sprintf(statement, value, cliVariable, name, shorthands, usage)
	default:
		panic("unsupported type")
	}

	return fmt.Sprintf(statement, cliVariable, name, shorthands, value, usage)
}

func (p *Parameter) DefaultValueOnSource() string {
	switch p.Type {
	case schema.TypeInvalid:
		panic("invalid parameter type")
	case schema.TypeBool:
		res := `false`
		if v, ok := p.DefaultValue.(bool); ok {
			res = fmt.Sprintf("%t", v)
		}
		return res
	case schema.TypeInt:
		res := `0`
		if v, ok := p.DefaultValue.(int); ok {
			res = fmt.Sprintf("%d", v)
		}
		return res
	case schema.TypeInt64:
		res := `0`
		if v, ok := p.DefaultValue.(int64); ok {
			res = fmt.Sprintf("%d", v)
		}
		return res
	case schema.TypeFloat:
		res := `0`
		if v, ok := p.DefaultValue.(float64); ok {
			res = fmt.Sprintf("%f", v)
		}
		return res
	case schema.TypeString:
		res := `""`
		if v, ok := p.DefaultValue.(string); ok {
			res = fmt.Sprintf(`"%s"`, v)
		}
		return res
	case schema.TypeIntList:
		res := `[]int{}`
		if v, ok := p.DefaultValue.([]int); ok {
			strVal := ""
			for _, v := range v {
				strVal += fmt.Sprintf("%d,", v)
			}
			res = fmt.Sprintf(`[]int{%s}`, strVal)
		}
		return res
	case schema.TypeStringList:
		res := `[]string{}`
		if v, ok := p.DefaultValue.([]string); ok {
			strVal := ""
			for _, v := range v {
				strVal += fmt.Sprintf(`"%s",`, v)
			}
			res = fmt.Sprintf(`[]string{%s}`, strVal)
		}
		return res
	case schema.TypeId:
		res := `0`
		if v, ok := p.DefaultValue.(sacloud.ID); ok {
			res = fmt.Sprintf("sacloud.ID(%d)", v)
		}
		return res
	case schema.TypeIdList:
		res := `[]sacloud.ID{}`
		if v, ok := p.DefaultValue.([]sacloud.ID); ok {
			strVal := ""
			for _, v := range v {
				strVal += fmt.Sprintf("sacloud.ID(%d),", v)
			}
			res = fmt.Sprintf(`[]sacloud.ID{%s}`, strVal)
		}
		return res
	default:
		panic("unsupported type")
	}
}

func (p *Parameter) FieldName() string {
	return ToCamelCaseName(p.Name)
}

func (p *Parameter) FieldTypeName() string {
	switch p.Type {
	case schema.TypeBool:
		return "bool"
	case schema.TypeInt:
		return "int"
	case schema.TypeInt64:
		return "int64"
	case schema.TypeFloat:
		return "float64"
	case schema.TypeString:
		return "string"
	case schema.TypeIntList:
		return "[]int"
	case schema.TypeStringList:
		return "[]string"
	case schema.TypeId:
		return "sacloud.ID"
	case schema.TypeIdList:
		return "[]sacloud.ID"
	}
	panic("invalid type")
}

func (p *Parameter) FieldTag() string {
	return "" // TODO 当面不要
}

func (p *Parameter) SetEmptyStatement() string {
	switch p.Type {
	case schema.TypeBool:
		return "false"
	case schema.TypeInt:
		return "0"
	case schema.TypeInt64:
		return "0"
	case schema.TypeFloat:
		return "0"
	case schema.TypeString:
		return `""`
	case schema.TypeIntList:
		return "[]int{0}"
	case schema.TypeStringList:
		return `[]string{""}`
	case schema.TypeId:
		return "sacloud.ID(0)"
	case schema.TypeIdList:
		return "[]sacloud.ID{}"
	}

	panic("invalid type")
}

const commonValidatorTemplate = `{
	validator := {{.ValidateFunc}}
	errs := validator("{{.FlagName}}" , p.{{.Name}} )
	if errs != nil {
		errors = append(errors , errs...)
	}
}
`

var validateConflictsWithTemplate = `{
	errs := validation.ConflictsWith("{{.FlagName}}" , p.{{.Name}} , map[string]interface{}{
		{{range $k,$v := .ConflictsWith}}
		"{{$k}}": p.{{$v}},{{end}}
	})
	if errs != nil {
		errors = append(errors , errs...)
	}
}
`
var validateBetweenTemplate = `{
	errs := validation.SliceLenBetween("{{.FlagName}}" , p.{{.Name}} , {{.Min}} , {{.Max}})
	if errs != nil {
		errors = append(errors , errs...)
	}
}
`

func (p *Parameter) ValidatorStatements() string {
	var err error
	validatorBuf := bytes.NewBufferString("")

	commonTmpl := template.New("t1")
	template.Must(commonTmpl.Parse(commonValidatorTemplate))
	conflictsWithTmpl := template.New("t2")
	template.Must(conflictsWithTmpl.Parse(validateConflictsWithTemplate))
	betweenTmpl := template.New("t3")
	template.Must(betweenTmpl.Parse(validateBetweenTemplate))

	// required validator
	if p.Required {
		validatorName := "validateRequired"
		err = commonTmpl.Execute(validatorBuf, map[string]interface{}{
			"FlagName":     p.FlagNameWithDash(),
			"Name":         p.FieldName(),
			"ValidateFunc": validatorName,
		})
		if err != nil {
			panic(err)
		}
	}

	// sakuraID(number-only,12 digit) validator
	if p.SakuraID {
		validatorName := "validateSakuraID"
		err = commonTmpl.Execute(validatorBuf, map[string]interface{}{
			"FlagName":     p.FlagNameWithDash(),
			"Name":         p.FieldName(),
			"ValidateFunc": validatorName,
		})
		if err != nil {
			panic(err)
		}
	}

	// custom validator
	if p.ValidateFunc != nil {
		validatorName := fmt.Sprintf(`define.Resources["%s"].Commands["%s"].Params["%s"].ValidateFunc`,
			p.Command.Resource.Name, p.Command.Name, p.Name)
		err = commonTmpl.Execute(validatorBuf, map[string]interface{}{
			"FlagName":     p.FlagNameWithDash(),
			"Name":         p.FieldName(),
			"ValidateFunc": validatorName,
		})
		if err != nil {
			panic(err)
		}
	}

	// conflicts with
	if p.ConflictsWith != nil && len(p.ConflictsWith) > 0 {
		conflictsWith := map[string]interface{}{}

		for _, key := range p.ConflictsWith {
			k := ToCLIFlagName(key)
			v := ToCamelCaseName(key)
			conflictsWith[k] = v
		}

		err = conflictsWithTmpl.Execute(validatorBuf, map[string]interface{}{
			"FlagName":      p.FlagNameWithDash(),
			"Name":          p.FieldName(),
			"ConflictsWith": conflictsWith,
		})
		if err != nil {
			panic(err)
		}
	}

	// implements maxitem/minitem
	if p.MaxItems > 0 || p.MinItems > 0 {
		err = betweenTmpl.Execute(validatorBuf, map[string]interface{}{
			"FlagName": p.FlagNameWithDash(),
			"Name":     p.FieldName(),
			"Min":      p.MinItems,
			"Max":      p.MaxItems,
		})
		if err != nil {
			panic(err)
		}
	}

	return validatorBuf.String()
}

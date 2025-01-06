// Copyright 2017-2025 The sacloud/usacloud Authors
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
	"fmt"
	"reflect"
	"strings"

	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/usacloud/pkg/core"
)

type Command struct {
	*core.Command
	Resource *Resource
}

func NewCommand(r *Resource, c *core.Command) *Command {
	return &Command{Command: c, Resource: r}
}

func (c *Command) PackageDirName() string {
	return c.Resource.PackageDirName()
}

func (c *Command) PackageName() string {
	return c.Resource.PackageName()
}

func (c *Command) HasAliases() bool {
	if c.Command.ParameterInitializer == nil {
		return false
	}
	for _, f := range c.Fields() {
		if len(f.Aliases) > 0 {
			return true
		}
	}
	return false
}

func (c *Command) Fields() []Field {
	if c.Command.ParameterInitializer == nil {
		return nil
	}
	clitagFields, err := clitagParser.Parse(c.Command.ParameterInitializer())
	if err != nil {
		panic(err)
	}
	var fields []Field
	for _, f := range clitagFields {
		validateTag := f.StructField.Tag.Get("validate")
		fields = append(fields, Field{
			StructField: f,
			Required:    strings.Contains(validateTag, "required") && !strings.Contains(validateTag, "required_with"), // require_with/withoutは無視
		})
	}
	return fields
}

func (c *Command) CLIFlagDefinitionStatements(parameterVariableName, flagSetVariableName string) string {
	if c.Command.ParameterInitializer == nil {
		return ""
	}

	var statements []string
	for _, f := range c.Fields() {
		s := c.cliFlagDefinitionStatement(parameterVariableName, f)
		if s != "" {
			statements = append(statements, fmt.Sprintf("%s.%s", flagSetVariableName, s))
		}
	}
	return strings.Join(statements, "\n")
}

func (c *Command) cliFlagDefinitionStatement(parameterVariableName string, field Field) string {
	fieldVar := fmt.Sprintf("%s.%s", parameterVariableName, field.FieldName)
	fieldPointerVar := fieldVar
	fieldType := dereferencePtrType(field.Type)
	if field.Type.Kind() == reflect.Ptr {
		switch fieldType.Kind() {
		case reflect.Bool:
			fieldVar = "false"
		case reflect.Int, reflect.Int64, reflect.Float64:
			fieldVar = "0"
		case reflect.String:
			fieldVar = `""`
		case reflect.Slice:
			fieldVar = "nil"
		default:
			panic(fmt.Sprintf("unsupported type: field: %s, type: %s", field.FieldName, fieldType.Kind().String()))
		}
	} else {
		fieldPointerVar = "&" + fieldPointerVar
	}

	name := field.FlagName
	shorthands := field.Shorthand
	// value := p.DefaultValueOnSource()
	usage := field.LongDescription()

	statement := ""
	if isLibsacloudIDType(fieldType) {
		statement = `VarP(core.NewIDFlag(%s, %s), "%s", "%s", "%s")`
		return fmt.Sprintf(statement, fieldPointerVar, fieldPointerVar, name, shorthands, usage)
	}
	switch fieldType.Kind() {
	case reflect.Bool:
		statement = `BoolVarP(%s, "%s", "%s", %s, "%s")`
	case reflect.Int:
		statement = `IntVarP(%s, "%s", "%s", %s, "%s")`
	case reflect.Int64:
		statement = `Int64VarP(%s, "%s", "%s", %s, "%s")`
	case reflect.Float64:
		statement = `Float64VarP(%s, "%s", "%s", %s, "%s")`
	case reflect.String:
		statement = `StringVarP(%s, "%s", "%s", %s, "%s")`
	case reflect.Slice:
		if isLibsacloudIDType(fieldType.Elem()) {
			statement = `VarP(core.NewIDSliceFlag(%s, %s), "%s", "%s", "%s")`
			return fmt.Sprintf(statement, fieldPointerVar, fieldPointerVar, name, shorthands, usage)
		}
		switch fieldType.Elem().Kind() {
		case reflect.Int64:
			statement = `Int64SliceVarP(%s, "%s", "%s", %s, "%s")`
		case reflect.String:
			statement = `StringSliceVarP(%s, "%s", "%s", %s, "%s")`
		default:
			panic(fmt.Sprintf("unsupported type: field: %s, type: []%s", field.FieldName, fieldType.Elem().Kind().String()))
		}
	default:
		panic(fmt.Sprintf("unsupported type: field: %s, type: %s", field.FieldName, fieldType.Kind().String()))
	}

	return fmt.Sprintf(statement, fieldPointerVar, name, shorthands, fieldVar, usage)
}

func (c *Command) CLIFlagInitializePointerStatement(parameterVariableName, flagSetVariableName string) string {
	if c.Command.ParameterInitializer == nil {
		return ""
	}

	var statements []string
	for _, f := range c.Fields() {
		s := c.cliFlagInitializePointerStatement(parameterVariableName, f)
		if s != "" {
			statements = append(statements, s)
		}
	}
	return strings.Join(statements, "\n")
}

func (c *Command) cliFlagInitializePointerStatement(parameterVariableName string, field Field) string {
	fieldVar := fmt.Sprintf("%s.%s", parameterVariableName, field.FieldName)
	fieldType := dereferencePtrType(field.Type)
	if field.Type.Kind() != reflect.Ptr {
		return ""
	}

	srcTemplate := `if %s == nil {
	%s = %s
}`

	var statement string
	if isLibsacloudIDType(fieldType) {
		srcTemplate = `if %s == nil {
	v := types.ID(0)
	%s = &v
}`
		return fmt.Sprintf(srcTemplate, fieldVar, fieldVar)
	}

	switch fieldType.Kind() {
	case reflect.Bool:
		statement = "pointer.NewBool(false)"
	case reflect.Int:
		statement = "pointer.NewInt(0)"
	case reflect.Int64:
		statement = "pointer.NewInt64(int64(0))"
	case reflect.Float64:
		statement = "pointer.NewFloat64(float64(0))"
	case reflect.String:
		statement = `pointer.NewString("")`
	case reflect.Slice:
		if isLibsacloudIDType(fieldType.Elem()) {
			srcTemplate = `if %s == nil {
	v := []types.ID{}
	%s = &v
}`
			return fmt.Sprintf(srcTemplate, fieldVar, fieldVar)
		}
		switch fieldType.Elem().Kind() {
		case reflect.Int64:
			statement = "pointer.NewInt64Slice([]int64{})"
		case reflect.String:
			statement = "pointer.NewStringSlice([]string{})"
		default:
			panic(fmt.Sprintf("unsupported type: field: %s, type: []%s", field.FieldName, fieldType.Elem().Kind().String()))
		}
	default:
		panic(fmt.Sprintf("unsupported type: field: %s, type: %s", field.FieldName, fieldType.Kind().String()))
	}

	return fmt.Sprintf(srcTemplate, fieldVar, fieldVar, statement)
}

func (c *Command) CLIFlagCleanupEmptyStatement(parameterVariableName, flagSetVariableName string) string {
	if c.Command.ParameterInitializer == nil {
		return ""
	}

	var statements []string
	for _, f := range c.Fields() {
		s := c.cliFlagCleanupEmptyStatement(parameterVariableName, flagSetVariableName, f)
		if s != "" {
			statements = append(statements, s)
		}
	}
	return strings.Join(statements, "\n")
}

func (c *Command) cliFlagCleanupEmptyStatement(parameterVariableName, flagSetVariableName string, field Field) string {
	fieldVar := fmt.Sprintf("%s.%s", parameterVariableName, field.FieldName)
	if field.Type.Kind() != reflect.Ptr {
		return ""
	}

	srcTemplate := `if !%s.Changed("%s") {
	%s = nil
}`

	return fmt.Sprintf(srcTemplate, flagSetVariableName, field.FlagName, fieldVar)
}

func isLibsacloudIDType(t reflect.Type) bool {
	return reflect.TypeOf(types.ID(0)) == t
}

func dereferencePtrType(t reflect.Type) reflect.Type {
	if t.Kind() != reflect.Ptr {
		return t
	}
	return dereferencePtrType(t.Elem())
}

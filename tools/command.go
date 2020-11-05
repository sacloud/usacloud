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
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/fatih/structs"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/schema"
	"github.com/sacloud/usacloud/pkg/util"
	"github.com/sacloud/usacloud/tools/clitag"
	"github.com/sacloud/usacloud/tools/utils"
)

// Command コード生成時に利用するコマンド定義
type Command struct {
	*schema.Command

	Name              string
	Resource          *Resource
	Category          *schema.Category
	Params            []*Parameter
	OwnParams         []*Parameter
	CategorizedParams []*CategorizedParameters
}

type CategorizedParameters struct {
	*schema.Category
	Params []*Parameter
}

type CategorizedParameterFields struct {
	*schema.Category
	Fields []clitag.StructField
}

func NewCommand(name string, command *schema.Command, category *schema.Category, resource *Resource) *Command {
	c := &Command{
		Command:  command,
		Name:     name,
		Resource: resource,
		Category: category,
	}

	var params []*Parameter
	for _, p := range c.Command.BuiltParams() {
		params = append(params, NewParameter(p.ParamKey, p.Param, p.Category, c))
	}
	c.Params = params

	var ownParams []*Parameter
	for k, p := range c.Command.Params {
		if p.HandlerType != schema.HandlerNoop {
			ownParams = append(ownParams, NewParameter(k, p, c.Command.ParamCategory(p.Category), c))
		}
	}
	sort.Slice(ownParams, func(i, j int) bool {
		ti := ownParams[i]
		tj := ownParams[j]
		if ti.Order == tj.Order {
			return ti.Name < tj.Name
		}
		return ti.Order < tj.Order
	})
	c.OwnParams = ownParams

	c.buildCategorizedParams()
	return c
}

func (c *Command) buildCategorizedParams() {
	m := map[string]*CategorizedParameters{}
	for _, p := range c.Params {
		c := p.Category
		cp, ok := m[c.Key]
		if !ok {
			cp = &CategorizedParameters{
				Category: c,
			}
		}
		cp.Params = append(cp.Params, p)
		m[c.Key] = cp
	}
	c.CategorizedParams = []*CategorizedParameters{}
	for _, cat := range m {
		c.CategorizedParams = append(c.CategorizedParams, cat)
	}
	sort.Slice(c.CategorizedParams, func(i, j int) bool {
		if c.CategorizedParams[i].Order == c.CategorizedParams[j].Order {
			return c.CategorizedParams[i].Key < c.CategorizedParams[j].Key
		}
		return c.CategorizedParams[i].Order < c.CategorizedParams[j].Order
	})
}

func (c *Command) IsGlobal() bool {
	return c.Resource.IsGlobal
}

func (c *Command) ExperimentWarning() string {
	if c.Command.ExperimentWarning != "" {
		return c.Command.ExperimentWarning
	}
	return c.Resource.ExperimentWarning
}

func (c *Command) Usage() string {
	usage := c.Command.Usage
	if usage == "" {
		usage = fmt.Sprintf("%s %s", utils.ToCamelCaseName(c.Name), utils.ToCamelCaseName(c.Resource.Name))
	}
	if c.Resource.DefaultCommand == c.Name {
		usage = fmt.Sprintf("%s (default)", usage)
	}
	return usage
}

func (c *Command) ArgsUsage() string {
	argsUsage := c.Command.ArgsUsage
	if argsUsage == "" && c.Type.IsRequiredIDType() {
		t := c.Type
		switch {
		case t.IsNeedIDOnlyType():
			argsUsage = "<ID>"
		case t.IsNeedSingleIDType():
			argsUsage = "<ID or Name(only single target)>"
		default:
			argsUsage = "<ID or Name(allow multiple target)>"
		}
	}
	return argsUsage
}

func (c *Command) AliasesLiteral() string {
	return utils.FlattenStringList(c.Command.Aliases)
}

func (c *Command) HasIDParam() bool {
	for _, p := range c.Params {
		if p.Name == "id" && p.Type == schema.TypeId {
			return true
		}
	}
	return false
}

func (c *Command) HasOutputOption() bool {
	return !c.NoOutput
}

func (c *Command) CLIName() string {
	return utils.ToDashedName(c.Name)
}

func (c *Command) CLIVariableFuncName() string {
	return fmt.Sprintf("%s%sCmd", utils.ToCamelWithFirstLower(c.Resource.Name), utils.ToCamelCaseName(c.Name))
}

func (c *Command) CLIv2CommandsFileName() string {
	return fmt.Sprintf("%s_gen.go", utils.ToSnakeCaseName(c.Resource.Name))
}

func (c *Command) CLINormalizeFlagsFuncName() string {
	return fmt.Sprintf("%s%sNormalizeFlagNames", utils.ToCamelWithFirstLower(c.Resource.Name), utils.ToCamelCaseName(c.Name))
}

func (c *Command) InputParameterVariable() string {
	return fmt.Sprintf("%s%sParam", utils.ToCamelWithFirstLower(c.Resource.Name), utils.ToCamelCaseName(c.Name))
}

func (c *Command) InputParameterTypeName() string {
	return fmt.Sprintf("%s%sParam", utils.ToCamelCaseName(c.Name), utils.ToCamelCaseName(c.Resource.Name))
}

func (c *Command) FunctionName() string {
	return utils.ToCamelCaseName(c.Name)
}

func (c *Command) NeedConfirm() bool {
	return c.Type.IsNeedConfirmType() && !c.NeedlessConfirm
}

func (c *Command) ConfirmMessage() string {
	if c.Command.ConfirmMessage == "" {
		return utils.ToDashedName(c.Name)
	}
	return c.Command.ConfirmMessage
}

func (c *Command) RequireID() bool {
	return c.Command.Type.IsRequiredIDType()
}

func (c *Command) SingleArgToIdParam() bool {
	return c.Command.Type.IsNeedIDOnlyType()
}

func (c *Command) MultipleArgToIdParams() bool {
	return c.RequireID() && !c.SingleArgToIdParam()
}

func (c *Command) ArgToIdFunc() string {
	return fmt.Sprintf("find%s%sTargets", utils.ToCamelCaseName(c.Resource.Name), utils.ToCamelCaseName(c.Name))
}

func (c *Command) FlagOrderFunc() string {
	return fmt.Sprintf("%s%sFlagOrder", utils.ToCamelWithFirstLower(c.Resource.Name), utils.ToCamelCaseName(c.Name))
}

func (c *Command) TargetAPIName() string {
	return util.FirstNonEmptyString(c.AltResource, c.Resource.AltResource, utils.ToCamelCaseName(c.Resource.Name))
}

func (c *Command) FindResultFieldName() string {
	return util.FirstNonEmptyString(c.ListResultFieldName, c.Resource.ListResultFieldName, utils.ToCamelCaseName(c.Resource.Name)+"s")
}

func (c *Command) RequireSingleID() bool {
	return c.Type.IsNeedSingleIDType()
}

func (c *Command) HasLongAliases() bool {
	for _, p := range c.Params {
		aliases := p.LongAliases()
		if len(aliases) > 0 {
			return true
		}
	}
	return false
}

func (c *Command) CommandFileName() string {
	format := "zz_%s_gen.go"
	if c.UseCustomCommand {
		format = "%s.go"
	}
	return fmt.Sprintf(format, utils.ToSnakeCaseName(c.Name))
}

func (c *Command) ResourceName() string {
	return util.FirstNonEmptyString(c.AltResource, c.Resource.AltResource, utils.ToCamelCaseName(c.Resource.Name))
}

func (c *Command) InputModelTypeName() string {
	return fmt.Sprintf("%s%sParam", utils.ToCamelCaseName(c.Name), utils.ToCamelCaseName(c.Resource.Name))
}

func (c *Command) APIRequestTypeName() string {
	switch c.Type {
	case schema.CommandList:
		return "FindCondition"
	default:
		return fmt.Sprintf("%s%sRequest", c.ResourceName(), utils.ToCamelCaseName(c.Name))
	}
}

func (c *Command) PackageDirName() string {
	return c.Resource.PackageDirName()
}

/*
  TODO gen-cli-command-v1用、あとでschema.Commandを整理する際に移動するかも
*/

func (c *Command) CLICommandGeneratedSourceFile() string {
	return fmt.Sprintf("%s_gen.go", utils.ToSnakeCaseName(c.Name))
}

func (c *Command) CLICommandParameterTypeName() string {
	if c.Command.Parameters == nil {
		return ""
	}
	return structs.Name(c.Command.Parameters)
}

func (c *Command) ServiceRequestTypeName() string {
	name := c.Name
	if c.ServiceFuncAltName != "" {
		name = c.ServiceFuncAltName
	}
	return fmt.Sprintf("%sRequest", utils.ToCamelCaseName(name))
}

func (c *Command) ServiceFuncName() string {
	name := c.Name
	if c.ServiceFuncAltName != "" {
		name = c.ServiceFuncAltName
	}
	return fmt.Sprintf("%sWithContext", utils.ToCamelCaseName(name))
}

func (c *Command) ServiceCommandFuncName() string {
	return fmt.Sprintf("run%sService", utils.ToCamelCaseName(c.Name))
}

func (c *Command) CategorizedParameterFields() []*CategorizedParameterFields {
	if c.Parameters == nil {
		return nil
	}

	m := map[string]*CategorizedParameterFields{}
	for _, f := range c.Fields() {
		cp, ok := m[f.Category]
		if !ok {
			cp = &CategorizedParameterFields{
				Category: c.Command.ParamCategory(f.Category),
			}
		}
		cp.Fields = append(cp.Fields, f)
		m[f.Category] = cp
	}
	var categorizedFields []*CategorizedParameterFields
	for _, cat := range m {
		categorizedFields = append(categorizedFields, cat)
	}
	sort.Slice(categorizedFields, func(i, j int) bool {
		if categorizedFields[i].Order == categorizedFields[j].Order {
			return categorizedFields[i].Key < categorizedFields[j].Key
		}
		return categorizedFields[i].Order < categorizedFields[j].Order
	})

	return categorizedFields
}

func (c *Command) HasAliases() bool {
	if c.Command.Parameters == nil {
		return false
	}
	for _, f := range c.Fields() {
		if len(f.Aliases) > 0 {
			return true
		}
	}
	return false
}

func (c *Command) Fields() []clitag.StructField {
	fields, err := clitag.Parse(c.Command.Parameters)
	if err != nil {
		panic(err)
	}
	return fields
}

func (c *Command) CLIFlagDefinitionStatements(parameterVariableName, flagSetVariableName string) string {
	if c.Command.Parameters == nil {
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

func (c *Command) cliFlagDefinitionStatement(parameterVariableName string, field clitag.StructField) string {
	fieldVar := fmt.Sprintf("%s.%s", parameterVariableName, field.Name)
	fieldPointerVar := fieldVar
	if field.Type.Kind() == reflect.Ptr {
		fieldVar = "*" + fieldVar
	} else {
		fieldPointerVar = "&" + fieldPointerVar
	}

	name := field.FlagName
	shorthands := field.Shorthand
	//value := p.DefaultValueOnSource()
	usage := field.Description
	if len(field.Aliases) > 0 {
		usage = fmt.Sprintf("%s (aliases: %s)", usage, strings.Join(c.Aliases, ", "))
	}

	statement := ""
	fieldType := dereferencePtrType(field.Type)
	if isLibsacloudIDType(fieldType) {
		statement = `VarP(base.NewIDFlag(%s, %s), "%s", "%s", "%s")`
		return fmt.Sprintf(statement, fieldPointerVar, fieldPointerVar, name, shorthands, usage)
	} else {
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
				statement = `VarP(base.NewIDSliceFlag(%s, %s), "%s", "%s", "%s")`
				return fmt.Sprintf(statement, fieldPointerVar, fieldPointerVar, name, shorthands, usage)
			} else {
				switch fieldType.Elem().Kind() {
				case reflect.Int64:
					statement = `Int64SliceVarP(%s, "%s", "%s", %s, "%s")`
				case reflect.String:
					statement = `StringSliceVarP(%s, "%s", "%s", %s, "%s")`
				default:
					panic(fmt.Sprintf("unsupported type: field: %s, type: []%s", field.Name, fieldType.Elem().Kind().String()))
				}
			}
		default:
			panic(fmt.Sprintf("unsupported type: field: %s, type: %s", field.Name, fieldType.Kind().String()))
		}
	}

	return fmt.Sprintf(statement, fieldPointerVar, name, shorthands, fieldVar, usage)
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

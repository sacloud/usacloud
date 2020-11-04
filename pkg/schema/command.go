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

package schema

import (
	"fmt"
	"sort"
	"strings"

	"github.com/sacloud/usacloud/pkg/output"
)

type Command struct {
	Type      CommandType
	Aliases   []string
	Usage     string
	ArgsUsage string

	Category string
	Order    int

	// TODO v0向け、あとで消す
	Params          map[string]*Schema
	ParamCategories []Category

	// v1向け
	Parameters interface{} // cmd/xxx配下の各コマンドパラメータstruct

	AltResource         string // 空の場合はResourceのキーをCamelizeしてsacloud.XXXを対象とする。
	ListResultFieldName string
	SkipAuth            bool
	SkipAfterSecondArgs bool // trueの場合、2番目以降の引数(1番目はID or 名称)の解析を行わない
	UseCustomCommand    bool
	NeedlessConfirm     bool
	ConfirmMessage      string
	ExperimentWarning   string

	NoOutput bool

	NoSelector bool // --selectorオプション利用有無、trueの場合利用しない

	TableType          output.TableType
	IncludeFields      []string           // for output.TableDetail
	ExcludeFields      []string           // for output.TableDetail
	TableColumnDefines []output.ColumnDef // for output.TableSimple
}

func (c *Command) ParamCategory(key string) *Category {
	switch key {
	case "":
		return DefaultParamCategory
	case "output":
		return OutputParamCategory
	case "input":
		return InputParamCategory
	case "common":
		return CommonParamCategory
	case "sort":
		return SortParamCategory
	case "limit-offset":
		return LimitOffsetParamCategory
	case "filter":
		return FilterParamCategory
	default:
		if len(c.ParamCategories) == 0 {
			return &Category{
				Key:         key,
				DisplayName: fmt.Sprintf("%s options", strings.Title(key)),
				Order:       1,
			}
		}

		for _, cat := range c.ParamCategories {
			if cat.Key == key {
				return &cat
			}
		}
		return nil
	}
}

func (c *Command) BuiltParams() SortableParams {
	// Notice: ここで追加されるパラメータはdefine.Resourcesからは見えない。
	//         (コード生成時に追加されるため)
	//         このため、ランタイムでdefine.Resourcesを参照する必要のある
	//         ValidatorやConflictsWithは利用できない。
	//         Validatorを利用したい場合はコード生成時に手動で呼び出すコードを出力する。
	//         例: command.validateOutputOption(o output.Option)の呼び出し部分など

	params := make(map[string]*Schema)
	for k, v := range c.Params {
		params[k] = v
	}

	// add ID param
	if c.Type.IsRequiredIDType() {
		params["id"] = &Schema{
			Type:        TypeId,
			HandlerType: HandlerPathThrough,
			Description: "Set target ID",
			SakuraID:    true,
			Hidden:      true,
		}
	}
	if c.Type.CanUseSelector() && !c.NoSelector {
		params["selector"] = &Schema{
			Type:        TypeStringList,
			HandlerType: HandlerNoop,
			Description: "Set target filter by tag",
			Category:    "filter",
			Order:       10,
		}
	}

	if c.Type.IsNeedConfirmType() && !c.NeedlessConfirm {
		params["assumeyes"] = &Schema{
			Type:        TypeBool,
			HandlerType: HandlerNoop,
			Description: "Assume that the answer to any question which would be asked is yes",
			Category:    "input",
			Order:       10,
			Aliases:     []string{"y"},
		}
	}

	params["parameters"] = &Schema{
		Type:        TypeString,
		HandlerType: HandlerNoop,
		Description: "Set input parameters from JSON string",
		Category:    "input",
		Order:       21,
	}
	params["parameter-file"] = &Schema{
		Type:        TypeString,
		HandlerType: HandlerNoop,
		Description: "Set input parameters from file",
		Category:    "input",
		Order:       31,
	}
	params["generate-skeleton"] = &Schema{
		Type:        TypeBool,
		HandlerType: HandlerNoop,
		Description: "Output skelton of parameter JSON",
		Category:    "input",
		Order:       40,
	}

	if !c.NoOutput {
		params["output-type"] = &Schema{
			Type:        TypeString,
			HandlerType: HandlerNoop,
			Aliases:     []string{"out", "o"},
			Description: "Output type [table/json/csv/tsv]",
			Category:    "output",
			Order:       10,
		}
		params["column"] = &Schema{
			Type:        TypeStringList,
			HandlerType: HandlerNoop,
			Aliases:     []string{"col"},
			Description: "Output columns(using when '--output-type' is in [csv/tsv] only)",
			Category:    "output",
			Order:       20,
		}
		params["quiet"] = &Schema{
			Type:        TypeBool,
			HandlerType: HandlerNoop,
			Aliases:     []string{"q"},
			Description: "Only display IDs",
			Category:    "output",
			Order:       30,
		}
		params["format"] = &Schema{
			Type:        TypeString,
			HandlerType: HandlerNoop,
			Aliases:     []string{"fmt"},
			Description: "Output format(see text/template package document for detail)",
			Category:    "output",
			Order:       40,
		}
		params["format-file"] = &Schema{
			Type:        TypeString,
			HandlerType: HandlerNoop,
			Description: "Output format from file(see text/template package document for detail)",
			Category:    "output",
			Order:       50,
		}
		params["query"] = &Schema{
			Type:        TypeString,
			HandlerType: HandlerNoop,
			Description: "JMESPath query(using when '--output-type' is json only)",
			Category:    "output",
			Order:       60,
		}
		params["query-file"] = &Schema{
			Type:        TypeString,
			HandlerType: HandlerNoop,
			Description: "JMESPath query from file(using when '--output-type' is json only)",
			Category:    "output",
			Order:       65,
		}
	}

	res := SortableParams{}
	for k, v := range params {
		res = append(res, SortableParam{
			ParamKey: k,
			Param:    v,
			Category: c.ParamCategory(v.Category),
		})
	}

	sort.Sort(res)
	return res
}

func (c *Command) Validate() []error {
	var errors []error

	if c.Type == CommandInvalid {
		errors = append(errors, fmt.Errorf("command#Type: command type is invalid: (%#v)", c))
	}

	if c.Type == CommandList && c.TableType != output.TableSimple {
		errors = append(errors, fmt.Errorf("command#TableType: need output.TableSimple when Command#Type is CommandList"))
	}

	if c.TableType == output.TableSimple && len(c.TableColumnDefines) == 0 {
		errors = append(errors, fmt.Errorf("command#TableColumnDefines: required when Command#TableType is output.TableSimple"))
	}

	if c.NoSelector && !c.Type.CanUseSelector() {
		errors = append(errors, fmt.Errorf("command#NoSelector: NoSelector isnot used with CommandType[%s]", c.Type.String()))
	}

	if len(c.ParamCategories) > 0 && len(c.Params) > 0 {
		for _, category := range c.ParamCategories {
			exists := false
			for _, param := range c.Params {
				if param.Category == category.Key {
					exists = true
					break
				}
			}
			if !exists {
				err := fmt.Errorf("command#Categories: category(%s) isnot used by any Params", category.Key)
				errors = append(errors, err)
			}
		}
	}

	for k, v := range c.Params {
		errs := v.Validate(k)
		errors = append(errors, errs...)

		if len(v.ConflictsWith) > 0 {
			for _, s := range v.ConflictsWith {
				if _, ok := c.Params[s]; !ok {
					err := fmt.Errorf("command#%s.%q: not exists in Command#Params(%s)", k, "ConflictsWith", s)
					errors = append(errors, err)
				}
			}
		}

		// combination of command.Type and params.HandlerType is valid?
		if v.HandlerType.IsWhenListOnly() && c.Type != CommandList {
			err := fmt.Errorf("command#%s.%q: cannot use HandlerType(%v) when Command#Type is CommandList", k, "HandlerType", v.HandlerType)
			errors = append(errors, err)
		}

		//if v.Category != "" && v.Category != "output" && v.Category != "input" && len(c.ParamCategories) > 0 {
		//	exists := false
		//	// category is defined on command?
		//	for _, category := range c.ParamCategories {
		//		if category.Key == v.Category {
		//			exists = true
		//			break
		//		}
		//	}
		//	if !exists {
		//		err := fmt.Errorf("command#%s.%q: category(%s) isn't defined, but is used by %q", k, "Category", v.Category, k)
		//		errors = append(errors, err)
		//	}
		//}
	}

	return errors
}

type SortableParam struct {
	Category *Category
	Param    *Schema
	ParamKey string
}
type SortableParams []SortableParam

func (s SortableParams) Len() int {
	return len(s)
}

func (s SortableParams) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortableParams) Less(i, j int) bool {
	if s[i].Category.Order == s[j].Category.Order {
		if s[i].Param.Order == s[j].Param.Order {
			return s[i].ParamKey < s[j].ParamKey
		}
		return s[i].Param.Order < s[j].Param.Order
	}
	return s[i].Category.Order < s[j].Category.Order
}

func (s SortableParams) Get(key string) *SortableParam {
	for _, param := range s {
		if param.ParamKey == key {
			return &param
		}
	}
	return nil
}

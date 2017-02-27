package schema

import (
	"fmt"
	"github.com/sacloud/usacloud/output"
	"sort"
)

type Command struct {
	Type      CommandType
	Aliases   []string
	Usage     string
	ArgsUsage string

	Category string
	Order    int

	Params              map[string]*Schema
	ParamCategories     []Category
	AltResource         string // 空の場合はResourceのキーをCamelizeしてsacloud.XXXを対象とする。
	ListResultFieldName string
	SkipAuth            bool
	UseCustomCommand    bool

	TableType          output.OutputTableType
	IncludeFields      []string           // for output.TableDetail
	ExcludeFields      []string           // for output.TableDetail
	TableColumnDefines []output.ColumnDef // for output.TableSimple
}

func (c *Command) ParamCategory(key string) *Category {

	if key == "" {
		return DefaultParamCategory
	}

	for _, cat := range c.ParamCategories {
		if cat.Key == key {
			return &cat
		}
	}

	return nil
}

func (c *Command) SortedParams() SortableParams {

	params := SortableParams{}
	for k, v := range c.Params {
		params = append(params, SortableParam{
			ParamKey: k,
			Param:    v,
			Category: c.ParamCategory(v.Category),
		})
	}

	sort.Sort(params)
	return params
}

func (c *Command) Validate() []error {
	errors := []error{}

	if c.Type == CommandInvalid {
		errors = append(errors, fmt.Errorf("command#Type: command type is invalid: (%#v)", c))
	}

	if c.Type == CommandList && c.ListResultFieldName == "" {
		errors = append(errors, fmt.Errorf("command#ListResultFieldName: required when Command#Type is CommandList"))
	} else if c.Type != CommandList && c.ListResultFieldName != "" {
		errors = append(errors, fmt.Errorf("command#ListResultFieldName: can set only when Command#Type is CommandList"))
	}

	if c.Type == CommandList && c.TableType != output.TableSimple {
		errors = append(errors, fmt.Errorf("command#TableType: need output.TableSimple when Command#Type is CommandList"))
	}

	if c.TableType == output.TableSimple && len(c.TableColumnDefines) == 0 {
		errors = append(errors, fmt.Errorf("command#TableColumnDefines: required when Command#TableType is output.TableSimple"))
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

		if v.Category != "" {
			exists := false
			// category is defined on command?
			for _, category := range c.ParamCategories {
				if category.Key == v.Category {
					exists = true
					break
				}
			}
			if !exists {
				err := fmt.Errorf("command#%s.%q: category(%s) isn't defined, but is used by %q", k, "Category", v.Category, k)
				errors = append(errors, err)
			}
		}
	}

	// TODO IsRequiredIDTypeの場合にidパラメータがあるか(id+nameにするかも)

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
		} else {
			return s[i].Param.Order < s[j].Param.Order
		}

	} else {
		return s[i].Category.Order < s[j].Category.Order
	}
}

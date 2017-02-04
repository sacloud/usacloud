package schema

import (
	"fmt"
	"github.com/sacloud/usacloud/output"
)

type Command struct {
	Type                CommandType
	Aliases             []string
	Usage               string
	Params              map[string]*Schema
	AltResource         string // 空の場合はResourceのキーをCamelizeしてsacloud.XXXを対象とする。
	ListResultFieldName string
	SkipAuth            bool
	UseCustomCommand    bool

	TableType          output.OutputTableType
	IncludeFields      []string           // for output.TableDetail
	ExcludeFields      []string           // for output.TableDetail
	TableColumnDefines []output.ColumnDef // for output.TableSimple
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

		// Is the combination of command.Type and params.HandlerType correct?
		if v.HandlerType.IsWhenListOnly() && c.Type != CommandList {
			err := fmt.Errorf("command#%s.%q: cannot use HandlerType(%v) when Command#Type is CommandList", k, "HandlerType", v.HandlerType)
			errors = append(errors, err)
		}
	}

	// TODO IsRequiredIDTypeの場合にidパラメータがあるか

	return errors
}

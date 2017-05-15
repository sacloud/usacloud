package schema

import (
	"fmt"
	"github.com/sacloud/libsacloud/api"
)

type Schema struct {
	Type ValueType

	DefaultValue interface{}
	DefaultText  string
	EnvVars      []string
	InputDefault interface{}
	Aliases      []string
	Description  string // Usage -> cli
	Hidden       bool

	Category string
	Order    int

	Required      bool
	SakuraID      bool
	ConflictsWith []string
	ValidateFunc  ValidateFunc

	MaxItems int
	MinItems int

	HandlerType     HandlerType
	DestinationProp string
	CustomHandler   ValueHandlerFunc

	CompleteFunc CompletionFunc
}

type CompletionContext interface {
	GetAPIClient() *api.Client
	Args() []string
}

type CompletionFunc func(ctx CompletionContext, currentValue string) []string

type ValueHandlerFunc func(name string, src interface{}, dest interface{})

type ValidateFunc func(string, interface{}) []error

func (s *Schema) Validate(name string) []error {
	errs := []error{}

	// Type
	if s.Type == TypeInvalid {
		errs = append(errs, fmt.Errorf("schema#%s.%q: invalid type", name, "Type"))
	}

	// EnvVars
	if len(s.EnvVars) > 0 && s.hasDuplicateValue(s.EnvVars) {
		errs = append(errs, fmt.Errorf("schema#%s.%q: duplicate value", name, "EnvVars"))
	}

	// DefaultValue
	if s.DefaultValue != nil {
		var err error
		tmpl := "schema#%s.\"DefaultValue\": invalid type , need %s"
		switch s.Type {
		case TypeBool:
			if _, ok := s.DefaultValue.(bool); !ok {
				err = fmt.Errorf(tmpl, name, "bool")
			}
		case TypeInt:
			if _, ok := s.DefaultValue.(int); !ok {
				err = fmt.Errorf(tmpl, name, "int")
			}
		case TypeInt64:
			if _, ok := s.DefaultValue.(int64); !ok {
				err = fmt.Errorf(tmpl, name, "int64")
			}

		case TypeFloat:
			if _, ok := s.DefaultValue.(float64); !ok {
				err = fmt.Errorf(tmpl, name, "float64")
			}
		case TypeString:
			if _, ok := s.DefaultValue.(string); !ok {
				err = fmt.Errorf(tmpl, name, "string")
			}
		case TypeIntList:
			if _, ok := s.DefaultValue.([]int64); !ok {
				err = fmt.Errorf(tmpl, name, "[]int64")
			}
		case TypeStringList:
			if _, ok := s.DefaultValue.([]string); !ok {
				err = fmt.Errorf(tmpl, name, "[]string")
			}
		}

		if err != nil {
			errs = append(errs, err)
		}
	}

	// InputDefault
	if s.InputDefault != nil {
		var err error
		tmpl := "schema#%s.\"InputDefault\": invalid type , need %s"
		switch s.Type {
		case TypeBool:
			if _, ok := s.InputDefault.(bool); !ok {
				err = fmt.Errorf(tmpl, name, "bool")
			}
		case TypeInt:
			if _, ok := s.InputDefault.(int); !ok {
				err = fmt.Errorf(tmpl, name, "int")
			}
		case TypeInt64:
			if _, ok := s.InputDefault.(int64); !ok {
				err = fmt.Errorf(tmpl, name, "int64")
			}

		case TypeFloat:
			if _, ok := s.InputDefault.(float64); !ok {
				err = fmt.Errorf(tmpl, name, "float64")
			}
		case TypeString:
			if _, ok := s.InputDefault.(string); !ok {
				err = fmt.Errorf(tmpl, name, "string")
			}
		case TypeIntList:
			if _, ok := s.InputDefault.([]int64); !ok {
				err = fmt.Errorf(tmpl, name, "[]int64")
			}
		case TypeStringList:
			if _, ok := s.InputDefault.([]string); !ok {
				err = fmt.Errorf(tmpl, name, "[]string")
			}
		}

		if err != nil {
			errs = append(errs, err)
		}
	}

	// Aliases
	if len(s.Aliases) > 0 && s.hasDuplicateValue(s.Aliases) {
		errs = append(errs, fmt.Errorf("schema#%s.%q: duplicate value", name, "Aliases"))
	}

	// ConflictsWith
	if len(s.ConflictsWith) > 0 {

		if s.Required {
			errs = append(errs, fmt.Errorf("schema#%s.%q: cannot set with Required=true", name, "ConflictsWith"))
		} else if s.hasDuplicateValue(s.ConflictsWith) {
			errs = append(errs, fmt.Errorf("schema#%s.%q: duplicate value", name, "ConflictsWith"))
		}
	}

	// MaxItems/MinItems
	isList := (s.Type == TypeIntList || s.Type == TypeStringList)
	if isList {
		if s.MaxItems < 0 {
			errs = append(errs, fmt.Errorf("schema#%s.%q: cannot set less than 0", name, "MaxItems"))
		}
		if s.MinItems < 0 {
			errs = append(errs, fmt.Errorf("schema#%s.%q: cannot set less than 0", name, "MinItems"))
		}
		if s.MaxItems > 0 && s.MaxItems < s.MinItems {
			errs = append(errs, fmt.Errorf("schema#%s.%q: cannot set less than MinItems", name, "MaxItems"))
		}
	} else {
		if s.MaxItems != 0 {
			errs = append(errs, fmt.Errorf("schema#%s.%q: cannot set without TypeIntList or TypeStringList", name, "MaxItems"))
		}
		if s.MinItems != 0 {
			errs = append(errs, fmt.Errorf("schema#%s.%q: cannot set without TypeIntList or TypeStringList", name, "MinItems"))
		}

	}

	// ValueType and HandlerType
	if s.HandlerType.IsNeedSliceValue() && !s.Type.IsSliceType() {
		errs = append(errs, fmt.Errorf("schema#%s.%q: needs SliceType(TypeIntList or TypeStringList) in ValueType`", name, "HandlerType"))
	}

	if s.HandlerType == HandlerCustomFunc && s.CustomHandler == nil {
		errs = append(errs, fmt.Errorf("schema#%s.%q: is required when HandlerType is HandlerCustomFunc`", name, "CustomHandler"))
	}

	// DestinationProp
	if s.DestinationProp != "" && !s.HandlerType.CanSetDestinationProp() {
		errs = append(errs, fmt.Errorf("schema#%s.%q: cannot set without HandlerType or `CanSetDestinationProp() == true`", name, "DestinationProp"))
	}

	return errs
}

func (s *Schema) hasDuplicateValue(values []string) bool {
	l := len(values)

	tmp := map[string]bool{}
	for _, v := range values {
		tmp[v] = true
	}

	return l != len(tmp)
}

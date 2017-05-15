package command

import (
	"fmt"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
	"reflect"
	"strings"
	"time"
)

var numericZeros = []interface{}{
	int(0),
	int8(0),
	int16(0),
	int32(0),
	int64(0),
	uint(0),
	uint8(0),
	uint16(0),
	uint32(0),
	uint64(0),
	float32(0),
	float64(0),
}

// isEmpty is copied from github.com/stretchr/testify/assert/assetions.go
func isEmpty(object interface{}) bool {

	if object == nil {
		return true
	} else if object == "" {
		return true
	} else if object == false {
		return true
	}

	for _, v := range numericZeros {
		if object == v {
			return true
		}
	}

	objValue := reflect.ValueOf(object)

	switch objValue.Kind() {
	case reflect.Map:
		fallthrough
	case reflect.Slice, reflect.Chan:
		{
			return (objValue.Len() == 0)
		}
	case reflect.Struct:
		switch object.(type) {
		case time.Time:
			return object.(time.Time).IsZero()
		}
	case reflect.Ptr:
		{
			if objValue.IsNil() {
				return true
			}
			switch object.(type) {
			case *time.Time:
				return object.(*time.Time).IsZero()
			default:
				return false
			}
		}
	}
	return false
}

func validateSakuraID(fieldName string, object interface{}) []error {
	return schema.ValidateSakuraID()(fieldName, object)
}

func validateInStrValues(fieldName string, object interface{}, allowValues ...string) []error {
	return schema.ValidateInStrValues(allowValues...)(fieldName, object)
}

func validateRequired(fieldName string, object interface{}) []error {
	if isEmpty(object) {
		return []error{fmt.Errorf("%q: is required", fieldName)}
	}
	return []error{}
}

func validateSetProhibited(fieldName string, object interface{}) []error {
	if !isEmpty(object) {
		return []error{fmt.Errorf("%q: can't set on current context", fieldName)}
	}
	return []error{}
}

func validateConflicts(fieldName string, object interface{}, values map[string]interface{}) []error {

	if !isEmpty(object) {
		for _, v := range values {
			if !isEmpty(v) {
				keys := []string{}
				for k := range values {
					keys = append(keys, fmt.Sprintf("%q", k))
				}
				return []error{fmt.Errorf("%q: is conflict with %s", fieldName, strings.Join(keys, " or "))}
			}
		}
	}
	return []error{}

}

func validateConflictValues(fieldName string, object interface{}, values map[string]interface{}) []error {

	if !isEmpty(object) {
		for _, v := range values {
			if !isEmpty(v) {
				keys := []string{}
				for k := range values {
					keys = append(keys, fmt.Sprintf("%q", k))
				}
				return []error{fmt.Errorf("%q(%#v): is conflict with %s", fieldName, object, strings.Join(keys, " or "))}
			}
		}
	}
	return []error{}

}

func validateBetween(fieldName string, object interface{}, min int, max int) []error {

	if object == nil {
		object = []int64{}
	}

	isSlice := func(object interface{}) bool {
		_, ok1 := object.([]int64)
		_, ok2 := object.([]string)

		return ok1 || ok2
	}

	if isSlice(object) {
		sliceLen := 0
		if s, ok := object.([]int64); ok {
			sliceLen = len(s)
		} else {
			s := object.([]string)
			sliceLen = len(s)
		}

		if max <= 0 {
			if sliceLen < min {
				return []error{fmt.Errorf("%q: slice length must be %d or more", fieldName, min)}
			}
		} else {
			if !(min <= sliceLen && sliceLen <= max) {
				return []error{fmt.Errorf("%q: slice length must be beetween %d and %d", fieldName, min, max)}
			}

		}
	}

	return []error{}
}

func validateOutputOption(o output.Option) []error {

	outputType := o.GetOutputType()
	columns := o.GetColumn()
	format := o.GetFormat()
	quiet := o.GetQuiet()

	if outputType != "" && format != "" {
		return []error{fmt.Errorf("%q: can't set with --output-format", "--format")}
	}

	if outputType != "" && quiet {
		return []error{fmt.Errorf("%q: can't set with --output-format", "--quiet")}
	}

	if outputType != "tsv" && outputType != "csv" && len(columns) > 0 {
		return []error{fmt.Errorf("%q: can't set when --output-format is csv/tsv", "column")}
	}

	return []error{}

}

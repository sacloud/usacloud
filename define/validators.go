package define

import (
	"fmt"
	"github.com/sacloud/usacloud/schema"
	"net"
	"os"
	"strings"
	"unicode/utf8"
)

func validateMulti(validators ...schema.SchemaValidateFunc) schema.SchemaValidateFunc {
	return func(fieldName string, object interface{}) []error {
		errors := []error{}

		for _, v := range validators {
			errs := v(fieldName, object)
			errors = append(errors, errs...)
		}
		return errors
	}
}

func validateStringSlice(validator schema.SchemaValidateFunc) schema.SchemaValidateFunc {
	return func(fieldName string, object interface{}) []error {
		res := []error{}

		// if target is nil , return OK(Use required attr if necessary)
		if object == nil {
			return res
		}

		if values, ok := object.([]string); ok {
			if len(values) == 0 {
				return res
			}
			for _, v := range values {
				errs := validator(fieldName, v)
				if errs != nil {
					res = append(res, errs...)
				}
			}
		}

		return res
	}
}

func validateIntSlice(validator schema.SchemaValidateFunc) schema.SchemaValidateFunc {
	return func(fieldName string, object interface{}) []error {
		res := []error{}

		// if target is nil , return OK(Use required attr if necessary)
		if object == nil {
			return res
		}

		if values, ok := object.([]int64); ok {
			if len(values) == 0 {
				return res
			}

			for _, v := range values {
				errs := validator(fieldName, v)
				if errs != nil {
					res = append(res, errs...)
				}
			}
		}

		return res
	}
}

// validateStrLen return function is implemented schema.SchemaValidateFunc
func validateStrLen(min int, max int) schema.SchemaValidateFunc {
	return func(fieldName string, object interface{}) []error {
		res := []error{}

		// if target is nil , return OK(Use required attr if necessary)
		if object == nil {
			return res
		}

		if s, ok := object.(string); ok {
			if s == "" {
				return res
			}
			strlen := utf8.RuneCountInString(s)
			if max == 0 {
				if strlen < min {
					res = append(res, fmt.Errorf("%q: String length must be %d or more", fieldName, min))
				}
			} else {
				if !(min <= strlen && strlen <= max) {
					res = append(res, fmt.Errorf("%q: String length must be between %d and %d", fieldName, min, max))
				}
			}
		}

		return res
	}
}

// validateIntRange return function is implementd schema.SchemaValidateFunc
func validateIntRange(min int, max int) schema.SchemaValidateFunc {
	return func(fieldName string, object interface{}) []error {
		res := []error{}

		// if target is nil , return OK(Use required attr if necessary)
		if object == nil {
			return res
		}

		if i, ok := object.(int); ok {
			if i == 0 {
				return res
			}

			if max == 0 {
				if i < min {
					res = append(res, fmt.Errorf("%q: must be %d or more", fieldName, min))
				}
			} else {
				if !(min <= i && i <= max) {
					res = append(res, fmt.Errorf("%q: must be between %d and %d", fieldName, min, max))
				}
			}
		}

		return res
	}
}

func validateInStrValues(allows ...string) schema.SchemaValidateFunc {
	return func(fieldName string, object interface{}) []error {
		res := []error{}

		// if target is nil , return OK(Use required attr if necessary)
		if object == nil {
			return res
		}

		if v, ok := object.(string); ok {
			if v == "" {
				return res
			}

			exists := false
			for _, allow := range allows {
				if v == allow {
					exists = true
					break
				}
			}
			if !exists {
				err := fmt.Errorf("%q: must be in [%s]", fieldName, strings.Join(allows, ","))
				res = append(res, err)
			}
		}
		return res
	}
}

func validateInIntValues(allows ...int) schema.SchemaValidateFunc {
	return func(fieldName string, object interface{}) []error {
		res := []error{}

		// if target is nil , return OK(Use required attr if necessary)
		if object == nil {
			return res
		}

		if v, ok := object.(int); ok {
			if v == 0 {
				return res
			}

			exists := false
			for _, allow := range allows {
				if v == allow {
					exists = true
					break
				}
			}
			if !exists {
				strAllows := []string{}
				for _, allow := range allows {
					strAllows = append(strAllows, fmt.Sprintf("%d", allow))
				}
				err := fmt.Errorf("%q: must be in [%s]", fieldName, strings.Join(strAllows, ","))
				res = append(res, err)
			}
		}
		return res
	}
}

func validateSakuraID() schema.SchemaValidateFunc {
	return func(fieldName string, object interface{}) []error {
		res := []error{}
		idLen := 12

		// if target is nil , return OK(Use required attr if necessary)
		if object == nil {
			return res
		}

		if id, ok := object.(int64); ok {
			if id == 0 {
				return res
			}
			s := fmt.Sprintf("%d", id)
			strlen := utf8.RuneCountInString(s)
			if id < 0 || strlen != idLen {
				res = append(res, fmt.Errorf("%q: Resource ID must be a %d digit number", fieldName, idLen))
			}
		}

		return res
	}
}

func validateSakuraShortID(digit int) schema.SchemaValidateFunc {
	return func(fieldName string, object interface{}) []error {
		res := []error{}
		idLen := digit

		// if target is nil , return OK(Use required attr if necessary)
		if object == nil {
			return res
		}

		if id, ok := object.(int64); ok {
			if id == 0 {
				return res
			}
			s := fmt.Sprintf("%d", id)
			strlen := utf8.RuneCountInString(s)
			if id < 0 || strlen != idLen {
				res = append(res, fmt.Errorf("%q: Resource ID must be a %d digit number", fieldName, idLen))
			}
		}

		return res
	}
}

func mergeParameterMap(params ...map[string]*schema.Schema) map[string]*schema.Schema {
	dest := map[string]*schema.Schema{}
	for _, m := range params {
		for k, v := range m {
			dest[k] = v
		}
	}
	return dest
}

func validateFileExists() schema.SchemaValidateFunc {
	return func(fieldName string, object interface{}) []error {
		res := []error{}

		// if target is nil , return OK(Use required attr if necessary)
		if object == nil {
			return res
		}

		if path, ok := object.(string); ok {
			if path == "" {
				return res
			}

			_, err := os.Stat(path)
			if err != nil {
				res = append(res, fmt.Errorf("%q: File must be exists", fieldName))
			}
		}

		return res
	}
}

func validateIPv4Address() schema.SchemaValidateFunc {
	return func(fieldName string, object interface{}) []error {
		res := []error{}

		// if target is nil , return OK(Use required attr if necessary)
		if object == nil {
			return res
		}

		if value, ok := object.(string); ok {
			if value == "" {
				return res
			}

			ip := net.ParseIP(value)
			if ip == nil || ip.To4() == nil {
				res = append(res, fmt.Errorf("%q: Invalid IPv4 address format", fieldName))
			}
		}

		return res
	}
}

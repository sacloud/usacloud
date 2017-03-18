package schema

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

func ValidateMulti(validators ...SchemaValidateFunc) SchemaValidateFunc {
	return func(fieldName string, object interface{}) []error {
		errors := []error{}

		for _, v := range validators {
			errs := v(fieldName, object)
			errors = append(errors, errs...)
		}
		return errors
	}
}

func ValidateStringSlice(validator SchemaValidateFunc) SchemaValidateFunc {
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

func ValidateIntSlice(validator SchemaValidateFunc) SchemaValidateFunc {
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

func ValidateStrLen(min int, max int) SchemaValidateFunc {
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

func ValidateIntRange(min int, max int) SchemaValidateFunc {
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

func ValidateInStrValues(allows ...string) SchemaValidateFunc {
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

func ValidateInIntValues(allows ...int) SchemaValidateFunc {
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

func ValidateSakuraID() SchemaValidateFunc {
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
				res = append(res, fmt.Errorf("%q: Resource ID must be a %d digits number", fieldName, idLen))
			}
		}

		return res
	}
}

func ValidateSakuraShortID(digit int) SchemaValidateFunc {
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
			if id < 0 || strlen > idLen {
				res = append(res, fmt.Errorf("%q: Resource ID must be less than %d digits", fieldName, idLen))
			}
		}

		return res
	}
}

func ValidateMemberCD() SchemaValidateFunc {
	return func(fieldName string, object interface{}) []error {
		res := []error{}

		// if target is nil , return OK(Use required attr if necessary)
		if object == nil {
			return res
		}

		if cd, ok := object.(string); ok {
			if cd == "" {
				return res
			}

			r := regexp.MustCompile(`^[\w]+[\w\-]*[\w]+$`)
			if !r.MatchString(cd) {
				res = append(res, fmt.Errorf("%q: memberCD must be [0-9A-Za-z_]", fieldName))
			}
		}

		return res
	}
}

func ValidateFileExists() SchemaValidateFunc {
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

func ValidateIPv4Address() SchemaValidateFunc {
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

func ValidateDateTimeString() SchemaValidateFunc {
	return func(fieldName string, object interface{}) []error {
		res := []error{}

		allowDatetimeFormatList := []string{
			time.RFC3339,
		}

		// if target is nil , return OK(Use required attr if necessary)
		if object == nil {
			return res
		}

		if value, ok := object.(string); ok {
			if value == "" {
				return res
			}

			for _, format := range allowDatetimeFormatList {
				_, err := time.Parse(format, value)
				if err == nil {
					// success
					return res
				}
			}

			res = append(res, fmt.Errorf("%q: Invalid Datetime format", fieldName))

		}

		return res
	}
}

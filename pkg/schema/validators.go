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
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func ValidateMulti(validators ...ValidateFunc) ValidateFunc {
	return func(fieldName string, object interface{}) []error {
		errors := []error{}

		for _, v := range validators {
			errs := v(fieldName, object)
			errors = append(errors, errs...)
		}
		return errors
	}
}

func ValidateMultiOr(validators ...ValidateFunc) ValidateFunc {
	return func(fieldName string, object interface{}) []error {
		errors := []error{}

		for _, v := range validators {
			errs := v(fieldName, object)
			errors = append(errors, errs...)
			if len(errors) > 0 {
				return errors
			}
		}
		return errors
	}
}

func ValidateStringSlice(validator ValidateFunc) ValidateFunc {
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

func ValidateIntSlice(validator ValidateFunc) ValidateFunc {
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

func ValidateStrLen(min int, max int) ValidateFunc {
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

func ValidateIntRange(min int, max int) ValidateFunc {
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

func ValidateInStrValues(allows ...string) ValidateFunc {
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

func ValidateInIntValues(allows ...int) ValidateFunc {
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

func ValidateSakuraID() ValidateFunc {
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

func ValidateSakuraShortID(digit int) ValidateFunc {
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

func ValidateMemberCD() ValidateFunc {
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

func ValidateSlackWebhookURL() ValidateFunc {
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

			r := regexp.MustCompile(`^https://hooks.slack.com/services/\w+/\w+/\w+$`)
			if !r.MatchString(cd) {
				res = append(res, fmt.Errorf(`%q: slack webhook url must be ^https://hooks.slack.com/services/\w+/\w+/\w+$`, fieldName))
			}
		}

		return res
	}
}

func ValidateFileExists() ValidateFunc {
	return func(fieldName string, object interface{}) []error {
		res := []error{}

		// if target is nil , return OK(Use required attr if necessary)
		if object == nil {
			return res
		}

		if path, ok := object.(string); ok {
			if path == "" || path == "-" {
				return res
			}

			s, err := os.Stat(path)
			if err == nil {
				if s.Size() == 0 {
					res = append(res, fmt.Errorf("%q: File must not be empty", fieldName))
				}
			} else {
				res = append(res, fmt.Errorf("%q: File must be exists", fieldName))
			}
		}

		return res
	}
}

func ValidateStdinExists() ValidateFunc {
	return func(fieldName string, object interface{}) []error {
		res := []error{}

		// if target is nil , return OK(Use required attr if necessary)
		if object == nil {
			return res
		}

		if path, ok := object.(string); ok {
			if path != "" {
				return res
			}

			fi, err := os.Stdin.Stat()
			if err != nil {
				res = append(res, fmt.Errorf("Opening STDIN is failed"))
			} else {
				if fi.Size() == 0 && fi.Mode()&os.ModeNamedPipe == 0 {
					res = append(res, fmt.Errorf("%q: must not be empty", "STDIN"))
				}
			}
		}

		return res
	}
}

func ValidateIPv4Address() ValidateFunc {
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
			if ip == nil || !strings.Contains(value, ".") {
				res = append(res, fmt.Errorf("%q: Invalid IPv4 address format", fieldName))
			}
		}

		return res
	}
}

func ValidateIPv6Address() ValidateFunc {
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
			if ip == nil || !strings.Contains(value, ":") {
				res = append(res, fmt.Errorf("%q: Invalid IPv6 address format", fieldName))
			}
		}

		return res
	}
}

func ValidateIPv4AddressWithPrefixOption() ValidateFunc {
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

			tokens := strings.Split(value, "/")

			if len(tokens) > 2 {
				res = append(res, fmt.Errorf("%q: Invalid IPv4[Prefix] format", fieldName))
				return res
			}

			ip := net.ParseIP(tokens[0])
			if ip == nil || ip.To4() == nil {
				res = append(res, fmt.Errorf("%q: Invalid IPv4[Prefix] format", fieldName))
				return res
			}

			if len(tokens) == 2 {
				i, e := strconv.Atoi(tokens[1])
				if e != nil {
					res = append(res, fmt.Errorf("%q: Invalid IPv4[Prefix] format", fieldName))
					return res
				}
				if !(1 <= i && i <= 32) {
					res = append(res, fmt.Errorf("%q: Invalid IPv4[Prefix] format", fieldName))
					return res
				}
			}
		}

		return res
	}
}

func ValidateIPv4AddressWithPrefix() ValidateFunc {
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

			tokens := strings.Split(value, "/")

			if len(tokens) != 2 {
				res = append(res, fmt.Errorf("%q: Invalid IPv4+Prefix format", fieldName))
				return res
			}

			ip := net.ParseIP(tokens[0])
			if ip == nil || ip.To4() == nil {
				res = append(res, fmt.Errorf("%q: Invalid IPv4+Prefix format", fieldName))
				return res
			}

			i, e := strconv.Atoi(tokens[1])
			if e != nil {
				res = append(res, fmt.Errorf("%q: Invalid IPv4+Prefix format", fieldName))
				return res
			}
			if !(1 <= i && i <= 32) {
				res = append(res, fmt.Errorf("%q: Invalid IPv4+Prefix format", fieldName))
				return res
			}
		}

		return res
	}
}

func ValidateMACAddress() ValidateFunc {
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

			_, err := net.ParseMAC(value)
			if err != nil {
				res = append(res, fmt.Errorf("%q: Invalid MAC address format", fieldName))
			}
		}

		return res
	}
}

func ValidateDateTimeString() ValidateFunc {
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

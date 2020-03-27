package validation

import (
	"fmt"

	"github.com/sacloud/usacloud/pkg/utils"
)

func IntBetween(fieldName string, object interface{}, min int, max int) []error {
	// if target is empty, return OK(Use required attr if necessary)
	if utils.IsEmpty(object) {
		return []error{}
	}

	v, ok := object.(int)
	if !ok {
		return []error{fmt.Errorf("%q: must be int", fieldName)}
	}

	if !(min <= v && v <= max) {
		return []error{fmt.Errorf("%q: must be between %d and %d", fieldName, min, max)}
	}

	return []error{}
}

func Int64Between(fieldName string, object interface{}, min int64, max int64) []error {
	// if target is empty, return OK(Use required attr if necessary)
	if utils.IsEmpty(object) {
		return []error{}
	}

	v, ok := object.(int64)
	if !ok {
		return []error{fmt.Errorf("%q: must be int64", fieldName)}
	}

	if !(min <= v && v <= max) {
		return []error{fmt.Errorf("%q: must be between %d and %d", fieldName, min, max)}
	}

	return []error{}
}

package validation

import (
	"fmt"
	"strings"

	"github.com/sacloud/usacloud/pkg/utils"
)

func Required(fieldName string, object interface{}) []error {
	if utils.IsEmpty(object) {
		return []error{fmt.Errorf("%q: is required", fieldName)}
	}
	return []error{}
}

func StringInSlice(fieldName string, object interface{}, valid []string) []error {
	var res []error

	// if target is nil , return OK(Use required attr if necessary)
	if object == nil {
		return res
	}

	if v, ok := object.(string); ok {
		if v == "" {
			return res
		}

		exists := false
		for _, allow := range valid {
			if v == allow {
				exists = true
				break
			}
		}
		if !exists {
			err := fmt.Errorf("%q: must be in [%s]", fieldName, strings.Join(valid, ","))
			res = append(res, err)
		}
	}
	return res
}

package validation

import (
	"fmt"

	"github.com/sacloud/usacloud/pkg/utils"
)

func Required(fieldName string, object interface{}) []error {
	if utils.IsEmpty(object) {
		return []error{fmt.Errorf("%q: is required", fieldName)}
	}
	return []error{}
}

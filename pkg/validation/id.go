package validation

import (
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/utils"
)

func ValidID(fieldName string, object interface{}) []error {
	var res []error

	// if target is empty, return OK(Use required attr if necessary)
	if utils.IsEmpty(object) {
		return res
	}

	var id types.ID
	switch v := object.(type) {
	case int64:
		id = types.ID(v)
	case string:
		id = types.StringID(v)
	default:
		res = append(res, fmt.Errorf("%q: Resource ID must be valid format", fieldName))
		return res
	}

	if id.IsEmpty() {
		res = append(res, fmt.Errorf("%q: Resource ID must be valid format", fieldName))
	}

	return res
}

package validation

import (
	"fmt"
	"strings"

	"github.com/sacloud/usacloud/pkg/utils"
)

func ConflictsWith(fieldName string, object interface{}, values map[string]interface{}) []error {
	if !utils.IsEmpty(object) {
		for _, v := range values {
			if !utils.IsEmpty(v) {
				var keys []string
				for k := range values {
					keys = append(keys, fmt.Sprintf("%q", k))
				}
				return []error{fmt.Errorf("%q: is conflict with %s", fieldName, strings.Join(keys, " or "))}
			}
		}
	}
	return []error{}

}

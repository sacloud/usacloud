package validation

import "fmt"

func SliceLenBetween(fieldName string, object interface{}, min int, max int) []error {

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

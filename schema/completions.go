package schema

import (
	"fmt"
)

func CompleteInStrValues(values ...string) CompletionFunc {
	return func(ctx CompletionContext, currentValue string) []string {
		return values
	}
}

func CompleteInIntValues(values ...int) CompletionFunc {
	return func(ctx CompletionContext, currentValue string) []string {
		res := []string{}
		for _, v := range values {
			res = append(res, fmt.Sprintf("%d", v))
		}
		return res
	}
}

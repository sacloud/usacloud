// Copyright 2017-2019 The Usacloud Authors
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

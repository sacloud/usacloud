// Copyright 2016-2019 The Libsacloud Authors
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

package builder

import (
	"fmt"
	"strings"
)

type baseBuilder struct {
	client APIClient
	errors []error
}

func (b *baseBuilder) toStringList(values []int64) []string {
	keys := []string{}
	for _, k := range values {
		keys = append(keys, fmt.Sprintf("%d", k))
	}
	return keys
}

func (b *baseBuilder) getFlattenErrors() error {
	var list = make([]string, 0)
	for _, str := range b.errors {
		list = append(list, str.Error())
	}
	return fmt.Errorf(strings.Join(list, "\n"))
}

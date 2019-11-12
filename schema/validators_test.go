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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateIPv4(t *testing.T) {

	expects := map[string]bool{
		"192.168.0.1":                  true,
		"192.168.0.1.1":                false,
		"192.168.0.0/24":               false,
		"2401:2500:10a:100e::1":        false,
		"fe::1":                        false,
		"2401:2500:10a:100e::xxxxxxxx": false,
	}

	for target, expect := range expects {

		errs := ValidateIPv4Address()("TEST", target)
		actual := len(errs) == 0

		assert.Equal(t, actual, expect, "Target: %s", target)
	}

}

func TestValidateIPv6(t *testing.T) {

	expects := map[string]bool{
		"192.168.0.1":                  false,
		"2401:2500:10a:100e::1":        true,
		"fe::1":                        true,
		"2401:2500:10a:100e::xxxxxxxx": false,
	}

	for target, expect := range expects {

		errs := ValidateIPv6Address()("TEST", target)
		actual := len(errs) == 0

		assert.Equal(t, actual, expect, "Target: %s", target)
	}

}

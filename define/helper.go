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

package define

import (
	"fmt"

	"github.com/sacloud/usacloud/schema"
)

func validateMulti(validators ...schema.ValidateFunc) schema.ValidateFunc {
	return schema.ValidateMulti(validators...)
}
func validateMultiOr(validators ...schema.ValidateFunc) schema.ValidateFunc {
	return schema.ValidateMultiOr(validators...)
}
func validateStringSlice(validator schema.ValidateFunc) schema.ValidateFunc {
	return schema.ValidateStringSlice(validator)
}

func validateIntSlice(validator schema.ValidateFunc) schema.ValidateFunc {
	return schema.ValidateIntSlice(validator)
}

func validateStrLen(min int, max int) schema.ValidateFunc {
	return schema.ValidateStrLen(min, max)
}

func validateIntRange(min int, max int) schema.ValidateFunc {
	return schema.ValidateIntRange(min, max)
}

func validateInStrValues(allows ...string) schema.ValidateFunc {
	return schema.ValidateInStrValues(allows...)
}

func validateInIntValues(allows ...int) schema.ValidateFunc {
	return schema.ValidateInIntValues(allows...)
}

func validateSakuraID() schema.ValidateFunc {
	return schema.ValidateSakuraID()
}

func validateSakuraShortID(digit int) schema.ValidateFunc {
	return schema.ValidateSakuraShortID(digit)
}

func validateMemberCD() schema.ValidateFunc {
	return schema.ValidateMemberCD()
}

func validateSlackWebhookURL() schema.ValidateFunc {
	return schema.ValidateSlackWebhookURL()
}

func validateFileExists() schema.ValidateFunc {
	return schema.ValidateFileExists()
}

func validateExistsFileOrStdIn() schema.ValidateFunc {
	return schema.ValidateMultiOr(schema.ValidateFileExists(), schema.ValidateStdinExists())
}

func validateIPv4Address() schema.ValidateFunc {
	return schema.ValidateIPv4Address()
}

func validateIPv4AddressWithPrefixOption() schema.ValidateFunc {
	return schema.ValidateIPv4AddressWithPrefixOption()
}

func validateIPv4AddressWithPrefix() schema.ValidateFunc {
	return schema.ValidateIPv4AddressWithPrefix()
}

func validateMACAddress() schema.ValidateFunc {
	return schema.ValidateMACAddress()
}

func validateDateTimeString() schema.ValidateFunc {
	return schema.ValidateDateTimeString()
}

func validateBackupTime() schema.ValidateFunc {
	timeStrings := []string{}

	minutes := []int{0, 15, 30, 45}

	// create list [00:00 ,,, 23:45]
	for hour := 0; hour <= 23; hour++ {
		for _, minute := range minutes {
			timeStrings = append(timeStrings, fmt.Sprintf("%02d:%02d", hour, minute))
		}
	}

	return schema.ValidateInStrValues(timeStrings...)
}

func mergeParameterMap(params ...map[string]*schema.Schema) map[string]*schema.Schema {
	dest := map[string]*schema.Schema{}
	for _, m := range params {
		for k, v := range m {
			dest[k] = v
		}
	}
	return dest
}

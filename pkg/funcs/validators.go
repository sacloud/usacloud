// Copyright 2017-2020 The Usacloud Authors
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

package funcs

import (
	"github.com/sacloud/usacloud/pkg/cli"
)

func validateInStrValues(fieldName string, object interface{}, allowValues ...string) []error {
	return cli.ValidateInStrValues(fieldName, object, allowValues...)
}

func validateRequired(fieldName string, object interface{}) []error {
	return cli.ValidateRequired(fieldName, object)
}

func validateSetProhibited(fieldName string, object interface{}) []error {
	return cli.ValidateSetProhibited(fieldName, object)
}

func validateConflictValues(fieldName string, object interface{}, values map[string]interface{}) []error {
	return cli.ValidateConflictValues(fieldName, object, values)
}

func validateIPv4AddressArgs(ipaddr string) []error {
	return cli.ValidateIPv4Address("Args", ipaddr)
}

func validateIPv6AddressArgs(ipaddr string) []error {
	return cli.ValidateIPv6Address("Args", ipaddr)
}

func validateExistsFileOrStdIn(fieldName string, filePath string) []error {
	return cli.ValidateExistsFileOrStdIn(fieldName, filePath)
}

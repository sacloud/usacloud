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

package params

import (
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/output"
)

func validateSakuraID(fieldName string, object interface{}) []error {
	return cli.ValidateSakuraID(fieldName, object)
}

func validateRequired(fieldName string, object interface{}) []error {
	return cli.ValidateRequired(fieldName, object)
}

func validateOutputOption(o output.Option) []error {
	return cli.ValidateOutputOption(o)
}

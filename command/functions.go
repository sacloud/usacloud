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

package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/utils"
)

func FlattenErrors(errors []error) error {
	if len(errors) == 0 {
		return nil
	}
	var list = make([]string, 0)
	for _, str := range errors {
		list = append(list, str.Error())
	}
	return fmt.Errorf(strings.Join(list, "\n"))
}

func FlattenErrorsWithPrefix(errors []error, pref string) error {
	var list = make([]string, 0)
	for _, str := range errors {
		list = append(list, fmt.Sprintf("[%s] : %s", pref, str.Error()))
	}
	return fmt.Errorf(strings.Join(list, "\n"))

}

func Confirm(in *os.File, msg string) bool {
	fi, err := in.Stat()
	if err != nil {
		panic(err)
	}
	if fi.Size() > 0 {
		return true
	}
	fmt.Printf("\n%s(y/n) [n]: ", msg)

	var input string
	fmt.Fscanln(in, &input)
	return input == "y" || input == "yes"
}

func ConfirmContinue(in *os.File, target string, ids ...sacloud.ID) bool {
	if len(ids) == 0 {
		return Confirm(in, fmt.Sprintf("Are you sure you want to %s?", target))
	}

	strIDs := utils.StringIDs(ids)
	msg := fmt.Sprintf("Target resource IDs => [\n\t%s\n]", strings.Join(strIDs, ",\n\t"))
	return Confirm(in, fmt.Sprintf("%s\nAre you sure you want to %s?", msg, target))
}

func IsSetOr(ctx Context, targetes ...string) bool {
	for _, target := range targetes {
		if ctx.IsSet(target) {
			return true
		}
	}
	return false
}

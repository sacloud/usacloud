// Copyright 2017-2025 The sacloud/usacloud Authors
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

package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type In interface {
	Stat() (os.FileInfo, error)
	io.Reader
}

func Confirm(msg string, in In, out io.Writer) (bool, error) {
	fi, err := in.Stat()
	if err != nil {
		return false, err
	}
	if fi.Size() > 0 {
		return true, nil
	}
	_, err = fmt.Fprintf(out, "\n%s(y/n) [n]: ", msg)
	if err != nil {
		return false, err
	}

	scanner := bufio.NewScanner(in)
	scanner.Scan()
	input := scanner.Text()

	return input == "y" || input == "yes", nil
}

func ConfirmContinue(target string, in In, out io.Writer, ids ...string) (bool, error) {
	msg := fmt.Sprintf("Are you sure you want to %s?", target)
	if len(ids) > 0 {
		msg = fmt.Sprintf("Target resource IDs => [\n\t%s\n]\nAre you sure you want to %s?",
			strings.Join(ids, ",\n\t"),
			target,
		)
	}
	return Confirm(msg, in, out)
}

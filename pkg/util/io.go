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

import "os"

// FileOrStdin ファイルパス、または標準入力をオープンする
//
// pathに空文字または"-"を指定した場合に標準入力が利用される
func FileOrStdin(path string) (file *os.File, deferFunc func(), err error) {
	if path == "" || path == "-" {
		file = os.Stdin
		deferFunc = func() {}
	} else {
		file, err = os.Open(path)
		if err != nil {
			return
		}
		deferFunc = func() {
			file.Close()
		}
	}
	return
}

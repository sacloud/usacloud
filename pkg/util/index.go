// Copyright 2017-2022 The sacloud/usacloud Authors
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

import "fmt"

// ValidIndex インデックスが有効範囲内にあるか
//
// TODO v0では1開始だったが、v1では0開始とする(GH-552)
func ValidIndex(index, length int) error {
	if index <= 0 || index-1 >= length {
		return fmt.Errorf("index(%d) is out of range", index)
	}
	return nil
}

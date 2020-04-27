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

package weekday

import "github.com/sacloud/libsacloud/v2/sacloud/types"

func FromStrings(ss []string) []types.EBackupSpanWeekday {
	var strWeekdays []string
	exists := false
	for _, v := range ss {
		if v == "all" {
			exists = true
			break
		}
	}
	if exists {
		strWeekdays = types.BackupWeekdayStrings
	}

	var weekdays []types.EBackupSpanWeekday
	for _, s := range strWeekdays {
		weekdays = append(weekdays, StrToWeekday(s))
	}
	return weekdays
}

func StrToWeekday(s string) types.EBackupSpanWeekday {
	return types.EBackupSpanWeekday(s)
}

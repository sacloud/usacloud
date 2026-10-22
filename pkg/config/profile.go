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

package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"

	"github.com/sacloud/saclient-go"
	"github.com/spf13/pflag"
)

func (o *Config) loadFromProfile(flags *pflag.FlagSet, errW io.Writer) {
	if flags.Changed("profile") {
		v, err := flags.GetString("profile")
		if err != nil {
			fmt.Fprintf(errW, "[WARN] reading value of %q flag is failed: %s\n", "profile", err)
			return
		}
		o.Profile = v
	}

	op, err := saclient.NewProfileOp(os.Environ())
	if err != nil {
		fmt.Fprintf(errW, "[WARN] loading profile %q is failed: %s\n", o.Profile, err)
		return
	}

	profileName := o.Profile
	if profileName == "" {
		current, err := op.GetCurrentName()
		switch {
		case errors.Is(err, fs.ErrNotExist):
			// プロファイルディレクトリまたは current ファイルが存在しない場合は
			// default プロファイルを使用する
			profileName = DefaultProfileName
		case err != nil:
			fmt.Fprintf(errW, "[WARN] loading profile %q is failed: %s\n", profileName, err)
			return
		case current == "":
			// current ファイルが空または空白のみの場合も default を使用する
			profileName = DefaultProfileName
		default:
			profileName = current
		}
	}

	loaded, err := op.Read(profileName)
	switch {
	case errors.Is(err, fs.ErrNotExist):
		// default プロファイルはファイルが存在しなくてもエラーにしない
		if profileName != DefaultProfileName {
			fmt.Fprintf(errW, "[WARN] loading profile %q is failed: profile %q is not exists\n", profileName, profileName)
		}
		o.Profile = profileName
		return
	case err != nil:
		fmt.Fprintf(errW, "[WARN] loading profile %q is failed: %s\n", profileName, err)
		o.Profile = profileName
		return
	}
	// プロファイルの内容(JSON)をConfigへ反映する
	if buf, err := json.Marshal(loaded.Attributes); err != nil {
		fmt.Fprintf(errW, "[WARN] loading profile %q is failed: %s\n", profileName, err)
		return
	} else if err := json.Unmarshal(buf, o); err != nil {
		fmt.Fprintf(errW, "[WARN] loading profile %q is failed: %s\n", profileName, err)
		return
	}
	o.Profile = profileName
}

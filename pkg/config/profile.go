// Copyright 2017-2021 The Usacloud Authors
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

// +build !wasm

package config

import (
	"fmt"
	"io"

	"github.com/sacloud/libsacloud/v2/sacloud/profile"
	"github.com/spf13/pflag"
)

func (o *Config) loadFromProfile(flags *pflag.FlagSet, errW io.Writer) {
	if flags.Changed("profile") {
		v, err := flags.GetString("profile")
		if err != nil {
			fmt.Fprintf(errW, "[WARN] reading value of %q flag is failed: %s", "profile", err) // nolint
			return
		}
		o.Profile = v
	}

	profileName := o.Profile
	if profileName == "" {
		current, err := profile.CurrentName()
		if err != nil {
			fmt.Fprintf(errW, "[WARN] loading profile %q is failed: %s", profileName, err) // nolint
			return
		}
		profileName = current
	}
	if err := profile.Load(profileName, o); err != nil {
		fmt.Fprintf(errW, "[WARN] loading profile %q is failed: %s", profileName, err) // nolint
		return
	}
}

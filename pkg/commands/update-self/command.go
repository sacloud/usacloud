//go:build !wasm
// +build !wasm

// Copyright 2017-2022 The Usacloud Authors
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

package updateself

import (
	"fmt"
	"os"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"github.com/sacloud/usacloud/pkg/version"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:                   "update-self",
	Short:                 "Update Usacloud to latest-stable version",
	DisableFlagsInUseLine: true,
	Args:                  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if err := updateSelf(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func updateSelf() error {
	rl, ok, err := selfupdate.DetectLatest("sacloud/usacloud")
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("no release detected. %s is considered up-to-date", version.Version)
	}
	current := semver.MustParse(version.Version)

	if rl.Version.Equals(current) {
		fmt.Println("Usacloud is up-to-date with ", version.Version)
		return nil
	}

	cmdPath, err := os.Executable()
	if err != nil {
		return err
	}

	if err := selfupdate.UpdateTo(rl.AssetURL, cmdPath); err != nil {
		return err
	}
	fmt.Println("Successfully updated to version", rl.Version.String())
	return nil
}

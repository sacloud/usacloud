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

package migration

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/helper/printer"
	"github.com/sacloud/usacloud/pkg/profile"
)

const OldConfigFileName = ".usacloud_config"

var (
	confirmMigrateFunc   = confirmMigrate
	confirmOverwriteFunc = confirmOverwrite
)

var (
	migrateInfoMsg = `
================================================================================
Warning: old-style configuration file(%q) found.
         please run 'usacloud config migrate' command.
================================================================================

`
	confirmMigrateMsg = `Are you sure you want to migrate old-style configuration file?`

	confirmOverwriteMsg = `Default profile config[%q] already exists.
Do you want to replace it?`
)

func getOldConfigPath() (string, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, OldConfigFileName), nil
}

func confirmMigrate() bool {
	return cli.Confirm(os.Stdin, confirmMigrateMsg) // TODO ビルドを通すための仮実装
}

func confirmOverwrite(configFilePath string) bool {
	return cli.Confirm(os.Stdin, fmt.Sprintf(confirmOverwriteMsg, configFilePath)) // TODO ビルドを通すための仮実装
}

func CheckConfigVersion() error {
	src, err := getOldConfigPath()
	if err != nil {
		return fmt.Errorf("Getting old config path is failed: %s", err)
	}

	// exists?
	if _, err := os.Stat(src); err == nil {
		fmt.Fprintf(os.Stdout, migrateInfoMsg, src) // TODO ビルドを通すための仮実装
	}

	return nil
}

func MigrateConfig() error {
	src, err := getOldConfigPath()
	if err != nil {
		return fmt.Errorf("Getting old config path is failed: %s", err)
	}

	// exists?
	if _, err := os.Stat(src); err != nil {
		// not exists, noop
		return nil
	}

	if !confirmMigrateFunc() {
		return nil
	}

	// profile[default] exists?
	dest, err := profile.GetConfigFilePath(profile.DefaultProfileName)
	if err != nil {
		return fmt.Errorf("Getting new config path is failed: %s", err)
	}
	if _, err := os.Stat(dest); err == nil {
		// exists, show confirm dialog
		if !confirmOverwriteFunc(dest) {
			return nil
		}
	}

	// prepare dest dir
	destDir := filepath.Dir(dest)
	if _, err := os.Stat(destDir); err != nil {
		err := os.MkdirAll(destDir, 0755)
		if err != nil {
			return fmt.Errorf("Can't create destination dir[%q]: %s", destDir, err)
		}
	}

	// move src to dest
	if err := os.Rename(src, dest); err != nil {
		return fmt.Errorf("Migrating [%q] to [%q] is failed: %s", src, dest, err)
	}

	printer.Fprintf(os.Stdout, color.New(color.FgGreen), "\nMigrated: [%q] to [%q]\n", src, dest) // TODO ビルドを通すための仮実装
	return nil
}
